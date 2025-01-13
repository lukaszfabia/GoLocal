import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/auth/bloc/auth_bloc.dart';

class SignUpPage extends StatelessWidget {
  SignUpPage({super.key});
  final _formKey = GlobalKey<FormState>();
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  final TextEditingController _firstNameController = TextEditingController();
  final TextEditingController _lastNameController = TextEditingController();
  final TextEditingController _retypedPasswordController =
      TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Row(
              children: [
                Flexible(
                  child: TextFormField(
                    controller: _firstNameController,
                    expands: false,
                    decoration: InputDecoration(
                      border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(10.0)),
                      prefixIcon: Icon(Icons.person),
                      labelText: 'First Name',
                    ),
                  ),
                ),
                SizedBox(width: 10.0),
                Flexible(
                  child: TextFormField(
                    controller: _lastNameController,
                    expands: false,
                    decoration: InputDecoration(
                      border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(10.0)),
                      prefixIcon: Icon(Icons.person),
                      labelText: 'Last Name',
                    ),
                  ),
                ),
              ],
            ),
            SizedBox(height: 10.0),
            TextFormField(
              controller: _emailController,
              expands: false,
              decoration: InputDecoration(
                border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(10.0)),
                prefixIcon: Icon(Icons.email),
                labelText: 'Email',
              ),
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
            ),
            SizedBox(height: 10.0),
            TextFormField(
              controller: _retypedPasswordController,
              expands: false,
              decoration: InputDecoration(
                border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(10.0)),
                prefixIcon: Icon(Icons.lock),
                labelText: 'Retype Password',
              ),
              validator: (value) {
                if (value != _passwordController.text) {
                  return 'Passwords do not match';
                }
                return null;
              },
            ),
            SizedBox(height: 10.0),
            OutlinedButton(
              onPressed: () {
                BlocProvider.of<AuthBloc>(context).add(AuthSignUp(
                    _emailController.text,
                    _firstNameController.text,
                    _lastNameController.text,
                    _passwordController.text));
              },
              child: Text('Sign Up'),
            ),
          ],
        ),
      ),
    );
  }
}
