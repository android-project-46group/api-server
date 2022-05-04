package db

import (
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/android-project-46group/api-server/db/my_models"
)

func (q *SqlQuerier) ExistGroup(groupName string) bool {

	_, err := models.Groups(qm.Where("group_name = ?", groupName)).One(q.ctx, q.DB)

	return err == nil
}

func (q *SqlQuerier) FindGroupByName(groupName string) (*models.Group, error) {

	g, err := models.Groups(qm.Where("group_name = ?", groupName)).One(q.ctx, q.DB)

	return g, err
}
