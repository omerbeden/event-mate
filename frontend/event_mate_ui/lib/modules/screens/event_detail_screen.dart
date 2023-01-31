import 'package:event_mate/widgets/event_card_widget.dart';
import 'package:flutter/material.dart';

class EventDetailScreen extends StatefulWidget {
  const EventDetailScreen({
    required this.id,
    Key? key,
  }) : super(key: key);

  final int id;
  @override
  State<StatefulWidget> createState() => EventDetailScreenState();
}

class EventDetailScreenState extends State<EventDetailScreen> {
  final TextEditingController textEditingController = TextEditingController();
  final ScrollController scrollController = ScrollController();

  List<dynamic> listMessages = ["message1", "message2"];

  String groupChatId = '';

  Future<bool> onBackPress() {
    return Future.value(false);
  }

  void onSendMessage(String content) {}

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Column(
      children: [
        const EventCard(
            title: "title",
            description: "description",
            location: "location",
            duration: "duration"),
        Container(
            decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(10),
                color: Theme.of(context).secondaryHeaderColor),
            padding: const EdgeInsets.symmetric(
              vertical: 4,
              horizontal: 16,
            ).copyWith(right: 0),
            child: Row(children: [
              Expanded(
                  child: Padding(
                padding: const EdgeInsets.only(
                  left: 8,
                ),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  crossAxisAlignment: CrossAxisAlignment.center,
                  // ignore: prefer_const_literals_to_create_immutables
                  children: [
                    Text(
                      "Rules",
                      style: TextStyle(
                        fontWeight: FontWeight.bold,
                      ),
                    )
                  ],
                ),
              ))
            ])),

        Wrap(
          alignment: WrapAlignment.spaceAround,
          children: [
            for (var rule in [1, 2, 3]) Chip(label: Text("rule")),
          ],
        ),
        //ChatSection

        Expanded(
            child: Container(
                decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(10),
                    color: Colors.amber),
                child: SafeArea(
                  child: WillPopScope(
                    child: Stack(
                      children: [
                        Column(
                          children: [
                            //list
                            Flexible(
                                child: ListView.builder(
                              padding: const EdgeInsets.all(10),
                              itemCount: 10,
                              reverse: true,
                              controller: scrollController,
                              itemBuilder: (context, index) => Text("data"),
                            )),

                            //Input
                            SizedBox(
                              width: double.infinity,
                              height: 50,
                              child: Row(
                                children: [
                                  Flexible(
                                      child: TextField(
                                    decoration: const InputDecoration(
                                      focusedBorder: OutlineInputBorder(
                                        borderSide: BorderSide(
                                            color: Colors.greenAccent,
                                            width: 5.0),
                                      ),
                                      enabledBorder: OutlineInputBorder(
                                        borderSide: BorderSide(
                                            color: Colors.red, width: 5.0),
                                      ),
                                      contentPadding: EdgeInsets.symmetric(
                                          horizontal: 20, vertical: 15),
                                      hintText: 'Write here',
                                      hintStyle: TextStyle(
                                        color: Colors.grey,
                                        fontSize: 16,
                                      ),
                                      prefixIcon: Icon(
                                        Icons.title,
                                        color: Colors.grey,
                                      ),
                                    ),
                                    textInputAction: TextInputAction.send,
                                    keyboardType: TextInputType.text,
                                    textCapitalization:
                                        TextCapitalization.sentences,
                                    controller: textEditingController,
                                    onSubmitted: (value) {
                                      onSendMessage(textEditingController.text);
                                    },
                                  )),
                                  Container(
                                    margin: const EdgeInsets.only(left: 4),
                                    decoration: BoxDecoration(
                                      color: Colors.blue,
                                      borderRadius: BorderRadius.circular(30),
                                    ),
                                    child: IconButton(
                                        onPressed: () {
                                          onSendMessage(
                                              textEditingController.text);
                                        },
                                        icon: const Icon(Icons.send_rounded),
                                        color: Colors.white),
                                  ),
                                ],
                              ),
                            )
                          ],
                        )
                      ],
                    ),
                    onWillPop: onBackPress,
                  ),
                ))),
      ],
    ));
  }

  Widget buildItem(String text) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.end,
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.end,
          children: [Text(text)],
        )
      ],
    );
  }
}
