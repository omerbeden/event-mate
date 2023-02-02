import 'package:event_mate/utils/user_preference.dart';
import 'package:event_mate/widgets/appbar_widget.dart';
import 'package:event_mate/widgets/event_card_widget.dart';
import 'package:flutter/material.dart';

import '../models/user.dart';

class EventsScreen extends StatefulWidget {
  const EventsScreen({Key? key}) : super(key: key);

  @override
  State<StatefulWidget> createState() => _EventsScreenState();
}

class _EventsScreenState extends State<EventsScreen> {
  late Future<List<Event>> events;

  Future<void> _refreshData() {
    return Future.delayed(
      // This is just an arbitrary delay that simulates some network activity.
      const Duration(seconds: 2),
      () => setState(() {}),
    );
  }

  Future<List<Event>> getAttendedEvents() async {
    return UserPreferences.myUser.attandedEvents;
  }

  @override
  void initState() {
    super.initState();
    events = getAttendedEvents();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: buildEventScreenAppBar(context),
      body: RefreshIndicator(
          onRefresh: _refreshData,
          child: FutureBuilder<List<Event>>(
              future: events,
              builder: (context, snapshot) {
                if (snapshot.connectionState == ConnectionState.waiting) {
                  return const Center(
                    child: CircularProgressIndicator(),
                  );
                }
                return ListView.builder(
                    padding: const EdgeInsets.symmetric(vertical: 12),
                    itemCount: snapshot.data?.length,
                    itemBuilder: (context, index) {
                      return EventCard(
                        title: snapshot.data![index].title,
                        description: snapshot.data![index].description,
                        location: snapshot.data![index].location,
                        duration: snapshot.data![index].duration,
                      );
                    });
              })),
    );
  }
}
