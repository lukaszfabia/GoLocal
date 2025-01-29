import 'package:flutter/widgets.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/app.dart';
import 'package:golocal/src/auth/bloc/auth_bloc.dart';
import 'package:golocal/src/auth/ui/auth_screen.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/manage_page/event_create_page.dart';
import 'package:golocal/src/event/promote_page/promote_event_page.dart';
import 'package:golocal/src/event/report_page/report_event_page.dart';
import 'package:golocal/src/event/events_page/ui/event_detail_page.dart';
import 'package:golocal/src/event/events_page/ui/events_map.dart';
import 'package:golocal/src/event/events_page/ui/events_view_page.dart';
import 'package:golocal/src/preference_survey/ui/preference_survey_page.dart';
import 'package:golocal/src/shared/scaffold_shell.dart';
import 'package:golocal/src/shared/streamtolistenable.dart';
import 'package:golocal/src/user/ui/profile_page.dart';
import 'package:golocal/src/vote/ui/votes_page.dart';
import 'package:golocal/src/vote/ui/votes_for_event_page.dart';

// if you want to add new routes eg. /home/addevent or /home/event/:id/report you can do it here
// just add another route to the enum and go to the statefull shelll branch

abstract class AppRouter {
  static GoRouter? _router;

  static final GlobalKey<NavigatorState> rootNavigatorKey =
      GlobalKey<NavigatorState>(debugLabel: "root");
  static final GlobalKey<NavigatorState> shellNavigatorKeyHome =
      GlobalKey<NavigatorState>(debugLabel: 'home');
  static final GlobalKey<NavigatorState> shellNavigatorKeyProfile =
      GlobalKey<NavigatorState>(debugLabel: 'profile');
  static final GlobalKey<NavigatorState> shellNavigatorKeyVotes =
      GlobalKey<NavigatorState>(debugLabel: 'votes');

  static GoRouter router(AuthBloc authBloc) {
    _router ??= GoRouter(
      debugLogDiagnostics: true,
      navigatorKey: rootNavigatorKey,
      initialLocation: AppRoute.splash.path,
      routes: [
        GoRoute(
            path: '/',
            redirect: (context, state) {
              bool isLoggedIn = authBloc.state is Authenticated;
              return isLoggedIn ? AppRoute.events.path : AppRoute.auth.path;
            }),
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
          builder: (context, state, navigationShell) {
            return ScaffoldShell(navigationShell: navigationShell, title: null);
          },
          branches: [
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
            // events
            StatefulShellBranch(
              navigatorKey: shellNavigatorKeyHome,
              routes: [
                GoRoute(
                  path: AppRoute.events.path,
                  builder: (context, state) {
                    return const EventsViewPage();
                  },
                  routes: [
                    GoRoute(
                      path: 'create',
                      builder: (context, state) {
                        return EventCreatePage();
                      },
                    ),
                    GoRoute(
                      path: 'edit',
                      builder: (context, state) {
                        final event = state.extra as Event;
                        return EventCreatePage(event: event);
                      },
                    ),
                    GoRoute(
                      path: AppRoute.eventDetail.path,
                      builder: (context, state) {
                        final event = state.extra as Event;
                        return EventDetailPage(event: event);
                      },
                      routes: [
                        GoRoute(
                          path: "promote",
                          builder: (context, state) {
                            final event = state.extra as Event;

                            return PromoteEventPage(event);
                          },
                        ),
                        GoRoute(
                          path: 'report',
                          builder: (context, state) {
                            final event = state.extra as Event;
                            return ReportEventPage(event: event);
                          },
                        ),
                        GoRoute(
                          path: "votes",
                          builder: (context, state) {
                            final event = state.extra as Event;
                            return VotesForEventPage(event);
                          },
                        ),
                      ],
                    ),
                  ],
                ),
              ],
            ),
            StatefulShellBranch(navigatorKey: shellNavigatorKeyVotes, routes: [
              GoRoute(
                  path: AppRoute.votes.path,
                  builder: (context, state) {
                    return const VotesPage();
                  })
            ]),
            StatefulShellBranch(
              navigatorKey: shellNavigatorKeyProfile,
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
                  ],
                ),
              ],
            ),
          ],
        ),
      ],
      redirect: (context, state) {
        final bool isLoggingIn = (state.matchedLocation == AppRoute.auth.path ||
            state.matchedLocation == AppRoute.splash.path);
        final AuthState authState = authBloc.state;
        final bool isLoggedIn = authState is Authenticated;

        if (isLoggedIn && isLoggingIn) {
          return AppRoute.events.path;
        } else if (authState is Unathenticated) {
          return AppRoute.auth.path;
        }
        return null;
      },
      // Listen to the authBloc stream to refresh the current screen (e.g., if the user logs in or out)
      refreshListenable: StreamToListenable([authBloc.stream]),
    );
    return _router!;
  }
}

// Much easier to maintain and add new routes
enum AppRoute {
  splash(path: '/splash', name: 'splash'),

  events(path: '/events', name: 'events', title: "Events"),
  eventDetail(path: '/:id', name: 'eventDetail'),

  profile(path: '/profile', name: 'profile'),
  survey(path: '/survey', name: 'survey', title: "Preference Survey"),
  map(path: '/map', name: 'map'),
  votes(path: '/votes', name: 'votes'),
  auth(path: '/auth', name: 'auth');

  const AppRoute({required this.path, required this.name, this.title});

  final String path;
  final String name;
  final String? title;
}
