package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/android-project-46group/api-server/db/my_models"
)

type MemberBlogBind struct {
	models.Blog   `boil:",bind"`
	models.Member `boil:",bind"`
}

func (q *SqlQuerier) GetAllBlogs(groupName string) ([]MemberBlogBind, error) {

	g, _ := q.FindGroupByName(groupName)

	var mBlog []MemberBlogBind

	err := models.Blogs(
		qm.Select("blogs.*", "members.*"),
		qm.InnerJoin("members on members.member_id = blogs.member_id"),
		qm.Where("members.group_id = ?", g.GroupID),
	).Bind(q.ctx, q.DB, &mBlog)

	return mBlog, fmt.Errorf("GetAllBlogs: %w", err)
}
