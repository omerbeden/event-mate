CREATE TABLE IF NOT EXISTS activities(
    id SERIAL PRIMARY KEY ,
    title VARCHAR(20),
	category VARCHAR(20),
	created_by INT ,
	quota INT
);

CREATE TABLE IF NOT EXISTS activity_locations(
	activity_id  INT PRIMARY KEY REFERENCES activities(id) ON DELETE CASCADE,
	city VARCHAR(20) NOT NULL,
	district VARCHAR(50) NOT NULL,
	description TEXT,
	latitude REAL NOT NULL,
	longitude REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS user_profiles(
	id  SERIAL PRIMARY KEY ,
	name VARCHAR(20),
	last_name VARCHAR(20),
	about VARCHAR(100),
	profile_image_url TEXT,
	external_id TEXT UNIQUE,
	user_name VARCHAR(20) UNIQUE;
);

CREATE TABLE IF NOT EXISTS user_profile_addresses(
	profile_id  INT PRIMARY KEY REFERENCES user_profiles(id) ON DELETE CASCADE ,
	City VARCHAR(20)
);


CREATE TABLE IF NOT EXISTS user_profile_stats(
	profile_id  INT PRIMARY KEY REFERENCES user_profiles(id) ON DELETE CASCADE,
	point REAL,
	attanded_activities INT
);


CREATE TABLE IF NOT EXISTS participants (
	activity_id INT REFERENCES activities(id) ON DELETE CASCADE,
	user_id INT REFERENCES user_profiles(id)ON DELETE CASCADE,
	
	CONSTRAINT participants_pk PRIMARY KEY(activity_id,user_id));


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
