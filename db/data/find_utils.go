package data

import (
	"fmt"

	models "github.com/android-project-46group/api-server/db/my_models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// https://kokoichi0206.mydns.jp/imgs/nogi/iwamotorenka.jpeg
func FindMemberByNameEn(nameEn string) (int, error) {
	IMG_URL_BASE := "https://kokoichi0206.mydns.jp/imgs/"

	member_info, err := models.MemberInfos(
		qm.Select("*"),
		qm.Where("img_url=?", fmt.Sprintf("%s%s/%s.jpeg", IMG_URL_BASE, "nogi", nameEn)),
		qm.Or("img_url=?", fmt.Sprintf("%s%s/%s.jpeg", IMG_URL_BASE, "sakura", nameEn)),
		qm.Or("img_url=?", fmt.Sprintf("%s%s/%s.jpeg", IMG_URL_BASE, "hinata", nameEn)),
	).One(Ctx, DB)
	if err != nil {
		return 0, err
	}

	return member_info.MemberID, err
}
