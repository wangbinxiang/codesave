package memory

import (
	"github.com/astaxie/beego/cache"
	"log"
)

var Memory cache.Cache

func init() {
	var err error
	if Memory == nil {
		Memory, err = cache.NewCache("memory", `{"interval":60}`)
		if err != nil {
			log.Println(err)
		}
	}
}
