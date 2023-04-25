package main 

import (
    "github.com/t-nobuyuki0330/mbs-back/controller"
    "github.com/gin-gonic/gin"
    _ "github.com/gin-contrib/cors"
    _ "time"
)

func main() {
    router := gin.Default()

    // config := cors.DefaultConfig()
    // config.AllowAllOrigins = true
    // router.Use(cors.New(config))

    router.POST("/funbook/api/search", controller.SearchFunctions)

    router.Run(":9669")
}
