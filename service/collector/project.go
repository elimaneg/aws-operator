package collector

import (
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
}

func (p *Project) Collect(ch chan<- prometheus.Metric) error {
	ch <- prometheus.MustNewConstMetric(
		projectInfo,
		prometheus.GaugeValue,
		project.Bundle_version(),
		project.GitSHA(),
		project.Name(),
		project.Version(),
	)

	return nil
}

func (p *Project) Describe(ch chan<- *prometheus.Desc) error {
	ch <- projectDesc
	return nil
}
