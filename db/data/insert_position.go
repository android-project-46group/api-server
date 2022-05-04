package data

import (
	"fmt"

	models "github.com/android-project-46group/api-server/db/my_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func InsertPositions() {
	formations := LoadFormationFile("hinatazaka")

	for _, formation := range formations {

		s_id, err := FindSongIdByTitle(formation.Title)
		if err != nil {
			continue
		}

		for _, position := range formation.Position {
			memberName := nameConversions[position.NameJa]
			m_id, mErr := FindMemberIdByName(memberName)
			if mErr != nil {
				fmt.Println(memberName + " が DB 上に見つかりませんでした。")
				continue
			}

			m := &models.Position{
				SongID:   s_id,
				MemberID: m_id,
				Position: position.Position,
			}
			m.IsCenter.Bool = IsCenter(formation.Title, memberName)
			m.IsCenter.Valid = true
			m.Insert(Ctx, DB, boil.Infer())
		}
	}
}

func FindSongIdByTitle(title string) (int, error) {
	s, err := models.Songs(qm.Where("title = ?", title)).One(Ctx, DB)
	if err != nil {
		return 0, err
	}
	return s.SongID, nil
}

func FindMemberIdByName(name string) (int, error) {
	s, err := models.Members(qm.Where("name_ja = ?", name)).One(Ctx, DB)
	if err != nil {
		return 0, err
	}
	return s.MemberID, nil
}

func IsCenter(title string, name string) bool {

	val, ok := centerLists[title]
	if !ok {
		return false
	}
	if val == name {
		return true
	}
	return false
}

var centerLists = map[string]string{
	"ってか": "金村 美玖",
	"思いがけないダブルレインボー":    "金村 美玖",
	"アディショナルタイム":        "金村 美玖",
	"何度でも何度でも":          "上村 ひなの",
	"君しか勝たん":            "加藤 史帆",
	"膨大な夢に押し潰されて":       "加藤 史帆",
	"アザトカワイイ":           "佐々木 美玲",
	"ソンナコトナイヨ":          "小坂 菜緒",
	"こんなに好きになっちゃっていいの？": "小坂 菜緒",
	"ドレミソラシド":           "小坂 菜緒",
	"キュン":               "小坂 菜緒",
}

var nameConversions = map[string]string{
	"潮紗理菜":  "潮 紗理菜",
	"影山優佳":  "影山 優佳",
	"加藤史帆":  "加藤 史帆",
	"齊藤京子":  "齊藤 京子",
	"佐々木久美": "佐々木 久美",
	"佐々木美玲": "佐々木 美玲",
	"高瀬愛奈":  "高瀬 愛奈",
	"高本彩花":  "高本 彩花",
	"東村芽依":  "東村 芽依",
	"金村美玖":  "金村 美玖",
	"河田陽菜":  "河田 陽菜",
	"小坂菜緒":  "小坂 菜緒",
	"富田鈴花":  "富田 鈴花",
	"丹生明里":  "丹生 明里",
	"濱岸ひより": "濱岸 ひより",
	"松田好花":  "松田 好花",
	"宮田愛萌":  "宮田 愛萌",
	"渡邉美穂":  "渡邉 美穂",
	"上村ひなの": "上村 ひなの",
	"髙橋未来虹": "髙橋 未来虹",
	"森本茉莉":  "森本 茉莉",
	"山口陽世":  "山口 陽世",
	"柿崎芽実":  "柿崎 芽実",
	"井口眞緒":  "井口 眞緒",
}
