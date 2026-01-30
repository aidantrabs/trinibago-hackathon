package handler

import (
    "github.com/aidantrabs/trinbago-hackathon/backend/internal/db"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
    queries *db.Queries
    pool    *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Handler {
    return &Handler{
        queries: db.New(pool),
        pool:    pool,
    }
}
