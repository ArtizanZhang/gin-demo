package v1

import "C"
import (
	"fmt"
	"github.com/ArtizanZhang/gin-demo/models"
	"github.com/ArtizanZhang/gin-demo/pkg/e"
	"github.com/ArtizanZhang/gin-demo/pkg/setting"
	"github.com/ArtizanZhang/gin-demo/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	state := c.Query("state")

	maps := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	if state != "" {
		maps["state"] = com.StrTo(state).MustInt()
	}

	data := models.GetTags(util.GetPage(c), setting.PageSize, maps)
	fmt.Println(data)
	total := models.GetTagTotal(maps)

	res := map[string]interface{}{
		"lists": data,
		"total": total,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": res,
	})
}

// AddTag 新增
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.Query("state")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名字不能为空")
	valid.MaxSize(name, 100, "name").Message("名字不能超过100个字符")
	valid.Required(createdBy, "createdBy").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "createdBy").Message("创建人不能超过100个字符")
	valid.Range(state, 0, 1, "state").Message("状态只能是0和1")

	var code int

	switch {
	case valid.HasErrors():
		code = e.INVALID_PARAMS
	case models.ExitTagName(name):
		code = e.ERROR_EXIST_TAG
	default:
		code = e.SUCCESS
		models.AddTag(name, state, createdBy)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditTag 修改
func EditTag(c *gin.Context) {

}

// DeleteTag 删除
func DeleteTag(c *gin.Context) {

}
