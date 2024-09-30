package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DatabaseInterface interface {
	Connect() error
	Close()
	GetConnection() *pgx.Conn
	GetContext() context.Context
}

type database struct {
	connectionString string
	connection       *pgx.Conn
	context          context.Context
}

func NewDatabase() *database {
	return &database{
		connectionString: "postgres://user:password@localhost:5432/authenticator",
		context:          context.Background(),
	}
}

func (database *database) Connect() error {
	conn, err := pgx.Connect(database.context, database.connectionString)
	if err != nil {
		return err
	}
	database.connection = conn
	return nil
}

func (database *database) Close() {
	database.connection.Close(database.context)
}

func (database *database) GetConnection() *pgx.Conn {
	return database.connection
}

func (database *database) GetContext() context.Context {
	return database.context
}
