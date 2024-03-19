package userApi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"open-api-backend/model"
)

func (UserApi) GetUserById(c *gin.Context) {
	id := c.Query("id")
	// 校验

	//
	var user model.UserInfo
	affected := db.Take(&model.User{}, id).Scan(&user).RowsAffected
	if affected == 0 {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	c.JSON(http.StatusOK, user)
}
