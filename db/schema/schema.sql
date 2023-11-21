
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  auth0_user_id varchar(100) UNIQUE NOT NULL,
  username varchar(50) UNIQUE NOT NULL,
  email varchar(100) UNIQUE NOT NULL,
  created_at timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE videos (
  id BIGSERIAL PRIMARY KEY,
  title varchar(255) NOT NULL,
  description varchar NOT NULL,
  file_url varchar(255) NOT NULL,
  thumbnail_url varchar(255) NOT NULL,
  user_id varchar(100) NOT NULL,
  created_at timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE comments (
  id BIGSERIAL PRIMARY KEY,
  content varchar NOT NULL,
  video_id bigint,
  user_id varchar(100) NOT NULL,
  created_at timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE likes (
  id BIGSERIAL PRIMARY KEY,
  video_id bigint,
  user_id varchar(100) NOT NULL,
  created_at timestamp DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE categories (
  id BIGSERIAL PRIMARY KEY,
  name varchar(50) UNIQUE NOT NULL
);

CREATE TABLE video_categories (
  video_id bigint NOT NULL,
  category_id bigint NOT NULL
);
