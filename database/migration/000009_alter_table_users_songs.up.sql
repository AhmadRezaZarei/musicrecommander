ALTER TABLE users_songs 
    ADD identified_song_id BIGINT NULL, ADD INDEX (identified_song_id);

ALTER TABLE users_songs 
    ADD CONSTRAINT fk_users_songs_identified_song_id_identified_songs_id FOREIGN KEY (identified_song_id) REFERENCES identified_songs(id) ON DELETE RESTRICT ON UPDATE RESTRICT;