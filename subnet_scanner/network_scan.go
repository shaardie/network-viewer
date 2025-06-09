package SubnetScanner

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	probing "github.com/prometheus-community/pro-bing"
	"github.com/shaardie/network-viewer/database"
)

const workerCount = 16

func (s SubnetScanner) scanNetwork(ctx context.Context, subnet *database.Subnet) error {
	_, ipNet, err := net.ParseCIDR(subnet.Subnet)
	if err != nil {
		return fmt.Errorf("unable to parse subnet %v, %w", subnet.Subnet, err)
	}

	// start writer
	results, writerDone := s.writer(ctx)

	// start iterator
	jobs := IPIterator(ipNet)

	// start workers
	workerDone := make([]chan struct{}, workerCount)
	for i := range workerCount {
		workerDone[i] = make(chan struct{})
		go s.worker(ctx, i, subnet, jobs, results, workerDone[i])
	}

	// wait for workers
	for _, wd := range workerDone {
		<-wd
	}

	slog.Info("network scan done", "network", ipNet.String())
	// wait for writer
	<-writerDone

	slog.Info("network scan done", "network", ipNet.String())
	return nil
}

func (s SubnetScanner) worker(ctx context.Context, id int, subnet *database.Subnet, jobs <-chan net.IPNet, results chan<- *database.IP, done chan<- struct{}) {
	for ipWithMask := range jobs {
		// cancel
		if ctx.Err() != nil {
			break
		}

		slog.Debug("ping", "worker id", id, "ip", ipWithMask)

		ip := database.IP{}
		found := true
		if err := s.db.Where(&database.IP{IP: ipWithMask.String()}).Find(&ip); err != nil {
			found = false
		}

		ip.SubnetID = subnet.ID
		ip.IP = ipWithMask.String()

		pinger, err := probing.NewPinger(ipWithMask.IP.String())
		if err != nil {
			slog.Error("unable to create pinger", "ip", ipWithMask, "error", err)
			continue
		}
		pinger.Count = 1
		pinger.Timeout = time.Millisecond * 500
		err = pinger.Run()
		if err != nil {
			slog.Error("failed to ping", "ip", ipWithMask, "error", err)
			continue
		}
		r := pinger.Statistics()

		// nobody there and we do not know this IP
		if !found && r.PacketsRecv == 0 {
			continue
		}

		ip.Online = false
		if r.PacketsRecv != 0 {
			ip.Online = true
			ip.RTT = r.AvgRtt
			ip.Hops = pinger.TTL - int(pinger.Statistics().TTLs[0])

			hostnames, _ := net.LookupAddr(ipWithMask.IP.String())
			if len(hostnames) > 0 {
				ip.Hostname = hostnames[0] // currently use only the first one
			}
		}

		results <- &ip
	}
	done <- struct{}{}
}

func (s SubnetScanner) writer(ctx context.Context) (chan<- *database.IP, <-chan struct{}) {
	results := make(chan *database.IP)
	done := make(chan struct{})
	go func() {
		for ip := range results {
			// Cancel
			if ctx.Err() != nil {
				break
			}
			if err := s.db.Save(ip).Error; err != nil {
				slog.Error("failed to save", "error", err)
			}
		}
		done <- struct{}{}
	}()
	return results, done
}
