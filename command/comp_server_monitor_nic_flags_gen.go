// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ServerMonitorNicCompleteFlags(ctx Context, params *MonitorNicServerParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "end":
		comp = define.Resources["Server"].Commands["monitor-nic"].Params["end"].CompleteFunc
	case "id":
		comp = define.Resources["Server"].Commands["monitor-nic"].Params["id"].CompleteFunc
	case "index":
		comp = define.Resources["Server"].Commands["monitor-nic"].Params["index"].CompleteFunc
	case "key-format":
		comp = define.Resources["Server"].Commands["monitor-nic"].Params["key-format"].CompleteFunc
	case "start":
		comp = define.Resources["Server"].Commands["monitor-nic"].Params["start"].CompleteFunc
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}