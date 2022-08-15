package convert

import (
	"fmt"
	"v1_gorm_direct/graph/model"
	"v1_gorm_direct/model/tablemodel"
)

func User(userModel tablemodel.User) *model.User {
	var user model.User
	user.ID = fmt.Sprint(userModel.ID)
	user.Name = userModel.Name
	user.Email = userModel.Email
	user.Password = userModel.Password
	return &user
}

func Users(t []*tablemodel.User) []*model.User {
	res := make([]*model.User, 0, len(t))
	for _, v := range t {
		res = append(res, User(*v))
	}

	return res
}
