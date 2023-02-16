package controllers

import (
	"database/sql"
	"fmt"
)

func Transfer(db *sql.DB, AccountId int, Akun_ID_Penerima int, Jumlah_Transfer int, Nama_Bank string) (saldo int, status bool) {
	var query = "INSERT into Transfer(AccountID, Akun_ID_Penerima, Jumlah_Transfer, Nama_Bank) values (?, ?, ?, ?) "

	statement, _ := db.Prepare(query)
	statement.Exec(AccountId, Akun_ID_Penerima, Jumlah_Transfer, Nama_Bank)

	_, err := db.Exec("update Users set Balance = Balance - ? where ID = ? ", Jumlah_Transfer, AccountId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	_, err2 := db.Exec("update Users set Balance = Balance + ? where ID = ? ", Jumlah_Transfer, Akun_ID_Penerima)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	fmt.Println("update success!")

	return saldo, status

}
