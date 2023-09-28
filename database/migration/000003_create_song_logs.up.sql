CREATE TABLE song_logs (
  id int NOT NULL,
  user_id bigint NOT NULL,
  song_id bigint NOT NULL,
  created_at timestamp NOT NULL,
  duration_played int NOT NULL,
  KEY idx_song_logs_user_id (user_id),
  KEY idx_song_logs_song_id (song_id),
  CONSTRAINT fk_song_logs_song_id_users_songs_id FOREIGN KEY (song_id) REFERENCES users_songs (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_song_logs_user_id_users_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);