package main

import (
	"database/sql"
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
)

var (
	Db DB
)

type DB struct {
	*sql.DB
}

// InitalizeTables creates the tables needed in the database if they
// do not already exist.
func (db DB) InitalizeTables() (err error) {
	_, err = db.Query(`CREATE TABLE IF NOT EXISTS purchases (
id INT PRIMARY KEY,
title VARCHAR(255) NOT NULL,
desc VARCHAR(255) NOT NULL,
price DECIMAL(19,4) NOT NULL,
date INT NOT NULL);`)

	// We don't have to check if err != nil because we'd just be
	// returning it anyways and we only have one table.
	return
}