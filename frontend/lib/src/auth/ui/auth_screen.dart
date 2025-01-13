import 'package:flutter/material.dart';
import 'package:golocal/src/auth/ui/signin_form.dart';
import 'package:golocal/src/auth/ui/signup_form.dart';

class AuthScreen extends StatefulWidget {
  const AuthScreen({super.key});

  @override
  State<AuthScreen> createState() => _AuthScreenState();
}

class _AuthScreenState extends State<AuthScreen> {
  final PageController _pageController = PageController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        decoration: BoxDecoration(
          gradient: LinearGradient(
            colors: [
              Colors.blue.shade100,
              Colors.red.shade200,
            ],
            begin: Alignment.topCenter,
            end: Alignment.bottomCenter,
          ),
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            Padding(
              padding: const EdgeInsets.all(24.0),
              child: Text("GoLocal", style: TextStyle(fontSize: 48.0)),
            ),
            Flexible(
              flex: 2,
              child: PageView(
                controller: _pageController,
                children: [
                  SignInPage(),
                  SignUpPage(),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
