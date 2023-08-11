package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/juancarbajal/gocache/utils"
)

// Registro de info del control de cache
type TCacheControlRecord struct {
	expiration uint32
	created    time.Time
	object     string
}

type TCacheControlRecordList map[string]TCacheControlRecord

// Control de cache
type TCacheControl struct {
	data TCacheControlRecordList
}

// Constructor TCacheControl
func NewTCacheControl() *TCacheControl {
	return &TCacheControl{
		data: make(TCacheControlRecordList),
	}
}

// Add data to cache
func (d TCacheControl) Add(key string, expiration uint32) string {
	objectName := utils.GetMD5Hash(key)
	_, ok := d.data[key]
	if !ok {
		r := &TCacheControlRecord{
			expiration: expiration,
			created:    time.Now(),
			object:     objectName,
		}
		d.data[key] = *r
	}
	return objectName
}

func (d TCacheControl) GetObject(key string) (string, bool) {
	r, exists := d.data[key]
	return r.object, exists
}

// Delete a key from cache control
func (d TCacheControl) Remove(key string) string {
	delete(d.data, key)
	return utils.GetMD5Hash(key)
}

// Validate if expired a key
func (d TCacheControl) IsExpired(key string) bool {
	r, exists := d.data[key]
	if !exists {
		return true
	}
	time.Now().Sub(r.created) // TODO pending implementation
	// time.Duration(r.expiration)
	return true
}

// Save cache control to file
func (d TCacheControl) ToFile(filename string) (bool, error) {
	bytes, err := json.Marshal(d.data)
	if err != nil {
		return false, err
	}
	ioutil.WriteFile(filename, bytes, 0644)
	return true, nil
}

// Get cache control from a file
func (d TCacheControl) FromFile(filename string) (bool, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}
	json.Unmarshal(b, &d.data)
	return true, nil
}
