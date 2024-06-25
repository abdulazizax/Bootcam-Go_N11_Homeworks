package main

import (
	"context"
	"log"
	"net"

	pb "translator-gprc/translator"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTranslatorServer
}

var translations = map[string]string{
	"salom":    "hello",
	"dunyo":    "world",
	"xayr":     "goodbye",
	"olma":     "apple",
	"banan":    "banana",
	"toshbaqa": "turtle",
}

func (s *server) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	translatedText, ok := translations[req.GetText()]
	if !ok {
		translatedText = "Tarjima topilmadi"
	}

	return &pb.TranslateResponse{TranslatedText: translatedText}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTranslatorServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
