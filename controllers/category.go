package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/sinksmell/lanblog/models"
	"strconv"
	"strings"
)

// CategoryController operations for Category
type CategoryController struct {
	beego.Controller
}

// URLMapping ...
func (c *CategoryController) URLMapping() {
	c.Mapping("Create", c.Create)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Update", c.Update)
	c.Mapping("Delete", c.Delete)
}

// Create ...
// @Title Create
// @Description create Category
// @Param	body		body 	models.Category	true		"body for Category content"
// @Success 201 {int} models.Category
// @Failure 403 body is empty
// @router /create [post]
func (c *CategoryController) Create() {
	result := models.NewCommResult()
	var v models.Category
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddCategory(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			result.Msg = "OK"
			//c.Data["json"] = v
		} else {
			result.Msg = err.Error()
		}
	} else {
		result.Msg = err.Error()
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Category by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Category
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CategoryController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCategoryById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Category
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Category
// @Failure 403
// @router /list [get]
func (c *CategoryController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllCategory(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Update ...
// @Title Update
// @Description update the Category
// @Param	body		body 	models.Category	true		"body for Category content"
// @Success 201 {int} models.Category
// @Failure 403 body is empty
// @router /update [post]
func (c *CategoryController) Update() {
	result := models.NewCommResult()
	v := models.Category{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateCategoryById(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			result.Msg = "OK"
		} else {
			result.Msg = err.Error()
		}
	} else {
		result.Msg = err.Error()
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// Remove ...
// @Title Remove
// @Description delete the Category
// @Param	body		body 	models.Category	true		"body for Category content"
// @Success 201 {int} models.Category
// @Failure 403 body is empty
// @router /delete [post]
func (c *CategoryController) Delete() {

	result := models.NewCommResult()
	v := models.Category{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.DeleteCategory(v.Id); err == nil {
			c.Ctx.Output.SetStatus(201)
			result.Msg = "OK"
		} else {
			result.Msg = err.Error()
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}
