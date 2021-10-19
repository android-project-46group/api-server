package db

import (
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "web/db/my_models"
)

func ExistGroup(groupName string) bool {

	_, err := models.Groups(qm.Where("group_name = ?", groupName)).One(Ctx, DB)

	return err == nil	
}

func FindGroupByName(groupName string) (*models.Group, error) {

	g, err := models.Groups(qm.Where("group_name = ?", groupName)).One(Ctx, DB)

	return g, err
}
