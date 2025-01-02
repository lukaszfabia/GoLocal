import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/event/ui/events_view_page.dart';
import 'package:golocal/src/user/ui/profile_page.dart';
import 'package:golocal/src/routing/router.dart';

class ScaffoldShell extends StatelessWidget {
  final StatefulNavigationShell navigationShell;
  const ScaffoldShell({Key? key, required this.navigationShell})
      : super(key: key ?? const ValueKey("ScaffoldShell"));

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: navigationShell,
      bottomNavigationBar: NavigationBar(
        selectedIndex: navigationShell.currentIndex,
        labelBehavior: NavigationDestinationLabelBehavior.onlyShowSelected,
        destinations: [
          NavigationDestination(icon: Icon(Icons.map), label: 'Map'),
          NavigationDestination(icon: Icon(Icons.event), label: 'Events'),
          NavigationDestination(icon: Icon(Icons.person), label: 'Profile'),
        ],
        onDestinationSelected: _goBranch,
      ),
    );
  }

  void _goBranch(int index) {
    navigationShell.goBranch(index,
        initialLocation: index == navigationShell.currentIndex);
  }
}
