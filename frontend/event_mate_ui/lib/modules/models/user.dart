class User {
  final String userName;
  final String name;
  final String lastName;
  final String about;
  final String imagePath;
  final List<Event> attandedEvents;
  final UserProfileAdress adress;
  final String job;
  final bool isDarkMode;
  final UserProfileStat stat;

  const User(
      {required this.userName,
      required this.name,
      required this.lastName,
      required this.about,
      required this.imagePath,
      required this.attandedEvents,
      required this.adress,
      required this.job,
      required this.isDarkMode,
      required this.stat});
}

class UserProfileStat {
  final int followers;
  final int following;
  final int attandedEvents;
  final double points;

  const UserProfileStat(
      {required this.followers,
      required this.following,
      required this.attandedEvents,
      required this.points});
}

class UserProfileAdress {
  final String city;
  const UserProfileAdress({required this.city});
}

//user profiledaki gösterilecek olan event
class Event {
  final String title;
  final String description;
  final String location;
  final String duration;

  const Event(
      {required this.title,
      required this.description,
      required this.location,
      required this.duration});
}
