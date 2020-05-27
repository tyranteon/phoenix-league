CREATE TABLE USERS (
	user_id BIGINT PRIMARY KEY,
	profile_json TEXT,
	mmr INTEGER,
	is_vetted BOOLEAN,
	override_vetting BOOLEAN
)