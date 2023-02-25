package data

import (
	"fmt"
	"os"

	models "github.com/android-project-46group/api-server/db/my_models"
	_ "github.com/lib/pq"
)

func InsertMain() {
	DB, err := DbInit()
	if err != nil {
		panic("connection error")
	}
	defer DB.Close()

	if err := DB.Ping(); err != nil {
		panic("db PingError")
	}

	if !shouldInsert() {
		fmt.Println("Not need to insert")
		os.Exit(0)
	}

	// groups table
	InsertGroups()

	// members table
	fmt.Println("Insert members")
	InsertMembersNogi()
	InsertMembersSakura()
	InsertMembersHinata()
	UpdateGraduatedMembersNogi()

	// locales
	InsertLocales()

	// member_infos table
	fmt.Println("Insert member_infos")
	InsertMemberInfosHinata()
	InsertMemberInfosNogi()
	InsertMemberInfosSakura()
	InsertGraduatedMemberInfos()
	UpdateGraduatedMembersSakura()

	// tags table
	InsertTags()

	// formations table
	InsertFormations()

	// songs table
	InsertSongs("hinatazaka")

	// positoins table
	InsertPositions()

	// tags table
	InsertMemberTags()

	// blogs table
	fmt.Println("Insert blogs")
	InsertBlogs()

	// Default API key
	InsertApiKey()

	fmt.Println("Insert Done!")
}

func shouldInsert() bool {
	count, err := models.Members().Count(Ctx, DB)
	if err != nil {
		return false
	}
	return count == 0
}
