import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/navigation/navigation_bloc.dart';

class NavBar extends StatelessWidget {
  const NavBar({super.key});

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      items: const <BottomNavigationBarItem>[
        BottomNavigationBarItem(
          icon: Icon(Icons.event),
          label: 'Events',
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.person),
          label: 'Profile',
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.settings),
          label: 'Settings',
        ),
      ],
      onTap: (index) {
        switch (index) {
          case 0:
            BlocProvider.of<NavigationBloc>(context).add(GoToEventsPage());
            break;
          case 1:
            BlocProvider.of<NavigationBloc>(context).add(GoToProfilePage());
            break;
          case 2:
            BlocProvider.of<NavigationBloc>(context)
                .add(GoToPreferenceSurvey());
            break;
        }
      },
    );
  }
}
