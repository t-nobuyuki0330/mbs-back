package main 

import (
    "github.com/t-nobuyuki0330/mbs-back/controller"
    "github.com/t-nobuyuki0330/mbs-back/funbook_db"
    "github.com/gin-gonic/gin"

    "os"
    "github.com/joho/godotenv"
)

func main() {
    router := gin.Default()

    router.POST( "/funbook/api/search", controller.SearchFunctions )

    funbook_db.Init ()

    err := godotenv.Load()
    if err != nil {
        panic( "Error loading .env file" )
    } else {
        router.Run( ":" + os.Getenv( "APP_PORT" ) )
    }
}