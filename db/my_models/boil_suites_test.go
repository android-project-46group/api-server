// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("APIKeys", testAPIKeys)
	t.Run("Blogs", testBlogs)
	t.Run("Formations", testFormations)
	t.Run("Groups", testGroups)
	t.Run("MemberInfos", testMemberInfos)
	t.Run("MemberTags", testMemberTags)
	t.Run("Members", testMembers)
	t.Run("Positions", testPositions)
	t.Run("SchemaMigrations", testSchemaMigrations)
	t.Run("Songs", testSongs)
	t.Run("Tags", testTags)
}

func TestDelete(t *testing.T) {
	t.Run("APIKeys", testAPIKeysDelete)
	t.Run("Blogs", testBlogsDelete)
	t.Run("Formations", testFormationsDelete)
	t.Run("Groups", testGroupsDelete)
	t.Run("MemberInfos", testMemberInfosDelete)
	t.Run("MemberTags", testMemberTagsDelete)
	t.Run("Members", testMembersDelete)
	t.Run("Positions", testPositionsDelete)
	t.Run("SchemaMigrations", testSchemaMigrationsDelete)
	t.Run("Songs", testSongsDelete)
	t.Run("Tags", testTagsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("APIKeys", testAPIKeysQueryDeleteAll)
	t.Run("Blogs", testBlogsQueryDeleteAll)
	t.Run("Formations", testFormationsQueryDeleteAll)
	t.Run("Groups", testGroupsQueryDeleteAll)
	t.Run("MemberInfos", testMemberInfosQueryDeleteAll)
	t.Run("MemberTags", testMemberTagsQueryDeleteAll)
	t.Run("Members", testMembersQueryDeleteAll)
	t.Run("Positions", testPositionsQueryDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsQueryDeleteAll)
	t.Run("Songs", testSongsQueryDeleteAll)
	t.Run("Tags", testTagsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("APIKeys", testAPIKeysSliceDeleteAll)
	t.Run("Blogs", testBlogsSliceDeleteAll)
	t.Run("Formations", testFormationsSliceDeleteAll)
	t.Run("Groups", testGroupsSliceDeleteAll)
	t.Run("MemberInfos", testMemberInfosSliceDeleteAll)
	t.Run("MemberTags", testMemberTagsSliceDeleteAll)
	t.Run("Members", testMembersSliceDeleteAll)
	t.Run("Positions", testPositionsSliceDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceDeleteAll)
	t.Run("Songs", testSongsSliceDeleteAll)
	t.Run("Tags", testTagsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("APIKeys", testAPIKeysExists)
	t.Run("Blogs", testBlogsExists)
	t.Run("Formations", testFormationsExists)
	t.Run("Groups", testGroupsExists)
	t.Run("MemberInfos", testMemberInfosExists)
	t.Run("MemberTags", testMemberTagsExists)
	t.Run("Members", testMembersExists)
	t.Run("Positions", testPositionsExists)
	t.Run("SchemaMigrations", testSchemaMigrationsExists)
	t.Run("Songs", testSongsExists)
	t.Run("Tags", testTagsExists)
}

func TestFind(t *testing.T) {
	t.Run("APIKeys", testAPIKeysFind)
	t.Run("Blogs", testBlogsFind)
	t.Run("Formations", testFormationsFind)
	t.Run("Groups", testGroupsFind)
	t.Run("MemberInfos", testMemberInfosFind)
	t.Run("MemberTags", testMemberTagsFind)
	t.Run("Members", testMembersFind)
	t.Run("Positions", testPositionsFind)
	t.Run("SchemaMigrations", testSchemaMigrationsFind)
	t.Run("Songs", testSongsFind)
	t.Run("Tags", testTagsFind)
}

func TestBind(t *testing.T) {
	t.Run("APIKeys", testAPIKeysBind)
	t.Run("Blogs", testBlogsBind)
	t.Run("Formations", testFormationsBind)
	t.Run("Groups", testGroupsBind)
	t.Run("MemberInfos", testMemberInfosBind)
	t.Run("MemberTags", testMemberTagsBind)
	t.Run("Members", testMembersBind)
	t.Run("Positions", testPositionsBind)
	t.Run("SchemaMigrations", testSchemaMigrationsBind)
	t.Run("Songs", testSongsBind)
	t.Run("Tags", testTagsBind)
}

func TestOne(t *testing.T) {
	t.Run("APIKeys", testAPIKeysOne)
	t.Run("Blogs", testBlogsOne)
	t.Run("Formations", testFormationsOne)
	t.Run("Groups", testGroupsOne)
	t.Run("MemberInfos", testMemberInfosOne)
	t.Run("MemberTags", testMemberTagsOne)
	t.Run("Members", testMembersOne)
	t.Run("Positions", testPositionsOne)
	t.Run("SchemaMigrations", testSchemaMigrationsOne)
	t.Run("Songs", testSongsOne)
	t.Run("Tags", testTagsOne)
}

func TestAll(t *testing.T) {
	t.Run("APIKeys", testAPIKeysAll)
	t.Run("Blogs", testBlogsAll)
	t.Run("Formations", testFormationsAll)
	t.Run("Groups", testGroupsAll)
	t.Run("MemberInfos", testMemberInfosAll)
	t.Run("MemberTags", testMemberTagsAll)
	t.Run("Members", testMembersAll)
	t.Run("Positions", testPositionsAll)
	t.Run("SchemaMigrations", testSchemaMigrationsAll)
	t.Run("Songs", testSongsAll)
	t.Run("Tags", testTagsAll)
}

func TestCount(t *testing.T) {
	t.Run("APIKeys", testAPIKeysCount)
	t.Run("Blogs", testBlogsCount)
	t.Run("Formations", testFormationsCount)
	t.Run("Groups", testGroupsCount)
	t.Run("MemberInfos", testMemberInfosCount)
	t.Run("MemberTags", testMemberTagsCount)
	t.Run("Members", testMembersCount)
	t.Run("Positions", testPositionsCount)
	t.Run("SchemaMigrations", testSchemaMigrationsCount)
	t.Run("Songs", testSongsCount)
	t.Run("Tags", testTagsCount)
}

func TestHooks(t *testing.T) {
	t.Run("APIKeys", testAPIKeysHooks)
	t.Run("Blogs", testBlogsHooks)
	t.Run("Formations", testFormationsHooks)
	t.Run("Groups", testGroupsHooks)
	t.Run("MemberInfos", testMemberInfosHooks)
	t.Run("MemberTags", testMemberTagsHooks)
	t.Run("Members", testMembersHooks)
	t.Run("Positions", testPositionsHooks)
	t.Run("SchemaMigrations", testSchemaMigrationsHooks)
	t.Run("Songs", testSongsHooks)
	t.Run("Tags", testTagsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("APIKeys", testAPIKeysInsert)
	t.Run("APIKeys", testAPIKeysInsertWhitelist)
	t.Run("Blogs", testBlogsInsert)
	t.Run("Blogs", testBlogsInsertWhitelist)
	t.Run("Formations", testFormationsInsert)
	t.Run("Formations", testFormationsInsertWhitelist)
	t.Run("Groups", testGroupsInsert)
	t.Run("Groups", testGroupsInsertWhitelist)
	t.Run("MemberInfos", testMemberInfosInsert)
	t.Run("MemberInfos", testMemberInfosInsertWhitelist)
	t.Run("MemberTags", testMemberTagsInsert)
	t.Run("MemberTags", testMemberTagsInsertWhitelist)
	t.Run("Members", testMembersInsert)
	t.Run("Members", testMembersInsertWhitelist)
	t.Run("Positions", testPositionsInsert)
	t.Run("Positions", testPositionsInsertWhitelist)
	t.Run("SchemaMigrations", testSchemaMigrationsInsert)
	t.Run("SchemaMigrations", testSchemaMigrationsInsertWhitelist)
	t.Run("Songs", testSongsInsert)
	t.Run("Songs", testSongsInsertWhitelist)
	t.Run("Tags", testTagsInsert)
	t.Run("Tags", testTagsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("BlogToMemberUsingMember", testBlogToOneMemberUsingMember)
	t.Run("MemberInfoToMemberUsingMember", testMemberInfoToOneMemberUsingMember)
	t.Run("MemberTagToMemberUsingMember", testMemberTagToOneMemberUsingMember)
	t.Run("MemberTagToTagUsingTag", testMemberTagToOneTagUsingTag)
	t.Run("MemberToGroupUsingGroup", testMemberToOneGroupUsingGroup)
	t.Run("PositionToMemberUsingMember", testPositionToOneMemberUsingMember)
	t.Run("PositionToSongUsingSong", testPositionToOneSongUsingSong)
	t.Run("SongToFormationUsingFormation", testSongToOneFormationUsingFormation)
	t.Run("SongToGroupUsingGroup", testSongToOneGroupUsingGroup)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("FormationToSongs", testFormationToManySongs)
	t.Run("GroupToMembers", testGroupToManyMembers)
	t.Run("GroupToSongs", testGroupToManySongs)
	t.Run("MemberToBlogs", testMemberToManyBlogs)
	t.Run("MemberToMemberInfos", testMemberToManyMemberInfos)
	t.Run("MemberToMemberTags", testMemberToManyMemberTags)
	t.Run("MemberToPositions", testMemberToManyPositions)
	t.Run("SongToPositions", testSongToManyPositions)
	t.Run("TagToMemberTags", testTagToManyMemberTags)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("BlogToMemberUsingBlogs", testBlogToOneSetOpMemberUsingMember)
	t.Run("MemberInfoToMemberUsingMemberInfos", testMemberInfoToOneSetOpMemberUsingMember)
	t.Run("MemberTagToMemberUsingMemberTags", testMemberTagToOneSetOpMemberUsingMember)
	t.Run("MemberTagToTagUsingMemberTags", testMemberTagToOneSetOpTagUsingTag)
	t.Run("MemberToGroupUsingMembers", testMemberToOneSetOpGroupUsingGroup)
	t.Run("PositionToMemberUsingPositions", testPositionToOneSetOpMemberUsingMember)
	t.Run("PositionToSongUsingPositions", testPositionToOneSetOpSongUsingSong)
	t.Run("SongToFormationUsingSongs", testSongToOneSetOpFormationUsingFormation)
	t.Run("SongToGroupUsingSongs", testSongToOneSetOpGroupUsingGroup)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("FormationToSongs", testFormationToManyAddOpSongs)
	t.Run("GroupToMembers", testGroupToManyAddOpMembers)
	t.Run("GroupToSongs", testGroupToManyAddOpSongs)
	t.Run("MemberToBlogs", testMemberToManyAddOpBlogs)
	t.Run("MemberToMemberInfos", testMemberToManyAddOpMemberInfos)
	t.Run("MemberToMemberTags", testMemberToManyAddOpMemberTags)
	t.Run("MemberToPositions", testMemberToManyAddOpPositions)
	t.Run("SongToPositions", testSongToManyAddOpPositions)
	t.Run("TagToMemberTags", testTagToManyAddOpMemberTags)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("APIKeys", testAPIKeysReload)
	t.Run("Blogs", testBlogsReload)
	t.Run("Formations", testFormationsReload)
	t.Run("Groups", testGroupsReload)
	t.Run("MemberInfos", testMemberInfosReload)
	t.Run("MemberTags", testMemberTagsReload)
	t.Run("Members", testMembersReload)
	t.Run("Positions", testPositionsReload)
	t.Run("SchemaMigrations", testSchemaMigrationsReload)
	t.Run("Songs", testSongsReload)
	t.Run("Tags", testTagsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("APIKeys", testAPIKeysReloadAll)
	t.Run("Blogs", testBlogsReloadAll)
	t.Run("Formations", testFormationsReloadAll)
	t.Run("Groups", testGroupsReloadAll)
	t.Run("MemberInfos", testMemberInfosReloadAll)
	t.Run("MemberTags", testMemberTagsReloadAll)
	t.Run("Members", testMembersReloadAll)
	t.Run("Positions", testPositionsReloadAll)
	t.Run("SchemaMigrations", testSchemaMigrationsReloadAll)
	t.Run("Songs", testSongsReloadAll)
	t.Run("Tags", testTagsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("APIKeys", testAPIKeysSelect)
	t.Run("Blogs", testBlogsSelect)
	t.Run("Formations", testFormationsSelect)
	t.Run("Groups", testGroupsSelect)
	t.Run("MemberInfos", testMemberInfosSelect)
	t.Run("MemberTags", testMemberTagsSelect)
	t.Run("Members", testMembersSelect)
	t.Run("Positions", testPositionsSelect)
	t.Run("SchemaMigrations", testSchemaMigrationsSelect)
	t.Run("Songs", testSongsSelect)
	t.Run("Tags", testTagsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("APIKeys", testAPIKeysUpdate)
	t.Run("Blogs", testBlogsUpdate)
	t.Run("Formations", testFormationsUpdate)
	t.Run("Groups", testGroupsUpdate)
	t.Run("MemberInfos", testMemberInfosUpdate)
	t.Run("MemberTags", testMemberTagsUpdate)
	t.Run("Members", testMembersUpdate)
	t.Run("Positions", testPositionsUpdate)
	t.Run("SchemaMigrations", testSchemaMigrationsUpdate)
	t.Run("Songs", testSongsUpdate)
	t.Run("Tags", testTagsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("APIKeys", testAPIKeysSliceUpdateAll)
	t.Run("Blogs", testBlogsSliceUpdateAll)
	t.Run("Formations", testFormationsSliceUpdateAll)
	t.Run("Groups", testGroupsSliceUpdateAll)
	t.Run("MemberInfos", testMemberInfosSliceUpdateAll)
	t.Run("MemberTags", testMemberTagsSliceUpdateAll)
	t.Run("Members", testMembersSliceUpdateAll)
	t.Run("Positions", testPositionsSliceUpdateAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceUpdateAll)
	t.Run("Songs", testSongsSliceUpdateAll)
	t.Run("Tags", testTagsSliceUpdateAll)
}
