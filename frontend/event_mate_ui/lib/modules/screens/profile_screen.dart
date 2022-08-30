import 'package:event_mate/widgets/event_card_widget.dart';
import 'package:flutter/material.dart';
import 'package:event_mate/modules/models/user.dart';
import 'package:event_mate/utils/user_preference.dart';
import 'package:event_mate/widgets/appbar_widget.dart';
import 'package:event_mate/widgets/profile_widget.dart';
import 'package:event_mate/modules/screens/edit_profile_screen.dart';

import 'event_detail_screen.dart';

class ProfileScreen extends StatefulWidget {
  const ProfileScreen({Key? key}) : super(key: key);

  @override
  _ProfileScreenState createState() => _ProfileScreenState();
}

class _ProfileScreenState extends State<ProfileScreen> {
  late Future<List<Event>> events;
  final user = UserPreferences.myUser;

  @override
  void initState() {
    super.initState();
    events = getAttendedEvents();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: buildAppBar(context),
      body: ListView(
        physics: const BouncingScrollPhysics(),
        children: [
          ProfileWidget(
            imagePath: user.imagePath,
            onClicked: () async {
              await Navigator.of(context).push(
                MaterialPageRoute(builder: (context) => EditProfileScreen()),
              );
            },
          ),
          const SizedBox(height: 24),
          buildName(user),
          const SizedBox(height: 48),
          buildStats(user),
          const SizedBox(height: 48),
          buildAbout(user),
          const SizedBox(height: 48),
          const Text('Attanded Events',
              style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold)),
          buildAttendedEvents(),
          const SizedBox(height: 48),
        ],
      ),
    );
  }

  Widget buildName(User user) => Column(
        children: [
          Text(
            user.userName,
            style: const TextStyle(fontWeight: FontWeight.bold, fontSize: 24),
          ),
          const SizedBox(height: 4),
          Text(
            user.name,
            style: const TextStyle(color: Colors.grey),
          ),
          Text(
            user.lastName,
            style: const TextStyle(color: Colors.grey),
          ),
          const SizedBox(height: 4),
          Text(
            user.job,
            style: const TextStyle(color: Colors.grey),
          )
        ],
      );

  Widget buildAbout(User user) => Container(
        padding: const EdgeInsets.symmetric(horizontal: 48),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Text(
              'About',
              style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 16),
            Text(
              user.about,
              style: const TextStyle(fontSize: 16, height: 1.4),
            ),
          ],
        ),
      );

  Widget buildAttendedEvents() {
    return FutureBuilder<List<Event>>(
        future: events,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            return GridView.builder(
              shrinkWrap: true,
              physics: const NeverScrollableScrollPhysics(),
              padding: const EdgeInsets.all(0.5),
              itemCount: snapshot.data?.length,
              gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 4,
                childAspectRatio: 1.0,
                mainAxisSpacing: 1.5,
                crossAxisSpacing: 1.5,
              ),
              itemBuilder: (context, index) {
                return EventCard(
                    title: snapshot.data![index].title,
                    description: snapshot.data![index].description,
                    location: snapshot.data![index].location,
                    duration: snapshot.data![index].duration);
              },
            );
          } else {
            return const Text('waiting data');
          }
        });
  }

  Widget buildStats(User user) {
    var upstats = Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Column(
          children: [
            const Text(
              'Following',
              style: TextStyle(fontSize: 12),
            ),
            Text(
              user.stat.following.toString(),
              style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
            ),
          ],
        ),
        VerticalDivider(),
        Column(
          children: [
            const Text(
              'Followers',
              style: TextStyle(fontSize: 12),
            ),
            Text(
              user.stat.attandedEvents.toString(),
              style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
            ),
          ],
        ),
        VerticalDivider(),
        Column(
          children: [
            const Text(
              'Attanded Events',
              style: TextStyle(fontSize: 12),
            ),
            Text(
              user.stat.following.toString(),
              style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
            ),
          ],
        ),
        VerticalDivider(),
      ],
    );

    return Container(
      padding: EdgeInsets.symmetric(vertical: 20),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        mainAxisAlignment: MainAxisAlignment.end,
        children: [
          upstats,
          SizedBox(height: 24),
          Column(
            mainAxisSize: MainAxisSize.min,
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              const Text(
                'Points',
                style: TextStyle(fontSize: 12),
              ),
              Text(
                user.stat.points.toString(),
                style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
              ),
            ],
          )
        ],
      ),
    );
  }

  Future<List<Event>> getAttendedEvents() async {
    //TODO:Get this from api
    return user.attandedEvents;
  }
}
