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
	InsertMemberInfosFromGroupNameEn("hinatazaka")
}

func InsertMemberInfosNogi() {
	InsertMemberInfosFromGroupName("nogizaka")
	InsertMemberInfosFromGroupNameEn("nogizaka")
}

func InsertMemberInfosSakura() {
	InsertMemberInfosFromGroupName("sakurazaka")
	InsertMemberInfosFromGroupNameEn("sakurazaka")
}

func findLocale(name string) (*models.Locale, error) {
	s, err := models.Locales(qm.Where("name=?", name)).One(Ctx, DB)
	if err != nil {
		return nil, fmt.Errorf("FindApiKeyByName: %w", err)
	}
	return s, nil
}

func InsertMemberInfosFromGroupName(group string) {
	infos := LoadMemberInfoFile(group)
	locale, err := findLocale("ja")
	if err != nil {
		fmt.Println("cannot find locale for ja")
		return
	}

	for _, info := range infos {

		member_db, _ := FindUserByName(info.Name)

		m := &models.MemberInfo{
			MemberID:   member_db.MemberID,
			Birthday:   info.Birthday,
			BloodType:  info.BloodType,
			Height:     info.Height,
			Generation: info.Generation,
			LocaleID:   locale.LocaleID,
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

func InsertMemberInfosFromGroupNameEn(group string) {
	infos := LoadMemberInfoFile(fmt.Sprintf("%s_en", group))
	locale, err := findLocale("en")
	if err != nil {
		fmt.Println("cannot find locale for en")
		return
	}

	for _, info := range infos {

		member_db, _ := FindUserByName(info.Name)

		m := &models.MemberInfo{
			MemberID:   member_db.MemberID,
			Birthday:   info.Birthday,
			BloodType:  info.BloodType,
			Height:     info.Height,
			Generation: info.Generation,
			LocaleID:   locale.LocaleID,
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
