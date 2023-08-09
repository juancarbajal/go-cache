package cache

import (
	"database/sql"
	"log"
)

type TCacheSqlite struct {
	db sql.DB
}

var tcsqlitei *TCacheSqlite = nil

func (cs TCacheSqlite) GetInstance() *TCacheSqlite {
	if tcsqlitei == nil {
		tcsqlitei = new(TCacheSqlite)
		tcsqlitei.open()
	}
	return tcsqlitei
}
func (cs TCacheSqlite) Add(key string, value string, expiration int64) error {

}
func (cs TCacheSqlite) Remove(key string) {

}
func (cs TCacheSqlite) Find(key string) (string, error) {

}

func (cs TCacheSqlite) open() error {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	cs.db := db
	return err
}
