package memcache

import (
	"github.com/astaxie/beego/cache"
	"log"
)

var Memcache cache.Cache

func init() {
	var err error
	if Memcache == nil {
		// Memcache, err = cache.NewCache("memcache", `{"conn":"127.0.0.1:11211"}`)
		Memcache, err = cache.NewCache("memory", `{"interval":60}`)
		if err != nil {
			log.Println(err)
		}
	}
}
