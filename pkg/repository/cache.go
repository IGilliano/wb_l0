package repository

import (
	"fmt"
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
	fmt.Println(data.Id, data.Order)
	c.data[data.Id] = data.Order
}
