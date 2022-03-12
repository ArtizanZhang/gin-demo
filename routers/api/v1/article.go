package v1

import (
	"github.com/ArtizanZhang/gin-demo/models"
	"github.com/ArtizanZhang/gin-demo/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

func GetArticle(c *gin.Context) {

}

func GetArticles(c *gin.Context) {

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

}

func DeleteArticle(c *gin.Context) {

}
