package main

import (
    "github.com/t-nobuyuki0330/mbs-back/controller"
    "github.com/gin-gonic/gin"
    "os"
    "github.com/joho/godotenv"
)

const DEBUG = true

func main() {
    router := gin.Default()

    router.POST("/funbook/api/search", controller.SearchFunctions )

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
