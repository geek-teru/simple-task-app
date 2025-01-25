package testdata

import (
	"time"

	"github.com/geek-teru/simple-task-app/ent"
)

var t = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

var TaskTestData = []*ent.Task{
	&ent.Task{
		// 登録済み
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
		ID: 1, Title: "task01 updated", Description: "task1 description updated updated", Status: "TODO", DueDate: &t, UserID: 1,
	},
}
