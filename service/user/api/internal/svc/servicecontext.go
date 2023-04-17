package svc

import (
	"tourBooking/service/user/api/internal/config"
	"tourBooking/service/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(sqlx.NewMysql(c.DataSource)),
	}
}
