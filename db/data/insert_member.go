package data

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	// "./db/data"
	models "github.com/android-project-46group/api-server/db/my_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	// "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func InsertMembersHinata() {
	groupName := "hinatazaka"
	infos := LoadMemberInfoFile(groupName)
	groupId, err := FindGroupIdByName(groupName)
	if err != nil {
		fmt.Println("Group name " + groupName + " does not exist.")
		return
	}

	for _, info := range infos {
		// 1, 2 期生は日向の結成日
		date := time.Date(2019, 2, 11, 9, 0, 0, 0, time.Local)
		if info.Generation == "3期生" && info.Name != "上村ひなの" {
			date = time.Date(2019, 9, 7, 9, 0, 0, 0, time.Local)
		}
		// 今回はひなたなので２
		m := &models.Member{GroupID: groupId, NameJa: info.Name}
		m.JoinedAt.Time = date
		m.JoinedAt.Valid = true

		m.Insert(Ctx, DB, boil.Infer())
	}

	// 脱退したメンバーを挿入する
	k := &models.Member{GroupID: groupId, NameJa: "柿崎 芽実"}
	k.JoinedAt.Time = time.Date(2019, 2, 11, 9, 0, 0, 0, time.Local)
	k.JoinedAt.Valid = true
	k.LeftAt.Time = time.Date(2019, 8, 1, 9, 0, 0, 0, time.Local)
	k.LeftAt.Valid = true
	k.Insert(Ctx, DB, boil.Infer())

	i := &models.Member{GroupID: groupId, NameJa: "井口 眞緒"}
	i.JoinedAt.Time = time.Date(2019, 2, 11, 9, 0, 0, 0, time.Local)
	i.JoinedAt.Valid = true
	i.LeftAt.Time = time.Date(2020, 3, 2, 9, 0, 0, 0, time.Local)
	i.LeftAt.Valid = true
	i.Insert(Ctx, DB, boil.Infer())
}

func InsertMembersNogi() {
	groupName := "nogizaka"
	infos := LoadMemberInfoFile(groupName)
	groupId, err := FindGroupIdByName(groupName)
	if err != nil {
		fmt.Println("Group name " + groupName + " does not exist.")
		return
	}

	new_4th := []string{"弓木 奈於", "松尾 美佑", "林 瑠奈", "佐藤 璃果", "黒見 明香"}

	for _, info := range infos {
		date := time.Date(2011, 8, 21, 9, 0, 0, 0, time.Local)
		if info.Generation == "2期生" {
			date = time.Date(2013, 3, 28, 9, 0, 0, 0, time.Local)
		} else if info.Generation == "3期生" {
			date = time.Date(2016, 9, 4, 9, 0, 0, 0, time.Local)
		} else if info.Generation == "4期生" {
			date = time.Date(2018, 12, 3, 9, 0, 0, 0, time.Local)
		}

		for _, b := range new_4th {
			if info.Name == b {
				date = time.Date(2020, 2, 16, 9, 0, 0, 0, time.Local)
			}
		}

		m := &models.Member{GroupID: groupId, NameJa: info.Name}
		m.JoinedAt.Time = date
		m.JoinedAt.Valid = true

		m.Insert(Ctx, DB, boil.Infer())
	}

	// 脱退したメンバーを挿入する
	k := &models.Member{GroupID: groupId, NameJa: "大園 桃子"}
	k.JoinedAt.Time = time.Date(2016, 9, 4, 9, 0, 0, 0, time.Local)
	k.JoinedAt.Valid = true
	k.LeftAt.Time = time.Date(2021, 9, 4, 9, 0, 0, 0, time.Local)
	k.LeftAt.Valid = true
	k.Insert(Ctx, DB, boil.Infer())

	w := &models.Member{GroupID: groupId, NameJa: "渡辺 みり愛"}
	w.JoinedAt.Time = time.Date(2013, 3, 28, 9, 0, 0, 0, time.Local)
	w.JoinedAt.Valid = true
	w.LeftAt.Time = time.Date(2021, 8, 31, 9, 0, 0, 0, time.Local)
	w.LeftAt.Valid = true
	w.Insert(Ctx, DB, boil.Infer())

	i := &models.Member{GroupID: groupId, NameJa: "伊藤 純奈"}
	i.JoinedAt.Time = time.Date(2013, 3, 28, 9, 0, 0, 0, time.Local)
	i.JoinedAt.Valid = true
	i.LeftAt.Time = time.Date(2021, 8, 29, 9, 0, 0, 0, time.Local)
	i.LeftAt.Valid = true
	i.Insert(Ctx, DB, boil.Infer())

	m := &models.Member{GroupID: groupId, NameJa: "松村 沙友理"}
	m.JoinedAt.Time = time.Date(2011, 8, 21, 9, 0, 0, 0, time.Local)
	m.JoinedAt.Valid = true
	m.LeftAt.Time = time.Date(2021, 8, 27, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Insert(Ctx, DB, boil.Infer())

	m = &models.Member{GroupID: groupId, NameJa: "堀未 央奈"}
	m.JoinedAt.Time = time.Date(2013, 3, 28, 9, 0, 0, 0, time.Local)
	m.JoinedAt.Valid = true
	m.LeftAt.Time = time.Date(2021, 3, 28, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Insert(Ctx, DB, boil.Infer())

	m = &models.Member{GroupID: groupId, NameJa: "白石 麻衣"}
	m.JoinedAt.Time = time.Date(2011, 8, 21, 9, 0, 0, 0, time.Local)
	m.JoinedAt.Valid = true
	m.LeftAt.Time = time.Date(2020, 10, 28, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Insert(Ctx, DB, boil.Infer())
}

func InsertMembersSakura() {
	groupName := "sakurazaka"
	infos := LoadMemberInfoFile(groupName)
	groupId, err := FindGroupIdByName(groupName)
	if err != nil {
		fmt.Println("Group name " + groupName + " does not exist.")
		return
	}

	for _, info := range infos {
		// 1, 2 期生は櫻の結成日
		date := time.Date(2020, 10, 14, 9, 0, 0, 0, time.Local)

		m := &models.Member{GroupID: groupId, NameJa: info.Name}
		m.JoinedAt.Time = date
		m.JoinedAt.Valid = true

		m.Insert(Ctx, DB, boil.Infer())
	}
}
