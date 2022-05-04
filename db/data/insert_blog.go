package data

import (
	"fmt"

	models "web/db/my_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func InsertBlogs() {

	groups := []string{"nogizaka", "sakurazaka", "hinatazaka"}
	for _, group := range groups {
		blogs := LoadBlogInfoFile(group)
		for _, blog := range blogs {
			name := blog.Name
			if blog.Name == "saitoasuka" {
				name = "saitouasuka"
			}
			mId, err := FindMemberByNameEn(name)
			if err != nil {
				fmt.Println(blog.Name + "が DB に見つかりませんでした")
				continue
			}

			b := &models.Blog{
				MemberID:      mId,
				BlogURL:       blog.BlogUrl,
				LastBlogImg:   blog.LastBlogImg,
				LastUpdatedAt: blog.LastUpdatedAt,
			}
			b.Insert(Ctx, DB, boil.Infer())
		}
	}
}
