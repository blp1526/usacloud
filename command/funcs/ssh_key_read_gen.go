// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SSHKeyRead(ctx command.Context, params *params.ReadSSHKeyParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSSHKeyAPI()

	// set params

	// call Read(id)
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("SSHKeyRead is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
