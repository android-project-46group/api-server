package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	models "github.com/android-project-46group/api-server/db/my_models"
	"github.com/android-project-46group/api-server/util"
)

// Custom struct using two generated structs
type MemberInfoBind struct {
	models.MemberInfo `boil:",bind"`
	models.Member     `boil:",bind"`
	models.Group      `boil:",bind"`
}

func (q *SqlQuerier) GetAllMemberInfos(ctx context.Context, groupName string, locale int, sortKey util.SortKey, desc bool) ([]MemberInfoBind, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.GetAllMemberInfos")
	defer span.Finish()

	column := keyToDBColumn(sortKey)
	order := "ASC"
	if desc {
		order = "DESC"
	}

	q.logger.Debugf(ctx, "GetAllMemberInfos with groupName %s and locale %d and sortKey %s and order %s", groupName, locale, sortKey, desc)

	var jMember []MemberInfoBind
	err := models.Members(
		qm.Select("member_infos.*", "members.*", "groups.*"),
		qm.InnerJoin("member_infos on members.member_id = member_infos.member_id"),
		qm.InnerJoin("locales on locales.locale_id = member_infos.locale_id AND locales.locale_id = ?", locale),
		qm.InnerJoin("groups on groups.group_id = members.group_id"),
		qm.Where("CASE WHEN ? = '' THEN '1' ELSE groups.group_name = ? END", groupName, groupName),
		qm.OrderBy(column+" "+order),
	).Bind(ctx, q.DB, &jMember)
	if err != nil {
		return nil, fmt.Errorf("GetAllMemberInfos: %w", err)
	}

	q.logger.Debugf(ctx, "GetAllMemberInfos got: ", jMember)
	return jMember, nil
}

func keyToDBColumn(sortKey util.SortKey) string {
	switch sortKey {
	case util.MemberID:
		return models.MemberInfoTableColumns.MemberID
	case util.Birthday:
		return models.MemberInfoTableColumns.Birthday
	case util.Height:
		return models.MemberInfoTableColumns.Height
	default:
		return models.MemberInfoTableColumns.MemberID
	}
}
