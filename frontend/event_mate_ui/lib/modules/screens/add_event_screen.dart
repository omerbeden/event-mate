import 'package:event_mate/widgets/date_picker_widget.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import '../../widgets/time_picker_widget.dart';

class AddEventScreen extends StatefulWidget {
  const AddEventScreen({Key? key}) : super(key: key);

  @override
  AddEventScreenState createState() {
    return AddEventScreenState();
  }
}

class AddEventScreenState extends State<AddEventScreen> {
  final TextEditingController titleController = TextEditingController();
  final TextEditingController categoryController = TextEditingController();
  String selectedCity = "";
  final List<String> cities = <String>['One', 'Two', 'Three', 'Four'];

  @override
  Widget build(BuildContext context) {
    // Build a Form widget using the _formKey created above.
    return Scaffold(
      body: SafeArea(
          child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          const SizedBox(height: 40),
          Container(
            height: 40,
            width: MediaQuery.of(context).size.width / 1.12,
            decoration: BoxDecoration(
              color: Colors.grey[200],
              borderRadius: BorderRadius.circular(15),
            ),
            child: TextFormField(
              controller: titleController,
              keyboardType: TextInputType.text,
              decoration: const InputDecoration(
                border: InputBorder.none,
                contentPadding:
                    EdgeInsets.symmetric(horizontal: 20, vertical: 15),
                hintText: 'Event Title',
                hintStyle: TextStyle(
                  color: Colors.grey,
                  fontSize: 16,
                ),
                prefixIcon: Icon(
                  Icons.title,
                  color: Colors.grey,
                ),
              ),
              inputFormatters: [
                LengthLimitingTextInputFormatter(20),
              ],
              onChanged: (value) {
                var text = value.replaceAll(RegExp(r'\s+\b|\b\s'), ' ');
                setState(() {
                  titleController.value = titleController.value.copyWith(
                      text: text,
                      selection: TextSelection.collapsed(offset: text.length),
                      composing: TextRange.empty);
                });
              },
            ),
          ),
          const SizedBox(height: 12),
          Container(
            height: 55,
            width: MediaQuery.of(context).size.width / 1.12,
            decoration: BoxDecoration(
              color: Colors.grey[200],
              borderRadius: BorderRadius.circular(15),
            ),
            child: TextFormField(
              controller: categoryController,
              keyboardType: TextInputType.name,
              inputFormatters: [
                LengthLimitingTextInputFormatter(5),
              ],
              decoration: const InputDecoration(
                border: InputBorder.none,
                contentPadding:
                    EdgeInsets.symmetric(horizontal: 20, vertical: 15),
                hintText: 'Category',
                hintStyle: TextStyle(
                  color: Colors.grey,
                  fontSize: 16,
                ),
                prefixIcon: Icon(
                  Icons.category,
                  color: Colors.grey,
                ),
              ),
              onChanged: (value) {
                setState(() {
                  categoryController.value = categoryController.value.copyWith(
                      text: value,
                      selection: TextSelection.collapsed(offset: value.length),
                      composing: TextRange.empty);
                });
              },
            ),
          ),
          const SizedBox(height: 12),
          Container(
            height: 55,
            width: MediaQuery.of(context).size.width / 1.12,
            decoration: BoxDecoration(
              color: Colors.grey[200],
              borderRadius: BorderRadius.circular(15),
            ),
            child: DropdownButtonFormField(
              decoration: const InputDecoration(
                border: InputBorder.none,
                contentPadding:
                    EdgeInsets.symmetric(horizontal: 20, vertical: 15),
                hintText: 'Location',
                hintStyle: TextStyle(
                  color: Colors.grey,
                  fontSize: 16,
                ),
                prefixIcon: Icon(
                  Icons.location_city,
                  color: Colors.grey,
                ),
              ),
              items: cities.map<DropdownMenuItem<String>>((String value) {
                return DropdownMenuItem<String>(
                  value: value,
                  child: Text(value),
                );
              }).toList(),
              onChanged: (String? value) {
                setState(() {
                  selectedCity = value!;
                });
              },
            ),
          ),
          const SizedBox(
            height: 12,
          ),
          ElevatedButton(
              onPressed: () {
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(
                      content: Text('Processing Data, ana ekrana yönlendir')),
                );
                Navigator.pop(context);
              },
              child: const Text("Create"))
        ],
      )),
    );
  }
}
