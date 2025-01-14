package netlink

import "github.com/vishvananda/netlink"

type Addr = netlink.Addr

func (n *NetLink) AddrList(link Link, family int) (
	addresses []Addr, err error) {
	return netlink.AddrList(link, family)
}

func (n *NetLink) AddrReplace(link Link, addr *Addr) error {
	return netlink.AddrReplace(link, addr)
}
