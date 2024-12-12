import 'package:flutter/material.dart';
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
            child: LoginForm(),
          )
        ],
      ),
    );
  }
}

class LoginForm extends StatelessWidget {
  const LoginForm({super.key});
  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        const Text('Login'),
        const SizedBox(height: 20),
        TextField(
          autofillHints: [AutofillHints.email],
          decoration: InputDecoration(
            labelText: 'Email',
          ),
        ),
        const SizedBox(height: 20),
        const TextField(
          decoration: InputDecoration(
            labelText: 'Password',
          ),
        ),
        const SizedBox(height: 20),
        ElevatedButton(
          onPressed: () {},
          child: const Text('Login'),
        ),
      ],
    );
  }
}
