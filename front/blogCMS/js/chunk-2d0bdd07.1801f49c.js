(window.webpackJsonp=window.webpackJsonp||[]).push([["chunk-2d0bdd07"],{"2e31":function(t,e,l){"use strict";l.r(e);var a=l("9bd2"),o={data:function(){return{topicId:1,topic:{id:"1",title:"topic demo",summary:"上海市普陀区金沙江路 1518 弄",cate:{name:"test"},labels:[{id:1,name:"test01"},{id:2,name:"test02"}],author:"sinksmell",views:2,content:"## default text"},newContent:"",cates:[],labels:[]}},methods:{handleChange:function(){var e=this;a.a.get("/topic/"+this.topicId).then(function(t){e.topic=t,e.newContent="this.topic.content"})},updateTopic:function(){var e=this;a.a.post("/topic/update",this.topic).then(function(t){"OK"===t.msg&&e.$message({message:"修改成功",tags:"success"})})}},mounted:function(){var e=this;a.a.get("/category/list").then(function(t){e.cates=t}),a.a.get("/label/list").then(function(t){e.labels=t})}},c=l("2877"),i=Object(c.a)(o,function(){var l=this,t=l.$createElement,a=l._self._c||t;return a("d2-container",[a("template",{slot:"header"},[a("el-form",{attrs:{inline:!0}},[a("el-form-item",[a("el-button",{attrs:{type:"success"}},[l._v("\n                    选择文章序号\n                ")])],1),a("el-form-item",[a("el-input-number",{attrs:{min:1,max:1<<26,label:"描述文字"},on:{change:l.handleChange},model:{value:l.topicId,callback:function(t){l.topicId=t},expression:"topicId"}})],1)],1)],1),a("el-card",[a("el-form",{attrs:{model:l.topic,"label-width":"80px"}},[a("el-form-item",{attrs:{label:"文章标题"}},[a("el-col",{attrs:{span:10}},[a("el-input",{model:{value:l.topic.title,callback:function(t){l.$set(l.topic,"title",t)},expression:"topic.title"}})],1)],1),a("el-form-item",{attrs:{label:"文章作者"}},[a("el-col",{attrs:{span:6}},[a("el-input",{model:{value:l.topic.author,callback:function(t){l.$set(l.topic,"author",t)},expression:"topic.author"}})],1)],1),a("el-form-item",{attrs:{label:"当前分类"}},[a("el-tag",{attrs:{type:"error"}},[l._v(l._s(l.topic.cate.name))])],1),a("el-form-item",{attrs:{label:"文章分类"}},[a("el-select",{attrs:{placeholder:"请选择文章分类"},model:{value:l.topic.cate,callback:function(t){l.$set(l.topic,"cate",t)},expression:"topic.cate"}},l._l(l.cates,function(t,e){return a("div",{key:e},[a("el-option",{attrs:{label:t.name,value:t}},[l._v(l._s(t.name)+"\n                        ")])],1)}),0)],1),a("el-form-item",{attrs:{label:"当前标签"}},[a("el-row",[a("el-col",[l._l(l.topic.labels,function(t){return a("el-tag",{key:t.id,attrs:{type:"success"}},[l._v(l._s(t.name))])}),l._v("  \n                    ")],2)],1)],1),a("el-form-item",{attrs:{label:"文章标签"}},[a("el-checkbox-group",{model:{value:l.topic.labels,callback:function(t){l.$set(l.topic,"labels",t)},expression:"topic.labels"}},l._l(l.labels,function(t,e){return a("el-checkbox-button",{key:e,attrs:{size:"medium",label:t}},[l._v(l._s(t.name))])}),1)],1),a("el-form-item",{attrs:{label:"图片资源"}},[a("el-col",{attrs:{span:12}},[a("el-input",{model:{value:l.topic.url,callback:function(t){l.$set(l.topic,"url",t)},expression:"topic.url"}})],1)],1),a("el-form-item",[a("el-card",[a("el-row",{attrs:{type:"flex",justify:"center"}},[a("img",{staticStyle:{width:"945px",height:"345px"},attrs:{src:l.topic.url}})])],1)],1),a("el-form-item",{attrs:{label:"文章摘要"}},[a("el-col",{attrs:{span:12}},[a("el-input",{attrs:{type:"textarea"},model:{value:l.topic.summary,callback:function(t){l.$set(l.topic,"summary",t)},expression:"topic.summary"}})],1)],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:l.updateTopic}},[l._v("确认")]),a("el-button",[l._v("取消")])],1)],1)],1),a("br"),a("el-card",{staticClass:"d2-card",staticStyle:{"padding-top":"50px"}},[a("d2-mde",{staticClass:"mde",model:{value:l.topic.content,callback:function(t){l.$set(l.topic,"content",t)},expression:"topic.content"}})],1),a("br"),a("el-card",{staticClass:"d2-card",staticStyle:{"padding-top":"50px"},attrs:{shadow:"never"}},[a("pre",[l._v(l._s(l.topic.content))])])],2)},[],!1,null,"70a2bb18",null);i.options.__file="index.vue";e.default=i.exports}}]);