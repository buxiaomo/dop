package models

type Role struct {
	Id			int			`orm:"column(id);pk;auto" description:"id" json:"id"`
	Role		string		`orm:"column(role);size(255)" description:"role" json:"role"`
}