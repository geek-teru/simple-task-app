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

func TestCreateUser(t *testing.T) {

	// テストケース
	tests := []struct {
		name    string
		args    service.UserRequest
		want    service.UserResponse
		wanterr error
	}{
		{
			// 正常系
			name:    "case: Success",
			args:    testdata.UserReqTestData[1],
			want:    testdata.UserResTestData[1],
			wanterr: nil,
		},
		// {
		// 	// 異常系: 一意制約違反
		// 	name:    "case: Duplicate error",
		// 	args:    testdata.UserReqTestData[0],
		// 	want:    nil,
		// 	wanterr: fmt.Errorf("failed to create user in repository: ent: constraint failed: pq: duplicate key value violates unique constraint \"users_email_key\""),
		// },
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepo := repository.NewMockUserRepositoryInterface(ctrl)
	userService := service.NewUserService(userRepo)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		userRepo.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(testdata.UserTestData[1], nil)

		got, goterr := userService.CreateUser(tt.args)
		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.User{})); diff != "" {
				t.Errorf("[FAIL]return mismatch\n got = %v,\n want= %v\n", got, tt.want)
			} else {
				fmt.Println("OK")
			}
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

// func TestGetUserById(t *testing.T) {

// 	// テストケース
// 	tests := []struct {
// 		name    string
// 		args    *service.UserRequest
// 		want    *service.UserRequest
// 		wanterr error
// 	}{
// 		{
// 			// 正常系
// 			name:    "case: Success",
// 			args:    testdata.UserReqTestData[0],
// 			want:    testdata.UserReqTestData[0],
// 			wanterr: nil,
// 		},
// 		{
// 			// 異常系: 存在しないデータ
// 			name:    "case: Duplicate error",
// 			args:    testdata.UserReqTestData[1],
// 			want:    nil,
// 			wanterr: fmt.Errorf("failed to get user by id (10001) in repository: ent: user not found"),
// 		},
// 	}

// 	repo := repository.NewUserRepository(testClient)

// 	for _, tt := range tests {
// 		fmt.Println(tt.name)

// 		// test対象メソッドの実行
// 		got, goterr := repo.GetUserById(context.Background(), tt.args.ID)
// 		// fmt.Println(got)

// 		// 結果の比較
// 		if tt.wanterr == nil && goterr == nil {
// 			// 正常
// 			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.User{})); diff != "" {
// 				t.Errorf("[FAIL]return mismatch\n got = %v,\n want= %v\n", got, tt.want)
// 			} else {
// 				fmt.Println("OK")
// 			}
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
