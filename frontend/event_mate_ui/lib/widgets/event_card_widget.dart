import 'dart:ui';

import 'package:event_mate/modules/screens/event_detail_screen.dart';
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
      padding: const EdgeInsets.symmetric(
        vertical: 10,
      ),
      child: Column(children: [
        eventCardHeader(title, context),

        //body
        Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(10),
            ),
            padding: const EdgeInsets.symmetric(
              vertical: 4,
              horizontal: 16,
            ),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.end,
                  // ignore: prefer_const_literals_to_create_immutables
                  children: [
                    const Icon(Icons.timelapse),
                    const Text("Start Date"),
                    const SizedBox(
                      width: 10.0,
                    ),
                    const Text("-",
                        style: TextStyle(
                          fontWeight: FontWeight.w600,
                          fontSize: 30.0,
                        )),
                    const SizedBox(
                      width: 10.0,
                    ),
                    const Text("End Date"),
                  ],
                ),
                GestureDetector(
                  onTap: () {
                    Navigator.push(
                        context,
                        MaterialPageRoute(
                            builder: (context) => EventDetailScreen(id: 1)));
                  },
                  child: const Text(
                      "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."),
                ),
                const SizedBox(
                  height: 15,
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceAround,
                  children: [
                    Column(
                      // ignore: prefer_const_literals_to_create_immutables
                      children: [
                        const Icon(
                          Icons.location_on,
                        ),
                        Text(location),
                      ],
                    ),
                    Row(
                      children: [
                        IconButton(
                            onPressed: () {},
                            icon: Icon(Icons.person_add_alt_rounded)),
                        const Text("10"),
                      ],
                    ),
                    IconButton(onPressed: () {}, icon: Icon(Icons.share))
                  ],
                ),
              ],
            )),
        Divider(color: Theme.of(context).primaryColor)
      ]),
    );
  }
}

Widget eventCardHeader(String title, BuildContext context) {
  return // Header Section
      Container(
    decoration: BoxDecoration(
      borderRadius: BorderRadius.circular(10),
    ),
    padding: const EdgeInsets.symmetric(
      vertical: 4,
      horizontal: 16,
    ).copyWith(right: 0),
    child: Row(
      children: [
        const Icon(Icons.pedal_bike),
        Expanded(
          child: Padding(
            padding: const EdgeInsets.only(
              left: 8,
            ),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              crossAxisAlignment: CrossAxisAlignment.start,
              // ignore: prefer_const_literals_to_create_immutables
              children: [
                Text(
                  title,
                  style: const TextStyle(
                    fontWeight: FontWeight.bold,
                    fontSize: 30,
                  ),
                ),
                const Text(
                  "Omer tarafından oluşturuldu",
                  style: TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 10,
                      color: Colors.grey),
                )
              ],
            ),
          ),
        ),
        IconButton(
          onPressed: () {
            showDialog(
                context: context,
                builder: (context) => Dialog(
                      child: ListView(
                        padding: const EdgeInsets.symmetric(
                          vertical: 16,
                        ),
                        shrinkWrap: true,
                        children: [
                          'this will be custom for the aspect of the person event creted and just viewing',
                        ]
                            .map((e) => InkWell(
                                  onTap: () async {},
                                  child: Container(
                                    padding: const EdgeInsets.symmetric(
                                      horizontal: 12,
                                      vertical: 16,
                                    ),
                                    child: Text(e),
                                  ),
                                ))
                            .toList(),
                      ),
                    ));
          },
          icon: const Icon(
            Icons.more_vert,
          ),
        ),
      ],
    ),
  );
}
