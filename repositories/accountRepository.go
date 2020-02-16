package repositories

import (
	"database/sql"
	"net/http"

	"github.com/ckmu32/GoBankingCore/models"
	_ "github.com/go-sql-driver/mysql"
)

// getConnection Gets a handler for the DB to use later. Make sure to close the connection after the call of this method.
/*
func getConnection() bool {
	db, err := sql.Open("mysql", "bankingCore:rIdo2XfXbJxcojSs@/go_banking_core")
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
*/

// Insert Inserts an account to the DB.
func Insert(account *models.Account, response *models.OperationResponse) bool {
	db, err := sql.Open("mysql", "bankingCore:rIdo2XfXbJxcojSs@/go_banking_core")
	defer db.Close()
	if err != nil {
		println(err.Error())
		response.Code = http.StatusInternalServerError
		response.Description = "Errors ocurred."
		response.Errors = append(response.Errors, "Error getting DB handler - accountRepository (Insert).")
		return false
	}

	// Check that the DB is up.
	err = db.Ping()
	if err != nil {
		println(err.Error())
		response.Code = http.StatusInternalServerError
		response.Description = "Errors ocurred."
		response.Errors = append(response.Errors, "Error while calling DB - accountRepository (Insert).")
		return false
	}

	// Prepare statement for inserting data.
	stmtIns, err := db.Prepare("INSERT INTO `ACCOUNT`(`NUMBER`, `CARD_HOLDER_ID`, `TYPE`, `EXTERNAL_ACCOUNT`, " +
		"`LEVEL`, `AVAILABLE_BALANCE`, `REAL_BALANCE`, `STATUS`, `CREATION_USER`, `MODIFICATION_USER`) " +
		"VALUES (?,?,?,?,?,?,?,?,?,?)")

	defer stmtIns.Close() // Close the statement when we leave this method().
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Description = "Errors ocurred."
		response.Errors = append(response.Errors, "Error on prepare statement - accountRepository (Insert).")
		println(err.Error())
		return false
	}

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
		response.Code = http.StatusInternalServerError
		response.Description = "Errors ocurred."
		response.Errors = append(response.Errors, "Error on inserting statement - accountRepository (Insert).")
		println(err.Error())
		return false
	}

	return true
}

// GetByNumber Returns an account by its number.
func GetByNumber(number string, response *models.OperationResponse) models.Account {
	account := models.Account{}
	db, err := sql.Open("mysql", "bankingCore:rIdo2XfXbJxcojSs@/go_banking_core")
	defer db.Close()
	if err != nil {
		println(err.Error())
		response.Code = http.StatusInternalServerError
		response.Description = "Errors ocurred."
		response.Errors = append(response.Errors, "Error getting DB handler - accountRepository (GetByNumber).")
		return account
	}

	// Check that the DB is up.
	err = db.Ping()
	if err != nil {
		println(err.Error())
		response.Code = http.StatusInternalServerError
		response.Description = "Errors ocurred."
		response.Errors = append(response.Errors, "Error while calling DB - accountRepository (GetByNumber).")
		return account
	}

	//defer db.Close()
	/*
		if !getConnection() {
			response.Errors = append(response.Errors, "Error on getConnection - accountRepository (GetByNumber).")
			return account
		}
	*/

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT `NUMBER`, `CARD_HOLDER_ID`, `TYPE`, `EXTERNAL_ACCOUNT`, `LEVEL`, " +
		"`AVAILABLE_BALANCE`, `REAL_BALANCE`, `STATUS`, `CREATION_DATE`, `CREATION_USER`, `MODIFICATION_DATE`, " +
		"`MODIFICATION_USER` FROM `ACCOUNT` WHERE `NUMBER` = ?")

	defer stmtOut.Close()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Description = "Errors ocurred."
		response.Errors = append(response.Errors, "Error on prepare statement - accountRepository (GetByNumber).")
		println(err.Error())
		return account
	}

	// Execute SELECT.
	err = stmtOut.QueryRow(number).Scan(
		&account.Number,
		&account.CardHolderID,
		&account.Type,
		&account.ExternalAccount,
		&account.Level,
		&account.AvailableBalance,
		&account.RealBalance,
		&account.Status,
		&account.CreationDate,
		&account.CreationUser,
		&account.ModificationDate,
		&account.ModificationUser)

	if err != nil {
		response.Code = http.StatusNotFound
		response.Description = "Errors ocurred."
		response.Errors = append(response.Errors, "Error on getting statement - accountRepository (GetByNumber).")
		println(err.Error())
	}

	return account
}
