package models

import (
	"log"
	db "web/database"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"Username" form:"Username"`
	Password string `json:"Password" form:"Password"`
}

//添加
func (p *User) Add() (id int64, err error) {

	rs, err := db.Conns.Exec("INSERT INTO user(Username, Password) VALUES (?, ?)", p.Username, p.Password)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

//获取多条
func (p *User) Gets() (users []User, err error) {

	users = make([]User, 0)
	rows, err := db.Conns.Query("SELECT id, Username, Password FROM user")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username, &user.Password)
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return
	}
	return
}

//获取单条
func (p *User) Get() (user User, err error) {
	if p.Id != 0 {
		err = db.Conns.QueryRow("SELECT id, Username, Password FROM user WHERE id=?", p.Id).Scan(
			&user.Id, &user.Username, &user.Password,
		)
	} else {
		err = db.Conns.QueryRow("SELECT id, Username, Password FROM user WHERE username=?", p.Username).Scan(
			&user.Id, &user.Username, &user.Password,
		)
	}
	return
}

//修改
func (p *User) Mod() (ra int64, err error) {
	stmt, err := db.Conns.Prepare("UPDATE user SET Username=?, Password=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.Username, p.Password, p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

//删除
func (p *User) Del() (ra int64, err error) {
	rs, err := db.Conns.Exec("DELETE FROM user WHERE id=?", p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
