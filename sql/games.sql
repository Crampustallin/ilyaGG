DROP TABLE IF EXISTS games;

CREATE TABLE games 
(
	game_title text PRIMARY KEY,
	game_art_url text
	add_date date DEFAULT CURRENT_DATE,
);

ALTER TABLE games OWNER TO "ilyaGG"
