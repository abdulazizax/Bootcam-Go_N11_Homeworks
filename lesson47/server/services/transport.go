package serveces

import (
	"context"
	pb "server/genproto/TransportService/transport"
	"server/storage/postgres"

	"github.com/jmoiron/sqlx"
)

type TransportRepo struct {
	db *postgres.TransportRepo
	pb.UnimplementedTransportServiceServer
}

func NewTransportRepo(db *sqlx.DB) *TransportRepo {
	return &TransportRepo{
		db: postgres.NewTransportRepo(db),
	}
}

func (t *TransportRepo) GetBusSchedule(ctx context.Context, in *pb.ScheduleRequest) (*pb.ScheduleResponse, error) {
	res, err := t.db.GetBusSchedule(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (t *TransportRepo) TrackBusLocation(ctx context.Context, in *pb.LocationRequest) (*pb.LocationResponse, error) {
	res, err := t.db.TrackBusLocation(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (t *TransportRepo) ReportTrafficJam(ctx context.Context, in *pb.TrafficJamReport) (*pb.ReportResponse, error) {
	res, err := t.db.ReportTrafficJam(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, err
}
