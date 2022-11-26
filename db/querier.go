package db

import (
	models "github.com/android-project-46group/api-server/db/my_models"
)

type Querier interface {
	InsertApiKey(key string) error
	FindApiKeyByName(key string) (*models.APIKey, error)
	GetAllBlogs(groupName string) ([]MemberBlogBind, error)
	GetAllFormations(groupName string) ([]PositionSongsBind, error)
	GetFormations(groupName string) ([]PositionSongsBind, error)
	FindGroupByName(groupName string) (*models.Group, error)
	FindLocaleByName(name string) (*models.Locale, error)
	GetAllMemberInfos(groupName string, localeId int) ([]MemberInfoBind, error)
	GetAllPositions(groupName string) ([]MemberInfoBind, error)
	GetPositionFromTitle(title string) ([]PositionMemberBind, error)
	GetAllSongs(groupName string) (models.SongSlice, error)
	GetCenter(title string) []string
}
