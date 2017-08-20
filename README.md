# safemap
map read and write safe

Install using:

    go get github.com/githubyang/safemap

# example

```golang
package main

import (
	"fmt"
	"safemap"
)

func main() {
	Map := safemap.NewSafeMap()

	for i := 0; i < 100; i++ {
		go writeMap(Map, i, i)
		go readMap(Map, i)
	}

}

func readMap(Map *safemap.SafeMap, key int) int {
	s, ok := (Map.Get(key)).(int)
	if ok {
		fmt.Println(s)
		return s
	} else {
		return 0
	}
}

func writeMap(Map *safemap.SafeMap, key int, value int) {
	Map.Set(key, value)
}
```
