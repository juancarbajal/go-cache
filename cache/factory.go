package cache

type ICacheManager interface {
	Add(key string, value string, expiration uint32) error
	Remove(key string)
	Find(key string) (string, error)
}

type TCacheManagerFactory struct {
}

var staticCacheFactory TCacheManagerFactory

func (cm TCacheManagerFactory) GetCache(t string, options map[string]string) (ICacheManager, error) {
	switch t {
	case "redis":
		r := TCacheRedis{}
		r.Init(options["host"], options["port"], options["password"])
		return r, nil
	case "sqlite":
		s := TCacheSqlite{}
		s.Init(nil)
		return s, nil
	default:
		f := NewTCacheFile(options["path"])
		return f, nil
	}
}
