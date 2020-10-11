package models

import (
	"DateCertproject/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
Id int 'form:"id"'
Phone string	'form:"phone"'
Password string	'form:"password"'
}

func (u User) SavaUser() {
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.password))
	psswordBytes := md5Hash.Sum(nil)
	u.password = hex.EncodeToString(psswordBytes)

	row, err := db_mysql.Db.Exec("insert into user(phone,password)"+
		"values(?,?) ", u.Phone, u.Password)
	if err != nil {
		return -1, err
	}
		id,err :=row.RowsAffected()
		if err !=nil{
			return -1 , err
		}
	return id, nil
	}
func (u User)QueryUser() (*User,error) {
	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?",
	u.phone,u.Password)
	var phone string
	err := row.Scan(&u.Phone)
	if err !=nil {
		return err
	}
	return &u,err
}