import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/auth/auth_repository.dart';
import 'package:golocal/src/event/bloc/events_bloc.dart';
import 'package:golocal/src/event/data/events_repository_dummy.dart';
import 'package:golocal/src/auth/bloc/auth_bloc.dart';
import 'package:golocal/src/routing/router.dart';

class GoLocalApp extends StatelessWidget {
  const GoLocalApp({super.key});

  @override
  Widget build(BuildContext context) {
    return RepositoryProvider(
      create: (context) => AuthRepository(),
      child: MultiBlocProvider(
        providers: [
          BlocProvider(
            lazy: false,
            create: (context) => AuthBloc(context.read<AuthRepository>())
              ..add(const AuthInitialCheck()),
          ),
          BlocProvider(
            create: (context) => EventsBloc(EventsRepositoryDummy()),
          ),
        ],
        child: Builder(
          builder: (context) => MaterialApp.router(
            debugShowCheckedModeBanner: false,
            routerConfig: AppRouter.router(context.read<AuthBloc>()),
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
