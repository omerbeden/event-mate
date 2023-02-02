import 'package:event_mate/modules/screens/login_screen.dart';
import 'package:event_mate/utils/Authentication.dart';
import 'package:event_mate/widgets/event_card_widget.dart';
import 'package:flutter/material.dart';
import 'package:event_mate/modules/models/user.dart';
import 'package:event_mate/utils/user_preference.dart';
import 'package:event_mate/widgets/appbar_widget.dart';
import 'package:event_mate/widgets/profile_widget.dart';
import 'package:event_mate/modules/screens/edit_profile_screen.dart';
import 'package:firebase_auth/firebase_auth.dart' as fbu;

import 'event_detail_screen.dart';

class ProfileScreen extends StatefulWidget {
  const ProfileScreen({Key? key, required this.firebaseUser}) : super(key: key);
  final fbu.User firebaseUser;

  @override
  _ProfileScreenState createState() => _ProfileScreenState();
}

class _ProfileScreenState extends State<ProfileScreen> {
  late Future<List<Event>> events;
  final user = UserPreferences.myUser;
  bool _isSigningOut = false;

  @override
  void initState() {
    super.initState();
    events = getAttendedEvents();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: buildAppBar(context),
      endDrawer: Drawer(
          child: SingleChildScrollView(
        padding: EdgeInsets.zero,
        child: SingleChildScrollView(
            child: Container(
          child: Column(children: [
            const DrawerHeader(
              decoration: BoxDecoration(
                color: Colors.blue,
              ),
              child: Text('Settings'),
            ),
            ListTile(
              title: const Text('Log Out'),
              onTap: () {
                setState(() {
                  _isSigningOut = true;
                });
                Authentication.signOut(context: context);
                Navigator.pop(context);
                Navigator.of(context, rootNavigator: true)
                    .pushReplacement(_routeToLogInSecreen());
              },
            )
          ]),
        )),
      )),
      body: ListView(
        physics: const BouncingScrollPhysics(),
        children: [
          ProfileWidget(
            imagePath: widget.firebaseUser.photoURL.toString(),
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
          Expanded(child: buildAttendedEvents()),
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
            return ListView.builder(
              shrinkWrap: true,
              physics: const NeverScrollableScrollPhysics(),
              padding: const EdgeInsets.all(0.5),
              itemCount: snapshot.data?.length,
              itemBuilder: (context, index) {
                return Wrap(
                  clipBehavior: Clip.antiAliasWithSaveLayer,
                  children: [
                    GestureDetector(
                      child:
                          eventCardHeader(snapshot.data![index].title, context),
                      onTap: () {
                        Navigator.push(
                            context,
                            MaterialPageRoute(
                                builder: (context) =>
                                    EventDetailScreen(id: 1)));
                      },
                    ),
                    Divider(color: Theme.of(context).primaryColor)
                  ],
                );
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

  Route _routeToLogInSecreen() {
    return PageRouteBuilder(
      pageBuilder: (context, animation, secondaryAnimation) => LogInScreen(),
      transitionsBuilder: (context, animation, secondaryAnimation, child) {
        var begin = Offset(-1.0, 0.0);
        var end = Offset.zero;
        var curve = Curves.ease;

        var tween =
            Tween(begin: begin, end: end).chain(CurveTween(curve: curve));

        return SlideTransition(
          position: animation.drive(tween),
          child: child,
        );
      },
    );
  }
}
