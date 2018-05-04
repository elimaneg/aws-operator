// +build k8srequired

package draining

import (
	"testing"

	"github.com/giantswarm/e2e-harness/pkg/framework"
	"github.com/giantswarm/e2eclients/aws"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/aws-operator/integration/setup"
)

var (
	g *framework.Guest
	h *framework.Host
	c *aws.Client
)

// TestMain allows us to have common setup and teardown steps that are run
// once for all the tests https://golang.org/pkg/testing/#hdr-Main.
func TestMain(m *testing.M) {
	var err error

	logger, err := micrologger.New(micrologger.Config{})
	if err != nil {
		panic(err.Error())
	}

	{
		c := framework.GuestConfig{
			Logger: logger,
		}

		g, err = framework.NewGuest(c)
		if err != nil {
			panic(err.Error())
		}
	}

	{
		c := framework.HostConfig{}

		h, err = framework.NewHost(c)
		if err != nil {
			panic(err.Error())
		}
	}

	c = aws.NewClient()

	setup.WrapTestMain(c, g, h, m)
}