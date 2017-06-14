package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func LoadBalancerListCompleteArgs(ctx command.Context, params *params.ListLoadBalancerParam, cur, prev, commandName string) {

}

func LoadBalancerCreateCompleteArgs(ctx command.Context, params *params.CreateLoadBalancerParam, cur, prev, commandName string) {

}

func LoadBalancerReadCompleteArgs(ctx command.Context, params *params.ReadLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerUpdateCompleteArgs(ctx command.Context, params *params.UpdateLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerDeleteCompleteArgs(ctx command.Context, params *params.DeleteLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerBootCompleteArgs(ctx command.Context, params *params.BootLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerShutdownCompleteArgs(ctx command.Context, params *params.ShutdownLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerShutdownForceCompleteArgs(ctx command.Context, params *params.ShutdownForceLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerResetCompleteArgs(ctx command.Context, params *params.ResetLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerWaitForBootCompleteArgs(ctx command.Context, params *params.WaitForBootLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerWaitForDownCompleteArgs(ctx command.Context, params *params.WaitForDownLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerVipInfoCompleteArgs(ctx command.Context, params *params.VipInfoLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerVipAddCompleteArgs(ctx command.Context, params *params.VipAddLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerVipUpdateCompleteArgs(ctx command.Context, params *params.VipUpdateLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerVipDeleteCompleteArgs(ctx command.Context, params *params.VipDeleteLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerServerInfoCompleteArgs(ctx command.Context, params *params.ServerInfoLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerServerAddCompleteArgs(ctx command.Context, params *params.ServerAddLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerServerUpdateCompleteArgs(ctx command.Context, params *params.ServerUpdateLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerServerDeleteCompleteArgs(ctx command.Context, params *params.ServerDeleteLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LoadBalancerMonitorCompleteArgs(ctx command.Context, params *params.MonitorLoadBalancerParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLoadBalancerAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.LoadBalancers {
		fmt.Println(res.LoadBalancers[i].ID)
		var target interface{} = &res.LoadBalancers[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}
