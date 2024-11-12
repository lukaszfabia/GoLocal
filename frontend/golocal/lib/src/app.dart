import 'package:flutter/material.dart';
import 'package:golocal/src/login/login_screen.dart';

class GoLocalApp extends StatelessWidget {
  const GoLocalApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: LoginScreen(),
    );
  }
}
