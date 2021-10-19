package data

import (
    "fmt"

	_ "github.com/lib/pq"
	models "web/db/my_models"

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
	// groupId := 3
	
	for _, info := range infos {

		member_db, _ := FindUserByName(info.Name)
		
		fmt.Printf("%s\n", info.Generation)
		fmt.Printf("%d: %s\n", member_db.MemberID, info.Name)
		
		m := &models.MemberInfo{
			MemberID: member_db.MemberID,
			Birthday: info.Birthday,
			BloodType: info.BloodType,
			Height: info.Height,
			Generation: info.Generation,
		}
		m.BlogURL.String = info.BlogUrl
		m.BlogURL.Valid = true
		m.ImgURL.String = info.ImgUrl
		m.ImgURL.Valid = true
		
		m.Insert(Ctx, DB, boil.Infer())
	}
}
