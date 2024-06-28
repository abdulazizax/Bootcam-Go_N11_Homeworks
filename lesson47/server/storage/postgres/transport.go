package postgres

import (
	"context"
	pb "server/genproto/TransportService/transport"
	"time"

	"github.com/jmoiron/sqlx"
)

type TransportRepo struct {
	db *sqlx.DB
}

func NewTransportRepo(db *sqlx.DB) *TransportRepo {
	return &TransportRepo{db: db}
}

func (t *TransportRepo) GetBusSchedule(ctx context.Context, in *pb.ScheduleRequest) (*pb.ScheduleResponse, error) {
	query := `
		SELECT 
			stop_name,
			arrival_time
		FROM 
			schedule 
		WHERE 
			route_number = $1
	`

	rows, err := t.db.QueryxContext(ctx, query, in.RouteNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []*pb.Schedule
	for rows.Next() {
		var arrivalTime time.Time
		var schedule pb.Schedule
		err = rows.Scan(
			&schedule.StopName,
			&arrivalTime,
		)
		if err != nil {
			return nil, err
		}
		schedule.ArrivalTime = arrivalTime.Format("2006-01-02 15:04:05")
		schedules = append(schedules, &schedule)
	}

	return &pb.ScheduleResponse{
		Schedule: schedules,
	}, nil
}

func (t *TransportRepo) TrackBusLocation(ctx context.Context, in *pb.LocationRequest) (*pb.LocationResponse, error) {
	query := `
		SELECT
			latitude,
			longitude,
			timestamp
		FROM
			bus_location
		WHERE
			bus_id = $1
		ORDER BY
			timestamp DESC
		LIMIT 1
	`

	var location pb.Location
	var timeStamp time.Time

	rows := t.db.QueryRowxContext(ctx, query, in.BusId)
	err := rows.Scan(
		&location.Latitude,
		&location.Longitude,
		&timeStamp,
	)
	if err != nil {
		return nil, err
	}
	location.Timestamp = timeStamp.Format("2006-01-02 15:04:05")

	return &pb.LocationResponse{
		Location: &location,
	}, err
}

func (t *TransportRepo) ReportTrafficJam(ctx context.Context, in *pb.TrafficJamReport) (*pb.ReportResponse, error) {
	query := `
		INSERT INTO traffic_jam_report (
			location,
			severity,
			description) 
		VALUES($1, $2, $3)
	`

	_, err := t.db.ExecContext(ctx, query, in.Location, in.Severity, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.ReportResponse{Status: "Report received successfully"}, nil
}
