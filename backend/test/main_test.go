package repository_test

import (
	"fmt"
	"os"
	"testing"

	testfixtures "github.com/go-testfixtures/testfixtures/v3"

	"github.com/geek-teru/simple-task-app/db"
	"github.com/geek-teru/simple-task-app/ent"
)

var testClient *ent.Client

func TestMain(m *testing.M) {
	err := setUp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	code := m.Run()
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

func loadFixture(t *testing.T, path string) {
	t.Helper()

	testDB, err := db.NewDB()
	if err != nil {
		t.Fatal("failed to create test database")
	}

	fixtures, err := testfixtures.New(
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.Database(testDB),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(path),
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}
