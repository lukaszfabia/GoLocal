import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/auth/data/auth_repository.dart';
import 'package:golocal/src/auth/data/iauth_repository.dart';
import 'package:golocal/src/event/events_page/bloc/events_bloc.dart';
import 'package:golocal/src/event/data/impl/events_repository_impl.dart';
import 'package:golocal/src/auth/bloc/auth_bloc.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/vote/data/ivotes_repository.dart';
import 'package:golocal/src/vote/bloc/vote_bloc.dart';
import 'package:golocal/src/survey/data/ipreference_survey_repository.dart';
import 'package:golocal/src/survey/data/preference_survey_repository.dart';
import 'package:golocal/src/survey/bloc/preference_survey_bloc.dart';
import 'package:golocal/src/routing/router.dart';
import 'package:golocal/src/vote/data/votes_repository_impl.dart';

class GoLocalApp extends StatelessWidget {
  const GoLocalApp({super.key});

  IEventsRepository get eventsRepository => EventsRepositoryImpl();
  IAuthRepository get authRepository => AuthRepository();
  IVotesRepository get votesRepository => VotesRepositoryImpl();
  IPreferenceSurveyRepository get preferenceSurveyRepository =>
      PreferenceSurveyRepository();

  @override
  Widget build(BuildContext context) {
    return MultiRepositoryProvider(
      providers: [
        RepositoryProvider(
          create: (context) => authRepository,
        ),
        RepositoryProvider(
          create: (context) => eventsRepository,
        ),
        RepositoryProvider(
          create: (context) => votesRepository,
        ),
        RepositoryProvider(
          create: (context) => preferenceSurveyRepository,
        ),
      ],
      child: MultiBlocProvider(
        providers: [
          BlocProvider(
            lazy: false,
            create: (context) => AuthBloc(context.read<IAuthRepository>())
              ..add(const AuthInitialCheck()),
          ),
          BlocProvider(
            create: (context) => EventsBloc(context.read<IEventsRepository>()),
          ),
          BlocProvider(
            create: (context) => VoteBloc(context.read<IVotesRepository>()),
          ),
          BlocProvider(
              create: (context) => PreferenceSurveyBloc(
                  context.read<IPreferenceSurveyRepository>())),
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
