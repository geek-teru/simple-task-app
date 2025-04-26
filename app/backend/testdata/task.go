package testdata

import (
	"time"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/service"
)

var t = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

var TaskTestData = []*ent.Task{
	&ent.Task{
		// 登録済み 1
		ID: 1, Title: "task01", Description: "task01 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
	&ent.Task{
		// 未登録
		ID: 10001, Title: "task99", Description: "task99 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
	&ent.Task{
		// 必須項目なし
		ID: 1, Title: "", Description: "", Status: "", DueDate: &t, UserID: 1,
	},
	&ent.Task{
		// 更新後データ
		ID: 1, Title: "task01 updated", Description: "task1 description updated", Status: "TODO", DueDate: &t, UserID: 1,
	},
	&ent.Task{
		// 登録済み 2
		ID: 2, Title: "task02", Description: "task02 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
	&ent.Task{
		// 登録済み 3
		ID: 3, Title: "task03", Description: "task03 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
	&ent.Task{
		// 登録済み 4
		ID: 4, Title: "task04", Description: "task04 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
}

var TaskReqTestData = []*service.TaskRequest{
	&service.TaskRequest{
		// 登録済み
		Title: "task01", Description: "task01 description", Status: "TODO", DueDate: &t,
	},
	&service.TaskRequest{
		// 未登録
		Title: "task99", Description: "task99 description", Status: "TODO", DueDate: &t,
	},
	&service.TaskRequest{
		// 必須項目なし
		Title: "", Description: "task01 description", Status: "TODO", DueDate: &t,
	},
	&service.TaskRequest{
		// 更新後データ
		Title: "task01 updated", Description: "task1 description updated", Status: "TODO", DueDate: &t,
	},
}

var TaskResTestData = []service.TaskResponse{
	service.TaskResponse{
		// 登録済み
		ID: 1, Title: "task01", Description: "task01 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
	service.TaskResponse{
		// 未登録
		ID: 10001, Title: "task99", Description: "task99 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
	service.TaskResponse{
		// 必須項目なし
		ID: 1, Title: "", Description: "", Status: "", DueDate: &t, UserID: 1,
	},
	service.TaskResponse{
		// 更新後データ
		ID: 1, Title: "task01 updated", Description: "task1 description updated", Status: "TODO", DueDate: &t, UserID: 1,
	},
	service.TaskResponse{
		// 登録済み 2
		ID: 2, Title: "task02", Description: "task02 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
	service.TaskResponse{
		// 登録済み 3
		ID: 3, Title: "task03", Description: "task03 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
	service.TaskResponse{
		// 登録済み 4
		ID: 4, Title: "task04", Description: "task04 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
}
