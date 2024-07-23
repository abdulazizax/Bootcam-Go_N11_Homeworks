package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

type PriceUpdate struct {
	Company string  `json:"company"`
	Product string  `json:"product"`
	Price   float64 `json:"price"`
}

var companies = []string{"Apple", "Samsung", "Google", "Microsoft", "Amazon"}
var products = []string{"Phone", "Laptop", "Tablet", "Smartwatch", "Headphones"}

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for {
		for _, company := range companies {
			update := PriceUpdate{
				Company: company,
				Product: products[r.Intn(len(products))],
				Price:   float64(r.Intn(901) + 100),
			}

			jsonUpdate, err := json.Marshal(update)
			if err != nil {
				log.Printf("Error marshalling JSON: %v", err)
				continue
			}

			err = rdb.Publish(ctx, "price_updates", string(jsonUpdate)).Err()
			if err != nil {
				log.Printf("Error publishing message: %v", err)
			}
		}

		time.Sleep(2 * time.Second)
	}
}
