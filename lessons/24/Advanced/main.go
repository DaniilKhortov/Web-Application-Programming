package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func simulateWork() {
	start := time.Now()
	sum := 0.0
	for i := 0; i < 10_000_000; i++ {
		sum += math.Sqrt(float64(i)) * math.Sin(float64(i))
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Sprintf("Result: %.2f", sum)
	fmt.Println("Work time:", time.Since(start))
}

func handler(w http.ResponseWriter, r *http.Request) {
	simulateWork()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Done! Simulated heavy work completed.")
}

func main() {

	http.HandleFunc("/work", handler)

	fmt.Println("Server launched on http://localhost:8080")
	fmt.Println("Pprof is available on http://localhost:8080/debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
