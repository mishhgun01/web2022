package storage

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Storage - структура драйвера для СУБД Postgres.
type Storage struct {
	mu *sync.Mutex

	// пул соединений.
	pool *pgxpool.Pool
}

func New(conn string) (*Storage, error) {
	p, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &Storage{pool: p, mu: &sync.Mutex{}}, nil
}

// intToInt32Array преобразует массив int в массив int32.
func intToInt32Array(in []int) []int32 {
	var out []int32
	for _, val := range in {
		out = append(out, int32(val))
	}
	return out
}

// Инициализация БД.
func (s *Storage) Init() {
	wd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(wd + "/internal/SQL/schema.sql")
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = s.pool.Exec(context.Background(), string(dat))
	if err != nil {
		log.Fatal(err.Error())
	}
}
