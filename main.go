package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"wanghaojun.whj/whjFirstPriject/common"
	"wanghaojun.whj/whjFirstPriject/router"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = router.CollectRouter(r)

	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}
