CREATE TABLE IF NOT EXISTS events(
    id serial primary key ,
    title varchar(20),
	category varchar(20),
	created_user_id int
);

CREATE TABLE IF NOT EXISTS event_locations(
	event_id  int primary key references events(id) ON DELETE CASCADE,
	city varchar(20)
);

CREATE TABLE IF NOT EXISTS user_profiles(
	id  serial primary key ,
	name varchar(20),
	last_name varchar(20)
);

CREATE TABLE IF NOT EXISTS user_profile_addresses(
	profile_id  int primary key references user_profiles(id) ,
	City varchar(20)
);


CREATE TABLE IF NOT EXISTS user_profile_stats(
	profile_id  int primary key references user_profiles(id) ,
	points varchar(20)
);


CREATE TABLE IF NOT EXISTS participants (
	event_id int references events(id) ON DELETE CASCADE,
	user_id int references user_profiles(id)ON DELETE CASCADE,
	
	CONSTRAINT participants_pk PRIMARY KEY(event_id,user_id));
