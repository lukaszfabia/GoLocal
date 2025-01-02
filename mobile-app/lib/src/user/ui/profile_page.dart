import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/auth/bloc/auth_bloc.dart';
import 'package:golocal/src/preference_survey/ui/preference_survey_page.dart';
import 'package:golocal/src/routing/router.dart';
import 'package:golocal/src/shared/scaffold_shell.dart';

class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        ElevatedButton(
          onPressed: () {
            context.read<AuthBloc>().add(const AuthLogout());
          },
          child: const Text('Log out'),
        ),
        ElevatedButton(
          onPressed: () {
            context.push(AppRoute.profile.path + AppRoute.survey.path);
          },
          child: const Text('Fill out the preference survey'),
        ),
      ],
    );
  }
}
