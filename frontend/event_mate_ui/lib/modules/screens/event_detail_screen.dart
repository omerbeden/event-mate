import 'package:event_mate/widgets/event_card_widget.dart';
import 'package:flutter/material.dart';

class EventDetailScreen extends StatefulWidget {
  const EventDetailScreen({
    required this.id,
    Key? key,
  }) : super(key: key);

  final int id;
  @override
  State<StatefulWidget> createState() => EventDetailScreenState();
}

class EventDetailScreenState extends State<EventDetailScreen> {
  Future<bool> onBackPress() {
    return Future.value(false);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Container(
            padding: const EdgeInsets.symmetric(
              vertical: 10,
            ),
            child: Column(
              // ignore: prefer_const_literals_to_create_immutables
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                const EventCard(
                    title: "title",
                    description: "description",
                    location: "location",
                    duration: "duration"),
                Text("Rules"),
                Wrap(
                  alignment: WrapAlignment.spaceAround,
                  children: [
                    for (var rule in [1, 2, 3]) Chip(label: Text("rule")),
                  ],
                ),
                //ChatSection
                Container(
                    child: SafeArea(
                  child: WillPopScope(
                    child: Stack(
                      children: [
                        Column(
                          children: [
                            Text("list messages"),
                            Text("input message")
                          ],
                        )
                      ],
                    ),
                    onWillPop: onBackPress,
                  ),
                ))
              ],
            )));
  }
}
