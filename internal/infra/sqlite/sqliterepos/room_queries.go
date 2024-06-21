package sqliterepos

const (
	findAllRoomsQuery = `
  SELECT
    roo.id,
    roo.name,
    roo.code,
    roo.max_players,
    roo.min_players,
    count(rp.player_id) AS total_players,
    roo.created_at,
    roo.updated_at
  FROM rooms roo
  JOIN room_players rp
    ON roo.id = rp.room_id
  GROUP BY roo.id;
  `

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
