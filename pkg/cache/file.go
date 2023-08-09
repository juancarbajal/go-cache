package cache

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/juancarbajal/go-cache/pkg/utils"
)

// Cache File
type TCacheFile struct {
	folder string
}

func (cf TCacheFile) Add(key string, value string, expiration int64) error {
	if _, err := cf.existsPathOrCreate(cf.folder); err != nil { //creamos la carpeta si no existe
		return err
	}
	fileKey := utils.GetMD5Hash(key)
	byteValue := []byte(value)
	ioutil.WriteFile(filepath.Join(cf.folder, fileKey), byteValue, 0644)
	return nil
}

func (cf TCacheFile) Remove(key string) {

}
func (cf TCacheFile) Find(key string) (string, error) {
	fileKey := utils.GetMD5Hash(key)
	fileName := filepath.Join(cf.folder, fileKey)
	b, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	stringValue := string(b)
	return stringValue, nil
}

func (cf TCacheFile) existsPathOrCreate(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		if err2 := os.MkdirAll(cf.folder, os.ModePerm); err != nil {
			return false, err2
		} else {
			return true, nil
		}
	}
	return false, err
}
