package cache

type TCacheManager interface {
	Add(key string, value string, expiration int64) error
	Remove(key string)
	Find(key string) (string, error)
}

type TCacheManagerFactory struct {
	c TCacheManager
}

func (cm TCacheManagerFactory) create(t string) (TCacheManager, error) {
	if t == "file" {
		cm.c = TCacheFile{folder: "cache"}
		return cm.c, nil
	}
	return nil, nil
}
