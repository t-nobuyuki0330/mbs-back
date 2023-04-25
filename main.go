package main

import (
    "github.com/t-nobuyuki0330/mbs-back/controller"
    "github.com/t-nobuyuki0330/mbs-back/funbook_db"
    "github.com/gin-gonic/gin"
    "os"
    "github.com/joho/godotenv"
    "github.com/gin-contrib/cors"
)

const DEBUG = false

func main() {
    router := gin.Default()
    config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    router.Use(cors.New(config))

    router.POST("/funbook/api/search", controller.SearchFunctions )

    funbook_db.Init()

    err := godotenv.Load()
    if err != nil {
        panic( "Error loading .env file" )
    } else {
        if !DEBUG {
            router.RunTLS( ":" + os.Getenv( "APP_PORT" ), os.Getenv( "SERVER_PEM" ), os.Getenv( "SERVER_KEY" ) )
        } else {
            router.Run( ":" + os.Getenv( "APP_PORT" ) )
        }
    }
}
