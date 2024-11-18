import 'package:flutter/material.dart';
import 'package:golocal/src/login/ui/login_form.dart';
import 'package:golocal/src/shared/background.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: Stack(
        fit: StackFit.loose,
        alignment: Alignment.center,
        clipBehavior: Clip.hardEdge,
        children: [
          BackgroundWidget(),
          Center(
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 24),
              child: LoginForm(),
            ),
          )
        ],
      ),
    );
  }
}
