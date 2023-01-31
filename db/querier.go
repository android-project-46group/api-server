package db

import (
	"context"

	models "github.com/android-project-46group/api-server/db/my_models"
)

type Querier interface {
	InsertApiKey(ctx context.Context, key string) error
	FindApiKeyByName(ctx context.Context, key string) (*models.APIKey, error)
	GetAllGroups(ctx context.Context) ([]*models.Group, error)
	GetAllBlogs(ctx context.Context, groupName string) ([]MemberBlogBind, error)
	GetAllFormations(ctx context.Context, groupName string) ([]PositionSongsBind, error)
	GetFormations(ctx context.Context, groupName string) ([]PositionSongsBind, error)
	FindGroupByName(ctx context.Context, groupName string) (*models.Group, error)
	FindLocaleByName(ctx context.Context, name string) (*models.Locale, error)
	GetAllMemberInfos(ctx context.Context, groupName string, localeId int) ([]MemberInfoBind, error)
	GetAllPositions(ctx context.Context, groupName string) ([]MemberInfoBind, error)
	GetPositionFromTitle(ctx context.Context, title string) ([]PositionMemberBind, error)
	GetAllSongs(ctx context.Context, groupName string) (models.SongSlice, error)
	GetCenter(ctx context.Context, title string) []string
}
