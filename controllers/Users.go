package controllers

import (
	"be15/gp4/entities"
	"database/sql"
	"fmt"
	"log"
)

func Users(db *sql.DB, Nomor_Telpon string, Password string) {
	
	rows, errSelect := db.Query("SELECT ID, Akun, Nama, Nomor_Telepon, Password, Balance FROM Users WHERE Nomor_Telepon=? and Password=?", Nomor_Telpon, Password) 
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}
	var allUsers []entities.Users
	for rows.Next() {
		var datarow entities.Users
		errScan := rows.Scan(&datarow.ID, &datarow.Akun, &datarow.Nama, &datarow.Nomor_Telepon, &datarow.Password, &datarow.Balance)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		allUsers = append(allUsers, datarow)
		// fmt.Println(datarow)
	}
	if len(allUsers) == 0 {
		fmt.Println("Gagal masuk, cek kembali nomor telepon dan password")
		return
	}

	// fmt.Println(allGuru)
	for _, v := range allUsers {
		fmt.Println("Akun:", v.Akun, "Saldo:", v.Balance)
	}
	
}

func InsertUsers(db *sql.DB, newUsers entities.Users) {
	var query = "INSERT INTO Users(ID, Akun, Nama, Nomor_Telepon, Password, Balance) VALUES(?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
	}

	result, errInsert := statement.Exec(newUsers.ID, newUsers.Akun, newUsers.Nama, newUsers.Nomor_Telepon, newUsers.Password, newUsers.Balance)
	if errInsert != nil {
		log.Fatal("error exec insert", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("proses berhasil dijalankan")
		} else {
			fmt.Println("proses gagal")
		}
	}
}
