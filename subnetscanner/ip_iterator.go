package subnetscanner

import (
	"math/big"
	"net"
)

func IPIterator(ipNet *net.IPNet) <-chan net.IPNet {
	ch := make(chan net.IPNet, 255)

	go func() {
		defer close(ch)
		ip := ipNet.IP

		for ipNet.Contains(ip) {
			ch <- net.IPNet{
				IP:   ip,
				Mask: ipNet.Mask,
			}
			ip = incIP(ip)
		}
	}()

	return ch
}

func incIP(ip net.IP) net.IP {
	ip = ip.To4()
	ipInt := big.NewInt(0).SetBytes(ip)
	ipInt.Add(ipInt, big.NewInt(1))
	b := ipInt.Bytes()

	if len(b) < 4 {
		padding := make([]byte, 4-len(b))
		b = append(padding, b...)
	}
	return net.IP(b).To4()
}
