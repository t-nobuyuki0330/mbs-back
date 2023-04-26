package main

import (
    "github.com/t-nobuyuki0330/mbs-back/controller"
    "github.com/t-nobuyuki0330/mbs-back/funbook_db"
    "github.com/gin-gonic/gin"
    "os"
    "github.com/joho/godotenv"
    _ "github.com/gin-contrib/cors"
)

const DEBUG = true

func main() {
    router := gin.Default()

    // CORS設定
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        c.Next()
    })

    router.POST("/funbook/api/search", controller.SearchFunctions )

    funbook_db.Init()

    err := godotenv.Load()
    if err != nil {
        panic( "Error loading .env file" )
    } else {
        if !DEBUG {
            router.RunTLS( ":" + os.Getenv( "APP_PORT" ), os.Getenv( "SERVER_CRT" ), os.Getenv( "SERVER_KEY" ) )
        } else {
            router.Run( ":" + os.Getenv( "APP_PORT" ) )
        }
    }
}

