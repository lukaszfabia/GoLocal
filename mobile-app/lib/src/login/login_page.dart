import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/login/bloc/auth_bloc.dart';
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
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();

  LoginForm({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        const Text('Login'),
        const SizedBox(height: 20),
        TextField(
          controller: emailController,
          autofillHints: [AutofillHints.email],
          decoration: InputDecoration(
            labelText: 'Email',
          ),
        ),
        const SizedBox(height: 20),
        TextField(
          controller: passwordController,
          decoration: InputDecoration(
            labelText: 'Password',
          ),
        ),
        const SizedBox(height: 20),
        ElevatedButton(
          onPressed: () {
            final email = emailController.text;
            final password = passwordController.text;
            context
                .read<AuthBloc>()
                .add(SignInWithEmail(email: email, password: password));
          },
          child: const Text('Login'),
        ),
      ],
    );
  }
}
