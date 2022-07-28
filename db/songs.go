package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/android-project-46group/api-server/db/my_models"
)

func (q *SqlQuerier) GetAllSongs(groupName string) (models.SongSlice, error) {
	s, err := models.Songs(
		qm.InnerJoin("groups on groups.group_id = songs.group_id"),
		qm.Where("groups.group_name = ?", groupName),
	).All(q.ctx, q.DB)
	return s, fmt.Errorf("GetAllSongs: %w", err)
}

// Member まで持ってきているのは過剰かも。
// NameJa など string でやろうとしたらできなかった
type PositionMember struct {
	models.Position `boil:",bind"`
	models.Member   `boil:",bind"`
}

func (q *SqlQuerier) GetCenter(title string) []string {
	// センターは複数人になりうる
	var pms []PositionSongsBind
	var res []string
	err := models.Positions(
		qm.InnerJoin("songs on songs.song_id = positions.song_id"),
		qm.Where("positions.is_center = true"),
		qm.Where("groups.group_name = ?", title),
	).Bind(q.ctx, q.DB, &pms)
	if err != nil {
		return res
	}
	for _, pm := range pms {
		res = append(res, pm.Member.NameJa)
	}
	return res
}
