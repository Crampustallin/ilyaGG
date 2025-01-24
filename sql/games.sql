DROP TABLE IF EXISTS games;

CREATE TABLE games 
(
	game_title text PRIMARY KEY,
	game_art_url text,
);

ALTER TABLE games OWNER TO "ilyaGG"
