import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

import 'event_card_widget.dart';

class PostWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Wrap(
      children: [
        Padding(
          padding: EdgeInsets.all(8),
          child: Container(
            width: 44,
            height: 44,
            decoration: BoxDecoration(
              color: Colors.blue,
              shape: BoxShape.circle,
            ),
            child: ClipRect(
              child: Icon(Icons.portable_wifi_off_outlined),
            ),
          ),
        ),
        Padding(
          padding: EdgeInsets.all(8),
          child: Text("user name"),
        ),
        Container(
            height: 200,
            child: EventCard(
                title: 'Title',
                description: 'Description',
                location: 'location',
                duration: 'duration'))
      ],
    );
  }
}
