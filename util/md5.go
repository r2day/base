package util

import (
	"crypto/md5"
	"fmt"
)

const (
	signTpl = "%s_%s"
)

func SignWithMd5(payload []byte, secretKey string) string {
	data := fmt.Sprintf(signTpl, string(payload), secretKey)
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func VerifyPayload(payload []byte, secretKey string, hash string) (bool, error) {
	if SignWithMd5(payload, secretKey) == hash {
		return true, nil
	}
	return false, nil
}
