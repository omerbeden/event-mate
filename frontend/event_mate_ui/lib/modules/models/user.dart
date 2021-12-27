class User {
  final String imagePath;
  final String name;
  final String userName;
  final String job;
  final String email;
  final String about;
  final bool isDarkMode;

  const User({
    required this.imagePath,
    required this.name,
    required this.userName,
    required this.email,
    required this.job,
    required this.about,
    required this.isDarkMode,
  });
}
