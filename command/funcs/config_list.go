package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/command/profile"
)

func ConfigList(ctx command.Context, params *params.ListConfigParam) error {
	profiles, err := profile.List()
	if err != nil {
		return err
	}

	var cp string
	if !params.Brief {
		cp, err = profile.GetCurrentName()
		if err != nil {
			return err
		}
	}

	for _, p := range profiles {
		if params.Brief {
			fmt.Fprintf(command.GlobalOption.Out, "%s\n", p)
			continue
		}

		mark := " "
		if p == cp {
			mark = "*"
		}
		fmt.Fprintf(command.GlobalOption.Out, "%s %s\n", mark, p)
	}
	return nil
}
