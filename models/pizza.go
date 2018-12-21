package models

import (
	"log"
	db "web/database"
)

type Pizza struct {
	Id          int
	Name        string
	Description string
	Options     string
}

//添加
func (p *Pizza) Add() (id int64, err error) {

	rs, err := db.Conns.Exec("INSERT INTO pizza(Name, Description, Options) VALUES (?, ?, ?)", p.Name, p.Description, p.Options)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

//获取多条
func (p *Pizza) Gets() (pizzas []Pizza, err error) {

	pizzas = make([]Pizza, 0)
	rows, err := db.Conns.Query("SELECT id, Name, Description, Options FROM pizza")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var pizza Pizza
		rows.Scan(&pizza.Id, &pizza.Name, &pizza.Description, &pizza.Options)
		pizzas = append(pizzas, pizza)
	}

	if err = rows.Err(); err != nil {
		return
	}
	return
}

//获取单条
func (p *Pizza) Get() (pizza Pizza, err error) {
	if p.Id != 0 {
		err = db.Conns.QueryRow("SELECT id, Name, Description, Options FROM pizza WHERE id=?", p.Id).Scan(
			&pizza.Id, &pizza.Name, &pizza.Options, &pizza.Description,
		)
	} else {
		err = db.Conns.QueryRow("SELECT id, Name, Description, Options FROM user WHERE name=?", p.Name).Scan(
			&pizza.Id, &pizza.Name, &pizza.Options, &pizza.Description,
		)
	}
	return
}

//修改
func (p *Pizza) Mod() (ra int64, err error) {
	stmt, err := db.Conns.Prepare("UPDATE pizza SET Name=?, Description=? Options=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.Name, p.Description, p.Options, p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

//删除
func (p *Pizza) Del() (ra int64, err error) {
	rs, err := db.Conns.Exec("DELETE FROM pizza WHERE id=?", p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
