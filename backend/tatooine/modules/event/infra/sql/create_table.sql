create table events(
    id serial primary key ,
    Title varchar(20)
)

Create Table event_locations(
	event_id  int primary key references events(id) ,
	city varchar(20)
)

Create Table userProfile(
	id  serial primary,
	name varchar(20),
	last_name varchar(20)
)
Create Table userProfileAddress(
	profileId  int primary key references userProfile(id) ,
	City varchar(20)
)


Create Table userProfileStat(
	profileId  int primary key references userProfile(id) ,
	points varchar(20)
)


Create table Participants (
	eventId int references events(id) ,
	profileId int references userprofile(id),
	
	CONSTRAINT participants_pk PRIMARY KEY(eventId,profileId) )
