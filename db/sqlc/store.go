package db

import "github.com/jackc/pgx/v5/pgxpool"

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
}


type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}


func NewStore(conn *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: conn,
		Queries: New(conn),
	}
}
