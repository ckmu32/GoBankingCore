package repositories

import (
	"database/sql"

	"github.com/ckmu32/GoBankingCore/models"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// getConnection Gets a handler for the DB to use later. Make sure to close the connection after the call of this method.
func getConnection() bool {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		println(err.Error())
		return false
	}
	//defer db.Close()

	// Check that the DB is up.
	err = db.Ping()
	if err != nil {
		println(err.Error())
		return false
	}

	return true
}

// Insert Inserts an account to the DB.
func Insert(account *models.Account, response *models.OperationResponse) bool {
	defer db.Close()
	if !getConnection() {
		response.Errors = append(response.Errors, "Error on getConnection - accountRepository (Insert).")
		return false
	}

	// Prepare statement for inserting data.
	stmtIns, err := db.Prepare("INSERT INTO `ACCOUNT`(`NUMBER`, `CARD_HOLDER_ID`, `TYPE`, `EXTERNAL_ACCOUNT`, " +
		"`LEVEL`, `AVAILABLE_BALANCE`, `REAL_BALANCE`, `STATUS`, `CREATION_USER`, `MODIFICATION_USER`) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		response.Errors = append(response.Errors, "Error on prepare statement - accountRepository (Insert).")
		println(err.Error())
		return false
	}
	defer stmtIns.Close() // Close the statement when we leave this method().

	// Insert the model attributes.
	_, err = stmtIns.Exec(
		account.Number,
		account.CardHolderID,
		account.Type,
		account.ExternalAccount,
		account.Level,
		account.AvailableBalance,
		account.RealBalance,
		account.Status,
		account.CreationUser,
		account.ModificationUser)
	if err != nil {
		response.Errors = append(response.Errors, "Error on inserting statement - accountRepository (Insert).")
		println(err.Error())
		return false
	}

	return true
}

// GetByNumber Returns an account by its number.
func GetByNumber(number string, response *models.OperationResponse) models.Account {
	account := models.Account{}
	defer db.Close()
	if !getConnection() {
		response.Errors = append(response.Errors, "Error on getConnection - accountRepository (GetByNumber).")
		return account
	}

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT `NUMBER`, `CARD_HOLDER_ID`, `TYPE`, `EXTERNAL_ACCOUNT`, `LEVEL`, " +
		"`AVAILABLE_BALANCE`, `REAL_BALANCE`, `STATUS`, `CREATION_DATE`, `CREATION_USER`, `MODIFICATION_DATE`, " +
		"`MODIFICATION_USER` FROM `account` WHERE `NUMBER` = ?")
	if err != nil {
		response.Errors = append(response.Errors, "Error on prepare statement - accountRepository (GetByNumber).")
		println(err.Error())
		return account
	}
	defer stmtOut.Close()

	// Execute SELECT.
	err = stmtOut.QueryRow(number).Scan(
		account.Number,
		account.CardHolderID,
		account.Type,
		account.ExternalAccount,
		account.Level,
		account.AvailableBalance,
		account.RealBalance,
		account.Status,
		account.CreationDate,
		account.CreationUser,
		account.ModificationDate)

	if err != nil {
		response.Errors = append(response.Errors, "Error on getting statement - accountRepository (GetByNumber).")
		println(err.Error())
		return account
	}

	return account
}
