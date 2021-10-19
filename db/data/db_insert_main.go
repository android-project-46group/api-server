package data

import (
	"fmt"
	
	_ "github.com/lib/pq"
)

func InsertMain() {
	DB, err := DbInit()
	if err != nil {
		fmt.Println("connection error")
	}
	defer DB.Close()

	if err := DB.Ping(); err != nil {
		fmt.Println("PingError")
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
}
