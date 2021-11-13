package db

import (
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "web/db/my_models"
)

type MemberBlogBind struct {
	models.Blog   `boil:",bind"`
	models.Member `boil:",bind"`
}

func GetAllBlogs(groupName string) ([]MemberBlogBind, error) {

	g, _ := FindGroupByName(groupName)

	var mBlog []MemberBlogBind

	err := models.Blogs(
		qm.Select("blogs.*", "members.*"),
		qm.InnerJoin("members on members.member_id = blogs.member_id"),
		qm.Where("members.group_id = ?", g.GroupID),
	).Bind(Ctx, DB, &mBlog)

	return mBlog, err
}