package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	models "github.com/android-project-46group/api-server/db/my_models"
)

// Get all groups from Database.
func (q *SqlQuerier) GetAllGroups(ctx context.Context) ([]*models.Group, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.FindGroupByName")
	defer span.Finish()

	gs, err := models.Groups().All(ctx, q.DB)
	if err != nil {
		return nil, fmt.Errorf("FindGroupByName: %w", err)
	}

	q.logger.Debugf(ctx, "FindGroupByName got: ", gs)
	return gs, nil
}

func (q *SqlQuerier) FindGroupByName(ctx context.Context, groupName string) (*models.Group, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.FindGroupByName")
	defer span.Finish()

	g, err := models.Groups(qm.Where("group_name = ?", groupName)).One(ctx, q.DB)

	if err != nil {
		return nil, fmt.Errorf("FindGroupByName: %w", err)
	}

	q.logger.Debugf(ctx, "FindGroupByName got: ", g)
	return g, nil
}
