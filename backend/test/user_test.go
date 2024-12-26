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

func TestCreateUser(t *testing.T) {

	// fixturesの投入
	loadFixture(t)

	// テストケース
	tests := []struct {
		name    string
		args    *ent.User
		wanterr error
	}{
		{
			// 正常系
			name:    "case: Success",
			args:    testdata.UserTestData[1],
			wanterr: nil,
		},
		{
			// 異常系: 一意制約違反
			name:    "case: Duplicate error",
			args:    testdata.UserTestData[0],
			wanterr: fmt.Errorf("failed to create user in repository: ent: constraint failed: pq: duplicate key value violates unique constraint \"users_email_key\""),
		},
	}

	repo := repository.NewUserRepository(testClient)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		goterr := repo.CreateUser(context.Background(), tt.args)
		//fmt.Println(goterr)

		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			fmt.Println("OK")
		} else {
			if goterr.Error() != tt.wanterr.Error() {
				t.Errorf("[FAIL] return error mismatch\n goterr = %v,\n wanterr= %v\n", goterr, tt.wanterr)
			} else {
				fmt.Println("OK")
			}
		}
	}
}

func TestGetUserById(t *testing.T) {

	// fixturesの投入
	loadFixture(t)

	// テストケース
	tests := []struct {
		name    string
		args    *ent.User
		want    *ent.User
		wanterr error
	}{
		{
			// 正常系
			name:    "case: Success",
			args:    testdata.UserTestData[0],
			want:    testdata.UserTestData[0],
			wanterr: nil,
		},
		{
			// 異常系: 存在しないデータ
			name:    "case: Duplicate error",
			args:    testdata.UserTestData[1],
			want:    nil,
			wanterr: fmt.Errorf("failed to get user by id (99) in repository: ent: user not found"),
		},
	}

	repo := repository.NewUserRepository(testClient)

	for _, tt := range tests {
		fmt.Println(tt.name)

		// test対象メソッドの実行
		got, goterr := repo.GetUserById(context.Background(), tt.args.ID)
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
