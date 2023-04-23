package main 

import (
    "github.com/t-nobuyuki0330/mbs-back/controller"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.POST("/funbook/api/search", controller.SearchFunctions )

    router.Run(":9669")
}