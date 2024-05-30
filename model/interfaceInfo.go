package model

import "gorm.io/gorm"

type InterfaceInfo struct {
	gorm.Model

	UserId         uint   `json:"user_id" gorm:"comment:创建人的ID;not null"`
	Name           string `json:"name" gorm:"comment:接口名称;size:256;not null"`
	Description    string `json:"description" gorm:"comment:接口描述;size:256"`
	Url            string `json:"url" gorm:"comment:接口地址;size:512;not null"`
	Method         string `json:"method" gorm:"comment:请求类型;not null"`
	RequestHeader  string `json:"request_header" gorm:"comment:请求头"`
	ResponseHeader string `json:"response_header" gorm:"comment:响应头"`
	RequestParams  string `json:"request_params" gorm:"comment:请求参数"`
	Status         bool   `json:"status" gorm:"comment:接口状态：0关闭，1开启;not null"`
}

func (receiver InterfaceInfo) TableName() string {
	return "interface_info"
}
