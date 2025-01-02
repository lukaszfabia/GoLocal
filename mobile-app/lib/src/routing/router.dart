import 'package:flutter/widgets.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/app.dart';
import 'package:golocal/src/auth/bloc/auth_bloc.dart';
import 'package:golocal/src/auth/ui/auth_screen.dart';
import 'package:golocal/src/event/ui/events_map.dart';
import 'package:golocal/src/event/ui/events_view_page.dart';
import 'package:golocal/src/preference_survey/ui/preference_survey_page.dart';
import 'package:golocal/src/shared/scaffold_shell.dart';
import 'package:golocal/src/shared/streamtolistenable.dart';
import 'package:golocal/src/user/ui/profile_page.dart';

// if you want to add new routes eg. /home/addevent or /home/event/:id/report you can do it here
// just add another route to the enum and go to the statefull shelll branch

abstract class AppRouter {
  static final GlobalKey<NavigatorState> rootNavigatorKey =
      GlobalKey<NavigatorState>(debugLabel: "root");
  static final GlobalKey<NavigatorState> shellNavigatorKeyHome =
      GlobalKey<NavigatorState>();
  static final GlobalKey<NavigatorState> shellNavigatorKeyProfile =
      GlobalKey<NavigatorState>();

  static GoRouter router(AuthBloc authBloc) => GoRouter(
        navigatorKey: rootNavigatorKey,
        initialLocation: AppRoute.splash.path,
        routes: [
          GoRoute(
            path: AppRoute.auth.path,
            name: AppRoute.auth.name,
            builder: (context, state) => const AuthScreen(),
          ),
          GoRoute(
            path: AppRoute.splash.path,
            name: AppRoute.splash.name,
            builder: (context, state) => const SplashScreen(),
          ),
          // shell for main screens with navbar
          StatefulShellRoute.indexedStack(
            parentNavigatorKey: rootNavigatorKey,
            builder: (context, state, navigationShell) =>
                ScaffoldShell(navigationShell: navigationShell),
            branches: [
              // []
              StatefulShellBranch(
                routes: [
                  GoRoute(
                    path: AppRoute.map.path,
                    builder: (context, state) {
                      return const EventsMap();
                    },
                  ),
                ],
              ),
              StatefulShellBranch(
                routes: [
                  GoRoute(
                    path: AppRoute.home.path,
                    builder: (context, state) {
                      return const EventsViewPage();
                    },
                  ),
                ],
              ),
              StatefulShellBranch(
                routes: [
                  GoRoute(
                      path: AppRoute.profile.path,
                      builder: (context, state) {
                        return const ProfilePage();
                      },
                      routes: [
                        GoRoute(
                          path: AppRoute.survey.path,
                          builder: (context, state) {
                            return const PreferenceSurveyPage();
                          },
                        ),
                      ]),
                ],
              ),
            ],
          ),
        ],
        redirect: (context, state) {
          final bool isLoggingIn =
              (state.matchedLocation == AppRoute.auth.path ||
                  state.matchedLocation == AppRoute.splash.path);
          final AuthState authState = context.read<AuthBloc>().state;

          if (authState is AuthInitial) {
            return AppRoute.splash.path;
          } else if (authState is Authenticated && isLoggingIn) {
            return AppRoute.home.path;
          } else if (authState is Unathenticated) {
            return AppRoute.auth.path;
          }
          return null;
        },
        // Listen to the authBloc stream to refresh the current screen (e.g., if the user logs in or out)
        refreshListenable: StreamToListenable([authBloc.stream]),
      );
}

// Much easier to maintain and add new routes
enum AppRoute {
  splash(path: '/splash', name: 'splash'),
  home(path: '/home', name: 'home'),
  profile(path: '/profile', name: 'profile'),
  survey(path: '/profile/survey', name: 'survey'),
  map(path: '/map', name: 'map'),
  auth(path: '/auth', name: 'auth');

  const AppRoute({required this.path, required this.name});

  final String path;
  final String name;
}
