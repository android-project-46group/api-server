package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/android-project-46group/api-server/db/my_models"
)

// Custom struct using two generated structs
type PositionsInfoBind struct {
	models.MemberInfo `boil:",bind"`
	models.Member     `boil:",bind"`
}

func (q *SqlQuerier) GetAllPositions(groupName string) ([]MemberInfoBind, error) {
	var jMember []MemberInfoBind
	err := models.Members(
		qm.Select("member_infos.*", "members.*"),
		qm.InnerJoin("member_infos on members.member_id = member_infos.member_id"),
		qm.InnerJoin("groups on groups.group_id = members.group_id"),
		qm.Where("groups.group_name = ?", groupName),
	).Bind(q.ctx, q.DB, &jMember)
	return jMember, fmt.Errorf("GetAllPositions: %w", err)
}

type PositionMemberBind struct {
	models.Position   `boil:",bind"`
	models.Member     `boil:",bind"`
	models.MemberInfo `boil:",bind"`
}

func (q *SqlQuerier) GetPositionFromTitle(title string) ([]PositionMemberBind, error) {
	var pMs []PositionMemberBind
	err := models.Positions(
		qm.Select("positions.*", "members.*", "member_infos.*"),
		qm.InnerJoin("members on members.member_id = positions.member_id"),
		qm.InnerJoin("member_infos on member_infos.member_id = members.member_id"),
		qm.InnerJoin("songs on songs.song_id = positions.song_id"),
		qm.Where("songs.title = ?", title),
	).Bind(q.ctx, q.DB, &pMs)
	return pMs, fmt.Errorf("GetPositionFromTitle: %w", err)
}
