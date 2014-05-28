package memcache

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/memcache"
	"log"
)

var Memcache cache.Cache

func init() {
	var err error
	if Memcache == nil {
		Memcache, err = cache.NewCache("memcache", `{"conn":"127.0.0.1:11211"}`)
		if err != nil {
			log.Println(err)
		}
	}
}
