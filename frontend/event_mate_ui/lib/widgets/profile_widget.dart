import 'package:flutter/material.dart';

class ProfileWidget extends StatelessWidget {
  final String imagePath;
  final VoidCallback onClicked;

  const ProfileWidget({
    Key? key,
    required this.imagePath,
    required this.onClicked,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final color = Theme.of(context).colorScheme.primary;

    return Center(
      child: Stack(
        children: [
          buildImage(),
          Positioned(
            bottom: 0,
            right: 4,
            child: buildEditButton(color),
          ),
        ],
      ),
    );
  }

  Widget buildImage() {
    final image = NetworkImage(imagePath);

    return ClipOval(
      child: Material(
        color: Colors.transparent,
        child: Ink.image(
          image: image,
          fit: BoxFit.cover,
          width: 128,
          height: 128,
        ),
      ),
    );
  }

  Widget buildEditButton(Color color) => ElevatedButton(
      onPressed: onClicked,
      child: Icon(
        Icons.edit,
        size: 20,
        semanticLabel: 'edit',
      ),
      style: ButtonStyle(
        shape: MaterialStateProperty.all(CircleBorder()),
        padding: MaterialStateProperty.all(EdgeInsets.all(20)),
        backgroundColor: MaterialStateProperty.all(Colors.transparent),
        overlayColor: MaterialStateProperty.resolveWith<Color?>((states) {
          if (states.contains(MaterialState.pressed)) {
            return Colors.red;
          }
        }),
      ));
}
