package mmds

import (
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	log "github.com/sirupsen/logrus"
)

// Instance ..
type Instance struct {
	*ec2.Instance
}

// New MMDS client
func New() (i Instance, err error) {
	// Get an MDS client
	var mdsClient *ec2metadata.EC2Metadata
	if mdsClient, err = newAWSMDSClient(); err != nil {
		return
	}

	// Find instance region
	var region string
	if region, err = getInstanceRegion(mdsClient); err != nil {
		return
	}

	// Find instance ID
	var instanceID *string
	if instanceID, err = getInstanceID(mdsClient); err != nil {
		return
	}

	// Get an EC2 client
	var ec2Client *ec2.EC2
	if ec2Client, err = newAWSEC2Client(region); err != nil {
		return
	}

	// Find instance details
	var instance *ec2.Instance
	if instance, err = getInstance(ec2Client, instanceID); err != nil {
		return
	}

	if instance == nil {
		return i, fmt.Errorf("instance-id not found from EC2 API (nil object)")
	}

	i = Instance{instance}
	return
}

// GetPricingModel of the instance
func (i *Instance) GetPricingModel() string {
	if i.InstanceLifecycle != nil && *i.InstanceLifecycle == "spot" {
		return "spot"
	}

	return "on-demand"
}

func newAWSMDSClient() (c *ec2metadata.EC2Metadata, err error) {
	log.Debug("Starting AWS MDS API session")
	c = ec2metadata.New(session.New())

	if !c.Available() {
		err = fmt.Errorf("Unable to access the metadata service, are you running this binary from an AWS EC2 instance?")
	}

	return
}

func newAWSEC2Client(region string) (c *ec2.EC2, err error) {
	re := regexp.MustCompile("[a-z]{2}-[a-z]+-\\d")
	if !re.MatchString(region) {
		err = fmt.Errorf("Cannot start AWS EC2 client session with invalid region '%s'", region)
		return
	}

	log.Debug("Starting AWS EC2 Client session")
	c = ec2.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))
	return
}

func getInstanceAZ(c *ec2metadata.EC2Metadata) (az string, err error) {
	log.Debug("Fetching current AZ from MDS API")
	az, err = c.GetMetadata("placement/availability-zone")
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

func getInstanceRegion(c *ec2metadata.EC2Metadata) (region string, err error) {
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

func getInstanceID(c *ec2metadata.EC2Metadata) (*string, error) {
	log.Debug("Fetching current instance-id from MDS API")
	instanceID, err := c.GetMetadata("instance-id")
	log.Debugf("Found instance-id : '%s'", instanceID)
	return &instanceID, err
}

func getInstance(c *ec2.EC2, instanceID *string) (*ec2.Instance, error) {
	log.Debugf("Fetching instance object from instanceID '%s' from EC2 API", *instanceID)
	instances, err := c.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-id"),
				Values: []*string{
					instanceID,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(instances.Reservations) != 1 {
		return nil, fmt.Errorf("Unexpected amount of reservations retrieved : '%d',  expected 1", len(instances.Reservations))
	}

	if len(instances.Reservations[0].Instances) != 1 {
		return nil, fmt.Errorf("Unexpected amount of reservations retrieved : '%d',  expected 1", len(instances.Reservations[0].Instances))
	}

	return instances.Reservations[0].Instances[0], nil
}
