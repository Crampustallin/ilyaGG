DROP TABLE IF EXISTS users;

CREATE TABLE users
(
	user_id serial PRIMARY KEY,
	user_name text NOT NULL,
	pass_hash text NOT NULL,
	user_score int NOT NULL,
	user_pic BYTEA,
	user_status text
);

ALTER TABLE users OWNER TO "ilyaGG"

