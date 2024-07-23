package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type PriceUpdate struct {
	Company string  `json:"company"`
	Product string  `json:"product"`
	Price   float64 `json:"price"`
}

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pubsub := rdb.Subscribe(ctx, "price_updates")
	defer pubsub.Close()

	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ch := pubsub.Channel()
	for msg := range ch {
		var update PriceUpdate
		err := json.Unmarshal([]byte(msg.Payload), &update)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			continue
		}
		fmt.Printf("Company: %s, Product: %s, Price: $%.2f\n", update.Company, update.Product, update.Price)
	}
}
