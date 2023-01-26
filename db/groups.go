package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	models "github.com/android-project-46group/api-server/db/my_models"
)

func (q *SqlQuerier) FindGroupByName(ctx context.Context, groupName string) (*models.Group, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.FindGroupByName")
	defer span.Finish()

	g, err := models.Groups(qm.Where("group_name = ?", groupName)).One(ctx, q.DB)

	return g, fmt.Errorf("FindGroupByName: %w", err)
}
