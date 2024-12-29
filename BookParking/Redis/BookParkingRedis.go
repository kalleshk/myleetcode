package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Context for Redis operations
	ctx := context.Background()

	// Replace these with your Redis Cloud details
	const (
		redisAddr     = "redis-17124.c10.us-east-1-4.ec2.redns.redis-cloud.com:17124"
		redisPassword = "AqNcPrte50ZhQHirXb8qQJjZLN0ol36n"
	)

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // Leave empty for no password
		DB:       0,
	})

	// Check connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis Cloud!")

	// Perform CRUD operations
	// 1. Create (Set key-value)
	err = rdb.Set(ctx, "name", "GoLang Developer", 0).Err()
	if err != nil {
		log.Fatalf("Could not SET value: %v", err)
	}
	fmt.Println("Set key 'name'")

	// 2. Read (Get value)
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		log.Fatalf("Could not GET value: %v", err)
	}
	fmt.Printf("Got value for 'name': %s\n", val)

	// 3. Update (Modify key-value)
	err = rdb.Set(ctx, "name", "Updated Developer", 0).Err()
	if err != nil {
		log.Fatalf("Could not UPDATE value: %v", err)
	}
	fmt.Println("Updated key 'name'")

	// 4. Delete (Remove key)
	//err = rdb.Del(ctx, "name").Err()
	//if err != nil {
	//	log.Fatalf("Could not DELETE key: %v", err)
	//}
	//fmt.Println("Deleted key 'name'")
}
