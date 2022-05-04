package data

import (
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
		os.Exit(0)
	}

	// groups table
	InsertGroups()

	// members table
	InsertMembersNogi()
	InsertMembersSakura()
	InsertMembersHinata()

	// member_infos table
	InsertMemberInfosHinata()
	InsertMemberInfosNogi()
	InsertMemberInfosSakura()

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
	InsertBlogs()
}

func shouldInsert() bool {
	count, err := models.Members().Count(Ctx, DB)
	if err != nil {
		return false
	}
	return count == 0
}
