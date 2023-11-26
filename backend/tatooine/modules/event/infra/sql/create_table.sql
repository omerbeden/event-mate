create table events(
    id serial primary key ,
    title varchar(20),
	category varchar(20),
	created_user_id int
)

Create Table event_locations(
	event_id  int primary key references events(id) ON DELETE CASCADE,
	city varchar(20)
)

Create Table user_profile(
	id  serial primary key ,
	name varchar(20),
	last_name varchar(20)
)
Create Table user_profile_address(
	profileId  int primary key references userProfile(id) ,
	City varchar(20)
)


Create Table user_profile_stat(
	profileId  int primary key references userProfile(id) ,
	points varchar(20)
)


Create table participants (
	event_id int references events(id) ,
	user_id int references user_profile(id),
	
	CONSTRAINT participants_pk PRIMARY KEY(event_id,user_id)  ON DELETE CASCADE)
