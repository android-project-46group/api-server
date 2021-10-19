package db

import (
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "web/db/my_models"
	// "fmt"
)



// Member まで持ってきているのは過剰かも。
// NameJa など string でやろうとしたらできなかった
type PositionSongsBind struct {
	models.Position	`boil:",bind"`
	models.Song		`boil:",bind"`
	models.Member	`boil:",bind"`
}


func GetAllFormations(groupName string) ([]PositionSongsBind, error) {

	var sPositions []PositionSongsBind

	g, gErr := FindGroupByName(groupName)
	if gErr != nil {
		return sPositions, gErr
	}
	
	err := models.Positions(
		qm.Select("positions.*", "songs.*", "members.*"),
		qm.InnerJoin("songs on songs.song_id = positions.song_id"),
		qm.InnerJoin("members on members.member_id = positions.member_id"),
		qm.Where("songs.group_id = ?", g.GroupID),
		qm.OrderBy("songs.song_id DESC"),
	).Bind(Ctx, DB, &sPositions)
	return sPositions, err
}
