package api

import (
	"open-api-backend/api/interfaceApi"
	"open-api-backend/api/userApi"
)

type Api struct {
	UserApi      userApi.UserApi
	InterfaceApi interfaceApi.InterfaceApi
}

var RootApi = new(Api)
