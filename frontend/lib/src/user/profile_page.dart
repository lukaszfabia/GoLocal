import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/auth/bloc/auth_bloc.dart';

class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        const Text('Profile Page'),
        ElevatedButton(
          onPressed: () {
            context.read<AuthBloc>().add(const AuthLogout());
          },
          child: const Text('Logout'),
        ),
      ],
    );
  }
}
