package data

import (
	"fmt"
	"strconv"
	"strings"

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

		trimmed := strings.TrimRight(info.Height, "cm")
		heightCM, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			fmt.Println("failed to parse float: ", err)
		}

		m := &models.MemberInfo{
			MemberID:   member_db.MemberID,
			Birthday:   info.Birthday,
			BloodType:  info.BloodType,
			HeightCM:   heightCM,
			Generation: info.Generation,
			LocaleID:   locale.LocaleID,
		}
		m.BlogURL.String = info.BlogUrl
		m.BlogURL.Valid = true
		m.ImgURL.String = info.ImgUrl
		m.ImgURL.Valid = true
		err = m.Insert(Ctx, DB, boil.Infer())
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

		trimmed := strings.TrimRight(info.Height, "cm")
		heightCM, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			fmt.Println("failed to parse float: ", err)
		}

		m := &models.MemberInfo{
			MemberID:   member_db.MemberID,
			Birthday:   info.Birthday,
			BloodType:  info.BloodType,
			HeightCM:   heightCM,
			Generation: info.Generation,
			LocaleID:   locale.LocaleID,
		}
		m.BlogURL.String = info.BlogUrl
		m.BlogURL.Valid = true
		m.ImgURL.String = info.ImgUrl
		m.ImgURL.Valid = true
		err = m.Insert(Ctx, DB, boil.Infer())
		if err != nil {
			fmt.Println("Error" + info.Name)
			fmt.Println(err)
		}
	}
}

func InsertGraduatedMemberInfos() {
	// 脱退したメンバーを挿入する
	member_db, _ := FindUserByName("柿崎 芽実")
	m := &models.MemberInfo{
		MemberID:   member_db.MemberID,
		Birthday:   "2001年12月2日",
		BloodType:  "A 型",
		HeightCM:   157,
		Generation: "1期生",
		LocaleID:   2,
	}
	err := m.Insert(Ctx, DB, boil.Infer())
	if err != nil {
		fmt.Println("Error ", m)
		fmt.Println(err)
	}

	member_db, _ = FindUserByName("井口 眞緒")
	m = &models.MemberInfo{
		MemberID:   member_db.MemberID,
		Birthday:   "1995年11月10日",
		BloodType:  "AB 型",
		HeightCM:   163,
		Generation: "1期生",
		LocaleID:   2,
	}
	err = m.Insert(Ctx, DB, boil.Infer())
	if err != nil {
		fmt.Println("Error ", m)
		fmt.Println(err)
	}
}
