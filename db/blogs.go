package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	tracer "github.com/opentracing/opentracing-go"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/android-project-46group/api-server/db/my_models"
)

type MemberBlogBind struct {
	models.Blog   `boil:",bind"`
	models.Member `boil:",bind"`
}

func (q *SqlQuerier) GetAllBlogs(ctx context.Context, groupName string) ([]MemberBlogBind, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.GetAllBlogs")
	defer span.Finish()

	var mBlog []MemberBlogBind

	err := models.Blogs(
		qm.Select("blogs.*", "members.*"),
		qm.InnerJoin("members on members.member_id = blogs.member_id"),
		qm.InnerJoin("groups on groups.group_id = members.group_id"),
		qm.Where("CASE WHEN ? = '' THEN '1' ELSE groups.group_name = ? END", groupName, groupName),
	).Bind(ctx, q.DB, &mBlog)

	if err != nil {
		return nil, fmt.Errorf("failed to GetAllBlogs: %w", err)
	}

	q.logger.Debugf(ctx, "GetAllBlogs got: ", mBlog)
	return mBlog, nil
}
