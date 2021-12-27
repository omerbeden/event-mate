import 'dart:developer';


import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import 'profile_screen.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({Key? key, required this.title}) : super(key: key);
  final String title;

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
      body: PageView(controller: pageController, children: [
        Container(
          color: Colors.red,
        ),
        Container(color: Colors.blue),
        Container(
          color: Colors.white,
        ),
        Container(color: Colors.yellow),
        const ProfileScreen(),
      ]),
      bottomNavigationBar: CupertinoTabBar(
        backgroundColor: Colors.white,
        items: <BottomNavigationBarItem>[
          BottomNavigationBarItem(
              icon: Icon(Icons.home,
                  color: (_page == 0) ? Colors.black : Colors.grey),
              backgroundColor: Colors.white),
          BottomNavigationBarItem(
              icon: Icon(Icons.search,
                  color: (_page == 1) ? Colors.black : Colors.grey),
              backgroundColor: Colors.white),
          BottomNavigationBarItem(
              icon: Icon(Icons.add_circle,
                  color: (_page == 2) ? Colors.black : Colors.grey),
              backgroundColor: Colors.white),
          BottomNavigationBarItem(
              icon: Icon(Icons.mode_comment,
                  color: (_page == 3) ? Colors.black : Colors.grey),
              backgroundColor: Colors.white),
          BottomNavigationBarItem(
              icon: Icon(Icons.person,
                  color: (_page == 4) ? Colors.black : Colors.grey),
              backgroundColor: Colors.white),
        ],
        currentIndex: _page,
        onTap: navigationTapped,
        activeColor: Theme.of(context).primaryColor,
      ),
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
