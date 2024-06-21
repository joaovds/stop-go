-- for sqlite

CREATE TABLE "players" (
  "id" TEXT PRIMARY KEY,
  "nickname" VARCHAR(255) NOT NULL UNIQUE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "rooms" (
  "id" TEXT PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL UNIQUE,
  "code" VARCHAR(6) NOT NULL UNIQUE,
  "max_players" INTEGER NOT NULL DEFAULT 10,
  "min_players" INTEGER NOT NULL DEFAULT 2,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX "idx_rooms_code" ON "rooms" ("code");

CREATE TABLE "room_players" (
  "room_id" TEXT NOT NULL,
  "player_id" TEXT NOT NULL,
  "role" TEXT CHECK(role IN ('HOST', 'ADMIN', 'MEMBER')) NOT NULL DEFAULT 'MEMBER',
  "score" INTEGER NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("room_id", "player_id"),
  FOREIGN KEY ("room_id") REFERENCES "rooms" ("id"),
  FOREIGN KEY ("player_id") REFERENCES "players" ("id")
);

CREATE INDEX "idx_room_players_room_id" ON "room_players" ("room_id");
