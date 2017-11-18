package controllers

import (
	"github.com/moaible/sandbox-beego/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserController) Post() {
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	this.Data["json"] = map[string]string{"uid": uid}
	this.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (this *UserController) GetAll() {
	users := models.GetAllUsers()
	this.Data["json"] = users
	this.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (this *UserController) Get() {
	uid := this.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = user
		}
	}
	this.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (this *UserController) Put() {
	uid := this.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(this.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = uu
		}
	}
	this.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (this *UserController) Delete() {
	uid := this.GetString(":uid")
	models.DeleteUser(uid)
	this.Data["json"] = "delete success!"
	this.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (this *UserController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	if models.Login(username, password) {
		this.Data["json"] = "login success"
	} else {
		this.Data["json"] = "user not exist"
	}
	this.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (this *UserController) Logout() {
	this.Data["json"] = "logout success"
	this.ServeJSON()
}
