package models

type Provider struct {
	Id          		int 			`orm:"column(id);pk;auto" description:"id" json:"id"`
	Name        		string			`orm:"column(name);size(255);unique" description:"provider name" json:"name"`
}