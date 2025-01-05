package testdata

import (
	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/service"
)

var UserTestData = []*ent.User{
	&ent.User{
		// 登録済み
		ID: 1, Name: "user_a", Email: "user_a@example.com", Password: "passworda",
	},
	&ent.User{
		// 未登録
		ID: 10001, Name: "user_x", Email: "user_x@example.com", Password: "passwordx",
	},
}

var UserReqTestData = []*service.UserRequest{
	&service.UserRequest{
		// 登録済み
		Name: "user_a", Email: "user_a@example.com", Password: "passworda",
	},
	&service.UserRequest{
		// 未登録
		Name: "user_x", Email: "user_x@example.com", Password: "passwordx",
	},
}
