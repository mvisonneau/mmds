package mmds

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
)

func TestGetPricingModel(t *testing.T) {
	i := Instance{
		&ec2.Instance{
			InstanceLifecycle: pointy.String("boooo"),
		},
	}

	assert.Equal(t, "on-demand", i.GetPricingModel())

	i.InstanceLifecycle = pointy.String("spot")
	assert.Equal(t, "spot", i.GetPricingModel())
}
