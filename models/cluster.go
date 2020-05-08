package models

import (
	"time"
)

type Cluster struct {
	Id          		int 			`orm:"column(id);pk;auto" description:"id" json:"id"`
	Name        		string			`orm:"column(name);size(255);unique" description:"project name" json:"name"`
	Status        		string			`orm:"column(status);size(255)" description:"status" json:"status"`
	Server				string			`orm:"column(server);size(255)" description:"server" json:"server"`
	Provider			string			`orm:"column(provider);size(255)" description:"provider" json:"provider"`
	Kubeconfig        	string			`orm:"column(kubeconfig);size(255);unique" description:"kubeconfig file" json:"kubeconfig"`
	Description			string			`orm:"column(description)" description:"project description" json:"description"`
	CreatTime			time.Time		`orm:"auto_now_add;type(datetime)" description:"create time" json:"createtime"`
	Project				[]*Project		`orm:"reverse(many)" json:"project"`
	Registry			*Registry		`orm:"rel(fk);index" json:"registry"`
}