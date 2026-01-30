package main

import (
    "context"
    "log"

    "github.com/aidantrabs/trinbago-hackathon/backend/internal/config"
    "github.com/aidantrabs/trinbago-hackathon/backend/internal/db"
)

func main() {
    ctx := context.Background()

    cfg, err := config.Load()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    pool, err := db.Connect(ctx, cfg.DatabaseURL)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer pool.Close()

    log.Printf("Connected to database")
    log.Printf("Server configured on port %s", cfg.Port)
}
