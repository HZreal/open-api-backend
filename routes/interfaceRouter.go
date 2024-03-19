package routes

func addInterfaceRouter() {
	userGroup := apiGroup.Group("interface")
	{
		userGroup.GET("info/all")
		userGroup.GET("info/list")
		userGroup.GET("info/detail")
		userGroup.POST("info/update")
		userGroup.POST("info/update/passwd")
		userGroup.POST("info/delete")
	}
}
