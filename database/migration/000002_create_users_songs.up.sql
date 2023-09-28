CREATE TABLE users_songs (
  id bigint NOT NULL AUTO_INCREMENT,
  user_id bigint NOT NULL,
  song_id bigint NOT NULL,
  title varchar(255) NOT NULL,
  year int NOT NULL,
  duration bigint NOT NULL,
  data text NOT NULL,
  album_id bigint NOT NULL,
  album_name varchar(255) NOT NULL,
  artist_id bigint NOT NULL,
  artist_name varchar(255) NOT NULL,
  composer varchar(255) DEFAULT NULL,
  album_artist varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY idx_users_songs_user_id (user_id),
  CONSTRAINT fk_users_songs_users_id_users_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);