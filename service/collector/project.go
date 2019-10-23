package collector

import (
	"github.com/giantswarm/aws-operator/pkg/project"

	"github.com/prometheus/client_golang/prometheus"
)

var projectInfo = prometheus.NewDesc(
	prometheus.BuildFQName("giantswarm", "project", "info"),
	"A metric with a constant '1' value showing project informations.",
	[]string{
		"bundle_version",
		"gitSHA",
		"name",
		"version",
	},
	nil,
)

type Project struct {
}

func NewProject() *Project {
	return &Project{}
}

func (p *Project) Collect(ch chan<- prometheus.Metric) error {
	ch <- prometheus.MustNewConstMetric(
		projectInfo,
		prometheus.GaugeValue,
		1,
		project.BundleVersion(),
		project.GitSHA(),
		project.Name(),
		project.Version(),
	)

	return nil
}

func (p *Project) Describe(ch chan<- *prometheus.Desc) error {
	ch <- projectInfo
	return nil
}
