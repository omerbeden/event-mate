CREATE TABLE IF NOT EXISTS activities(
    id serial primary key ,
    title varchar(20),
	category varchar(20),
	created_by int 
);

CREATE TABLE IF NOT EXISTS activity_locations(
	activity_id  int primary key references activities(id) ON DELETE CASCADE,
	city varchar(20)
);

CREATE TABLE IF NOT EXISTS user_profiles(
	id  serial primary key ,
	name varchar(20),
	last_name varchar(20)
);

CREATE TABLE IF NOT EXISTS user_profile_addresses(
	profile_id  int primary key references user_profiles(id) ON DELETE CASCADE ,
	City varchar(20)
);


CREATE TABLE IF NOT EXISTS user_profile_stats(
	profile_id  int primary key references user_profiles(id) ON DELETE CASCADE,
	point real,
	attanded_activities int
);


CREATE TABLE IF NOT EXISTS participants (
	activity_id int references activities(id) ON DELETE CASCADE,
	user_id int references user_profiles(id)ON DELETE CASCADE,
	
	CONSTRAINT participants_pk PRIMARY KEY(activity_id,user_id));