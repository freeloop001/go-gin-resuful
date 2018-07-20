package models

import (
	"log"
	db "test/rest/database"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) Get() (person Person, err error) {
	row := db.SqlDB.QueryRow("SELECT id, first_name, last_name FROM person WHERE id=?", p.Id)
	err = row.Scan(&person.Id, &person.FirstName, &person.LastName)
	if err != nil {
		return
	}
	return
}

func (p Person) GetAll() (persons []Person, err error) {
	rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	return
}

func (p *Person) Add () (id int64, err error) {
	stmt, err := db.SqlDB.Prepare("INSERT INTO person(first_name, last_name) VALUES (?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	rs, err := stmt.Exec(p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (p Person) Update() (rows int64, err error) {
	stmt, err := db.SqlDB.Prepare("UPDATE person SET first_name=?, last_name=? WHERE id=?")
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rs, err := stmt.Exec(p.FirstName, p.LastName, p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	rows, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (p Person) Del() (rows int64, err error) {
	stmt, err := db.SqlDB.Prepare("DELETE FROM person WHERE id=?")
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rs, err := stmt.Exec(p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	rows, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	return
}
