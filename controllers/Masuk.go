package controllers

import (
	"database/sql"
	"fmt"
	"be15/gp4/entities"
)

func Masuk(db *sql.DB, ) {
	defer db.Close()
  var result entities.Masuk
  var Akun = "users1"
  err := db.
    QueryRow("select Akun, password from users where id = ?", Akun).
    Scan(&result.Akun, &result.Password)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Printf("Akun: %s\npassword: %s\n", result.Akun, result.Password)
	
}


