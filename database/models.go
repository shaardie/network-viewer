package database

import (
	"time"

	"gorm.io/gorm"
)

type subnetType string

const (
	SubnetTypeIPv4 subnetType = "ipv4"
	SubnetTypeIPv6 subnetType = "ipv6"
)

type Metadata struct {
	Comment string
}

type Subnet struct {
	gorm.Model
	Metadata

	Type   subnetType `gorm:"not null"`
	Subnet string     `gorm:"unique,not null"`

	ScannerInterval time.Duration
	ScannerEnabled  bool
	LastScan        time.Time

	IPs []IP
}

type IP struct {
	gorm.Model
	Metadata

	IP       string `gorm:"unique,not null"`
	RTT      time.Duration
	Hops     int
	MAC      string
	Online   bool
	Hostname string

	SubnetID uint
}
