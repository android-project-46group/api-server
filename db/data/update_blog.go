package data

import (
	"fmt"

	models "github.com/android-project-46group/api-server/db/my_models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func UpdateBlogs() {
	DB, err := DbInit()
	if err != nil {
		fmt.Println("connection error")
	}
	defer DB.Close()

	if err := DB.Ping(); err != nil {
		fmt.Println("PingError")
	}

	groups := []string{"nogizaka", "sakurazaka", "hinatazaka"}
	for _, group := range groups {
		blogs := LoadBlogInfoFile(group)
		for _, blog := range blogs {

			name := blog.Name
			// 名前が blog に使われているものと違う場合、ここで修正を加える
			if blog.Name == "saitoasuka" {
				name = "saitouasuka"
			}

			mId, err := FindMemberByNameEn(name)
			if err != nil {
				fmt.Println(blog.Name + "が DB に見つかりませんでした")
				continue
			}

			g, err := FindBlogByMemberId(mId)
			if err != nil {
				fmt.Printf("member_id %d の blog が見つかりませんでした？", mId)
				fmt.Println(err)
				continue
			}

			// Update information
			g.LastUpdatedAt = blog.LastUpdatedAt
			fmt.Println(g)
			if _, err = g.Update(Ctx, DB, boil.Infer()); err != nil {
				fmt.Printf("member_id %d の blog の更新に失敗しました", mId)
				fmt.Println(err)
			}

		}
	}
}

func FindBlogByMemberId(id int) (*models.Blog, error) {
	g, err := models.Blogs(
		qm.Where("member_id = ?", id),
	).One(Ctx, DB)

	return g, err
}
