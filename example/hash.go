package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/lidedong/cache"
	_ "github.com/lidedong/cache/redis"
	"os"
)

func main() {
	bm, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379", "key":"lidedong"}`)
	if err != nil {
		fmt.Println("NewCache err!", err.Error())
		os.Exit(-1)
	}

	err = bm.HMSet("coolblue", "name", "lide", "age", 14, "girl", false)

	if err != nil {
		fmt.Println("Put err")
		os.Exit(-1)
	}

	if !bm.IsExist("coolblue") {
		fmt.Println("check err")
	}
	if v, _ := redis.Values(bm.HMGet("coolblue", "name", "age", "girl"), err); v != nil {
		fmt.Println("coolblue name:", cache.GetString(v[0]))
		fmt.Println("coolblue age:", cache.GetInt(v[1]))
		fmt.Println("coolblue girl:", cache.GetBool(v[2]))
	}
	if bm.ClearAll() != nil {
		fmt.Println("clearall failed")
	}
}
