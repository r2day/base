package util

import (
	"crypto/md5"
	"encoding/hex"
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

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// ConvertToToken 转换为md5格式
func ConvertToToken(k string) string {
	token := getMD5Hash(k)
	return token
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

// GetOrderItemId 获取购物车id
func GetOrderItemId() string {
	return getId("OI")
}

// GetOrderId 获取购物车id
func GetOrderId() string {
	return getId("HLJ")
}

// TransactionId 交易id
func TransactionId() string {
	return getId("TX")
}

// SmsCode 短信验证码
func SmsCode() string {
	return "R-" + getId("")[13:]
}

// MerchantId 交易id
func MerchantId() string {
	return getId("M")
}

// MerchantKey 交易id
func MerchantKey() string {
	return ConvertToToken(getId("MX"))
}

// SessionId 通用session
func SessionId() string {
	return ConvertToToken(getId("SESSION"))
}

// ApplyId 申请回执
func ApplyId() string {
	return getId("")
}
