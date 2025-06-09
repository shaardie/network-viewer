package subnetscanner

import (
	"context"
	"log/slog"
	"time"

	"github.com/shaardie/network-viewer/database"
	"gorm.io/gorm"
)

type SubnetScanner struct {
	db     *gorm.DB
	cancel context.CancelFunc
}

func New(db *gorm.DB) SubnetScanner {
	return SubnetScanner{
		db: db,
	}
}

func (s SubnetScanner) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				subnets := []database.Subnet{}
				if err := s.db.Where(&database.Subnet{Type: database.SubnetTypeIPv4, ScannerEnabled: true}).Find(&subnets).Error; err != nil {
					slog.Error("failed to query database for subnets to scan", "error", err)
					continue
				}
				for _, subnet := range subnets {
					if subnet.LastScan.Add(subnet.ScannerInterval).After(time.Now()) {
						continue
					}
					slog.Info("start network scan", "subnet", subnet.Subnet)
					subnet.LastScan = time.Now()
					if err := s.db.Save(&subnet).Error; err != nil {
						slog.Error("failed to save new subnet state in database", "error", err)
						continue
					}
					go func() {
						err := s.scanNetwork(ctx, &subnet)
						if err != nil {
							slog.Error("failed to scan subnet", "error", err, "subnet", subnet.Subnet)
						}
					}()

				}
			case <-ctx.Done():
				slog.Info("scanner stopped")
				return
			}
		}
	}()
}

func (s SubnetScanner) Stop() {
	if s.cancel != nil {
		s.cancel()
	}
}
