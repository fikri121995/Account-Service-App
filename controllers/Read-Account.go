package controllers

import (
	"database/sql"
	"fmt"
	"be15/gp4/entities"
)

func Readacc(db *sql.DB, ID int) {
	var Users entities.Users 
	err := db.
	  QueryRow("select ID, Akun from Users where id = ?", ID).
	  Scan(&Users.ID, &Users.Akun)
	if err != nil {
	  fmt.Println(err.Error())
	  return
	}
  
	fmt.Printf("ID: %d\nAkun: %s\n", Users.ID, Users.Akun)
}
