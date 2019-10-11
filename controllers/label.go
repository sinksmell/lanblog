package controllers

import (
	"encoding/json"
	"errors"
	"github.com/sinksmell/lanblog/models"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
)

type labelResult struct {
	Success bool `json:"success"`
}

// LabelController operations for Label
type LabelController struct {
	beego.Controller
}

// URLMapping ...
func (c *LabelController) URLMapping() {
	c.Mapping("Create", c.Create)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Update", c.Update)
	c.Mapping("Delete", c.Delete)
}

// Create ...
// @Title Create
// @Description create Label
// @Param	body		body 	models.Label	true		"body for Label content"
// @Success 201 {int} models.Label
// @Failure 403 body is empty
// @router /create [post]
func (c *LabelController) Create() {
	var v models.Label
	result := models.NewCommResult()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddLabel(&v); err == nil {
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

// GetOne ...
// @Title Get One
// @Description get Label by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Label
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LabelController) GetOne() {
	//result := models.NewCommResult()
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetLabelById(id)
	if err != nil {
		//result.Msg=err.Error()
		c.Data["json"] = err.Error()
	} else {
		//result.Msg="OK"
		//result := &labelResult{}
		//result.Success=true
		c.Data["json"] = v
	}
	//	c.Data["json"]=result
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Label
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Label
// @Failure 403
// @router /list [get]
func (c *LabelController) GetAll() {
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

	l, err := models.GetAllLabel(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Update ...
// @Title Update
// @Description update the Label
// @Param	body		body 	models.Label	true		"body for Label content"
// @Success 201 {int} models.Label
// @Failure 403 body is empty
// @router /update [post]
func (c *LabelController) Update() {
	var v models.Label
	result := models.NewCommResult()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateLabelById(&v); err == nil {
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
// @Description delete the Label
// @Param	body		body 	models.Label	true		"body for Label content"
// @Success 201 {int} models.Label
// @Failure 403 body is empty
// @router /delete [post]
func (c *LabelController) Delete() {
	var v models.Label
	result := models.NewCommResult()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if err := models.DeleteLabel(v.Id); err == nil {
			c.Ctx.Output.SetStatus(201)
			result.Msg = "OK"
		} else {
			result.Msg = err.Error()
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}
