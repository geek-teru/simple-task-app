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
		name       string
		args       *service.UserRequest
		mockreturn *ent.User
		mockerr    error
		want       service.UserResponse
		wanterr    error
	}{
		{
			// 正常系
			name:       "case: Success",
			args:       testdata.UserReqTestData[1],
			mockreturn: testdata.UserTestData[1],
			mockerr:    nil,
			want:       testdata.UserResTestData[1],
			wanterr:    nil,
		},
		{
			// 異常系: データ重複
			name:       "case: Duplicate error",
			args:       testdata.UserReqTestData[0],
			mockreturn: nil,
			mockerr:    fmt.Errorf("[ERROR] failed to create user in repository: ent: constraint failed: pq: duplicate key value violates unique constraint \"users_email_key\""),
			want:       service.UserResponse{},
			wanterr:    fmt.Errorf("[ERROR] failed to create user in repository: ent: constraint failed: pq: duplicate key value violates unique constraint \"users_email_key\""),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepo := repository.NewMockUserRepositoryInterface(ctrl)
	userService := service.NewUserService(userRepo)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		userRepo.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(tt.mockreturn, tt.mockerr)

		got, goterr := userService.SignUp(tt.args)
		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.User{})); diff != "" {
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

func TestSignIn(t *testing.T) {

	// テストケース
	tests := []struct {
		name       string
		args       *service.UserRequest
		mockreturn *ent.User
		mockerr    error
		want       string
		wanterr    error
	}{
		{
			// 正常系
			name:       "case: Success",
			args:       testdata.UserReqTestData[0],
			mockreturn: testdata.UserTestData[0],
			mockerr:    nil,
			want:       "",
			wanterr:    nil,
		},
		{
			// 異常系: 未登録
			name:       "case: Not exist error",
			args:       testdata.UserReqTestData[1],
			mockreturn: nil,
			mockerr:    fmt.Errorf("[ERROR] failed to get user by id (10001) in repository: ent: user not found"),
			want:       "",
			wanterr:    fmt.Errorf("[ERROR] failed to get user by id (10001) in repository: ent: user not found"),
		},
		{
			// 異常系: password間違い
			name:       "case: Unauthorized error",
			args:       testdata.UserReqTestData[2],
			mockreturn: testdata.UserTestData[0],
			mockerr:    nil,
			want:       "",
			wanterr:    fmt.Errorf("[ERROR] failed to SignIn in service: crypto/bcrypt: hashedPassword is not the hash of the given password"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepo := repository.NewMockUserRepositoryInterface(ctrl)
	userService := service.NewUserService(userRepo)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		userRepo.EXPECT().GetUserByEmail(context.Background(), tt.args.Email).Return(tt.mockreturn, tt.mockerr)
		got, goterr := userService.SignIn(tt.args)
		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			// jwtの期待値は長さが1以上の文字列とする
			if len(got) > 0 {
				fmt.Println("OK")
			} else {
				t.Errorf("[FAIL] return mismatch\n got = %v,\n want= %v\n", got, tt.want)
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
