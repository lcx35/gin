package main

import (
	db "web/database"
	"web/libs"
	"web/router"
)

func main() {
	defer db.Conns.Close()
	r := router.InitRouter()
	r.Run(":" + libs.Conf.Read("site", "httpport"))
}
