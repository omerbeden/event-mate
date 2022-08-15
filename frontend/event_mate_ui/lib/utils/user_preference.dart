import 'package:event_mate/modules/models/user.dart';

class UserPreferences {
  static const myUser = User(
      imagePath:
          'https://images.unsplash.com/photo-1554151228-14d9def656e4?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=333&q=80',
      userName: 'Sarah123',
      name: 'Sarah Abs',
      adress: UserProfileAdress(city: "Sakarya"),
      attandedEvents: [
        Event(
            description: "description",
            title: "Bisiklet sürmece",
            location: 'Sakarya',
            duration: '2 days'),
        Event(
            description: "description",
            title: "Dağ tırmanısi",
            location: 'Sakarya',
            duration: '2 days'),
        Event(
            description: "description",
            title: "Bisiklet sürmece",
            location: 'Sakarya',
            duration: '2 days'),
        Event(
            description: "description",
            title: "piknik",
            location: 'Sakarya',
            duration: '2 days'),
        Event(
            description: "description",
            title: "piknik",
            location: 'Sakarya',
            duration: '2 days'),
        Event(
            description: "description",
            title: "piknik",
            location: 'Sakarya',
            duration: '2 days'),
      ],
      lastName: "beden",
      job: "Teacher",
      about:
          'Certified Personal Trainer and Nutritionist with years of experience in creating effective diets and training plans focused on achieving individual customers goals in a smooth way.',
      isDarkMode: false,
      stat: UserProfileStat(
          followers: 10, following: 20, attandedEvents: 5, points: 4.8));
}
