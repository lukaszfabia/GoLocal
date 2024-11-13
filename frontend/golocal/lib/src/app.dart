import 'package:flutter/material.dart';
import 'package:golocal/src/event/ui/events_view_page.dart';
import 'package:golocal/src/login/login_page.dart';

class GoLocalApp extends StatelessWidget {
  const GoLocalApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: EventsViewPage(),
    );
  }
}
