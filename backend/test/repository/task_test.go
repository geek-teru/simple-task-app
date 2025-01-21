package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/repository"
	"github.com/geek-teru/simple-task-app/testdata"
	cmp "github.com/google/go-cmp/cmp"
	cmpopts "github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateTask(t *testing.T) {

	// fixturesの投入
	loadFixture(t)

	// テストケース
	tests := []struct {
		name    string
		args    *ent.Task
		want    *ent.Task
		wanterr error
	}{
		{
			// 正常系
			name:    "case: Success",
			args:    testdata.TaskTestData[1],
			want:    testdata.TaskTestData[1],
			wanterr: nil,
		},
		{
			// 異常系: Not Empty制約違反
			name:    "case: Missing Required Error",
			args:    testdata.TaskTestData[0],
			want:    nil,
			wanterr: fmt.Errorf("[ERROR] failed to create task in repository: ent: validator failed for field \"Task.title\": value is less than the required length"),
		},
	}

	repo := repository.NewTaskRepository(testClient)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		got, goterr := repo.CreateTask(context.Background(), tt.args)
		//fmt.Println(goterr)

		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.Task{}, ent.TaskEdges{})); diff != "" {
				t.Errorf("[FAIL] return mismatch\n got = %v,\n want= %v\n", got, tt.want)
			} else {
				fmt.Println("OK")
			}
		} else if tt.wanterr == nil || goterr == nil {
			// 期待値と結果のどちらか片方がnil
			t.Errorf("[FAIL] return error mismatch\n goterr = %v,\n wanterr= %v\n", goterr, tt.wanterr)
		} else {
			// 異常
			if goterr.Error() != tt.wanterr.Error() {
				t.Errorf("[FAIL] return error mismatch\n goterr = %v,\n wanterr= %v\n", goterr, tt.wanterr)
			} else {
				fmt.Println("OK")
			}
		}
	}
}

// func TestGetTaskById(t *testing.T) {

// 	// fixturesの投入
// 	loadFixture(t)

// 	// テストケース
// 	tests := []struct {
// 		name    string
// 		args    *ent.Task
// 		want    *ent.Task
// 		wanterr error
// 	}{
// 		{
// 			// 正常系
// 			name:    "case: Success",
// 			args:    testdata.TaskTestData[0],
// 			want:    testdata.TaskTestData[0],
// 			wanterr: nil,
// 		},
// 		{
// 			// 異常系: 存在しないデータ
// 			name:    "case: Not exist error",
// 			args:    testdata.TaskTestData[1],
// 			want:    nil,
// 			wanterr: fmt.Errorf("[ERROR] failed to get task by id (10001) in repository: ent: task not found"),
// 		},
// 	}

// 	repo := repository.NewTaskRepository(testClient)

// 	for _, tt := range tests {
// 		fmt.Println(tt.name)

// 		// test対象メソッドの実行
// 		got, goterr := repo.GetTaskById(context.Background(), tt.args.ID)
// 		// fmt.Println(got)

// 		// 結果の比較
// 		if tt.wanterr == nil && goterr == nil {
// 			// 正常
// 			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.Task{}, ent.TaskEdges{})); diff != "" {
// 				t.Errorf("[FAIL] return mismatch\n got = %v,\n want= %v\n", got, tt.want)
// 			} else {
// 				fmt.Println("OK")
// 			}
// 		} else if tt.wanterr == nil || goterr == nil {
// 			// 期待値と結果のどちらか片方がnil
// 			t.Errorf("[FAIL] return error mismatch\n goterr = %v,\n wanterr= %v\n", goterr, tt.wanterr)
// 		} else {
// 			// 異常
// 			if goterr.Error() != tt.wanterr.Error() {
// 				t.Errorf("[FAIL] return error mismatch\n goterr = %v,\n wanterr= %v\n", goterr, tt.wanterr)
// 			} else {
// 				fmt.Println("OK")
// 			}
// 		}
// 	}
// }
