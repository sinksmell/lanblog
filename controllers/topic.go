package controllers

import (
	"encoding/json"
	"errors"
	"github.com/sinksmell/lanblog/models"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
)

// TopicController operations for Topic
type TopicController struct {
	beego.Controller
}

// URLMapping ...
func (c *TopicController) URLMapping() {
	c.Mapping("Create", c.Create)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Update", c.Update)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetByCate", c.GetByCate)
	c.Mapping("GetByLabel", c.GetByLabel)
	c.Mapping("GetByVagueName", c.GetByVagueName)
}

// Login ...
// @Title Login
// @Description create Topic
// @Param	body		body 	models.Topic	true		"body for Topic content"
// @Success 201 {int} models.Topic
// @Failure 403 body is empty
// @router /create [post]
func (c *TopicController) Create() {
	var v models.Topic
	result := models.NewCommResult()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddTopic(&v); err == nil {
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
// @Description get Topic by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Topic
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TopicController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTopicByID(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Topic
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Topic
// @Failure 403
// @router /list [get]
func (c *TopicController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 1 << 16
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

	l, err := models.GetAllTopics(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Topic
// @Param	body		body 	models.Topic	true		"body for Topic content"
// @Success 201 {int} models.Topic
// @Failure 403 body is empty
// @router /update [post]
func (c *TopicController) Update() {
	result := models.NewCommResult()
	v := models.Topic{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTopicByID(&v); err == nil {
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

// Delete ...
// @Title Delete
// @Description delete the Topic
// @Param	body		body 	models.Topic	true		"body for Topic content"
// @Success 201 {int} models.Topic
// @Failure 403 body is empty
// @router /delete [post]
func (c *TopicController) Delete() {
	var v models.Topic
	result := models.NewCommResult()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//beego.BeeLogger.Info("%+v\n", v)
		if err := models.DeleteTopic(v.Id); err == nil {
			c.Ctx.Output.SetStatus(201)
			result.Msg = "OK"
		} else {
			result.Msg = err.Error()
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// GetByCate ...
// @Title GetByCategory
// @Description get Topics by Category id
// @Param	id		path 	string	true		"category id"
// @Success 200 {object} models.TopicsResult
// @Failure 403 :id is empty
// @router /cate/:id [get]
func (c *TopicController) GetByCate() {
	result := models.TopicsResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	//传入的 category id 是有效值的话才能查到正确结果
	topics, err := models.GetTopicsByCateID(id)
	if err != nil {
		result.Msg = err.Error()
	} else {
		result.Msg = "OK"
		result.Topics = topics
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// GetByLabel ...
// @Title GetByLabel
// @Description get Topics by Label id
// @Param	id		path 	string	true		"Label id"
// @Success 200 {object} models.TopicsResult
// @Failure 403 :id is empty
// @router /label/:id [get]
func (c *TopicController) GetByLabel() {
	result := models.TopicsResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	//传入的 label id 是有效值的话才能查到正确结果
	topics, err := models.GetTopicsByLabelID(id)
	if err != nil {
		result.Msg = err.Error()
	} else {
		result.Msg = "OK"
		result.Topics = topics
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// GetByVagueName
// @Title GetByVagueName
// @Description get Topics by Vague Name
// @Param	vname		path 	string	true		"Vague Name"
// @Success 200 {object} models.TopicsResult
// @Failure 403 :vname is empty
// @router /find/:vname [get]
func (c *TopicController) GetByVagueName() {
	result := models.TopicsResult{}
	keyWord := c.Ctx.Input.Param(":vname")
	//传入的关键词作为模糊查找关键词
	topics, err := models.FastFind(keyWord)
	if err != nil {
		result.Msg = err.Error()
	} else {
		result.Msg = "OK"
		result.Topics = topics
	}
	c.Data["json"] = result
	c.ServeJSON()
}
