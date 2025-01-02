import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/event/ui/events_view_page.dart';
import 'package:golocal/src/user/ui/profile_page.dart';
import 'package:golocal/src/routing/router.dart';

class ScaffoldShell extends StatefulWidget {
  final StatefulNavigationShell navigationShell;
  final String? title;
  const ScaffoldShell({Key? key, required this.navigationShell, this.title})
      : super(key: key ?? const ValueKey("ScaffoldShell"));

  @override
  State<ScaffoldShell> createState() => _ScaffoldShellState();
}

class _ScaffoldShellState extends State<ScaffoldShell> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(child: widget.navigationShell),
      bottomNavigationBar: NavigationBar(
        selectedIndex: widget.navigationShell.currentIndex,
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

  Widget? _buildBackButton(BuildContext context) {
    if (Navigator.of(context).canPop()) {
      return IconButton(
        icon: const Icon(Icons.arrow_back),
        onPressed: () {
          Navigator.of(context).pop();
        },
      );
    } else {
      return null; // No back button if we can't pop the stack
    }
  }

  void _goBranch(int index) {
    widget.navigationShell.goBranch(index,
        initialLocation: index == widget.navigationShell.currentIndex);
  }
}
