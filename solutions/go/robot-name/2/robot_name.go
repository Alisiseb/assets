package robotname

import (
	"errors"
	"math/rand"
	"sync"
)

// Define the Robot type here.
type Robot struct {
	title string
}

var registry = make(map[string]bool, 0)
var lock sync.Mutex

func (r *Robot) Name() (string, error) {

	if len(registry) == (26 * 26 * 1000) {
		return "", errors.New("name capacity is full")
	}
	if (*r).title == "" {
		for {
			lock.Lock()
			name := sampleName()
			if !registry[name] {
				registry[name] = true
				(*r).title = name
				return name, nil
			}
			lock.Unlock()
		}
	} else {
		return (*r).title, nil
	}

}

func sampleName() string {

	name := ""
	for i := 0; i < 5; i++ {
		if len(name) < 2 {
			name += string(rune(rand.Intn(26) + 'A'))
		} else {
			name += string(rune(rand.Intn(10) + '0'))
		}
	}

	return name
}

func (r *Robot) Reset() {
	delete(registry, (*r).title)
	(*r).title = ""
}
