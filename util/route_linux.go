package util

import (
	"fmt"
	"net"
)

// AddNetRoute add subnet route
func AddNetRoute(tun string, subnet *net.IPNet) error {
	sargs := fmt.Sprintf("add -net %s dev %s", subnet, tun)
	return ExecCommand("route", sargs)
}

// AddHostRoute add host route
func AddHostRoute(tun string, host string) error {
	sargs := fmt.Sprintf("add -host %s dev %s", host, tun)
	return ExecCommand("route", sargs)
}
