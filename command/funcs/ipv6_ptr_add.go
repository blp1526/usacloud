package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func Ipv6PtrAdd(ctx command.Context, params *params.PtrAddIpv6Param) error {
	client := ctx.GetAPIClient()
	api := client.GetIPv6AddrAPI()

	targetIP, err := getIPv6AddrFromArgs(ctx.Args())
	if err != nil {
		return err
	}

	ip, err := api.Read(targetIP)
	if err == nil && ip.HostName != "" {
		return fmt.Errorf("PTR record has already been set for IPAddress %q", targetIP)
	}

	ip, err = api.Create(targetIP, params.Hostname)
	if err != nil {
		return fmt.Errorf("IPv6PtrAdd is failed: %s", err)
	}

	ipv6net, err := client.GetIPv6NetAPI().Read(ip.IPv6Net.ID)
	if err != nil {
		return fmt.Errorf("IPv6PtrAdd is failed: %s", err)
	}
	ip.IPv6Net = ipv6net

	return ctx.GetOutput().Print(ip)
}
