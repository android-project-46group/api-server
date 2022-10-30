package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/android-project-46group/api-server/db/my_models"
)

// Custom struct using two generated structs
type MemberInfoBind struct {
	models.MemberInfo `boil:",bind"`
	models.Member     `boil:",bind"`
}

func (q *SqlQuerier) GetAllMemberInfos(groupName string, locale int) ([]MemberInfoBind, error) {
	var jMember []MemberInfoBind
	err := models.Members(
		qm.Select("member_infos.*", "members.*"),
		qm.InnerJoin("member_infos on members.member_id = member_infos.member_id"),
		qm.InnerJoin("locales on locales.locale_id = member_infos.locale_id AND locales.locale_id = ?", locale),
		qm.InnerJoin("groups on groups.group_id = members.group_id"),
		qm.Where("groups.group_name = ?", groupName),
	).Bind(q.ctx, q.DB, &jMember)
	return jMember, fmt.Errorf("GetAllMemberInfos: %w", err)
}
