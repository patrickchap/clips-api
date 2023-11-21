CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "auth0_user_id" varchar(100) UNIQUE NOT NULL,
  "username" varchar(50) UNIQUE NOT NULL,
  "email" varchar(100) UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "videos" (
  "id" BIGSERIAL PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "description" varchar,
  "file_url" varchar(255) NOT NULL,
  "thumbnail_url" varchar(255) NOT NULL,
  "user_id" varchar(100),
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "comments" (
  "id" BIGSERIAL PRIMARY KEY,
  "content" varchar NOT NULL,
  "video_id" bigint,
  "user_id" varchar(100),
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "likes" (
  "id" BIGSERIAL PRIMARY KEY,
  "video_id" bigint,
  "user_id" varchar(100),
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE INDEX ON "users" ("auth0_user_id");

CREATE UNIQUE INDEX ON "users" ("username");

CREATE UNIQUE INDEX ON "users" ("email");

CREATE INDEX ON "videos" ("title");

CREATE INDEX ON "videos" ("user_id");

CREATE INDEX ON "comments" ("video_id");

CREATE INDEX ON "likes" ("video_id");

ALTER TABLE "videos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("auth0_user_id") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("video_id") REFERENCES "videos" ("id") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("auth0_user_id") ON DELETE CASCADE;

ALTER TABLE "likes" ADD FOREIGN KEY ("video_id") REFERENCES "videos" ("id") ON DELETE CASCADE;

ALTER TABLE "likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("auth0_user_id") ON DELETE CASCADE;

