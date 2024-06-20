// for sqlite

CREATE TABLE "players" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "nickname" VARCHAR(255) NOT NULL UNIQUE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);