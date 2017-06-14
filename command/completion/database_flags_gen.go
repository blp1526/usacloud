// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func DatabaseListCompleteFlags(ctx command.Context, params *params.ListDatabaseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		comp = define.Resources["Database"].Commands["list"].Params["name"].CompleteFunc
	case "id":
		comp = define.Resources["Database"].Commands["list"].Params["id"].CompleteFunc
	case "tags":
		comp = define.Resources["Database"].Commands["list"].Params["tags"].CompleteFunc
	case "from", "offset":
		comp = define.Resources["Database"].Commands["list"].Params["from"].CompleteFunc
	case "max", "limit":
		comp = define.Resources["Database"].Commands["list"].Params["max"].CompleteFunc
	case "sort":
		comp = define.Resources["Database"].Commands["list"].Params["sort"].CompleteFunc
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

func DatabaseCreateCompleteFlags(ctx command.Context, params *params.CreateDatabaseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "switch-id":
		comp = define.Resources["Database"].Commands["create"].Params["switch-id"].CompleteFunc
	case "plan":
		comp = define.Resources["Database"].Commands["create"].Params["plan"].CompleteFunc
	case "database", "db":
		comp = define.Resources["Database"].Commands["create"].Params["database"].CompleteFunc
	case "username":
		comp = define.Resources["Database"].Commands["create"].Params["username"].CompleteFunc
	case "password":
		comp = define.Resources["Database"].Commands["create"].Params["password"].CompleteFunc
	case "source-networks":
		comp = define.Resources["Database"].Commands["create"].Params["source-networks"].CompleteFunc
	case "enable-web-ui":
		comp = define.Resources["Database"].Commands["create"].Params["enable-web-ui"].CompleteFunc
	case "backup-time":
		comp = define.Resources["Database"].Commands["create"].Params["backup-time"].CompleteFunc
	case "port":
		comp = define.Resources["Database"].Commands["create"].Params["port"].CompleteFunc
	case "ipaddress1", "ip1", "ipaddress", "ip":
		comp = define.Resources["Database"].Commands["create"].Params["ipaddress1"].CompleteFunc
	case "nw-mask-len":
		comp = define.Resources["Database"].Commands["create"].Params["nw-mask-len"].CompleteFunc
	case "default-route":
		comp = define.Resources["Database"].Commands["create"].Params["default-route"].CompleteFunc
	case "name":
		comp = define.Resources["Database"].Commands["create"].Params["name"].CompleteFunc
	case "description", "desc":
		comp = define.Resources["Database"].Commands["create"].Params["description"].CompleteFunc
	case "tags":
		comp = define.Resources["Database"].Commands["create"].Params["tags"].CompleteFunc
	case "icon-id":
		comp = define.Resources["Database"].Commands["create"].Params["icon-id"].CompleteFunc
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

func DatabaseReadCompleteFlags(ctx command.Context, params *params.ReadDatabaseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["Database"].Commands["read"].Params["id"].CompleteFunc
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

func DatabaseUpdateCompleteFlags(ctx command.Context, params *params.UpdateDatabaseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "password":
		comp = define.Resources["Database"].Commands["update"].Params["password"].CompleteFunc
	case "port":
		comp = define.Resources["Database"].Commands["update"].Params["port"].CompleteFunc
	case "source-networks":
		comp = define.Resources["Database"].Commands["update"].Params["source-networks"].CompleteFunc
	case "enable-web-ui":
		comp = define.Resources["Database"].Commands["update"].Params["enable-web-ui"].CompleteFunc
	case "backup-time":
		comp = define.Resources["Database"].Commands["update"].Params["backup-time"].CompleteFunc
	case "name":
		comp = define.Resources["Database"].Commands["update"].Params["name"].CompleteFunc
	case "description", "desc":
		comp = define.Resources["Database"].Commands["update"].Params["description"].CompleteFunc
	case "tags":
		comp = define.Resources["Database"].Commands["update"].Params["tags"].CompleteFunc
	case "icon-id":
		comp = define.Resources["Database"].Commands["update"].Params["icon-id"].CompleteFunc
	case "id":
		comp = define.Resources["Database"].Commands["update"].Params["id"].CompleteFunc
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

func DatabaseDeleteCompleteFlags(ctx command.Context, params *params.DeleteDatabaseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "force", "f":
		comp = define.Resources["Database"].Commands["delete"].Params["force"].CompleteFunc
	case "id":
		comp = define.Resources["Database"].Commands["delete"].Params["id"].CompleteFunc
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

func DatabaseBootCompleteFlags(ctx command.Context, params *params.BootDatabaseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["Database"].Commands["boot"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DatabaseShutdownCompleteFlags(ctx command.Context, params *params.ShutdownDatabaseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["Database"].Commands["shutdown"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DatabaseShutdownForceCompleteFlags(ctx command.Context, params *params.ShutdownForceDatabaseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["Database"].Commands["shutdown-force"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DatabaseResetCompleteFlags(ctx command.Context, params *params.ResetDatabaseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["Database"].Commands["reset"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DatabaseWaitForBootCompleteFlags(ctx command.Context, params *params.WaitForBootDatabaseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["Database"].Commands["wait-for-boot"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func DatabaseWaitForDownCompleteFlags(ctx command.Context, params *params.WaitForDownDatabaseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["Database"].Commands["wait-for-down"].Params["id"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
