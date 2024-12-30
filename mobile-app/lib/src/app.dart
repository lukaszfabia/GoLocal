import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/auth/auth_service.dart';
import 'package:golocal/src/event/ui/events_view_page.dart';
import 'package:golocal/src/login/bloc/auth_bloc.dart';
import 'package:golocal/src/login/login_page.dart';
import 'package:golocal/src/navigation/navigation_bloc.dart';
import 'package:golocal/src/preference_survey/preference_survey_view.dart';
import 'package:golocal/src/shared/navbar.dart';

class GoLocalApp extends StatelessWidget {
  const GoLocalApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider(create: (context) => AuthBloc(authService: AuthService())),
        BlocProvider(create: (context) => NavigationBloc()),
      ],
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        home: BlocBuilder<AuthBloc, AuthState>(
          builder: (context, authState) {
            if (authState is AuthLoading) {
              return const Center(child: CircularProgressIndicator());
            } else if (authState is AuthSuccess) {
              return BlocBuilder<NavigationBloc, NavigationState>(
                builder: (context, navState) {
                  Widget page;
                  if (navState is EventsPageState) {
                    page = EventsViewPage();
                    // } else if (navState is ProfilePageState) {
                    //   page = ProfilePage();
                  } else if (navState is PreferenceSurveyState) {
                    page = PreferenceSurveyView();
                  } else {
                    page = EventsViewPage(); // Default fallback
                  }
                  return Scaffold(
                    appBar: AppBar(title: Text('GoLocal')),
                    body: page,
                    bottomNavigationBar: NavBar(),
                  );
                },
              );
            } else {
              return const LoginPage();
            }
          },
        ),
      ),
    );
  }
}
