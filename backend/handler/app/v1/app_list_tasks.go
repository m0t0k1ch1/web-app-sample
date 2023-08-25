package appv1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/pkg/errors"

	appv1 "github.com/m0t0k1ch1/web-app-sample/backend/gen/buf/app/v1"
	"github.com/m0t0k1ch1/web-app-sample/backend/gen/sqlc/mysql"
)

func (h *AppServiceHandler) ListTasks(ctx context.Context, req *connect.Request[appv1.ListTasksRequest]) (*connect.Response[appv1.ListTasksResponse], error) {
	qdb := mysql.New(h.env.DB)

	tasks, err := qdb.ListTasks(ctx)
	if err != nil {
		return nil, newUnknownError(errors.Wrap(err, "failed to list tasks"))
	}

	return connect.NewResponse(&appv1.ListTasksResponse{
		Tasks: newTasks(tasks),
	}), nil
}
