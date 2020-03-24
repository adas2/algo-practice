package custom

import (
    "fmt"
)

//MakeSet initialize the set
func makeSet() *customSet {
    return &customSet{
        container: make(map[string]struct{}),
    }
}

type customSet struct {
    container map[string]struct{}
}

func (c *customSet) Exists(key string) bool {
    _, exists := c.container[key]
    return exists
}

func (c *customSet) Add(key string) {
    c.container[key] = struct{}{}
}

func (c *customSet) Remove(key string) error {
    _, exists := c.container[key]
    if !exists {
        return fmt.Errorf("Remove Error: Item doesn't exist in set")
    }
    delete(c.container, key)
    return nil
}

func (c *customSet) Size() int {
    return len(c.container)
}
