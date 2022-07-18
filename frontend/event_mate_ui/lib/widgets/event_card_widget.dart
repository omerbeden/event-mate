import 'dart:ui';

import 'package:flutter/material.dart';

class EventCard extends StatelessWidget {
  final String title;
  final String description;
  final String location;
  final String duration;

  const EventCard({
    Key? key,
    required this.title,
    required this.description,
    required this.location,
    required this.duration,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Card(
        clipBehavior: Clip.antiAlias,
        margin: const EdgeInsets.all(10),
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
        child: ListView(
          children: [
            InkWell(
              splashColor: Colors.blue.withAlpha(50),
              onTap: () {
                debugPrint('Card tapped.');
              },
              child: ListTile(
                leading: const Icon(Icons.arrow_circle_down),
                title: Text(
                  title,
                  style: const TextStyle(
                      fontWeight: FontWeight.bold, fontSize: 16),
                ),
                subtitle: Text('created by'),
                dense: true,
              ),
            ),
            const SizedBox(
              width: 20,
            ),
            Text(description),
            const SizedBox(
              height: 20,
            ),
            Wrap(
              alignment: WrapAlignment.spaceBetween,
              clipBehavior: Clip.antiAlias,
              children: [
                Wrap(
                  children: [
                    const Icon(Icons.location_on),
                    Text(location),
                    const Icon(Icons.timelapse_sharp),
                    Text(duration),
                  ],
                ),
                Wrap(
                  children: [
                    IconButton(
                      onPressed: () {},
                      icon: const Icon(Icons.heart_broken),
                    ),
                    IconButton(
                      onPressed: () {},
                      icon: const Icon(Icons.share),
                    ),
                  ],
                ),
              ],
            ),
            const SizedBox(
              width: 20,
            ),
          ],
        ),
      ),
    );
  }
}
