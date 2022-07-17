import 'dart:ui';

import 'package:flutter/material.dart';

class EventCard extends StatelessWidget {
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
              child: const ListTile(
                leading: Icon(Icons.arrow_circle_down),
                title: Text(
                  'Card Title',
                  style: TextStyle(fontWeight: FontWeight.bold),
                ),
                subtitle: Text('created by'),
                dense: true,
                trailing: Icon(Icons.coffee_maker),
              ),
            ),
            const Text('data'),
            Wrap(
              children: [
                const Icon(Icons.location_on),
                const Text('location data'),
                const Icon(Icons.timelapse_sharp),
                const Text('duration data'),
                IconButton(
                  onPressed: () {},
                  icon: const Icon(Icons.heart_broken),
                ),
                IconButton(
                  onPressed: () {},
                  icon: const Icon(Icons.share),
                ),
              ],
            )
          ],
        ),
      ),
    );
  }
}
