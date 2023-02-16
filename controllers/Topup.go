package controllers

import(
	"database/sql"
	
)

func Topup(db *sql.DB, AccountID int, Jumlahtopup int) (int, bool) {
    // Retrieve current balance for account
    var saldo int
    err := db.QueryRow("SELECT Balance FROM Users WHERE ID = ?", AccountID).Scan(&saldo)
    if err != nil {
        return 0, false
    }

    // Add top-up amount to balance
    saldo += Jumlahtopup

    // Update balance in database
    _, err = db.Exec("UPDATE Users SET Balance = ? WHERE ID = ?", saldo, AccountID)
    if err != nil {
        return 0, false
    }

    return saldo, true
}
