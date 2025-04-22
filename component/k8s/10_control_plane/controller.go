package control_plane

import (
	"k-kube/component/common"
	"path/filepath"

	containersBases "github.com/tuxounet/k2-sdk/kernel/compute/providers/containers/bases"
	containersTypes "github.com/tuxounet/k2-sdk/kernel/compute/providers/containers/types"
	"github.com/tuxounet/k2-sdk/types"
)

type Controller struct {
	containersBases.BaseControllerContainer
}

const ControllerKey = "control_plane"

func NewController(component types.IAppComponent) types.IAppController {

	base := containersBases.NewBaseControllerContainer(component, &containersTypes.ContainerDefinition{
		Order: 1010,
		Name:  ControllerKey,
		Image: "docker.io/rancher/k3s:" + common.RancherVersion,
		Ports: []*containersTypes.ContainerDefinitionPort{
			{
				ContainerPort: "46443",
				Protocol:      "TCP",
				HostPort:      "46443",
				HostAddress:   "0.0.0.0",
			},
		},
		Env: map[string]string{
			"TZ":                    "Europe/Paris",
			"K3S_TOKEN":             common.KubeToken,
			"K3S_KUBECONFIG_OUTPUT": "/config/config",
			"K3S_KUBECONFIG_MODE":   "666",
		},
		Command: &[]string{
			"server",
			"--https-listen-port=46443",
			"--node-name=single-node",
			"--disable-cloud-controller",
			"--prefer-bundled-bin",
		},
		Volumes: []*containersTypes.ContainerDefinitionVolume{
			{
				Name:          "data",
				ContainerPath: "/var/lib/rancher/k3s",
				Binding: containersTypes.ContainerDefinitionVolumeBinding{
					Type:     containersTypes.ContainerDefinitionVolumeBindingTypeMount,
					HostPath: filepath.Join(component.GetApp().GetKernel().GetRunDirectory(), "var", "k3s"),
				},
			},
			{
				Name:          "config",
				ContainerPath: "/config",
				Binding: containersTypes.ContainerDefinitionVolumeBinding{
					Type:     containersTypes.ContainerDefinitionVolumeBindingTypeMount,
					HostPath: filepath.Join(component.GetApp().GetKernel().GetRunDirectory(), "home", ".kube"),
				},
			},
		},
		Security: &containersTypes.ContainerDefinitionSecurity{
			Privileged: true,
			RunAsUser:  "root",
		},
	}, nil)
	return &Controller{
		base,
	}
}
