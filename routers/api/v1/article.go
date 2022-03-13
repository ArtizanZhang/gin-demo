package v1

import (
	"github.com/ArtizanZhang/gin-demo/models"
	"github.com/ArtizanZhang/gin-demo/pkg/e"
	"github.com/ArtizanZhang/gin-demo/pkg/setting"
	"github.com/ArtizanZhang/gin-demo/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

func GetArticles(c *gin.Context) {
	maps := make(map[string]interface{})
	code := e.ERROR
	total := 0
	var lists []models.Article

	valid := validation.Validation{}

	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0和1")
		maps["state"] = state
	}

	if arg := c.Query("tag_id"); arg != "" {
		tagId := com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id").Message("tag_id必须大于1")
		maps["tag_id"] = tagId
	}

	if valid.HasErrors() {
		code = e.INVALID_PARAMS
	} else {
		total = models.GetArticleTotal(maps)
		if total > 0 {
			lists = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		}
		code = e.SUCCESS
	}

	data := map[string]interface{}{
		"total": total,
		"lists": lists,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

func GetArticle(c *gin.Context) {

	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("id最小为1")

	code := e.SUCCESS

	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": nil,
		})
		return
	}

	var article interface{}
	if arg := models.GetArticle(id); arg.ID == 0 {
		code = e.ERROR_NOT_EXIST_ARTICLE
		article = nil
	} else {
		code = e.SUCCESS
		article = arg
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": article,
	})
}

func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.Query("state")).MustInt() // 必须转Int

	valid := validation.Validation{}

	valid.Required(tagId, "tag_id").Message("tag_id 必填")
	valid.Min(tagId, 1, "tag_id").Message("tag_id 必须大于1")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.ERROR

	for {
		if valid.HasErrors() {
			code = e.INVALID_PARAMS
			for _, err := range valid.Errors {
				log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
			}
			break
		}

		if !models.ExistTagByID(tagId) {
			code = e.ERROR_NOT_EXIST_TAG
			break
		}

		code = e.SUCCESS
		models.AddArticle(map[string]interface{}{
			"tag_id":     tagId,
			"title":      title,
			"desc":       desc,
			"content":    content,
			"created_by": createdBy,
			"state":      state,
		})
		break
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  nil,
		"valid": valid.Errors,
	})

}

func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Required(id, "id").Message("id 必填")
	tagId := -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id").Message("tag_id 必须大于1")
	}

	code := e.ERROR
	article := models.GetArticle(id)
	if article.ID == 0 {
		code = e.ERROR_NOT_EXIST_ARTICLE
		c.JSON(http.StatusOK, gin.H{
			"code":  code,
			"msg":   e.GetMsg(code),
			"data":  nil,
			"valid": valid.Errors,
		})
		return

	}

	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code":  code,
			"msg":   e.GetMsg(code),
			"data":  nil,
			"valid": valid.Errors,
		})
		return
	}

	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.Query("state")).MustInt()

	data := make(map[string]interface{})
	if tagId > 0 {
		data["tag_id"] = tagId
	}

	if title != "" {
		data["title"] = title
	}

	if desc != "" {
		data["desc"] = desc
	}

	if content != "" {
		data["content"] = content
	}

	if createdBy != "" {
		data["createdBy"] = createdBy
	}
	if state != 0 {
		data["state"] = state
	}

	models.EditArticle(id, data)
	code = e.SUCCESS

	var res map[string]interface{}
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  make(map[string]string),
		"data2": res,
	})
}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Required(id, "id").Message("id 必填")

	code := e.ERROR
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code":  code,
			"msg":   e.GetMsg(code),
			"data":  nil,
			"valid": valid.Errors,
		})
		return
	}

	article := models.GetArticle(id)
	if article.ID == 0 {
		code = e.ERROR_NOT_EXIST_ARTICLE
		c.JSON(http.StatusOK, gin.H{
			"code":  code,
			"msg":   e.GetMsg(code),
			"data":  nil,
			"valid": valid.Errors,
		})
		return

	}

	models.DeleteTag(id)
	code = e.SUCCESS

	var res map[string]interface{}
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  make(map[string]string),
		"data2": res,
	})

}
