package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/android-project-46group/api-server/db/my_models"
	// "fmt"
)

// Member まで持ってきているのは過剰かも。
// NameJa など string でやろうとしたらできなかった
type PositionSongsBind struct {
	models.Position `boil:",bind"`
	models.Song     `boil:",bind"`
	models.Member   `boil:",bind"`
}

func (q *SqlQuerier) GetAllFormations(ctx context.Context, groupName string) ([]PositionSongsBind, error) {

	var sPositions []PositionSongsBind

	g, gErr := q.FindGroupByName(ctx, groupName)
	if gErr != nil {
		return sPositions, fmt.Errorf("GetAllFormations: %w", gErr)
	}

	err := models.Positions(
		qm.Select("positions.*", "songs.*", "members.*"),
		qm.InnerJoin("songs on songs.song_id = positions.song_id"),
		qm.InnerJoin("members on members.member_id = positions.member_id"),
		qm.Where("songs.group_id = ?", g.GroupID),
		qm.OrderBy("songs.song_id DESC"),
	).Bind(ctx, q.DB, &sPositions)
	return sPositions, fmt.Errorf("GetAllFormations: %w", err)
}

// Custom struct using two generated structs
type SongFormationsBind struct {
	models.Song      `boil:",bind"`
	models.Formation `boil:",bind"`
	Center           string `json:"center" boil:",bind"`
}

func (q *SqlQuerier) GetFormations(ctx context.Context, groupName string) ([]PositionSongsBind, error) {
	var jMember []PositionSongsBind
	err := models.Positions(
		qm.Select("positions.*", "songs.*", "members.name_ja"),
		qm.InnerJoin("songs on songs.song_id = positions.song_id"),
		qm.InnerJoin("members on members.member_id = positions.member_id"),
		qm.Where("groups.group_name = ?", groupName),
	).Bind(ctx, q.DB, &jMember)
	return jMember, fmt.Errorf("GetFormations: %w", err)
}
