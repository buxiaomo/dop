package models

type User struct {
	Id          			int 				`orm:"column(id);pk;auto" description:"id" json:"id"`
	Username        		string				`orm:"column(username);size(255);unique" description:"username" json:"username"`
	Password				string				`orm:"column(password);size(255)" description:"password" json:"password"`
	IsAdmin         		bool         		`orm:"default(false)" description:"IsAdmin" json:"isadmin"`
	Email					string				`orm:"column(email);size(255);unique" description:"email" json:"email"`
	Lock					bool				`orm:"column(lock);default(true)" description:"lock" json:"lock"`
	Project					[]*Project			`json:"project" orm:"reverse(many)"`
}