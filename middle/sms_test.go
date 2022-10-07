package middle

import (
	"os"
	"testing"
)

func TestSendSms(t *testing.T) {
	type args struct {
		secretId  string
		secretKey string
		smsAppId  string
		signName  string
		tplId     string
		params    []string
		phoneList []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test sms ",
			args{
				secretId:  os.Getenv("SECRET_ID"),
				secretKey: os.Getenv("SECRET_KEY"),
				smsAppId:  "1400733090",
				signName:  "我的程序员之路个人网",
				tplId:     "1564505",
				params:    []string{"123123", "20"},
				phoneList: []string{"17666116392"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendSms(tt.args.secretId, tt.args.secretKey, tt.args.smsAppId, tt.args.signName, tt.args.tplId, tt.args.params, tt.args.phoneList)
		})
	}
}
