package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Queue struct {
	mu     sync.Mutex
	nextID int
	List   []int
}

func (q *Queue) AddClient() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.nextID++
	q.List = append(q.List, q.nextID)
	return q.nextID
}

func (q *Queue) GetAll() []int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return append([]int{}, q.List...)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	queue := &Queue{List: make([]int, 0)}

	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		id := queue.AddClient()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"your_number": strconv.Itoa(id)})
	})

	http.HandleFunc("/queue", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(queue.GetAll())
	})

	log.Printf("✅ Server started on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
