package userApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"open-api-backend/common"
	"open-api-backend/model"
	dto "open-api-backend/model/dto"
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

func (UserApi) CreateUser(c *gin.Context) {
	var userCreate dto.UserCreateDTO
	err := c.ShouldBindJSON(&userCreate)
	if err != nil {
		fmt.Println("params error")
		c.JSON(http.StatusForbidden, common.Failed(common.ParamsError))
		return
	}
	user := &model.User{
		Model:        gorm.Model{},
		UserAccount:  userCreate.Username,
		UserName:     nil,
		AvatarUrl:    nil,
		Gender:       1,
		UserPassword: userCreate.Password,
		Phone:        nil,
		Email:        nil,
		UserStatus:   0,
		Role:         0,
		AccessKey:    "",
		SecretKey:    "",
	}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println("result.Error  ---->  ", result.Error)
		c.JSON(http.StatusOK, gin.H{"code": 99999, "msg": "create failed", "data": nil})
		return
	}
	fmt.Println("result  ---->  ", result)
	fmt.Println("result.RowsAffected  ---->  ", result.RowsAffected)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": user})
}
