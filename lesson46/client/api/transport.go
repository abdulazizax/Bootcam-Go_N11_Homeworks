package api

import (
	"context"
	"fmt"
	"lesson46/server/genproto/TransportService/transport"
	"log"
)

type transportService struct {
}

func NewTransportService() *transportService {
	return &transportService{}
}

func (t *transportService) GetBusSchedule(client transport.TransportServiceClient) {
	req := &transport.ScheduleRequest{RouteNumber: "1"}
	res, err := client.GetBusSchedule(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	log.Println("GetBusSchedule: ", res)
	fmt.Println()
}

func (t *transportService) TrackBusLocation(client transport.TransportServiceClient) {
	req := &transport.LocationRequest{
		BusId: "B001",
	}
	res, err := client.TrackBusLocation(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	log.Println("TrackBusLocation: ", res)
	fmt.Println()
}

func (t *transportService) ReportTrafficJam(client transport.TransportServiceClient) {
	req := &transport.TrafficJamReport{
		Location:    "Shahrisabz ko'chasi",
		Severity:    "Yuqorqi",
		Description: "Yengil avtomobillar sababli tirbandlik",
	}
	res, err := client.ReportTrafficJam(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	log.Println("ReportTrafficJam: ", res)
	fmt.Println()
}
