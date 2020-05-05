package service

import (
	"errors"
	"server-monitor-admin/global"
	"server-monitor-admin/model"
	"server-monitor-admin/utils"
)

// @title    注册
// @description   
// @auth    smalljop  2020/4/27 11:27
// @param     
// @return

func Register(user model.SysUser) (model.SysUser, error) {
	//先查询用户是否存在
	notRegister := global.DB.Where("email=?", user.Email).First(user).RecordNotFound()
	if !notRegister {
		return user, errors.New("邮箱已被注册")
	} else {
		//创建用户
		user.Password = utils.MD5V([]byte(user.Password))
		global.DB.Create(&user)
	}
	return user, nil
}

// @title    登录
// @description
// @auth    smalljop  2020/4/27 11:27
// @param
// @return

func Login(user model.SysUser) (model.SysUser, error) {
	//先查询用户是否存在
	user.Password = utils.MD5V([]byte(user.Password))
	err := global.DB.Where("user_name=? AND password=?", user.UserName, user.Password).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// @title    获取用户信息
// @description
// @auth    smalljop  2020/4/27 11:27
// @param
// @return

func GetUserInfo(userId int64) (model.SysUser, error) {
	var user model.SysUser
	err := global.DB.Where("id=?", userId).First(&user).Error
	return user, err
}
