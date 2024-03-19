package db

import (
    "database/sql"
    "log"
    "sync"
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
    "os"
    "strconv"
    "fmt"
)

var (
    DbCon    *sql.DB
    dbOnce sync.Once
)

// InitializeDB creates a global database connection using the Singleton pattern
func InitializeDB() error {
  if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
  }

  dbHost := os.Getenv("DB_HOST")
  dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
  dbUser := os.Getenv("DB_USER")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_NAME")
  poolSize, _ := strconv.Atoi(os.Getenv("DB_POOL_SIZE"))
  
  var err error
  dbOnce.Do(func() {
    mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
    DbCon, err = sql.Open("mysql", mysqlInfo)
    if err != nil {
      return
    }
    
    DbCon.SetMaxOpenConns(poolSize) // Set maximum number of open connections
    DbCon.SetMaxIdleConns(poolSize) // Set maximum number of idle connections
    err = DbCon.Ping()
  })

  return err
}

