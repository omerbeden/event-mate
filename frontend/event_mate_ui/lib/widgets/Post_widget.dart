import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

import 'event_card_widget.dart';

class PostWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Wrap(
      children: [
        const Padding(
            padding: EdgeInsets.all(8),
            child: CircleAvatar(
              backgroundColor: Colors.amber,
            )),
        const Padding(
          padding: EdgeInsets.all(8),
          child: Text("user name"),
        ),
        Container(
            height: 200,
            child: const EventCard(
                title: 'Title',
                description: 'Description',
                location: 'location',
                duration: 'duration'))
      ],
    );
  }
}
