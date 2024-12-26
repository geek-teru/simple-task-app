package repository_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/repository"
	"github.com/geek-teru/simple-task-app/testdata"
	cmp "github.com/google/go-cmp/cmp"
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
			wanterr: nil,
		},
	}

	repo := repository.NewUserRepository(testClient)

	for _, tt := range tests {
		fmt.Println(tt.name)
		//fmt.Println(tt.args) //debug
		goterr := repo.CreateUser(context.Background(), tt.args)
		fmt.Println(got)
		// 結果の比較
		if tt.wanterr == nil && goterr == nil {
			// 正常
			fmt.Println("normal")
		} else {
			if !errors.Is(goterr, tt.wanterr) {
				t.Errorf("[FAIL] return error mismatch\n goterr = %v,\n wanterr= %v\n", goterr, tt.wanterr)
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
			wanterr: nil,
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
			fmt.Println("normal")
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("[FAIL]return error mismatch\n goterr = %v,\n anterr= %v\n", goterr, tt.wanterr)
			}
		} else {
			// 異常
			if diff := cmp.Diff(goterr.Error(), tt.wanterr.Error()); diff != "" {
				fmt.Println("exception")
				t.Errorf("[FAIL]return error mismatch\n goterr = %v,\n anterr= %v\n", goterr, tt.wanterr)
			}
		}
		//fmt.Println("--------------------------------------------------------------------------------")
	}
}
