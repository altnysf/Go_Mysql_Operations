package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/golangdb")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	createStatement := "`users` (`ID` INT(11) NOT NULL AUTO_INCREMENT,`Username` varchar(45) NOT NULL,`Email` varchar(45) NOT NULL,`Password` varchar(45) NOT NULL,`FirstName` varchar(45) NOT NULL,`LastName` varchar(45) NOT NULL,`BirthDate` varchar(45) DEFAULT NULL,`IsActive` tinyint(1) DEFAULT NULL,PRIMARY KEY (`ID`),UNIQUE INDEX `ID_UNIQUE` (`ID` ASC) VISIBLE)	  ENGINE = InnoDBDEFAULT CHARACTER SET = utf8;"
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + createStatement)

	if err != nil {
		log.Fatal(err)
	}

	// ADD DATA
	res, err := db.Exec("INSERT INTO users(Username,Email,Password,FirstName,LastName,BirthDate,IsActive) VALUES('DenemeUser','webmail@deneme.com','1234+-','Yusuf','ALTUN','2017.1.1',1)")

	if err != nil {
		log.Fatal(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Inserted %d Rows", rowCount)

	// GET DATA

	var (
		ID        int
		Username  string
		Email     string
		Password  string
		FirstName string
		LastName  string
		BirthDate string
		IsActive  bool
	)

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Bulunan Satır İçeriği : %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))

	}
}
