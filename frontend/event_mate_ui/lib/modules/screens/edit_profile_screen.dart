import 'dart:io';

import 'package:flutter/material.dart';
import 'package:event_mate/modules/models/user.dart';
import 'package:event_mate/utils/user_preference.dart';
import 'package:event_mate/widgets/appbar_widget.dart';
import 'package:event_mate/widgets/profile_widget.dart';
import 'package:image_picker/image_picker.dart';

import '../../widgets/textfield_widget.dart';

class EditProfileScreen extends StatefulWidget {
  const EditProfileScreen({Key? key}) : super(key: key);

  @override
  _EditProfileScreenState createState() => _EditProfileScreenState();
}

class _EditProfileScreenState extends State<EditProfileScreen> {
  late User user;
  File? _photo;
  final ImagePicker _picker = ImagePicker();

  @override
  void initState() {
    super.initState();

    final user = UserPreferences.myUser;
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final user = UserPreferences.myUser;

    Future uploadFile() async {
      debugPrint('uploading to server via grpc');
      // if (_photo == null) return;
      // final fileName = basename(_photo!.path);
      // final destination = 'files/$fileName';

      // try {
      //   final ref = firebase_storage.FirebaseStorage.instance
      //       .ref(destination)
      //       .child('file/');
      //   await ref.putFile(_photo!);
      // } catch (e) {
      //   print('error occured');
      // }
    }

    Future imgFromGallery() async {
      final pickedFile = await _picker.pickImage(source: ImageSource.gallery);

      setState(() {
        if (pickedFile != null) {
          _photo = File(pickedFile.path);
          uploadFile();
        } else {
          print('No image selected.');
        }
      });
    }

    Future imgFromCamera() async {
      final pickedFile = await _picker.pickImage(source: ImageSource.camera);

      setState(() {
        if (pickedFile != null) {
          _photo = File(pickedFile.path);
          uploadFile();
        } else {
          print('No image selected.');
        }
      });
    }

    return Scaffold(
      appBar: buildAppBar(context),
      body: ListView(
        physics: const BouncingScrollPhysics(),
        padding: const EdgeInsets.symmetric(vertical: 32),
        children: [
          ProfileWidget(
            imagePath: user.imagePath,
            onClicked: () async {
              debugPrint('upload image code');
              showModalBottomSheet(
                  context: context,
                  builder: (BuildContext bc) {
                    return SafeArea(
                        child: Container(
                      child: Wrap(
                        children: <Widget>[
                          ListTile(
                              leading: Icon(Icons.photo_library),
                              title: Text('Gallery'),
                              onTap: () {
                                imgFromGallery();
                                Navigator.of(context).pop();
                              }),
                          ListTile(
                            leading: Icon(Icons.photo_camera),
                            title: Text('Camera'),
                            onTap: () {
                              imgFromCamera();
                              Navigator.of(context).pop();
                            },
                          )
                        ],
                      ),
                    ));
                  });
            },
          ),
          TextFieldWidget(
            label: 'User Name',
            text: user.name,
            onChanged: (name) {},
          ),
          const SizedBox(height: 24),
          TextFieldWidget(
            label: 'About',
            text: user.about,
            onChanged: (name) {},
          ),
          const SizedBox(height: 24),
          TextFieldWidget(
            label: 'City',
            text: user.adress.city,
            onChanged: (name) {},
          ),
          const SizedBox(height: 24),
          TextFieldWidget(
            label: 'Job',
            text: user.job,
            onChanged: (name) {},
          ),
          const SizedBox(height: 48),
          ElevatedButton(
              onPressed: () {},
              child: Icon(
                Icons.done,
                semanticLabel: 'Done',
              ),
              style: ButtonStyle(
                shape: MaterialStateProperty.all(CircleBorder()),
                padding: MaterialStateProperty.all(EdgeInsets.all(20)),
                backgroundColor: MaterialStateProperty.all(Colors.blue),
                overlayColor:
                    MaterialStateProperty.resolveWith<Color?>((states) {
                  if (states.contains(MaterialState.pressed)) {
                    return Colors.green;
                  } // <-- Splash color
                }),
              ))
        ],
      ),
    );
  }
}
