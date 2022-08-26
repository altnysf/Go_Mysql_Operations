package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@/golangdb")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	createStatement := "`users` (`ID` INT(11) NOT NULL AUTO_INCREMENT,`Username` varchar(45) NOT NULL,`Email` varchar(45) NOT NULL,`Password` varchar(45) NOT NULL,`FirstName` varchar(45) NOT NULL,`LastName` varchar(45) NOT NULL,`BirthDate` varchar(45) DEFAULT NULL,`IsActive` tinyint(1) DEFAULT NULL,PRIMARY KEY (`ID`),UNIQUE INDEX `ID_UNIQUE` (`ID` ASC) VISIBLE) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8;"
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + createStatement)

	if err != nil {
		log.Fatal(err)
	}

	// ADD DATA
	res, err := db.Exec("INSERT INTO users(Username,Email,Password,FirstName,LastName,BirthDate,IsActive) VALUES('DenemeUser7','webmail@deneme.com','1234+-','DenemeName7','DenemeSurname7','2017.1.1',1)")

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

	// UPDATE DATA

	upt, errUpt := db.Exec("UPDATE users set FirstName =?, LastName =? WHERE id = ?", "DenemeName", "DenemeSurname", 1)

	if errUpt != nil {
		log.Fatal(errUpt)
	}

	rowCount, errUpt1 := upt.RowsAffected()
	if err != nil {
		log.Fatal(errUpt1)
	}
	log.Printf("Inserted %d Rows", rowCount)

	// DELETE DATA

	del, errDel := db.Exec("DELETE FROM users WHERE id =?", 3)

	if errDel != nil {
		log.Fatal(errDel)
	}

	rowCount, errDel1 := del.RowsAffected()
	if err != nil {
		log.Fatal(errDel1)
	}
	log.Printf("Inserted %d Rows", rowCount)

}
