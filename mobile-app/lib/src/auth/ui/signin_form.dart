import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/auth/bloc/auth_bloc.dart';
import 'package:golocal/src/auth/ui/signup_form.dart';

class SignInPage extends StatelessWidget {
  SignInPage({super.key});
  final _formKey = GlobalKey<FormState>();
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            TextFormField(
              controller: _emailController,
              expands: false,
              decoration: InputDecoration(
                border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(10.0)),
                prefixIcon: Icon(Icons.email),
                labelText: 'Email',
              ),
              validator: (value) => value!.isEmpty ? 'Email is required' : null,
            ),
            SizedBox(height: 10.0),
            TextFormField(
              controller: _passwordController,
              expands: false,
              decoration: InputDecoration(
                border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(10.0)),
                prefixIcon: Icon(Icons.lock),
                labelText: 'Password',
              ),
              validator: (value) =>
                  value!.isEmpty ? 'Password is required' : null,
            ),
            OutlinedButton(
              onPressed: () {
                if (_formKey.currentState!.validate()) {
                  BlocProvider.of<AuthBloc>(context).add(AuthSignIn(
                      _emailController.text, _passwordController.text));
                }
              },
              child: Text("Sign in"),
            ),
          ],
        ),
      ),
    );
  }
}
