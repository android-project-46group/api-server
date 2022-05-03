package data

import (
	"fmt"

	models "github.com/android-project-46group/api-server/db/my_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func InsertGroups() {
	groups := []string{"nogizaka", "sakurazaka", "hinatazaka"}
	for _, group := range groups {
		fmt.Println(group)
		m := &models.Group{GroupName: group}
		m.Insert(Ctx, DB, boil.Infer())
	}
}

func FindGroupIdByName(group string) (int, error) {

	ss, err := models.Groups(qm.Where("group_name = ?", group)).One(Ctx, DB)
	if err != nil {
		return 0, err
	}
	return ss.GroupID, nil
}
