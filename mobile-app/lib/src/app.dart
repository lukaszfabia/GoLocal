import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/auth/auth_repository.dart';
import 'package:golocal/src/event/bloc/events_bloc.dart';
import 'package:golocal/src/event/data/events_repository_dummy.dart';
import 'package:golocal/src/event/ui/events_view_page.dart';
import 'package:golocal/src/auth/bloc/auth_bloc.dart';
import 'package:golocal/src/auth/ui/signin_page.dart';

class GoLocalApp extends StatelessWidget {
  const GoLocalApp({super.key});

  @override
  Widget build(BuildContext context) {
    return RepositoryProvider(
      create: (context) => AuthRepository(),
      child: MultiBlocProvider(
        providers: [
          BlocProvider(
            create: (context) => AuthBloc(context.read<AuthRepository>())
              ..add(const AuthInitialCheck()),
          ),
          BlocProvider(
            create: (context) => EventsBloc(EventsRepositoryDummy()),
          ),
        ],
        child: MaterialApp(
          debugShowCheckedModeBanner: false,
          home: BlocConsumer<AuthBloc, AuthState>(
            listener: (context, state) {
              if (state is AuthError) {
                ScaffoldMessenger.of(context)
                    .showSnackBar(SnackBar(content: Text(state.message)));
              }
              if (state is Authenticated) {
                BlocProvider.of<EventsBloc>(context).add(const FetchEvents());
              }
            },
            builder: (context, state) {
              if (state is AuthInitial) {
                return const SplashScreen();
              }
              if (state is Authenticated) {
                return EventsViewPage();
              }
              return SignInPage();
            },
          ),
        ),
      ),
    );
  }
}

class SplashScreen extends StatelessWidget {
  const SplashScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: const CircularProgressIndicator(),
      ),
    );
  }
}
