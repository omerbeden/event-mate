import 'package:event_mate/modules/screens/message_screen.dart';
import 'package:flutter/material.dart';

import '../modules/screens/search_secreen.dart';

AppBar buildAppBar(BuildContext context) {
  return AppBar(
    iconTheme: Theme.of(context).iconTheme,
    backgroundColor: Colors.white10,
    leading: BackButton(
      color: Colors.black12,
    ),
    elevation: 1,
  );
}

AppBar buildEvetDetailScreenAppBar(BuildContext context) {
  return AppBar(
    title: const Text("Event Mate"),
    iconTheme: Theme.of(context).iconTheme,
    backgroundColor: Colors.white10,
  );
}

AppBar buildEventScreenAppBar(BuildContext context) {
  return AppBar(
    title: const Text("Event Mate"),
    automaticallyImplyLeading: false,
    backgroundColor: Colors.white10,
    iconTheme: Theme.of(context).iconTheme,
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
