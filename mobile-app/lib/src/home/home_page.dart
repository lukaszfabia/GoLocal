import 'package:flutter/material.dart';
import 'package:golocal/src/event/ui/events_view_page.dart';
import 'package:golocal/src/user/profile_page.dart';
import 'package:golocal/src/preference_survey/ui/preference_survey_page.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  List<Widget> pages = [
    const Text('Map'),
    EventsViewPage(),
    ProfilePage(),
    PreferenceSurveyPage(),
  ];

  @override
  Widget build(BuildContext context) => DefaultTabController(
        length: pages.length,
        initialIndex: 1,
        child: Scaffold(
          bottomNavigationBar: TabBar(
            labelStyle: const TextStyle(fontSize: 12),
            unselectedLabelStyle: const TextStyle(fontSize: 0),
            overlayColor: WidgetStateProperty.all(Colors.transparent),
            indicator: BoxDecoration(borderRadius: BorderRadius.circular(20)),
            tabs: [
              const Tab(icon: Icon(Icons.map), text: 'Map'),
              const Tab(icon: Icon(Icons.event), text: 'Events'),
              const Tab(icon: Icon(Icons.person), text: 'Profile'),
              const Tab(icon: Icon(Icons.settings), text: 'Settings'),
            ],
          ),
          body: TabBarView(
            children: [
              Center(child: Text('Map')),
              EventsViewPage(),
              ProfilePage(),
              PreferenceSurveyPage(),
            ],
          ),
        ),
      );
}
