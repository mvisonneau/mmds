package mmds

import (
	"testing"

	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/stretchr/testify/assert"
)

func TestGetPricingModel(t *testing.T) {
	i := Instance{
		&ec2Types.Instance{
			InstanceLifecycle: ec2Types.InstanceLifecycleTypeCapacityBlock,
		},
	}

	assert.Equal(t, "on-demand", i.GetPricingModel())

	i.InstanceLifecycle = ec2Types.InstanceLifecycleTypeSpot
	assert.Equal(t, "spot", i.GetPricingModel())
}
