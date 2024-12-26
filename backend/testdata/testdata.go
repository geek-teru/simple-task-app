package testdata

import (
	"github.com/geek-teru/simple-task-app/ent"
)

var UserTestData = []*ent.User{
	&ent.User{
		// 登録済み
		ID: 1, Name: "user_a", Email: "user_a@example.com", Password: "password",
	},
	&ent.User{
		// 未登録
		ID: 99, Name: "user_x", Email: "user_x@example.com", Password: "password",
	},
}
