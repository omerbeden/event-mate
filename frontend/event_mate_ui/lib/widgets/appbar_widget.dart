import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

AppBar buildAppBar(BuildContext context) {
  return AppBar(
    leading: BackButton(
      color: Colors.black12,
    ),
    backgroundColor: Colors.transparent,
    elevation: 0,
  );
}
