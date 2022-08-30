import 'package:flutter/material.dart';

import '../../utils/Authentication.dart';
import '../../widgets/google_login_button.dart';

class LogInScreen extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => _LogInScreenState();
}

class _LogInScreenState extends State<LogInScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.amber,
      body: SafeArea(
          child: Padding(
        padding: const EdgeInsets.only(
          left: 16.0,
          right: 16.0,
          bottom: 20.0,
        ),
        child: Column(
          mainAxisSize: MainAxisSize.max,
          children: [
            Row(),
            Expanded(
                child: Column(
              mainAxisSize: MainAxisSize.min,
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Flexible(
                  child: Image.network(
                      "https://play-lh.googleusercontent.com/ahJtMe0vfOlAu1XJVQ6rcaGrQBgtrEZQefHy7SXB7jpijKhu1Kkox90XDuH8RmcBOXNn"),
                  flex: 1,
                ),
                const SizedBox(
                  height: 20,
                ),
                const Text("LOGIN TEXT"),
                const Text("data"),
              ],
            )),
            FutureBuilder(
                future: Authentication.initializeFirebase(context: context),
                builder: (context, snapshot) {
                  if (snapshot.hasError) {
                    return Text('Error initializing Firebase');
                  } else if (snapshot.connectionState == ConnectionState.done) {
                    return GoogleLogInButton();
                  }

                  return CircularProgressIndicator(
                      valueColor: AlwaysStoppedAnimation<Color>(Colors.amber));
                })
          ],
        ),
      )),
    );
  }
}
