package userApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"open-api-backend/common"
	"open-api-backend/model"
	dto "open-api-backend/model/dto"
	"open-api-backend/model/vo"
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

func (UserApi) GetUserList(c *gin.Context) {
	var query dto.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, common.Failed(common.ParamsError))
		return
	}

	var users []model.User
	var total int64
	offset := (query.Page - 1) * query.PageSize

	// 这里用GORM查询数据库，获取数据总数和分页数据
	db.Model(&model.User{}).Count(&total) // 假设db是*gorm.DB的实例
	db.Offset(offset).Limit(query.PageSize).Find(&users)

	// 计算总页数
	pages := int(total) / query.PageSize
	if int(total)%query.PageSize != 0 {
		pages++
	}

	// 返回分页数据
	c.JSON(http.StatusOK, common.SuccessWithData(vo.PaginationResult{
		Total:       int(total),
		Pages:       pages,
		CurrentPage: query.Page,
		PageSize:    query.PageSize,
		Records:     users,
	}))
}
