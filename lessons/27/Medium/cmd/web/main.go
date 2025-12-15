package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"queue-app/internal/config"
)

// Інформація про збірку
var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

var queue = 0

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	//Вивід збірки
	log.Printf(
		"%s started | env=%s | version=%s | build=%s | commit=%s",
		cfg.AppName, cfg.Env, Version, BuildTime, GitCommit,
	)

	//Обробник кореневого марщруту
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Електронна черга\nПоточний номер: %d\n", queue)
	})

	//Обробник шляху маршрутизації /next для додавання клієнтів
	http.HandleFunc("/next", func(w http.ResponseWriter, _ *http.Request) {
		queue++
		fmt.Fprintf(w, "Новий номер у черзі: %d\n", queue)
	})

	addr := ":" + cfg.Port
	log.Printf("Listening on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
