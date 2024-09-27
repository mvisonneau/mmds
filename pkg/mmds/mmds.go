package mmds

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	log "github.com/sirupsen/logrus"
)

// Instance ..
type Instance struct {
	*ec2Types.Instance
}

// New MMDS client.
func New() (i Instance, err error) {
	// Get an MDS client.
	var mdsClient *imds.Client
	if mdsClient, err = newAWSMDSClient(); err != nil {
		return
	}

	// Find instance region.
	var region string
	if region, err = getInstanceRegion(mdsClient); err != nil {
		return
	}

	// Find instance ID.
	var instanceID *string
	if instanceID, err = getInstanceID(mdsClient); err != nil {
		return
	}

	// Get an EC2 client.
	var ec2Client *ec2.Client
	if ec2Client, err = newAWSEC2Client(region); err != nil {
		return
	}

	// Find instance details.
	var instance *ec2Types.Instance
	if instance, err = getInstance(ec2Client, instanceID); err != nil {
		return
	}

	if instance == nil {
		return i, fmt.Errorf("instance-id not found from EC2 API (nil object)")
	}

	i = Instance{instance}
	return
}

// GetPricingModel of the instance.
func (i *Instance) GetPricingModel() string {
	if i.Instance.InstanceLifecycle == ec2Types.InstanceLifecycleTypeSpot {
		return "spot"
	}

	return "on-demand"
}

func newAWSMDSClient() (c *imds.Client, err error) {
	log.Debug("Starting AWS MDS API session")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	c = imds.NewFromConfig(cfg)

	_, err = c.GetMetadata(context.TODO(), &imds.GetMetadataInput{
		Path: "instance-id",
	})
	if err != nil {
		err = fmt.Errorf("Unable to access the metadata service, are you running this binary from an AWS EC2 instance?")
	}

	return
}

func newAWSEC2Client(region string) (c *ec2.Client, err error) {
	re := regexp.MustCompile("[a-z]{2}-[a-z]+-\\d")
	if !re.MatchString(region) {
		err = fmt.Errorf("Cannot start AWS EC2 client session with invalid region '%s'", region)
		return
	}

	log.Debug("Starting AWS EC2 Client session")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	c = ec2.NewFromConfig(cfg)
	return
}

func getInstanceAZ(c *imds.Client) (az string, err error) {
	log.Debug("Fetching current AZ from MDS API")
	output, err := c.GetMetadata(context.TODO(), &imds.GetMetadataInput{
		Path: "placement/availability-zone",
	})
	if err != nil {
		return "", err
	}

	buf := new(strings.Builder)
	if _, err = io.Copy(buf, output.Content); err != nil {
		return "", err
	}

	az = buf.String()
	log.Debugf("Found AZ: '%s'", az)
	return
}

func computeRegionFromAZ(az string) (region string, err error) {
	re := regexp.MustCompile("[a-z]{2}-[a-z]+-\\d[a-z]")
	if !re.MatchString(az) {
		err = fmt.Errorf("Cannot compute region from invalid availability-zone '%s'", az)
		return
	}

	region = az[:len(az)-1]
	log.Debugf("Computed region : '%s'", region)
	return
}

func getInstanceRegion(c *imds.Client) (region string, err error) {
	// Fetch current AZ
	var az string
	az, err = getInstanceAZ(c)
	if err != nil {
		return
	}

	// Compute region from AZ
	region, err = computeRegionFromAZ(az)
	return
}

func getInstanceID(c *imds.Client) (*string, error) {
	log.Debug("Fetching current instance-id from MDS API")
	output, err := c.GetMetadata(context.TODO(), &imds.GetMetadataInput{
		Path: "instance-id",
	})
	if err != nil {
		return nil, err
	}

	buf := new(strings.Builder)
	if _, err = io.Copy(buf, output.Content); err != nil {
		return nil, err
	}

	instanceID := buf.String()
	log.Debugf("Found instance-id : '%s'", instanceID)
	return &instanceID, nil
}

func getInstance(c *ec2.Client, instanceID *string) (*ec2Types.Instance, error) {
	log.Debugf("Fetching instance object from instanceID '%s' from EC2 API", *instanceID)
	instances, err := c.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
		Filters: []ec2Types.Filter{
			{
				Name:   aws.String("instance-id"),
				Values: []string{*instanceID},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(instances.Reservations) != 1 {
		return nil, fmt.Errorf("Unexpected amount of reservations retrieved : '%d', expected 1", len(instances.Reservations))
	}

	if len(instances.Reservations[0].Instances) != 1 {
		return nil, fmt.Errorf("Unexpected amount of instances retrieved : '%d', expected 1", len(instances.Reservations[0].Instances))
	}

	return &instances.Reservations[0].Instances[0], nil
}
