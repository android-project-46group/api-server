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
		} else if info.Generation == "5期生" {
			date = time.Date(2022, 2, 1, 9, 0, 0, 0, time.Local)
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

func UpdateGraduatedMembersNogi() {

	m, err := FindUserByName("相楽 伊織")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2018, 7, 16, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("矢田 里沙子")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2014, 10, 18, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("松井 玲奈")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2015, 5, 14, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("渡辺 みり愛")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2021, 8, 31, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("米徳 京花")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2014, 10, 18, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("山崎 怜奈")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 7, 17, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("寺田 蘭世")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2021, 12, 12, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("佐々木 琴子")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2020, 3, 31, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("伊藤 純奈")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2021, 8, 29, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("伊藤 かりん")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2019, 5, 24, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("新内 眞衣")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 2, 10, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("北野 日奈子")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 4, 30, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("堀 未央奈")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2021, 3, 28, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("秋元 真夏")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2023, 2, 26, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("和田 まあや")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 12, 4, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("若月 佑美")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2018, 11, 30, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("大和 里菜")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2014, 12, 15, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("宮澤 成良")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2013, 11, 17, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("松村 沙友理")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2021, 7, 13, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("星野 みなみ")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 2, 12, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("深川 麻衣")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2016, 6, 16, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("樋口 日奈")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 10, 31, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("畠中 清羅")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2015, 4, 4, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("橋本 奈々未")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2017, 2, 20, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("能條 愛未")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2018, 12, 15, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("西野 七瀬")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2018, 12, 31, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("永島 聖羅")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2016, 3, 20, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("中元 日芽香")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2017, 12, 31, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("中田 花奈")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2020, 10, 25, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("高山 一実")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2021, 11, 21, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("白石 麻衣")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2020, 10, 28, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("桜井 玲香")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2019, 9, 1, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("斉藤 優里")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2019, 6, 30, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("斎藤 ちはる")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2018, 7, 16, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("齋藤 飛鳥")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 12, 31, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("川村 真洋")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2018, 3, 31, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("川後 陽菜")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2018, 12, 20, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("柏 幸奈")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2013, 11, 17, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("衛藤 美彩")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2019, 3, 31, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("井上 小百合")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2020, 4, 27, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("伊藤 万理華")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2017, 12, 23, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("伊藤 寧々")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2014, 10, 19, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("市來 玲奈")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2014, 7, 21, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("岩瀬 佑美子")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2012, 11, 18, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("生駒 里奈")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2018, 5, 6, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("生田 絵梨花")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2021, 12, 31, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())
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

func UpdateGraduatedMembersSakura() {

	m, err := FindUserByName("渡邉 理佐")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 5, 22, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("渡邉 美穂")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 7, 31, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("原田 葵")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 8, 20, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("尾関 梨香")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 9, 11, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

	m, err = FindUserByName("菅井 友香")
	if err != nil {
		fmt.Println(err)
	}
	m.LeftAt.Time = time.Date(2022, 11, 9, 9, 0, 0, 0, time.Local)
	m.LeftAt.Valid = true
	m.Update(Ctx, DB, boil.Infer())

}
