package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	models "github.com/android-project-46group/api-server/db/my_models"
)

// Custom struct using two generated structs
type PositionsInfoBind struct {
	models.MemberInfo `boil:",bind"`
	models.Member     `boil:",bind"`
}

func (q *SqlQuerier) GetAllPositions(ctx context.Context, groupName string) ([]MemberInfoBind, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.GetAllPositions")
	defer span.Finish()

	var jMember []MemberInfoBind
	err := models.Members(
		qm.Select("member_infos.*", "members.*"),
		qm.InnerJoin("member_infos on members.member_id = member_infos.member_id"),
		qm.InnerJoin("groups on groups.group_id = members.group_id"),
		qm.Where("groups.group_name = ?", groupName),
	).Bind(ctx, q.DB, &jMember)
	return jMember, fmt.Errorf("GetAllPositions: %w", err)
}

type PositionMemberBind struct {
	models.Position   `boil:",bind"`
	models.Member     `boil:",bind"`
	models.MemberInfo `boil:",bind"`
}

func (q *SqlQuerier) GetPositionFromTitle(ctx context.Context, title string) ([]PositionMemberBind, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.GetPositionFromTitle")
	defer span.Finish()

	var pMs []PositionMemberBind
	err := models.Positions(
		qm.Select("positions.*", "members.*", "member_infos.*"),
		qm.InnerJoin("members on members.member_id = positions.member_id"),
		qm.InnerJoin("member_infos on member_infos.member_id = members.member_id"),
		qm.InnerJoin("songs on songs.song_id = positions.song_id"),
		qm.Where("songs.title = ?", title),
	).Bind(ctx, q.DB, &pMs)
	if err != nil {
		return nil, fmt.Errorf("failed to GetPositionFromTitle: %w", err)
	}

	q.logger.Debugf(ctx, "GetPositionFromTitle got: ", pMs)
	return pMs, nil
}
