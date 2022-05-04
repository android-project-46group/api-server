package data

import (
	"fmt"

	models "web/db/my_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func InsertSongs(groupName string) {
	formations := LoadFormationFile(groupName)
	g_id, err := FindGroupIdByName(groupName)
	if err != nil {
		return
	}
	fmt.Println(formations)
	for _, formation := range formations {
		
		fmt.Println(formation.Title)
		formationId, err := FindFormationIdByPositions(formation.Position)
		
		fmt.Println(formationId)

		if err != nil {
			continue
		}
		m := &models.Song{
			GroupID: g_id,
			FormationID: formationId,
			Title: formation.Title,
			Single: formation.Single,
		}
		m.Insert(Ctx, DB, boil.Infer())
	}
}

func FindFormationIdByPositions(positions []Position) (int, error) {
	f := MakeFormations(positions)
	fmt.Println(f)
	ss, err := models.Formations(
		qm.Where("first_row_num = ?", f[0]),
		qm.Where("second_row_num = ?", f[1]),
		qm.Where("third_row_num = ?", f[2]),
	).One(Ctx, DB)
	if err != nil {
		return 0, err
	}
	return ss.FormationID, nil
}
