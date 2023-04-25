package main 

import (
    "github.com/t-nobuyuki0330/mbs-back/controller"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "time"
)

func main() {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        // allow site
        // AllowOrigins: []string{
        //     "https://funbook.pages.dev/",
        // },
        AllowAllOrigins: true,
        // allow http method
        AllowMethods: []string{
            "POST",
            "GET",
        },
        // allow http request header
        AllowHeaders: []string{
            "Access-Control-Allow-Credentials",
            "Access-Control-Allow-Headers",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "Authorization",
        },
        // info
        AllowCredentials: true,
        // preflight request cache time
        MaxAge: 24 * time.Hour,
    }))

    router.POST("/funbook/api/search", controller.SearchFunctions)

    router.Run(":9669")
}
