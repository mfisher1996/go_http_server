package sql
import (
    _ "database/sql"
    "github.com/jmoiron/sqlx"
    "os"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

func Init() sqlx.DB {
    godotenv.Load(".env")
    db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        panic(err)
    }
    return *db

}

