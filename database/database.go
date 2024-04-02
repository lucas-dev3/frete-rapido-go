package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
)

var (
	once     sync.Once
	instance *pgx.Conn
)

func Connection() *pgx.Conn {
	once.Do(func() {
		conn, err := pgx.Connect(context.Background(), "postgresql://postgres:senhasupersegura@db:5432/frete_rapido")

		if err != nil {
			fmt.Printf("Não foi possível conectar ao banco de dados: %v\n", err)
			return
		}
		fmt.Println("Conexão com o banco de dados estabelecida")
		instance = conn
	})

	return instance
}
