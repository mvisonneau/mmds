package main

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Clients of AWS libs
type Clients struct {
	EC2 *ec2.EC2
	MDS *ec2metadata.EC2Metadata
}

// Values computed/generated
type Values struct {
	InstanceID   string
	PricingModel string
	Region       string
	AZ           string
}

var start time.Time

func run(ctx *cli.Context) error {
	start = time.Now()
	configureLogging(cfg.Log.Level, cfg.Log.Format)

	c := &Clients{}
	v := &Values{}

	// Configure MDS Client
	if err := c.getAWSMDSClient(); err != nil {
		return exit(cli.NewExitError(err.Error(), 1))
	}

	// Fetch current AZ
	var err error
	v.AZ, err = c.getInstanceAZ()
	if err != nil {
		return exit(cli.NewExitError(err.Error(), 1))
	}

	// Compute region from AZ
	v.Region, err = computeRegionFromAZ(v.AZ)
	if err != nil {
		return exit(cli.NewExitError(err.Error(), 1))
	}

	// Configure EC2 Client
	if err := c.getAWSEC2Client(v.Region); err != nil {
		return exit(cli.NewExitError(err.Error(), 1))
	}

	// Fetch instance ID
	v.InstanceID, err = c.getInstanceID()
	if err != nil {
		return exit(cli.NewExitError(err.Error(), 1))
	}

	// Retrieve instance object from AWS
	instance, err := c.getInstance(v.InstanceID)
	if err != nil {
		return exit(cli.NewExitError(analyzeEC2APIErrors(err), 1))
	}

	if instance.InstanceLifecycle != nil && *instance.InstanceLifecycle == "spot" {
		v.PricingModel = "spot"
	} else {
		v.PricingModel = "on-demand"
	}

	switch ctx.Command.FullName() {
	case "pricing-model":
		fmt.Println(v.PricingModel)
	default:
		return exit(cli.NewExitError(fmt.Sprintf("Function %v is not implemented", ctx.Command.FullName()), 1))
	}

	return exit(nil)
}

func (c *Clients) getAWSMDSClient() error {
	log.Debug("Starting AWS MDS API session")
	c.MDS = ec2metadata.New(session.New())

	if !c.MDS.Available() {
		return errors.New("Unable to access the metadata service, are you running this binary from an AWS EC2 instance?")
	}

	return nil
}

func (c *Clients) getAWSEC2Client(region string) (err error) {
	re := regexp.MustCompile("[a-z]{2}-[a-z]+-\\d")
	if !re.MatchString(region) {
		return fmt.Errorf("Cannot start AWS EC2 client session with invalid region '%s'", region)
	}

	log.Debug("Starting AWS EC2 Client session")
	c.EC2 = ec2.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))
	return
}

func (c *Clients) getInstanceAZ() (az string, err error) {
	log.Debug("Fetching current AZ from MDS API")
	az, err = c.MDS.GetMetadata("placement/availability-zone")
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

func (c *Clients) getInstanceID() (iid string, err error) {
	log.Debug("Fetching current instance-id from MDS API")
	iid, err = c.MDS.GetMetadata("instance-id")
	log.Debugf("Found instance-id : '%s'", iid)
	return
}

func (c *Clients) getInstance(instanceID string) (*ec2.Instance, error) {
	log.Debugf("Fetching instance object from instanceID '%s' from EC2 API", instanceID)
	instances, err := c.EC2.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-id"),
				Values: []*string{
					aws.String(instanceID),
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

func analyzeEC2APIErrors(err error) string {
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return aerr.Error()
		}
		return err.Error()
	}
	return ""
}

func exit(err error) error {
	log.Debugf("Executed in %s, exiting..", time.Since(start))
	return err
}
