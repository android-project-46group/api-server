package db

import (
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "web/db/my_models"
)

// Custom struct using two generated structs
type PositionsInfoBind struct {
	models.MemberInfo	`boil:",bind"`
	models.Member		`boil:",bind"`
}

func GetAllPositions(groupName string) ([]MemberInfoBind, error) {
	var jMember []MemberInfoBind
	err := models.Members(
		qm.Select("member_infos.*", "members.*"),
		qm.InnerJoin("member_infos on members.member_id = member_infos.member_id"),
		qm.InnerJoin("groups on groups.group_id = members.group_id"),
		qm.Where("groups.group_name = ?", groupName),
	).Bind(Ctx, DB, &jMember)
	return jMember, err
}
