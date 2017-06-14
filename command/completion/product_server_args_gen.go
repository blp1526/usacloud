package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProductServerListCompleteArgs(ctx command.Context, params *params.ListProductServerParam, cur, prev, commandName string) {

}

func ProductServerReadCompleteArgs(ctx command.Context, params *params.ReadProductServerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	if cur != "" && !isSakuraID(cur) {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProductServerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.ServerPlans {
		fmt.Println(res.ServerPlans[i].ID)

	}

}
