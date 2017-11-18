package controllers

import (
	"github.com/astaxie/beego"
	"github.com/moaible/sandbox-beego/models"
	"encoding/json"
)

// Operations about Rabbit
type RabbitController struct {
	beego.Controller
}

// @Title Create
// @Description create Rabbit
// @Param	body		body 	models.Rabbit	true		"The Rabbit content"
// @Success 200 {string} models.Rabbit.Id
// @Failure 403 body is empty
// @router / [post]
func (this *RabbitController) Post() {
	var ob models.Rabbit
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	objectid := models.AddRabbit(ob)
	this.Data["json"] = map[string]string{"ObjectId": objectid}
	this.ServeJSON()
}

// @Title Get
// @Description find object by Rabbit id
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (this *RabbitController) Get() {
	objectId := this.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = ob
		}
	}
	this.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (this *RabbitController) GetAll() {
	obs := models.GetAll()
	this.Data["json"] = obs
	this.ServeJSON()
}

// @Title Update
// @Description update the Rabbit
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (this *RabbitController) Put() {
	objectId := this.Ctx.Input.Param(":objectId")
	var ob models.Rabbit
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = "update success!"
	}
	this.ServeJSON()
}

// @Title Delete
// @Description delete the Rabbit
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (this *RabbitController) Delete() {
	objectId := this.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	this.Data["json"] = "delete success!"
	this.ServeJSON()

}
