import 'package:event_mate/modules/screens/login_screen.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(const EventMate());
}

class EventMate extends StatelessWidget {
  const EventMate({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'EventMate',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      builder: (context, child) {
        return CupertinoTheme(
          data: const CupertinoThemeData(),
          child: Material(child: child),
        );
      },
      home: LogInScreen(),
    );
  }
}
