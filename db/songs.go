package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	models "github.com/android-project-46group/api-server/db/my_models"
)

func (q *SqlQuerier) GetAllSongs(ctx context.Context, groupName string) (models.SongSlice, error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.GetAllSongs")
	defer span.Finish()

	s, err := models.Songs(
		qm.InnerJoin("groups on groups.group_id = songs.group_id"),
		qm.Where("groups.group_name = ?", groupName),
	).All(ctx, q.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to GetAllSongs: %w", err)
	}
	return s, nil
}

// Member まで持ってきているのは過剰かも。
// NameJa など string でやろうとしたらできなかった
type PositionMember struct {
	models.Position `boil:",bind"`
	models.Member   `boil:",bind"`
}

func (q *SqlQuerier) GetCenter(ctx context.Context, title string) []string {

	span, ctx := tracer.StartSpanFromContext(ctx, "db.GetCenter")
	defer span.Finish()

	// センターは複数人になりうる
	var pms []PositionSongsBind
	var res []string
	err := models.Positions(
		qm.InnerJoin("songs on songs.song_id = positions.song_id"),
		qm.Where("positions.is_center = true"),
		qm.Where("groups.group_name = ?", title),
	).Bind(ctx, q.DB, &pms)
	if err != nil {
		return res
	}
	for _, pm := range pms {
		res = append(res, pm.Member.NameJa)
	}

	q.logger.Debugf(ctx, "GetCenter got: ", res)
	return res
}
