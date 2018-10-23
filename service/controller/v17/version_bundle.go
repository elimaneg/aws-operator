package v17

import (
	"github.com/giantswarm/versionbundle"
)

func VersionBundle() versionbundle.Bundle {
	return versionbundle.Bundle{
		Changelogs: []versionbundle.Changelog{
			{
				Component:   "aws-operator",
				Description: "Add security group descriptions for master nodes.",
				Kind:        versionbundle.KindChanged,
			},
			{
				Component:   "cloudconfig",
				Description: "Removed CoreDNS related components (now managed by chart-operator).",
				Kind:        versionbundle.KindRemoved,
			},
			{
				Component:   "aws-operator",
				Description: "Set higher timeouts for NVME driver.",
				Kind:        versionbundle.KindAdded,
			},
			{
				Component:   "calico",
				Description: "Updated calico to 3.2.0.",
				Kind:        versionbundle.KindChanged,
			},
			{
				Component:   "kubernetes",
				Description: "Patched Kubernetes v1.11.1 with apimachinery fix to plug apiserver memory leak.",
				Kind:        versionbundle.KindChanged,
			},
			{
				Component:   "aws-operator",
				Description: "Set higher timeouts for NVME driver.",
				Kind:        versionbundle.KindAdded,
			},
		},
		Components: []versionbundle.Component{
			{
				Name:    "calico",
				Version: "3.2.0",
			},
			{
				Name:    "containerlinux",
				Version: "1745.4.0",
			},
			{
				Name:    "docker",
				Version: "18.03.1",
			},
			{
				Name:    "etcd",
				Version: "3.3.8",
			},
			{
				Name:    "kubernetes",
				Version: "1.11.1",
			},
		},
		Name:    "aws-operator",
		Version: "4.2.0",
	}
}