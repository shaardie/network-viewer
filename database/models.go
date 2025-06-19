package database

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"net"
	"time"

	"gorm.io/gorm"
)

type subnetType string

const (
	SubnetTypeIPv4 subnetType = "ipv4"
	SubnetTypeIPv6 subnetType = "ipv6"
)

type IPNet struct {
	*net.IPNet
}

func (ipNet *IPNet) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("wrong type")
	}
	_, in, err := net.ParseCIDR(s)
	if err != nil {
		return fmt.Errorf("unable to parse cidr, %w", err)
	}
	ipNet.IPNet = in
	return nil
}

func (ipNet IPNet) Value() (driver.Value, error) {
	return ipNet.String(), nil
}

type Subnet struct {
	gorm.Model

	Subnet IPNet

	ScannerInterval time.Duration
	ScannerEnabled  bool
	LastScan        time.Time

	IPs []IP

	Comment string
}

type IP struct {
	gorm.Model

	IP       IPNet
	RTT      time.Duration
	MAC      string
	Online   bool
	Hostname string

	SubnetID uint

	Comment string
}
