package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "blogadmin"
	dbPass := "password"
	dbName := "blogdb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func runMySqlDbTest() {
	db := dbConn()

	// Select test
	fmt.Println("SELECT TEST")
	selDB, err := db.Query("SELECT id, usr_email, password, ifnull(last_login, ''), ifnull(firstname, ''), " +
		"ifnull(lastname, '') FROM USERS ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id int
		var usrEmail, password, lastLogin, firstName, lastName string
		err = selDB.Scan(&id, &usrEmail, &password, &lastLogin, &firstName, &lastName)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("id : ", id)
		fmt.Println("usrEmail : ", usrEmail)
		fmt.Println("password : ", password)
		fmt.Println("lastLogin : ", lastLogin)
		fmt.Println("firstname : ", firstName)
		fmt.Println("lastname : ", lastName)
	}

	// TODO: Insert Test
	fmt.Println("INSERT TEST")
	usrEmail := "indra.nureska@gmail.com"
	password := "password"
	firstName := "Indra"
	lastName := "Nureska"

	dbInsert, err := db.Prepare("INSERT INTO USERS(usr_email, password, firstname, lastname) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertRes, err := dbInsert.Exec(usrEmail, password, firstName, lastName)
	if err != nil {
		panic(err.Error())
	} else {
		id, err := insertRes.LastInsertId()
		if err != nil {
			panic(err.Error())
		} else {
			// TODO: Update Test
			fmt.Println("UPDATE TEST")
			dbUpdate, err := db.Prepare("UPDATE USERS SET firstname=?, lastname=? WHERE id=?")
			if err != nil {
				panic(err.Error())
			}
			dbUpdate.Exec("ID", "INNUR", id)

			// TODO: Delete Test
			fmt.Println("DELETE TEST")
			delForm, err := db.Prepare("DELETE FROM USERS WHERE id=?")
			if err != nil {
				panic(err.Error())
			}
			delForm.Exec(id)
		}
	}

	defer db.Close()
}
