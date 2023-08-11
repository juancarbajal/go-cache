package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/juancarbajal/gocache/cache"
)

func main() {
	f := new(cache.TCacheManagerFactory)
	// options := map[string]string{"host": "172.17.0.3", "port": "6379", "password": ""}
	wd, err := os.Getwd()
	fmt.Println(wd)
	c, err := f.GetCache("file", map[string]string{"path": filepath.Join(wd, "_cache_")})
	if err != nil {
		log.Fatal(err)
	}

	errAdd := c.Add("f1", "hola esta es una prueba de cache", 30)
	if errAdd == nil {
		fmt.Println("ADD - OK")
	} else {
		fmt.Println(errAdd)
	}

	val, errFind := c.Find("f1")
	if errFind != nil {
		log.Fatal(err)
	}
	fmt.Println("GET - OK")
	fmt.Println(val)

}
