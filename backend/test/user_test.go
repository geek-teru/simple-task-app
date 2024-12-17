package repository_test

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/geek-teru/simple-task-app/ent"
	"github.com/geek-teru/simple-task-app/repository"
	cmp "github.com/google/go-cmp/cmp"
)

func TestCreateUser(t *testing.T) {
	// fixture
	absolutePath, err := filepath.Abs("testdata")
	if err != nil {
		t.Fatal("failed to get absolute path")
	}
	fmt.Println(absolutePath)

	loadFixture(t, absolutePath)

	type args struct {
		user *ent.User
	}

	// テストケース
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			// 正常系
			name: "case: Success",
			args: args{
				user: &ent.User{
					Name: "user_x", Email: "user_x@example.com", Password: "password",
				},
			},
			wantErr: nil,
		},
		{
			// 異常系
			name: "case: Duplicate error",
			args: args{
				user: &ent.User{
					Name: "user_y", Email: "user_a@example.com", Password: "password",
				},
			},
			wantErr: nil,
		},
	}

	repo := repository.NewUserRepository(testClient)

	for _, tt := range tests {
		fmt.Println(tt.name)

		// test対象メソッドの実行
		fmt.Println(tt.args.user)
		gotErr := repo.CreateUser(context.Background(), tt.args.user)
		fmt.Println(gotErr)
		// 結果の比較
		if tt.wantErr != nil || gotErr != nil { // wantとgotのどちらかがnilでない場合
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("[FAIL] return error mismatch\n gotErr = %v,\n wantErr= %v\n", gotErr, tt.wantErr)
			}
		}
		fmt.Println("--------------------------------------------------------------------------------")
	}
}

func TestGetUserByEmail(t *testing.T) {
	// fixture
	absolutePath, err := filepath.Abs("testdata")
	if err != nil {
		t.Fatal("failed to get absolute path")
	}

	loadFixture(t, absolutePath)

	type args struct {
		email string
	}

	// テストケース
	tests := []struct {
		name    string
		args    args
		want    *ent.User
		wantErr error
	}{
		{
			// 正常系
			name: "case: Success",
			args: args{
				email: "user_a@example.com",
			},
			want: &ent.User{
				ID:       1,
				Name:     "user_a",
				Email:    "user_a@example.com",
				Password: "passworda",
			},
			wantErr: nil,
		},
		{
			// 異常系
			name: "case: No data error",
			args: args{
				email: "user_z@example.com",
			},
			want: &ent.User{
				ID:       1,
				Name:     "user_a",
				Email:    "user_a@example.com",
				Password: "passworda",
			},
			wantErr: fmt.Errorf("failed to get user by email (user_z@example.com) in repository: ent: user not found"),
		},
	}

	repo := repository.NewUserRepository(testClient)

	for _, tt := range tests {
		fmt.Println(tt.name)

		// test対象メソッドの実行
		got, gotErr := repo.GetUserByEmail(context.Background(), tt.args.email)
		fmt.Println(got)
		// 結果の比較
		if tt.wantErr != nil || gotErr != nil {
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("[FAIL]return error mismatch\n gotErr = %v,\n wantErr= %v\n", gotErr, tt.wantErr)
			}
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("[FAIL]return error mismatch\n gotErr = %v,\n wantErr= %v\n", gotErr, tt.wantErr)
			}
		}
		fmt.Println("--------------------------------------------------------------------------------")
	}
}
