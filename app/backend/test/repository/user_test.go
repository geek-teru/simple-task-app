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

	cleanupUsersTable(t, testClient)

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
			args:    testdata.UserTestData[1],
			want:    testdata.UserTestData[1],
			wanterr: nil,
		},
		{
			// 異常系: 一意制約違反
			name:    "case: Duplicate Error",
			args:    testdata.UserTestData[0],
			want:    nil,
			wanterr: fmt.Errorf("[ERROR] failed to create user in repository: ent: constraint failed: pq: duplicate key value violates unique constraint \"users_email_key\""),
		},
	}

	repo := repository.NewUserRepository(testClient)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		got, goterr := repo.CreateUser(context.Background(), tt.args)
		//fmt.Println(goterr)

		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.User{}, ent.UserEdges{})); diff != "" {
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

func TestGetUserById(t *testing.T) {

	// fixturesの投入
	loadFixture(t)

	cleanupUsersTable(t, testClient)

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
			name:    "case: Not Exist Error",
			args:    testdata.UserTestData[2],
			want:    nil,
			wanterr: fmt.Errorf("[ERROR] failed to get user by id (999) in repository: ent: user not found"),
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
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.User{}, ent.UserEdges{})); diff != "" {
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

func TestGetUserByEmail(t *testing.T) {

	// fixturesの投入
	loadFixture(t)

	cleanupUsersTable(t, testClient)

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
			name:    "case: Not Exist Error",
			args:    testdata.UserTestData[2],
			want:    nil,
			wanterr: fmt.Errorf("[ERROR] failed to get user by email (david@example.com) in repository: ent: user not found"),
		},
	}

	repo := repository.NewUserRepository(testClient)

	for _, tt := range tests {
		fmt.Println(tt.name)

		// test対象メソッドの実行
		got, goterr := repo.GetUserByEmail(context.Background(), tt.args.Email)
		// fmt.Println(got)

		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.User{}, ent.UserEdges{})); diff != "" {
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

func TestUpdateUser(t *testing.T) {

	// fixturesの投入
	loadFixture(t)

	cleanupUsersTable(t, testClient)

	// テストケース
	tests := []struct {
		name    string
		args    *ent.User
		argsId  int
		want    *ent.User
		wanterr error
	}{
		{
			// 正常系
			name:    "case: Success",
			args:    testdata.UserTestData[3],
			argsId:  testdata.UserTestData[3].ID,
			want:    testdata.UserTestData[3],
			wanterr: nil,
		},
		{
			// 異常系: 一意制約違反
			name:    "case: Duplicate Error",
			args:    testdata.UserTestData[4],
			argsId:  testdata.UserTestData[4].ID,
			want:    nil,
			wanterr: fmt.Errorf("[ERROR] failed to update user in repository: ent: constraint failed: pq: duplicate key value violates unique constraint \"users_email_key\""),
		},
	}

	repo := repository.NewUserRepository(testClient)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		got, goterr := repo.UpdateUser(context.Background(), tt.args, tt.args.ID)
		//fmt.Println(goterr)

		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreUnexported(ent.User{}, ent.UserEdges{})); diff != "" {
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
