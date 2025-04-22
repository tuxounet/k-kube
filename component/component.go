package kube

import (
	"embed"

	folders "k-kube/component/k8s/00_folders"
	control_plane "k-kube/component/k8s/10_control_plane"

	runtimeBases "github.com/tuxounet/k2-sdk/bases"
	runtimeTypes "github.com/tuxounet/k2-sdk/types"
)

//go:embed config/*.yaml
var conf embed.FS

// @title			K-Kube
// @description	Un host Kubernetes
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	GPL-3.0
// @license.url	http://www.gnu.org/licenses/gpl-3.0.html
func NewComponent(app runtimeTypes.IApp) runtimeTypes.IAppComponent {
	return runtimeBases.NewBaseAppComponent(
		app,
		"kube",
		1000,
		nil,
		nil,
		&conf,
		runtimeTypes.AccessPolicyAdmin,
		[]runtimeTypes.AppControllerCtor{
			folders.NewController,
			control_plane.NewController,
		},
	)
}
