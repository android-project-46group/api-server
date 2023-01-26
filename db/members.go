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
type MemberInfoBind struct {
	models.MemberInfo `boil:",bind"`
	models.Member     `boil:",bind"`
}

func (q *SqlQuerier) GetAllMemberInfos(ctx context.Context, groupName string, locale int) ([]MemberInfoBind, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.GetAllMemberInfos")
	defer span.Finish()

	var jMember []MemberInfoBind
	err := models.Members(
		qm.Select("member_infos.*", "members.*"),
		qm.InnerJoin("member_infos on members.member_id = member_infos.member_id"),
		qm.InnerJoin("locales on locales.locale_id = member_infos.locale_id AND locales.locale_id = ?", locale),
		qm.InnerJoin("groups on groups.group_id = members.group_id"),
		qm.Where("groups.group_name = ?", groupName),
	).Bind(ctx, q.DB, &jMember)
	if err != nil {
		return nil, fmt.Errorf("GetAllMemberInfos: %w", err)
	}
	return jMember, nil
}
