package routes

import "open-api-backend/api"

func addUserRouter() {
	userApi := api.RootApi.UserApi

	userGroup := apiGroup.Group("user")
	{
		userGroup.GET("info/all")
		userGroup.POST("info/list", userApi.GetUserList)
		userGroup.GET("info/detail", userApi.GetUserById)
		userGroup.POST("info/create", userApi.CreateUser)
		userGroup.POST("info/update")
		userGroup.POST("info/update/passwd")
		userGroup.POST("info/delete")
	}
}
