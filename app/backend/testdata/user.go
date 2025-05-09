package testdata

import (
	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/service"
)

var UserTestData = []*ent.User{
	&ent.User{
		// 登録済み1
		ID: 1, Name: "alice", Email: "alice@example.com", Password: "$2a$10$IUjSMm7z8i6QaF5BfOc7wOKRkQqdDZ4TkmzutyAOe42vwteaKiqsO",
	},
	&ent.User{
		// 登録予定
		ID: 10001, Name: "carol", Email: "carol@example.com", Password: "xxxxxx",
	},
	&ent.User{
		// 未登録
		ID: 999, Name: "david", Email: "david@example.com", Password: "xxxxxx",
	},
	&ent.User{
		// 更新用
		ID: 1, Name: "alice", Email: "alice_alice@example.co.jp", Password: "$2a$10$IUjSMm7z8i6QaF5BfOc7wOKRkQqdDZ4TkmzutyAOe42vwteaKiqsO",
	},
	&ent.User{
		// 更新用(一意制約違反)
		ID: 1, Name: "alice", Email: "bob@example.com", Password: "$2a$10$IUjSMm7z8i6QaF5BfOc7wOKRkQqdDZ4TkmzutyAOe42vwteaKiqsO",
	},
}

var UserReqTestData = []*service.UserRequest{
	&service.UserRequest{
		// 登録済み
		Name: "alice", Email: "alice@example.com", Password: "alicepassword",
	},
	&service.UserRequest{
		// 登録予定
		Name: "carol", Email: "carol@example.com", Password: "xxxxxx",
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
		ID: 10001, Name: "carol",
	},
}
