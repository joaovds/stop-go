package sqliterepos

const (
	createRoomQuery = `
  INSERT INTO rooms (id, name, code, max_players, min_players, created_at, updated_at)
  VALUES ($1, $2, $3, $4, $5, $6, $7);
  `

	createRoomPlayerQuery = `
  INSERT INTO room_players (room_id, player_id, role, created_at, updated_at)
  VALUES ($1, $2, $3, $4, $5);
  `

	nameExistsQuery = `
  SELECT EXISTS(SELECT 1 FROM rooms WHERE name = $1);
  `
)
