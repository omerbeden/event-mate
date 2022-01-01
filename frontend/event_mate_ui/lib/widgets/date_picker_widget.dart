
import 'package:event_mate/widgets/button_widget.dart';
import 'package:flutter/material.dart';


class DatePickerWidget extends StatefulWidget {
  @override
  _DatePickerWidgetState createState() => _DatePickerWidgetState();
}

class _DatePickerWidgetState extends State<DatePickerWidget> {
  late DateTime date = DateTime.now();

  String getText() {
    if (date == null) {
      return 'Select Date';
    } else {
      return '${date.month}/${date.day}/${date.year}';
      // return '${date.month}/${date.day}/${date.year}';
    }
  }

  @override
  Widget build(BuildContext context) => ButtonWidget(
        text: getText(),
        onClicked: () => pickDate(context),
      );

  Future pickDate(BuildContext context) async {
    final initialDate = DateTime.now();
    final newDate = await showDatePicker(
      context: context,
      initialDate: date,
      firstDate: DateTime(DateTime.now().year - 5),
      lastDate: DateTime(DateTime.now().year + 5),
    );

    if (newDate == null) return;

    setState(() => date = newDate);
  }
}