package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SSHKeyListCompleteArgs(ctx command.Context, params *params.ListSSHKeyParam, cur, prev, commandName string) {

}

func SSHKeyCreateCompleteArgs(ctx command.Context, params *params.CreateSSHKeyParam, cur, prev, commandName string) {

}

func SSHKeyReadCompleteArgs(ctx command.Context, params *params.ReadSSHKeyParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSSHKeyAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.SSHKeys {
		fmt.Println(res.SSHKeys[i].ID)
		var target interface{} = &res.SSHKeys[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SSHKeyUpdateCompleteArgs(ctx command.Context, params *params.UpdateSSHKeyParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSSHKeyAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.SSHKeys {
		fmt.Println(res.SSHKeys[i].ID)
		var target interface{} = &res.SSHKeys[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SSHKeyDeleteCompleteArgs(ctx command.Context, params *params.DeleteSSHKeyParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSSHKeyAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.SSHKeys {
		fmt.Println(res.SSHKeys[i].ID)
		var target interface{} = &res.SSHKeys[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SSHKeyGenerateCompleteArgs(ctx command.Context, params *params.GenerateSSHKeyParam, cur, prev, commandName string) {

}
