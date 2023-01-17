import 'package:event_mate/modules/screens/add_event_screen.dart';
import 'package:event_mate/modules/screens/events_screen.dart';
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
    return Scaffold(
        body: PageView(
          children: [
            Container(
              color: Colors.white,
              child: const EventsScreen(),
            ),
            Container(
              color: Colors.white,
              child: const AddEventScreen(),
            ),
            Container(
                color: Colors.white,
                child: ProfileScreen(firebaseUser: widget.firebaseUser)),
          ],
          controller: pageController,
          physics: NeverScrollableScrollPhysics(),
          onPageChanged: onPageChanged,
        ),
        bottomNavigationBar: CupertinoTabBar(
          backgroundColor: Colors.white,
          // ignore: prefer_const_literals_to_create_immutables
          items: <BottomNavigationBarItem>[
            const BottomNavigationBarItem(
              label: 'Evets',
              icon: Icon(Icons.event),
            ),
            const BottomNavigationBarItem(
              label: 'Add',
              icon: Icon(Icons.add),
            ),
            const BottomNavigationBarItem(
              label: 'Profile',
              icon: Icon(Icons.person_off_outlined),
            ),
          ],
          onTap: navigationTapped,
          currentIndex: _page,
        ));
  }

  void navigationTapped(int page) {
    //Animating Page
    print(page);
    setState(() {
      this._page = page;
    });

    pageController.jumpToPage(page);
  }

  void onPageChanged(int page) {
    setState(() {
      this._page = page;
    });
  }
}
