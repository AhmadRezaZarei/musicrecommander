CREATE TABLE identified_songs (
  id bigint NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  meta json NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);