package folders

import (
	_ "embed"

	playbooksBases "github.com/tuxounet/k2-sdk/kernel/compute/providers/playbooks/bases"
	playbooksTypes "github.com/tuxounet/k2-sdk/kernel/compute/providers/playbooks/types"
	"github.com/tuxounet/k2-sdk/types"
)

type Controller struct {
	playbooksBases.BaseControllerPlaybook
}

//go:embed playbooks/provision.yaml
var provisionPlaybook string

const ControllerKey = "folders"

func NewController(component types.IAppComponent) types.IAppController {

	base := playbooksBases.NewBaseControllerPlaybook(component, &playbooksTypes.PlaybookDefinition{
		Order: 1000,
		Name:  ControllerKey,
		Start: provisionPlaybook,
	}, nil)
	return &Controller{
		base,
	}
}
