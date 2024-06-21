package sqliterepos

import (
	"context"

	"github.com/joaovds/stop-go/internal/domain/player"
	"github.com/joaovds/stop-go/internal/infra/sqlite"
	"github.com/joaovds/stop-go/pkg/errs"
)

type PlayerRepository struct {
	db *sqlite.DB
}

func NewPlayerRepository(db *sqlite.DB) *PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

// ----- ... -----

func (p *PlayerRepository) FindAll(ctx context.Context) ([]*player.Player, *errs.Error) {
	rows, err := p.db.Query(findAllPlayersQuery)
	if err != nil {
		return nil, errs.NewError("error finding all players")
	}
	defer rows.Close()

	players := make([]*player.Player, 0)
	for rows.Next() {
		var player player.Player
		err := rows.Scan(&player.ID, &player.Nickname, &player.CreatedAt, &player.UpdatedAt)
		if err != nil {
			return nil, errs.NewError("error scanning players")
		}

		players = append(players, &player)
	}

	return players, nil
}

// ----- ... -----

func (p *PlayerRepository) NicknameExists(ctx context.Context, nickname string) (bool, *errs.Error) {
	var exists bool
	err := p.db.QueryRow(nicknameExistsQuery, nickname).Scan(&exists)
	if err != nil {
		return false, errs.NewError("error checking if nickname exists")
	}

	return exists, nil
}

// ----- ... -----

func (p *PlayerRepository) Exists(ctx context.Context, id string) (bool, *errs.Error) {
	var exists bool
	err := p.db.QueryRow(existsPlayerQuery, id).Scan(&exists)
	if err != nil {
		return false, errs.NewError("error checking if player exists")
	}

	return exists, nil
}

// ----- ... -----

func (p *PlayerRepository) Create(ctx context.Context, player *player.Player) *errs.Error {
	_, err := p.db.Exec(createPlayerQuery, player.ID, player.Nickname, player.CreatedAt, player.UpdatedAt)
	if err != nil {
		return errs.NewError("error creating player")
	}

	return nil
}
