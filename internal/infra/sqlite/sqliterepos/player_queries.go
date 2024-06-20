package sqliterepos

const (
	findAllPlayersQuery = `
  SELECT id, nickname, created_at, updated_at
  FROM players;
  `

	nicknameExistsQuery = `
  SELECT EXISTS(SELECT 1 FROM players WHERE nickname = $1);
  `

	createPlayerQuery = `
  INSERT INTO players (nickname, created_at, updated_at)
  VALUES ($1, $2, $3);
  `
)
