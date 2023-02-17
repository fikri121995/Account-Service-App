package main

import (
	"be15/gp4/config"
	"be15/gp4/controllers"
	"be15/gp4/entities"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// establish database connection
	db := config.ConnectToDB()
	defer db.Close()

	var input int
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Add Account (register)")
		fmt.Println("2. Login")
		fmt.Println("3. Read Account")
		fmt.Println("4. Update Account")
		fmt.Println("5. Delete Account")
		fmt.Println("6. Top-up")
		fmt.Println("7. Transfer")
		fmt.Println("8. History Top-up")
		fmt.Println("9. History Transfer")
		fmt.Println("10. View Other User Profile")
		fmt.Println("0. Exit")
		fmt.Print("Input: ")
		fmt.Scan(&input)

		switch input {
		case 1:
			// prompt user for account information
			var nama string
			fmt.Println("masukan nama anda:")
			fmt.Scanf("%s\n", &nama)

			var Akunid string
			fmt.Println("masukan Akunid baru anda:")
			fmt.Scanf("%s\n", &Akunid)

			var Nomor_Telepon int
			fmt.Println("masukan nomor_telepon anda:")
			fmt.Scanf("%d\n", &Nomor_Telepon)

			var Password string
			fmt.Println("masukan Password baru anda:")
			fmt.Scanf("%v\n", &Password)

		case 2:
			// Login
			datarow := entities.Users{}
			fmt.Println("masukan nomor telepon:")
			fmt.Scanln(&datarow.Nomor_Telepon)
			fmt.Println("masukan password:")
			fmt.Scanln(&datarow.Password)
			controllers.Users(db, datarow.Nomor_Telepon, datarow.Password)
		case 3:
			// Read Account
			var Akun int
			fmt.Println("masukan Akun:")
			fmt.Scanln(&Akun)
			controllers.Readacc(db, Akun)
		case 4:
			// Update Account
			var id int
			fmt.Println("Masukan ID anda:")
			fmt.Scanln(&id)
			var Nama string
			fmt.Println("Masukan Nama:")
			fmt.Scanln(&Nama)
			var Nomor_Telepon int
			fmt.Println("Masukan Nomor Telepon:")
			fmt.Scanln(&Nomor_Telepon)
			controllers.Updateacc(db, id, Nama, Nomor_Telepon)
		case 5:
			// Delete Account
			fmt.Print("Masukkan ID akun yang akan dihapus: ")
			var id int
			fmt.Scanln(&id)
			controllers.DeleteAccount(db, id)
		case 6:
			// Top-up
			var AccountID int
			fmt.Println("Masukan id anda disini:")
			fmt.Scanln(&AccountID)
			fmt.Println("masukan nominal topup anda:")
			var Jumlahtopup int
			fmt.Scanln(&Jumlahtopup)
			controllers.Topup(db, AccountID, Jumlahtopup)
		case 7:
			//transfer
			var AccountId int
			fmt.Println("Masukan AccountID Anda:")
			fmt.Scanln(&AccountId)
			var Akun_ID_Penerima int
			fmt.Println("Masukan Akun id penerima:")
			fmt.Scanln(&Akun_ID_Penerima)
			var Jumlah_Transfer int
			fmt.Println("Masukan jumlah transfer:")
			fmt.Scanln(&Jumlah_Transfer)
			var Nama_Bank string
			fmt.Println("Masukan nama bank:")
			fmt.Scanln(&Nama_Bank)
			controllers.Transfer(db, AccountId, Akun_ID_Penerima, Jumlah_Transfer, Nama_Bank)

		case 8:
			// History Top-up
			var id int
			fmt.Println("Masukan id anda:")
			fmt.Scanln(&id)
			controllers.Historytopup(db, id)
		case 9:
			// History Transfer
			var id int
			fmt.Println("masukan id anda:")
			fmt.Scanln(&id)
			controllers.Historytransfer(db, id)
		case 10:
			// View Other User Profile
			var id string
			fmt.Println("masukan id anda:")
			fmt.Scanln(&id)
			controllers.Historyviewother(db, id)
		case 0:
			fmt.Println("Terima kasih telah bertransaksi")
			return
		default:
			fmt.Println("Input salah")
		}
	}
}
