package controllers

import(
	"fmt"
	"database/sql"
)

func Register(db *sql.DB, Akun, Nama string, Nomor_Telepon int, Password string) {
    var query = "INSERT into Users(Akun, Nama, Nomor_Telepon, Password, Balance) values (?, ?, ?, ?, 0)"
    statement, _ := db.Prepare(query)
    statement.Exec(Akun, Nama, Nomor_Telepon, Password)

    _, err := db.Exec("update Users set Nama = ?, Nomor_Telepon = ?, Password = ? where Akun = ?", Nama, Nomor_Telepon, Password, Akun)

    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("insert success!")
}
