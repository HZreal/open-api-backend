package routes

import "open-api-backend/api"

func addInterfaceRouter() {
	interfaceApi := api.RootApi.InterfaceApi
	userGroup := apiGroup.Group("interfaceInfo")
	{
		userGroup.GET("info/all")
		userGroup.GET("info/list", interfaceApi.GetInterfaceList)
		userGroup.GET("info/detail", interfaceApi.GetInterfaceById)
		userGroup.POST("info/update")
		userGroup.POST("info/update/passwd")
		userGroup.POST("info/delete")
	}
}
