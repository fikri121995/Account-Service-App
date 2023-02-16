package controllers

import (
	"database/sql"
	"fmt"
	"be15/gp4/entities"
)

func Historytransfer(db *sql.DB, AccountId int) {
	var Users entities.Account
	err := db.
	  QueryRow("select Jumlah_Transfer, Akun_ID_Penerima from Transfer where AccountId = ?", AccountId).
	  Scan(&Users.Jumlah_Transfer, &Users.Akun_ID_Penerima)
	if err != nil {
	  fmt.Println(err.Error())
	  return
	}
  
	fmt.Printf("Jumlah_Transfer: %d\nAkun_ID_Penerima: %d\n", Users.Jumlah_Transfer, Users.Akun_ID_Penerima)	
}
