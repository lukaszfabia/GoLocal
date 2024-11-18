import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/login/auth/auth_service.dart';
import 'package:golocal/src/event/ui/events_view_page.dart';
import 'package:golocal/src/login/bloc/auth_bloc.dart';
import 'package:golocal/src/login/login_page.dart';

class GoLocalApp extends StatelessWidget {
  GoLocalApp({super.key});
  final AuthService authService = AuthService();

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: BlocProvider(
          create: (context) => AuthBloc(authSerivce: authService),
          child: BlocBuilder<AuthBloc, AuthState>(
            builder: (context, state) {
              if (state is AuthSuccess) {
                return EventsViewPage();
              } else {
                return const LoginPage();
              }
            },
          )),
    );
  }
}
