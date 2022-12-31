package store

import "WB_L0/internal/models"

type Cache struct {
	buffer map[int]models.Order
	pos    int
}

func NewCache(store *Store) *Cache {
	cache := Cache{}
	cache.Init(store)
	return &cache
}

func (c *Cache) Init(store *Store) {

}

func (c *Cache) restoreCacheFromDB() {

}
