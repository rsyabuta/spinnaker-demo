package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

const version = "2.0.0"

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "demo-redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func main() {
	c := NewClient()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ns, err := c.Get("namespace").Result()
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, "Version: %s\n Namespace: %s", version, ns)
	})

	http.ListenAndServe(":8080", nil)
}
