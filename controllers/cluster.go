package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/buxiaomo/dop/models"
)

type ClusterController struct {
	BaseController
}

func (c *ClusterController) Get() {
	c.Data["Title"] = "集群管理"
	var clusters []models.Cluster
	orm.NewOrm().QueryTable(new(models.Cluster)).All(&clusters)
	c.Data["data"] = clusters
	c.Layout = "layout/layout.html"
	c.TplName = "cluster.html"
}

func (c *ClusterController) List(){
	var orm_clusters [] models.Cluster
	o := orm.NewOrm()
	o.QueryTable(new(models.Cluster)).RelatedSel("Registry").All(&orm_clusters)
	var Clusters []map[string]interface{}
	for _, m := range orm_clusters {
		_, _ = o.LoadRelated(&m, "Project")

		tmp := map[string]interface{}{
			"id": m.Id,
			"name": m.Name,
			"provider": m.Provider,
			"description": m.Description,
			"registry": fmt.Sprintf("%s:%v", m.Registry.Host, m.Registry.Port),
			"project": len(m.Project),
			"creattime": m.CreatTime,
		}
		Clusters = append(Clusters, tmp)
	}

	text := map[string]interface{} {
		"code": 0,
		"message": "",
		"count": len(Clusters),
		"data": Clusters,
	}

	c.Data["json"] = &text
	c.ServeJSON()
}
func (c *ClusterController) Delete(){
	id :=  c.GetString(":id")
	var (
		Cluster models.Cluster
	)
	o := orm.NewOrm()
	o.QueryTable(new(models.Cluster)).Filter("id", id).One(&Cluster)
	o.LoadRelated(&Cluster,"Project")
	if len(Cluster.Project) != 0 {
		c.Ctx.ResponseWriter.WriteHeader(202)
		c.Data["json"] = map[string]interface{}{
			"msg": "集群存在项目，无法删除",
		}
		c.ServeJSON()
		return
	}
	//o.Delete(&Cluster)
	c.Ctx.ResponseWriter.WriteHeader(204)
	c.ServeJSON()
}

//func (c *ClusterController) Info(){
//	id :=  c.GetString(":id")
//
//	var (
//		Cluster models.Cluster
//	)
//	kubeconfig := fmt.Sprintf("./kubeconfig/%v.kubeconfig", id)
//
//	o := orm.NewOrm()
//	o.QueryTable(new(models.Cluster)).Filter("id", id).One(&Cluster)
//	err := ioutil.WriteFile(kubeconfig, []byte(Cluster.Kubeconfig),0644)
//
//	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
//	if err != nil {
//		panic(err.Error())
//	}
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		panic(err.Error())
//	}
//	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
//	if err != nil {
//		panic(err.Error())
//	}
//
//	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
//	if err != nil {
//		panic(err.Error())
//	}
//
//	//c.Data["json"] = map[string]interface{}{
//	//	"node": len(nodes.Items),
//	//	"pods": len(pods.Items),
//	//}
//	c.Data["data"] = map[string]interface{}{
//		"node": len(nodes.Items),
//		"pods": len(pods.Items),
//	}
//	//c.ServeJSON()
//	c.TplName = "cluster_info.html"
//}