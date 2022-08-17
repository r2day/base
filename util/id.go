package util

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func getId(prefix string) string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return "-"
	}

	// Generate a snowflake ID.
	id := node.Generate()
	return fmt.Sprintf("%s%d", prefix, id)
}

// GetBrandId 获取品牌id
func GetBrandId() string {
	return getId("b")
}

// GetCategoryId 获取分组
func GetCategoryId() string {
	return getId("c")
}

// GetDishesId 获取菜品id
func GetDishesId() string {
	return getId("d")
}

// GetDepartmentId 部门id
func GetDepartmentId() string {
	return getId("dp")
}

// GetUnitId 规格id
func GetUnitId() string {
	return getId("u")
}

// GetAccountId 账号
func GetAccountId() string {
	return getId("AC")
}

// GetStoreId 账号
func GetStoreId() string {
	return getId("STO")
}

// GetItemId 获取物品id
func GetItemId() string {
	return getId("RI")
}

// GetCartItemId 获取物品项目id
func GetCartItemId() string {
	return getId("RCI")
}

// GetCartItemId 获取购物车id
func GetCartId() string {
	return getId("RC")
}
