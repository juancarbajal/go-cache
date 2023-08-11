package cache

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

const FILE_CACHE_CONTROL_FILENAME = "cache.db"

// Cache File
type TCacheFile struct {
	ICacheManager
	path    string
	control *TCacheControl
}

// Constructor TCacheFile
func NewTCacheFile(folder string) *TCacheFile {
	c := &TCacheFile{}
	c.Init(folder)
	return c
}

// Init configuration of CacheFile
func (cf *TCacheFile) Init(path string) {
	cf.path = path
	cf.control = NewTCacheControl()
	// Create the directory
	_, err := os.Stat(cf.path)
	if err == nil { // If exist directory load data
		cf.control.FromFile(filepath.Join(cf.path, FILE_CACHE_CONTROL_FILENAME))
	}
	if os.IsNotExist(err) { //if not exist folder, create the folder
		os.MkdirAll(cf.path, os.ModePerm)
	}
}

func (cf *TCacheFile) Add(key string, value string, expiration uint32) error {
	filename := cf.control.Add(key, expiration)
	byteValue := []byte(value)
	ioutil.WriteFile(filepath.Join(cf.path, filename), byteValue, 0644)
	return nil
}

func (cf *TCacheFile) Remove(key string) {
	filename := cf.control.Remove(key)
	os.Truncate(filepath.Join(cf.path, filename), 0)
}

// Extract value from cache
func (cf *TCacheFile) Find(key string) (string, error) {
	fileKey, exists := cf.control.GetObject(key)
	if !exists {
		return "", errors.New("Not exist key")
	}
	fileName := filepath.Join(cf.path, fileKey)
	b, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	stringValue := string(b)
	return stringValue, nil
}
