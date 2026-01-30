package main

import (
    "context"
    "log"
    "net/http"
    "strings"

    "github.com/aidantrabs/trinbago-hackathon/backend/internal/config"
    "github.com/aidantrabs/trinbago-hackathon/backend/internal/db"
    "github.com/aidantrabs/trinbago-hackathon/backend/internal/handler"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    ctx := context.Background()

    cfg, err := config.Load()
    if err != nil {
        log.Fatal("failed to load config:", err)
    }

    pool, err := db.Connect(ctx, cfg.DatabaseURL)
    if err != nil {
        log.Fatal("failed to connect to database:", err)
    }
    defer pool.Close()

    h := handler.New(pool)

    e := echo.New()
    e.HideBanner = true

    // middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins:     strings.Split(cfg.AllowedOrigins, ","),
        AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
        AllowHeaders:     []string{echo.HeaderContentType},
        AllowCredentials: true,
    }))

    // routes
    api := e.Group("/api")

    // festivals
    api.GET("/festivals", h.ListFestivals)
    api.GET("/festivals/upcoming", h.ListUpcomingFestivals)
    api.GET("/festivals/:slug", h.GetFestival)
    api.GET("/festivals/:slug/memories", h.ListMemoriesByFestival)

    // memories
    api.POST("/memories", h.CreateMemory)

    // subscriptions
    api.POST("/subscribe", h.Subscribe)
    api.GET("/subscribe/confirm/:token", h.ConfirmSubscription)
    api.GET("/unsubscribe/:token", h.Unsubscribe)

    log.Printf("server starting on port %s", cfg.Port)
    e.Logger.Fatal(e.Start(":" + cfg.Port))
}
