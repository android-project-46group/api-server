package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	models "github.com/android-project-46group/api-server/db/my_models"
)

type MemberBlogBind struct {
	models.Blog   `boil:",bind"`
	models.Member `boil:",bind"`
}

func (q *SqlQuerier) GetAllBlogs(ctx context.Context, groupName string) ([]MemberBlogBind, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.GetAllBlogs")
	defer span.Finish()

	g, _ := q.FindGroupByName(ctx , groupName)

	var mBlog []MemberBlogBind

	err := models.Blogs(
		qm.Select("blogs.*", "members.*"),
		qm.InnerJoin("members on members.member_id = blogs.member_id"),
		qm.Where("members.group_id = ?", g.GroupID),
	).Bind(ctx, q.DB, &mBlog)

	return mBlog, fmt.Errorf("GetAllBlogs: %w", err)
}
