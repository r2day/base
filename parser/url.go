package parser

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	defaultOffset = 0
	defaultLimit  = 24
)

const (
	OnePage = 1
)

// QueryRequest Binding from JSON
type QueryRequest struct {
	// FilterRequest 过滤器
	Filter []FilterRequest
}

// FilterRequest Binding from JSON
type FilterRequest struct {
	// id 唯一帐户id
	Id []uint `form:"id" json:"id" xml:"id"`

	// Status 状态
	Status string `form:"status" json:"status" xml:"status"`

	// CategoryId 分类id
	CategoryId uint `form:"category_id" json:"category_id" xml:"category_id"`

	// ProductId 商品id
	ProductId uint `form:"product_id" json:"product_id" xml:"product_id"`
}

// UrlParams 将url的参数统一进行解析
type UrlParams struct {
	Limit     int  `form:"limit" json:"limit" xml:"limit"`
	Offset    int  `form:"offset" json:"offset" xml:"offset"`
	HasFilter bool `form:"has_filter" json:"has_filter" xml:"has_filter"`
	Filter    FilterRequest
}

func ParserParams(c *gin.Context) UrlParams {
	params := UrlParams{}
	filter, ok := c.GetQueryArray("filter")
	rangeValue, ok := c.GetQueryArray("range")

	params.Limit = defaultLimit
	params.Offset = defaultOffset

	if len(rangeValue) == 1 {
		println("rangeValue-->", rangeValue[0], ok)
		rangeObj := make([]int, 2)
		err := json.Unmarshal([]byte(rangeValue[0]), &rangeObj)

		if err != nil {
			fmt.Println("json.Unmarshal failed-->", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return params
		}
		params.Offset = rangeObj[0]

		if rangeObj[1] > 0 {
			params.Limit = rangeObj[1]
		}
	}

	if ok && len(filter) != 0 {

		// 将过滤器中的所有参数都解析出来供
		// 业务查询进行使用
		filterInstance := FilterRequest{}

		err := json.Unmarshal([]byte(filter[0]), &filterInstance)
		if err != nil {
			// 如果是空的会解析失败
			return params
		}

		// 检查如果所有过滤字段都没有被解析到那么
		// 直接返回
		if filterInstance.Status == "" &&
			filterInstance.CategoryId == 0 &&
			filterInstance.ProductId == 0 {
			return params
		}
		params.HasFilter = true
		params.Filter = filterInstance
	}
	return params
}
