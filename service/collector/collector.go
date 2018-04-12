package collector

import (
	awsutil "github.com/giantswarm/aws-operator/client/aws"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

const (
	Namespace = "aws_operator"

	gaugeValue = float64(1)
)

type Config struct {
	Logger micrologger.Logger

	AwsConfig awsutil.Config
}

type Collector struct {
	logger micrologger.Logger

	awsClients awsutil.Clients
}

func New(config Config) (*Collector, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var emptyAwsConfig awsutil.Config
	if config.AwsConfig == emptyAwsConfig {
		return nil, microerror.Maskf(invalidConfigError, "%T.AwsConfig must not be empty", config)
	}

	awsClients := awsutil.NewClients(config.AwsConfig)

	c := &Collector{
		logger: config.Logger,

		awsClients: awsClients,
	}

	return c, nil
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- vpcs
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.logger.Log("level", "debug", "message", "collecting metrics")

	c.collectVPCs(ch)

	c.logger.Log("level", "debug", "message", "finished collecting metrics")
}