package funbook_db

import (
    "database/sql"
    "os"
    "fmt"
    "path/filepath"
    "io/ioutil"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

func Init() {
    db, err := ConnectDB()
    if err != nil {
        panic( err )
    }
    defer DisconnectDB( db )

    migrationFiles, err := filepath.Glob( "funbook_db/migrations/*.sql" )
    if err != nil {
        panic( err )
    }

    for _, file := range migrationFiles {
        fmt.Println( "Migrating:", file )

        migration, err := ioutil.ReadFile( file )
        if err != nil {
            panic( err )
        }

        _, err = db.Exec( string( migration ) )
        if err != nil {
            panic( err )
        }
    }

    fmt.Println( "Migration complete." )
}


func ConnectDB() ( *sql.DB, error ) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }
    // DBに接続
    connect_str := fmt.Sprintf( "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv( "DB_HOST" ), os.Getenv( "DB_PORT" ), os.Getenv( "DB_USER" ), os.Getenv( "DB_PASSWORD" ), os.Getenv( "DB_NAME" ) )
    db, err := sql.Open( "postgres", connect_str )
    if err != nil {
        return nil, err
    }
    return db, nil
}

func DisconnectDB( db *sql.DB ) error {
    defer db.Close()
    return nil
}