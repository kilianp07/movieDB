package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/kilianp07/movieDB/models"
)

// FilmController operations for Film
type FilmController struct {
	beego.Controller
}

// URLMapping ...
func (c *FilmController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Film
// @Param	body		body 	models.Film	true		"body for Film content"
// @Success 201 {object} models.Film
// @Failure 403 body is empty
// @router / [post]
func (c *FilmController) Post() {
	film := models.Film{}
	if err := c.ParseForm(&film); err != nil {
		c.Data["json"] = err.Error()
	} else {
		if _, err := models.AddFilm(&film); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = "Created Film"
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Film by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Film
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FilmController) GetOne() {
	filmID, _ := c.GetInt(":id")
	film, err := models.GetFilmById(int64(filmID))
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = film
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Film
// @Success 200 {object} models.Film
// @Failure 403
// @router / [get]
func (c *FilmController) GetAll() {
	filters := make(map[string]string)
	fields := make([]string, 0)
	sortby := make([]string, 0)
	order := make([]string, 0)
	limit := int64(10)
	offset := int64(0)

	films, err := models.GetAllFilm(filters, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = films
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Film
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Film	true		"body for Film content"
// @Success 200 {object} models.Film
// @Failure 403 :id is not int
// @router /:id [put]
func (c *FilmController) Put() {
	filmID, _ := c.GetInt(":id")
	film, err := models.GetFilmById(int64(filmID))
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}

	if err := c.ParseForm(film); err != nil {
		c.Data["json"] = err.Error()
	} else {
		if err := models.UpdateFilmById(film); err == nil {
			c.Data["json"] = "Film updated"
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Film
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FilmController) Delete() {
	filmID, _ := c.GetInt(":id")
	if err := models.DeleteFilm(int64(filmID)); err == nil {
		c.Data["json"] = "Film deleted"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
