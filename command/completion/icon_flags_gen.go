// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func IconListCompleteFlags(ctx command.Context, params *params.ListIconParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		comp = define.Resources["Icon"].Commands["list"].Params["name"].CompleteFunc
	case "id":
		comp = define.Resources["Icon"].Commands["list"].Params["id"].CompleteFunc
	case "scope":
		comp = define.Resources["Icon"].Commands["list"].Params["scope"].CompleteFunc
	case "tags":
		comp = define.Resources["Icon"].Commands["list"].Params["tags"].CompleteFunc
	case "from", "offset":
		comp = define.Resources["Icon"].Commands["list"].Params["from"].CompleteFunc
	case "max", "limit":
		comp = define.Resources["Icon"].Commands["list"].Params["max"].CompleteFunc
	case "sort":
		comp = define.Resources["Icon"].Commands["list"].Params["sort"].CompleteFunc
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

func IconCreateCompleteFlags(ctx command.Context, params *params.CreateIconParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "image":
		comp = define.Resources["Icon"].Commands["create"].Params["image"].CompleteFunc
	case "name":
		comp = define.Resources["Icon"].Commands["create"].Params["name"].CompleteFunc
	case "tags":
		comp = define.Resources["Icon"].Commands["create"].Params["tags"].CompleteFunc
	case "assumeyes", "y":
		comp = define.Resources["Icon"].Commands["create"].Params["assumeyes"].CompleteFunc
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

func IconReadCompleteFlags(ctx command.Context, params *params.ReadIconParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		comp = define.Resources["Icon"].Commands["read"].Params["id"].CompleteFunc
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

func IconUpdateCompleteFlags(ctx command.Context, params *params.UpdateIconParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		comp = define.Resources["Icon"].Commands["update"].Params["name"].CompleteFunc
	case "tags":
		comp = define.Resources["Icon"].Commands["update"].Params["tags"].CompleteFunc
	case "assumeyes", "y":
		comp = define.Resources["Icon"].Commands["update"].Params["assumeyes"].CompleteFunc
	case "id":
		comp = define.Resources["Icon"].Commands["update"].Params["id"].CompleteFunc
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

func IconDeleteCompleteFlags(ctx command.Context, params *params.DeleteIconParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "assumeyes", "y":
		comp = define.Resources["Icon"].Commands["delete"].Params["assumeyes"].CompleteFunc
	case "id":
		comp = define.Resources["Icon"].Commands["delete"].Params["id"].CompleteFunc
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
