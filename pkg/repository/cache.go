package repository

import (
	"log"
	"sync"
	"wb_l0"
)

type Cache struct {
	sync.RWMutex
	data map[int]wb_l0.Order
}

func (c *Cache) Set(data wb_l0.Data) {
	c.Lock()

	defer c.Unlock()
	c.data[data.Id] = data.Order
	log.Printf("Message deployed to cache, id: %d", data.Id)
}
