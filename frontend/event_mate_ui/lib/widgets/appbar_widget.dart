import 'package:event_mate/modules/screens/message_screen.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import '../modules/screens/search_secreen.dart';

AppBar buildAppBar(BuildContext context) {
  return AppBar(
    leading: BackButton(
      color: Colors.black12,
    ),
    backgroundColor: Colors.transparent,
    elevation: 1,
  );
}

AppBar buildEventScreenAppBar(BuildContext context) {
  return AppBar(
    title: Text("Event Mate"),
    automaticallyImplyLeading: false,
    elevation: 1,
    actions: [
      IconButton(
        icon: Icon(Icons.search),
        onPressed: () {
          Navigator.of(context)
              .push(MaterialPageRoute(builder: (_) => const SearchScreen()));
        },
      ),
      IconButton(
          onPressed: () {
            Navigator.of(context)
                .push(MaterialPageRoute(builder: (_) => MessageScreen()));
          },
          icon: Icon(Icons.message)),
    ],
  );
}

AppBar buildSearchScreenAppBar(
    BuildContext context, TextEditingController textController) {
  return AppBar(
      title: Container(
    width: double.infinity,
    height: 40,
    decoration: BoxDecoration(
        color: Colors.white, borderRadius: BorderRadius.circular(5)),
    child: Center(
      child: TextField(
        controller: textController,
        decoration: InputDecoration(
            prefixIcon: const Icon(Icons.search),
            suffixIcon: IconButton(
              icon: const Icon(Icons.clear),
              onPressed: () {
                textController.clear();
              },
            ),
            hintText: 'Search...',
            border: InputBorder.none),
      ),
    ),
  ));
}
