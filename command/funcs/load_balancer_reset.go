package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func LoadBalancerReset(ctx command.Context, params *params.ResetLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerReset is failed: %s", e)
	}

	err := internal.ExecWithProgress(
		fmt.Sprintf("Still resetting[ID:%d]...", params.Id),
		fmt.Sprintf("Reset load-balancer[ID:%d]", params.Id),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			_, err := api.RebootForce(params.Id)
			if err != nil {
				errChan <- err
				return
			}
			err = api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("LoadBalancerReset is failed: %s", err)
	}

	return nil
}
