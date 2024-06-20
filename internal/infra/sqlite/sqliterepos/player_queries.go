package sqliterepos

const (
	findAllPlayersQuery = `
  SELECT id, nickname, created_at, updated_at
  FROM players;
  `

	createPlayerQuery = `
  INSERT INTO players (nickname, created_at, updated_at)
  VALUES ($1, $2, $3);
  `
)
