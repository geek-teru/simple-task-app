package service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/geek-teru/simple-task-app/db"
	"github.com/geek-teru/simple-task-app/ent"
)

var testClient *ent.Client

func TestMain(m *testing.M) {
	err := setUp() // 前処理
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	code := m.Run()

	defer teardown() // 後処理
	os.Exit(code)
}

func setUp() error {
	client, err := db.NewClient()
	if err != nil {
		return err
	}

	testClient = client

	return nil
}

func teardown() {
	testClient.Close()
}
