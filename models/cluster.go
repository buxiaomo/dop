package models

type Cluster struct {
	Id          		int 			`orm:"column(id);pk;auto" description:"id" json:"id"`
	Name        		string			`orm:"column(name);size(255);unique" description:"project name" json:"name"`
	Kubeconfig        	string			`orm:"column(kubeconfig);size(255);unique" description:"kubeconfig file" json:"kubeconfig"`
	Description			string			`orm:"column(description)" description:"project description" json:"description"`
	Project				[]*Project		`orm:"reverse(many)" json:"project"`
}