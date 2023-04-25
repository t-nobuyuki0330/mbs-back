package main 

import (
    "github.com/t-nobuyuki0330/mbs-back/controller"
    "github.com/t-nobuyuki0330/mbs-back/funbook_db"
    "github.com/gin-gonic/gin"
    "os"
    "time"
    "github.com/joho/godotenv"
)

func main() {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        // allow site
        AllowOrigins: []string{
            "https://funbook.pages.dev/",
        },
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

    router.POST( "/funbook/api/search", controller.SearchFunctions )

    funbook_db.Init ()

    err := godotenv.Load()
    if err != nil {
        panic( "Error loading .env file" )
    } else {
        router.Run( ":" + os.Getenv( "APP_PORT" ) )
    }
}