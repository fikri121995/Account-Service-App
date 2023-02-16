package controllers

import(
	"fmt"
	"database/sql"
)

func Updateacc(db *sql.DB, id int, Nama string, Nomor_Telepon int)  {
	_, err := db.Exec("UPDATE Users set Nama = ?, Nomor_Telepon = ? WHERE ID =? ", Nama, Nomor_Telepon, id)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("Update success!")
}