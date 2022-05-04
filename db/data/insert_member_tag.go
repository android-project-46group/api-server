package data

import (
	"fmt"

	models "github.com/android-project-46group/api-server/db/my_models"
	_ "github.com/lib/pq"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func InsertMemberTags() {
	members, err := models.Members(qm.Select("*")).All(Ctx, DB)
	if err != nil {
		fmt.Println("Error while getting all members from db")
		return
	}

	// 1st tag
	// 今回は、全てのメンバーに可愛いタグを入れる
	// Tagid を入手する
	kTagId, kErr := GetTagIdByName("可愛い")
	if kErr != nil {
		fmt.Println("Error while getting 可愛い tag from db")
		return
	}

	for _, member := range members {

		// すでに存在する場合はスキップ
		if ExistMemberTag(member.MemberID, kTagId) {
			continue
		}

		m := &models.MemberTag{
			MemberID: member.MemberID,
			TagID:    kTagId,
		}

		m.Insert(Ctx, DB, boil.Infer())
	}

	// 2nd tag
	azatoMembers := []string{
		"宮田 愛萌", "渡邉 美穂", "齊藤 京子", "加藤 史帆", "影山 優佳",
	}
	// Tagid を入手する
	aTagId, aErr := GetTagIdByName("あざとい")
	if aErr != nil {
		fmt.Println("Error while getting 可愛い tag from db")
		return
	}
	for _, mem := range azatoMembers {
		mId, _ := FindMemberIdByName(mem)

		// すでに存在する場合はスキップ
		if ExistMemberTag(mId, aTagId) {
			continue
		}
		m := &models.MemberTag{
			MemberID: mId,
			TagID:    aTagId,
		}

		m.Insert(Ctx, DB, boil.Infer())
	}
}

func ExistMemberTag(mId int, tId int) bool {
	_, err := models.MemberTags(
		qm.Where("member_id = ?", mId),
		qm.Where("tag_id = ?", tId),
	).One(Ctx, DB)
	if err != nil {
		return false
	}
	return true
}
