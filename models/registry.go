package models

type Registry struct {
	Id          			int 			`orm:"column(id);pk;auto" description:"id" json:"id"`
	Name        			string			`orm:"column(name);size(255);unique" description:"project name" json:"name"`
	Host        			string			`orm:"column(host);size(255)" description:"status" json:"host"`
	Port        			int				`orm:"column(port);size(255);default(443)" description:"status" json:"port"`
	Username        		string			`orm:"column(username);size(255)" description:"registry username" json:"username"`
	Password				string			`orm:"column(password);size(255)" description:"registry password" json:"password"`
	Cluster					[]*Cluster    	`orm:"reverse(many)" json:"cluster"`
}