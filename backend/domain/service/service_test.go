package service_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/m0t0k1ch1-go/sqlutil"
	"github.com/m0t0k1ch1-go/timeutil/v4"
	"github.com/samber/oops"
	"github.com/stretchr/testify/require"

	"app/container"
	"app/domain/model"
	"app/internal/testutil"
)

var (
	clock    *timeutil.MockClock
	mysqlCtr *container.MySQLContainer
)

type Mocks struct{}

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	ctx := context.Background()

	clock = timeutil.NewMockClock(timeutil.Now())

	mysqlCtr = &container.MySQLContainer{}
	{
		db, dbTeardown, err := testutil.SetupMySQL(ctx)
		if err != nil {
			return failMain(oops.Wrapf(err, "failed to set up app mysql"))
		}
		defer dbTeardown()

		schemaPath, err := filepath.Abs("../../_schema/sql/app.sql")
		if err != nil {
			return failMain(oops.Wrapf(err, "failed to prepare app schema sql path"))
		}

		if err := sqlutil.ExecFile(ctx, db, schemaPath); err != nil {
			return failMain(oops.Wrapf(err, "failed to execute app schema sql"))
		}

		mysqlCtr.App = db
	}

	return m.Run()
}

func failMain(err error) int {
	fmt.Fprintln(os.Stderr, err.Error())
	return 1
}

func setup(t *testing.T) {
	t.Helper()
}

func teardown(t *testing.T) {
	t.Helper()

	ctx := context.Background()

	require.Nil(t, sqlutil.TruncateAll(ctx, mysqlCtr.App))
}

func encodePaginationCursor(t *testing.T, cursor model.PaginationCursor) string {
	t.Helper()

	encodedCursor, err := cursor.Encode()
	require.Nil(t, err)

	return encodedCursor
}
