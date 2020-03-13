package tccpn

import (
	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/aws-operator/service/controller/internal/changedetection"
)

const (
	// Name is the identifier of the resource.
	Name = "tccpn"
)

type Config struct {
	G8sClient versioned.Interface
	Detection *changedetection.TCCPN
	Logger    micrologger.Logger

	APIWhitelist     APIWhitelist
	EncrypterBackend string
	InstallationName string
}

// Resource implements the TCCPN resource, which stands for Tenant Cluster Control
// Plane Node. We manage a dedicated Cloud Formation stack for each node pool.
type Resource struct {
	g8sClient versioned.Interface
	detection *changedetection.TCCPN
	logger    micrologger.Logger

	apiWhitelist     APIWhitelist
	encrypterBackend string
	installationName string
}

func New(config Config) (*Resource, error) {
	if config.G8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.G8sClient must not be empty", config)
	}
	if config.Detection == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Detection must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.APIWhitelist.Private.Enabled && config.APIWhitelist.Private.SubnetList == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.APIWhitelist.Private.SubnetList must not be empty when %T.APIWhitelist.Private is enabled", config)
	}
	if config.APIWhitelist.Public.Enabled && config.APIWhitelist.Public.SubnetList == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.APIWhitelist.Public.SubnetList must not be empty when %T.APIWhitelist.Public is enabled", config)
	}
	if config.EncrypterBackend == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.EncrypterBackend must not be empty", config)
	}
	if config.InstallationName == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.InstallationName must not be empty", config)
	}

	r := &Resource{
		g8sClient: config.G8sClient,
		detection: config.Detection,
		logger:    config.Logger,

		apiWhitelist:     config.APIWhitelist,
		encrypterBackend: config.EncrypterBackend,
		installationName: config.InstallationName,
	}

	return r, nil
}

func (r *Resource) Name() string {
	return Name
}