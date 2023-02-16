package custom

import (
	"fmt"
)

//MakeSet initialize the set
func InitSet() *CustomSet {
	return &CustomSet{
		container: make(map[string]struct{}),
	}
}

type CustomSet struct {
	container map[string]struct{}
}

func (c *CustomSet) Exists(key string) bool {
	_, exists := c.container[key]
	return exists
}

func (c *CustomSet) Add(key string) {
	c.container[key] = struct{}{}
}

func (c *CustomSet) Remove(key string) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(c.container, key)
	return nil
}

func (c *CustomSet) Size() int {
	return len(c.container)
}
