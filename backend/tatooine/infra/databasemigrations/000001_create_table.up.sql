CREATE TABLE IF NOT EXISTS user_profiles(
	id  SERIAL PRIMARY KEY ,
	name VARCHAR(20),
	last_name VARCHAR(20),
	about VARCHAR(100),
	profile_image_url TEXT,
	external_id VARCHAR UNIQUE,
	user_name VARCHAR(20) UNIQUE,
	email VARCHAR(255) UNIQUE NOT NULL,
	is_verified Boolean NOT NULL
);


CREATE TABLE IF NOT EXISTS activities(
    id SERIAL PRIMARY KEY ,
    title VARCHAR(20),
	category VARCHAR(20),
	created_by INT ,
	quota INT,
	gender_composition VARCHAR(10),
	CONSTRAINT fk_activities_user_profiles FOREIGN KEY (created_by) REFERENCES user_profiles(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS activity_locations(
	activity_id  INT PRIMARY KEY REFERENCES activities(id) ON DELETE CASCADE,
	city VARCHAR(20) NOT NULL,
	district VARCHAR(50) NOT NULL,
	description TEXT,
	latitude REAL NOT NULL,
	longitude REAL NOT NULL
);



CREATE TABLE IF NOT EXISTS user_profile_addresses(
	profile_id  INT PRIMARY KEY REFERENCES user_profiles(id) ON DELETE CASCADE ,
	City VARCHAR(20)
);


CREATE TABLE IF NOT EXISTS user_profile_stats(
	profile_id  INT PRIMARY KEY REFERENCES user_profiles(id) ON DELETE CASCADE,
	average_point REAL DEFAULT 0 ,
	attanded_activities INT DEFAULT 0
);


CREATE TABLE IF NOT EXISTS participants (
	activity_id INT REFERENCES activities(id),
	user_id INT REFERENCES user_profiles(id),
	
	CONSTRAINT participants_pk PRIMARY KEY(activity_id,user_id),
	FOREIGN KEY (activity_id) REFERENCES activities(id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES user_profiles(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS activity_rules(
	id SERIAL PRIMARY KEY,
	activity_id INT NOT NULL,
	description TEXT NOT NULL,
	FOREIGN KEY (activity_id) REFERENCES activities(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS activity_flows(
	id SERIAL PRIMARY KEY,
	activity_id INT NOT NULL,
	description TEXT NOT NULL,
	FOREIGN KEY (activity_id) REFERENCES activities(id) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION insert_default_user_stats()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO user_profile_stats (profile_id, average_point, attanded_activities)
    VALUES (NEW.id, 0, 0);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_insert_default_user_stats
AFTER INSERT ON user_profiles
FOR EACH ROW
EXECUTE FUNCTION insert_default_user_stats();
