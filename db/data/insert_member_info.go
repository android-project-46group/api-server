package data

import (
	"fmt"

	models "github.com/android-project-46group/api-server/db/my_models"
	_ "github.com/lib/pq"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func FindUserByName(name string) (*models.Member, error) {
	member, err := models.Members(qm.Select("*"), qm.Where("name_ja=?", name)).One(Ctx, DB)

	if err != nil {
		return nil, err
	}
	return member, nil
}

func InsertMemberInfosHinata() {
	InsertMemberInfosFromGroupName("hinatazaka")
}

func InsertMemberInfosNogi() {
	InsertMemberInfosFromGroupName("nogizaka")
}

func InsertMemberInfosSakura() {
	InsertMemberInfosFromGroupName("sakurazaka")
}

func InsertMemberInfosFromGroupName(group string) {
	infos := LoadMemberInfoFile(group)

	for _, info := range infos {

		member_db, _ := FindUserByName(info.Name)

		m := &models.MemberInfo{
			MemberID:   member_db.MemberID,
			Birthday:   info.Birthday,
			BloodType:  info.BloodType,
			Height:     info.Height,
			Generation: info.Generation,
		}
		m.BlogURL.String = info.BlogUrl
		m.BlogURL.Valid = true
		m.ImgURL.String = info.ImgUrl
		m.ImgURL.Valid = true
		err := m.Insert(Ctx, DB, boil.Infer())
		if err != nil {
			fmt.Println("Error" + info.Name)
			fmt.Println(err)
		}
	}
}
