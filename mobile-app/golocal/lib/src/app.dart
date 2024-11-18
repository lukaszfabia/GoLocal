import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/auth/auth_service.dart';
import 'package:golocal/src/event/ui/events_view_page.dart';
import 'package:golocal/src/login/bloc/auth_bloc.dart';
import 'package:golocal/src/login/login_page.dart';

class GoLocalApp extends StatelessWidget {
  const GoLocalApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: BlocProvider(
          create: (context) => AuthBloc(authSerivce: AuthService()),
          child: BlocBuilder<AuthBloc, AuthState>(
            builder: (context, state) {
              if (state is AuthLoading) {
                return const Center(child: CircularProgressIndicator());
              } else if (state is AuthSuccess) {
                return EventsViewPage();
              } else {
                return const LoginPage();
              }
            },
          )),
    );
  }
}
