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

type PositionMemberBind struct {
	models.Position	`boil:",bind"`
	models.Member	`boil:",bind"`
}

func GetPositionFromTitle(title string) ([]PositionMemberBind, error) {
	var pMs []PositionMemberBind
	err := models.Positions(
		qm.Select("positions.*", "members.*"),
		qm.InnerJoin("members on members.member_id = positions.member_id"),
		qm.InnerJoin("songs on songs.song_id = positions.song_id"),
		qm.Where("songs.title = ?", title),
	).Bind(Ctx, DB, &pMs)
	return pMs, err
}
