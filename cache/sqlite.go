package cache

import (
	"database/sql"
	"log"
)

type TCacheSqlite struct {
	ICacheManager
	db sql.DB
}

func (cs TCacheSqlite) Add(key string, value string, expiration uint32) error {
	return nil
}
func (cs TCacheSqlite) Remove(key string) {

}
func (cs TCacheSqlite) Find(key string) (string, error) {
	return "", nil
}

func (cs TCacheSqlite) Init(options map[string]string) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	cs.db = *db
}
