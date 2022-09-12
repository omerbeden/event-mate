import 'package:event_mate/widgets/appbar_widget.dart';
import 'package:flutter/material.dart';

import '../../widgets/Post_widget.dart';
import 'event_detail_screen.dart';

class EventsScreen extends StatefulWidget {
  const EventsScreen({Key? key, this.androidDrawer}) : super(key: key);

  final Widget? androidDrawer;

  @override
  State<StatefulWidget> createState() => _EventsScreenState();
}

class _EventsScreenState extends State<EventsScreen> {
  final _androidRefreshKey = GlobalKey<RefreshIndicatorState>();
  static const _itemsLength = 50;

  Future<void> _refreshData() {
    return Future.delayed(
      // This is just an arbitrary delay that simulates some network activity.
      const Duration(seconds: 2),
      () => setState(() {}),
    );
  }

  Widget _listBuilder(BuildContext context, int index) {
    if (index >= _itemsLength) return Container();

    // Show a slightly different color palette. Show poppy-ier colors on iOS
    // due to lighter contrasting bars and tone it down on Android.

    return PostWidget();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      drawer: widget.androidDrawer,
      appBar: buildEventScreenAppBar(context),
      body: RefreshIndicator(
        onRefresh: _refreshData,
        child: ListView.builder(
          padding: const EdgeInsets.symmetric(vertical: 12),
          itemCount: _itemsLength,
          itemBuilder: _listBuilder,
        ),
      ),
    );
  }
}
