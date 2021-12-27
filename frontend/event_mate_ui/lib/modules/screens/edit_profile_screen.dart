import 'dart:io';

import 'package:flutter/material.dart';
import 'package:event_mate/modules/models/user.dart';
import 'package:event_mate/utils/user_preference.dart';
import 'package:event_mate/widgets/appbar_widget.dart';
import 'package:event_mate/widgets/profile_widget.dart';
import 'package:path/path.dart';

class EditProfileScreen extends StatefulWidget {
  const EditProfileScreen({Key? key}) : super(key: key);

  @override
  _EditProfileScreenState createState() => _EditProfileScreenState();
}

class _EditProfileScreenState extends State<EditProfileScreen> {
  late User user;

  @override
  void initState() {
    super.initState();

    final user = UserPreferences.myUser;
  }

  @override
  Widget build(BuildContext context) {
    final user = UserPreferences.myUser;

    return Scaffold(
      appBar: buildAppBar(context),
      body: ListView(
        physics: const BouncingScrollPhysics(),
        padding: EdgeInsets.symmetric(horizontal: 32),
        children: [
          ProfileWidget(
            imagePath: user.imagePath,
            onClicked: () async {},
          ),
          const Text(
            'About',
            style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
          ),
        ],
      ),
    );
  }
}
