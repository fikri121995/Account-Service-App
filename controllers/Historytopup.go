package controllers

import (
	"database/sql"
	"fmt"
	"be15/gp4/entities"
)

func Historytopup(db *sql.DB, AccountID int) {
	var Users entities.Topup 
	err := db.
	  QueryRow("select Tanggal_Topup, Jumlah_Topup from Topup where AccountID = ?", AccountID).
	  Scan(&Users.Tanggal_Topup, &Users.Jumlahtopup)
	if err != nil {
	  fmt.Println(err.Error())
	  return
	}
  
	fmt.Printf("Tanggal_Topup: %s\nJumlah_Topup: %d\n", Users.Tanggal_Topup, Users.Jumlahtopup)
	  
}
