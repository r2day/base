package resp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RenderList 列表展示
func RenderList(c *gin.Context, rangeTpl string, counter int64, obj any) {
	c.Header("Content-Range", rangeTpl)
	c.Header("X-Total-Count", fmt.Sprintf("%d", counter))
	c.JSON(http.StatusOK, obj)
}
