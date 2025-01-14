package testdata

import (
	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/service"
)

var UserTestData = []*ent.User{
	&ent.User{
		// 登録済み
		ID: 1, Name: "alice", Email: "alice@example.com", Password: "$2a$10$IUjSMm7z8i6QaF5BfOc7wOKRkQqdDZ4TkmzutyAOe42vwteaKiqsO",
	},
	&ent.User{
		// 未登録
		ID: 10001, Name: "bob", Email: "bob@example.com", Password: "$2a$10$ExzssGX4xS4joeZx7aO9SOpWXLBzhAQxjMBleRxf8ziC961FkJ7qq",
	},
}

var UserReqTestData = []*service.UserRequest{
	&service.UserRequest{
		// 登録済み
		Name: "alice", Email: "alice@example.com", Password: "alicepassword",
	},
	&service.UserRequest{
		// 未登録
		Name: "bob", Email: "bob@example.com", Password: "bobpassword",
	},
	&service.UserRequest{
		// パスワード間違い
		Name: "alice", Email: "alice@example.com", Password: "",
	},
}

var UserResTestData = []service.UserResponse{
	service.UserResponse{
		// 登録済み
		ID: 1, Name: "alice",
	},
	service.UserResponse{
		// 未登録
		ID: 10001, Name: "bob",
	},
}
