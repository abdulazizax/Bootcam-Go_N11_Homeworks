package main

import (
	"context"
	"log"
	"time"

	pb "translator-gprc/translator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTranslatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Translate(ctx, &pb.TranslateRequest{Text: "salom"})
	if err != nil {
		log.Fatalf("Could not translate: %v", err)
	}
	log.Printf("Translated text: %s", r.GetTranslatedText())
}
