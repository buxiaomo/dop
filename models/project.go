package models

type Project struct {
	Id          		int 				`orm:"column(id);pk;auto" description:"id" json:"id"`
	Name        		string				`orm:"column(name);size(255);unique" description:"project name" json:"name"`
	Description			string				`orm:"column(description)" description:"project description" json:"description" `
	Owner				*User				`orm:"rel(fk);index" json:"owner"`
	Cluster				*Cluster			`orm:"rel(fk);index" json:"cluster"`
}