CREATE TABLE "categories" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar(50) UNIQUE NOT NULL
);

CREATE TABLE "video_categories" (
  "video_id" bigint NOT NULL,
  "category_id" bigint NOT NULL
);

CREATE UNIQUE INDEX ON "categories" ("name");

CREATE INDEX ON "video_categories" ("video_id");

CREATE INDEX ON "video_categories" ("category_id");

ALTER TABLE "video_categories" ADD FOREIGN KEY ("video_id") REFERENCES "videos" ("id");

ALTER TABLE "video_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
