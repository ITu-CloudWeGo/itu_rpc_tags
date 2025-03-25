package main

import (
	tags_service "github.com/ITu-CloudWeGo/itu_rpc_tags/kitex_gen/tags_service/tagsservice"
	"log"
)

func main() {
	svr := tags_service.NewServer(new(TagsServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
