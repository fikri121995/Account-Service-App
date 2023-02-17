package controllers

import (
	"database/sql"
	"fmt"
	"be15/gp4/entities"
)

func Historyviewother(db *sql.DB, Akun string) {
	var Users entities.Users
	err := db.
		QueryRow("select Nama, Nomor_Telepon from users where id = ?", Akun).
		Scan(&Users.Nama, &Users.Nomor_Telepon)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Nama: %s\n Nomor_Telepon: %s\n", Users.Nama, Users.Nomor_Telepon)
}
