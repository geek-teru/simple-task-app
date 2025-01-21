package testdata

import (
	"time"

	"github.com/geek-teru/simple-task-app/ent"
)

var t = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

var TaskTestData = []*ent.Task{
	&ent.Task{
		// 登録済み
		ID: 1, Title: "", Description: "task1 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
	&ent.Task{
		// task未登録
		ID: 10001, Title: "task99", Description: "task99 description", Status: "TODO", DueDate: &t, UserID: 1,
	},
}
