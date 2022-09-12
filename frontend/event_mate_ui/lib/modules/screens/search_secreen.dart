import 'package:event_mate/widgets/appbar_widget.dart';
import 'package:flutter/material.dart';

class SearchScreen extends StatelessWidget {
  const SearchScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var _controller = TextEditingController();
    return Scaffold(appBar: buildSearchScreenAppBar(context, _controller));
  }
}
