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

func TestSignUp(t *testing.T) {

	// テストケース
	tests := []struct {
		name    string
		args    *service.UserRequest
		want    service.UserResponse
		wanterr error
	}{
		{
			// 正常系
			name:    "case: Success",
			args:    &testdata.UserReqTestData[1],
			want:    testdata.UserResTestData[1],
			wanterr: nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepo := repository.NewMockUserRepositoryInterface(ctrl)
	userService := service.NewUserService(userRepo)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		userRepo.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(testdata.UserTestData[1], nil)

		got, goterr := userService.SignUp(tt.args)
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

func TestSignIn(t *testing.T) {

	// テストケース
	tests := []struct {
		name    string
		args    *service.UserRequest
		want    string
		wanterr error
	}{
		{
			// 正常系
			name:    "case: Success",
			args:    &testdata.UserReqTestData[1],
			want:    "",
			wanterr: nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepo := repository.NewMockUserRepositoryInterface(ctrl)
	userService := service.NewUserService(userRepo)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		userRepo.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(testdata.UserTestData[1], nil)

		got, goterr := userService.SignIn(tt.args)
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

func TestGetUserProfile(t *testing.T) {

	// テストケース
	tests := []struct {
		name    string
		args    int
		want    service.UserResponse
		wanterr error
	}{
		{
			// 正常系
			name:    "case: Success",
			args:    testdata.UserTestData[0].ID,
			want:    testdata.UserResTestData[0],
			wanterr: nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepo := repository.NewMockUserRepositoryInterface(ctrl)
	userService := service.NewUserService(userRepo)

	for _, tt := range tests {
		fmt.Println(tt.name)
		userRepo.EXPECT().GetUserById(context.Background(), gomock.Any()).Return(testdata.UserTestData[0], nil)
		got, goterr := userService.GetUserProfile(tt.args)
		// fmt.Println(got)

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
