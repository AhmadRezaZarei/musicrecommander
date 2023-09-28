ALTER TABLE identified_songs ADD id_in_recognize_service VARCHAR(255) NOT NULL, 
    ADD INDEX idx_identified_songs_id_in_recognize_service (id_in_recognize_service);