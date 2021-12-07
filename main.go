package main

import (
	"github.com/gin-gonic/gin"
	"miyuki/tool"
)

func main()  {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	r.Run(80)
}