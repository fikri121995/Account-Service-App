package controllers

import (
	"fmt"
	"database/sql"
)

func Delete(db *sql.DB,)  {
	
_, err := db.Exec("delete from users where id = ?", "2")
if err != nil {
  fmt.Println(err.Error())
  return
}
fmt.Println("delete success!")  
}

func DeleteAccount(db *sql.DB, id int)  {
	db.Exec("delete from users where id = ?", id)
  
}