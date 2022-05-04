package data

import (
	"fmt"

	models "github.com/android-project-46group/api-server/db/my_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func InsertFormations() {
	formations := LoadFormationFile("hinatazaka")
	for _, formation := range formations {
		numArrays := MakeFormations(formation.Position)

		m := &models.Formation{}
		m.FirstRowNum.Int16 = int16(numArrays[0])
		m.FirstRowNum.Valid = true
		m.SecondRowNum.Int16 = int16(numArrays[1])
		m.SecondRowNum.Valid = true
		m.ThirdRowNum.Int16 = int16(numArrays[2])
		m.ThirdRowNum.Valid = true

		// 既に存在していたら追加しない
		_, err := FindFormationIdByPositions(formation.Position)
		if err == nil {
			fmt.Println("Found formation of " + formation.Title)
			continue
		}
		m.Insert(Ctx, DB, boil.Infer())
	}
}

func MakeFormations(positions []Position) []int {
	firstRowNum := 0
	secondRowNum := 0
	thirdRowNum := 0
	for _, position := range positions {
		pos := position.Position

		if pos[0:1] != "0" {
			thirdRowNum += 1
		} else if pos[1:2] != "0" {
			secondRowNum += 1
		} else if pos[2:3] != "0" {
			firstRowNum += 1
		}
	}
	return []int{firstRowNum, secondRowNum, thirdRowNum}
}
