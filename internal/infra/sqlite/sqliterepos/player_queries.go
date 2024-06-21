package sqliterepos

const (
	findAllPlayersQuery = `
  SELECT id, nickname, created_at, updated_at
  FROM players;
  `

	nicknameExistsQuery = `
  SELECT EXISTS(SELECT 1 FROM players WHERE nickname = $1);
  `

	existsPlayerQuery = `
  SELECT EXISTS(SELECT 1 FROM players WHERE id = $1);
  `

	createPlayerQuery = `
  INSERT INTO players (id, nickname, created_at, updated_at)
  VALUES ($1, $2, $3, $4);
  `
)
