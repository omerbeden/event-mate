import 'package:event_mate/modules/screens/add_event_screen.dart';
import 'package:event_mate/modules/screens/events_screen.dart';
import 'package:event_mate/modules/screens/login_screen.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import 'profile_screen.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({Key? key, required this.firebaseUser}) : super(key: key);
  final User firebaseUser;

  @override
  _HomeScreenState createState() => _HomeScreenState();
}

PageController pageController = PageController();

class _HomeScreenState extends State<HomeScreen> {
  int _page = 0;
  bool triedSilentLogin = false;
  bool setupNotifications = false;
  bool firebaseInitialized = false;

  @override
  void initState() {
    super.initState();
    pageController = PageController();
  }

  @override
  Widget build(BuildContext context) {
    return CupertinoTabScaffold(
      tabBar: CupertinoTabBar(
        items: const [
          BottomNavigationBarItem(
            label: 'Events',
            icon: Icon(Icons.music_note),
          ),
          BottomNavigationBarItem(
            label: 'Add',
            icon: Icon(Icons.add),
          ),
          BottomNavigationBarItem(
            label: 'Profile',
            icon: Icon(Icons.person_off_outlined),
          ),
        ],
      ),
      tabBuilder: (context, index) {
        switch (index) {
          case 0:
            return CupertinoTabView(
              defaultTitle: 'Events',
              builder: (context) => const EventsScreen(),
            );
          case 1:
            return CupertinoTabView(
              defaultTitle: 'Organize Event',
              builder: (context) => LogInScreen(),
            );
          case 2:
            return CupertinoTabView(
              defaultTitle: 'Profile',
              builder: (context) => const ProfileScreen(),
            );
          default:
            assert(false, 'Unexpected tab');
            return const SizedBox.shrink();
        }
      },
    );
  }

  void navigationTapped(int page) {
    //Animating Page
    print(page);
    setState(() {
      this._page = page;
    });

    pageController.jumpToPage(page);
  }
}
