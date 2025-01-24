DROP TABLE IF EXISTS user_games;

CREATE TABLE user_games 
(
	user_id int NOT NULL, -- primary key from users
	game_title text NOT NULL, -- primary key from games
	game_rating int, 
	game_review text
);

ALTER TABLE user_games OWNER TO "ilyaGG"
