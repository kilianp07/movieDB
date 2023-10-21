package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/kilianp07/movieDB/models"
)

// ActorController operations for Actor
type ActorController struct {
	beego.Controller
}

// URLMapping ...
func (c *ActorController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Actor
// @Param	body		body 	models.Actor	true		"body for Actor content"
// @Success 201 {object} models.Actor
// @Failure 403 body is empty
// @router / [post]
func (c *ActorController) Post() {
	actor := models.Actor{}
	if err := c.ParseForm(&actor); err != nil {
		c.Data["json"] = err.Error()
	} else {
		if _, err := models.AddActor(&actor); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = "Acteur créé avec succès"
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Actor by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Actor
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ActorController) GetOne() {
	actorID, _ := c.GetInt(":id")
	actor, err := models.GetActorById(int64(actorID))
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = actor
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Actor
// @Success 200 {object} models.Actor
// @Failure 403
// @router / [get]
func (c *ActorController) GetAll() {
	filters := make(map[string]string)
	fields := make([]string, 0)
	sortby := make([]string, 0)
	order := make([]string, 0)
	limit := int64(10)
	offset := int64(0)

	actors, err := models.GetAllActor(filters, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = actors
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Actor
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Actor	true		"body for Actor content"
// @Success 200 {object} models.Actor
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ActorController) Put() {
	actorID, _ := c.GetInt(":id")
	actor, err := models.GetActorById(int64(actorID))
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}

	if err := c.ParseForm(actor); err != nil {
		c.Data["json"] = err.Error()
	} else {
		if err := models.UpdateActorById(actor); err == nil {
			c.Data["json"] = "Acteur mis à jour avec succès"
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Actor
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ActorController) Delete() {
	actorID, _ := c.GetInt(":id")
	if err := models.DeleteActor(int64(actorID)); err == nil {
		c.Data["json"] = "Acteur supprimé avec succès"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
