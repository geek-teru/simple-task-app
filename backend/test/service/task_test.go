package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/geek-teru/simple-task-app/ent"
	repository "github.com/geek-teru/simple-task-app/repository/mock"
	"github.com/geek-teru/simple-task-app/service"
	"github.com/geek-teru/simple-task-app/testdata"
	gomock "github.com/golang/mock/gomock"
	cmp "github.com/google/go-cmp/cmp"
	cmpopts "github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateTask(t *testing.T) {

	// テストケース
	tests := []struct {
		name        string
		args        *service.TaskRequest
		args_userid int
		mockreturn  *ent.Task
		mockerr     error
		want        service.TaskResponse
		wanterr     error
	}{
		{
			// 正常系
			name:        "case: Success",
			args:        testdata.TaskReqTestData[1],
			args_userid: 1,
			mockreturn:  testdata.TaskTestData[1],
			mockerr:     nil,
			want:        testdata.TaskResTestData[1],
			wanterr:     nil,
		},
		{
			// 異常系: Not Empty制約違反
			name:       "case: Duplicate error",
			args:       testdata.TaskReqTestData[2],
			mockreturn: nil,
			mockerr:    fmt.Errorf("[ERROR] failed to create task in repository: ent: validator failed for field \"Task.title\": value is less than the required length"),
			want:       service.TaskResponse{},
			wanterr:    fmt.Errorf("[ERROR] failed to create task in repository: ent: validator failed for field \"Task.title\": value is less than the required length"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	TaskRepo := repository.NewMockTaskRepositoryInterface(ctrl)
	TaskService := service.NewTaskService(TaskRepo)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		TaskRepo.EXPECT().CreateTask(context.Background(), gomock.Any()).Return(tt.mockreturn, tt.mockerr)

		got, goterr := TaskService.CreateTask(tt.args, tt.args_userid)
		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.Task{})); diff != "" {
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

func TestListTask(t *testing.T) {

}

func TestGetTaskById(t *testing.T) {

	// テストケース
	tests := []struct {
		name        string
		args        int
		args_userid int
		mockreturn  *ent.Task
		mockerr     error
		want        service.TaskResponse
		wanterr     error
	}{
		{
			// 正常系
			name:        "case: Success",
			args:        testdata.TaskTestData[1].ID,
			args_userid: 1,
			mockreturn:  testdata.TaskTestData[1],
			mockerr:     nil,
			want:        testdata.TaskResTestData[1],
			wanterr:     nil,
		},
		{
			// 異常系: 存在しないデータ
			name:       "case: Duplicate error",
			args:       testdata.TaskTestData[1].ID,
			mockreturn: nil,
			mockerr:    fmt.Errorf("[ERROR] failed to get task by id (10001) in repository: ent: task not found"),
			want:       service.TaskResponse{},
			wanterr:    fmt.Errorf("[ERROR] failed to get task by id (10001) in repository: ent: task not found"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	TaskRepo := repository.NewMockTaskRepositoryInterface(ctrl)
	TaskService := service.NewTaskService(TaskRepo)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		TaskRepo.EXPECT().GetTaskById(context.Background(), gomock.Any(), gomock.Any()).Return(tt.mockreturn, tt.mockerr)

		got, goterr := TaskService.GetTaskById(tt.args, tt.args_userid)
		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.Task{})); diff != "" {
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

func TestUpdateTask(t *testing.T) {

	// テストケース
	tests := []struct {
		name        string
		args        *service.TaskRequest
		args_id     int
		args_userid int
		mockreturn  *ent.Task
		mockerr     error
		want        service.TaskResponse
		wanterr     error
	}{
		{
			// 正常系
			name:        "case: Success",
			args:        testdata.TaskReqTestData[3],
			args_id:     1,
			args_userid: 1,
			mockreturn:  testdata.TaskTestData[3],
			mockerr:     nil,
			want:        testdata.TaskResTestData[3],
			wanterr:     nil,
		},
		{
			// 異常系: Not Empty制約違反
			name:       "case: Duplicate error",
			args:       testdata.TaskReqTestData[2],
			mockreturn: nil,
			mockerr:    fmt.Errorf("[ERROR] failed to create task in repository: ent: validator failed for field \"Task.title\": value is less than the required length"),
			want:       service.TaskResponse{},
			wanterr:    fmt.Errorf("[ERROR] failed to create task in repository: ent: validator failed for field \"Task.title\": value is less than the required length"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	TaskRepo := repository.NewMockTaskRepositoryInterface(ctrl)
	TaskService := service.NewTaskService(TaskRepo)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		TaskRepo.EXPECT().UpdateTask(context.Background(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockreturn, tt.mockerr)

		got, goterr := TaskService.CreateTask(tt.args, tt.args_userid)
		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.Task{})); diff != "" {
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

func DeleteTask(t *testing.T) {

}
