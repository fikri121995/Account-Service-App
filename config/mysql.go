package config

import (
  "database/sql"
  "fmt"
  "log"
  "time"
)

func ConnectToDB() *sql.DB {
  // <username>:<password>@tcp(<hostname>:<port>)/<db_name>
  user := "root"
  pass := ""
  host := "localhost"
  port := "3306"
  dbname := "gp_4"

  connectionString := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
  db, errOpen := sql.Open("mysql", connectionString)
  if errOpen != nil {
    log.Fatal("Error open connection", errOpen.Error())
  }

  // See "Important settings" section.
  db.SetConnMaxLifetime(time.Minute * 3)
  db.SetMaxOpenConns(10)
  db.SetMaxIdleConns(10)

  errPing := db.Ping()
  if errPing != nil {
    log.Fatal("error connect to db", errPing.Error())
  } else {
    fmt.Println("berhasil")
  }
  fmt.Println("tambahan bari dari branch B")
  fmt.Println("tambahan dari branch A")
  return db

}