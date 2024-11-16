import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/login/bloc/auth_bloc.dart';

class LoginForm extends StatefulWidget {
  const LoginForm({
    super.key,
  });

  @override
  State<LoginForm> createState() => _LoginFormState();
}

class _LoginFormState extends State<LoginForm> {
  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  final _formKey = GlobalKey<FormState>();
  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: Colors.grey.withOpacity(0.3),
        borderRadius: BorderRadius.circular(10),
        boxShadow: [
          BoxShadow(
            color: Colors.white.withOpacity(0.1),
            blurRadius: 10,
            spreadRadius: 5,
          ),
        ],
      ),
      child: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Form(
            key: _formKey,
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                Text("Sign in",
                    style:
                        TextStyle(fontSize: 24, fontWeight: FontWeight.bold)),
                EmailFormField(emailController: emailController),
                SizedBox(height: 30),
                PasswordField(passwordController: passwordController),
                SizedBox(height: 30),
                TextButton(
                    onPressed: () {
                      _formKey.currentState!.validate()
                          ? BlocProvider.of<AuthBloc>(context).add(
                              SignInWithEmail(
                                  email: emailController.value.text,
                                  password: passwordController.value.text))
                          : null;
                    },
                    child: Text("Sign In"))
              ],
            )),
      ),
    );
  }

  @override
  void dispose() {
    emailController.dispose();
    passwordController.dispose();
    super.dispose();
  }
}

class PasswordField extends StatelessWidget {
  const PasswordField({
    super.key,
    required this.passwordController,
  });

  final TextEditingController passwordController;
  FormFieldValidator<String>? validator(String value) {
    return (value) {
      if (value == null || value.isEmpty) {
        return "Password cannot be empty";
      }
      if (value.length < 8) {
        return "Password must be at least 6 characters long";
      }
      return null;
    };
  }

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      controller: passwordController,
      expands: false,
      validator: validator(passwordController.value.text),
      keyboardType: TextInputType.visiblePassword,
      decoration: InputDecoration(
        hintText: "enter your password",
        prefixIcon: Icon(Icons.lock_outline),
      ),
      obscureText: true,
    );
  }
}

class EmailFormField extends StatelessWidget {
  const EmailFormField({
    super.key,
    required this.emailController,
  });

  final TextEditingController emailController;

  FormFieldValidator<String>? validator(String value) {
    return (value) {
      if (value == null || value.isEmpty) {
        return "Email cannot be empty";
      }
      if (!value.contains("@")) {
        return "Please provide a valid email";
      }
      return null;
    };
  }

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      expands: false,
      controller: emailController,
      keyboardType: TextInputType.emailAddress,
      validator: validator(emailController.value.text),
      decoration: InputDecoration(
        hintText: "enter you email",
        prefixIcon: Icon(Icons.email_outlined),
      ),
      enableSuggestions: true,
      autofillHints: [
        AutofillHints.email,
      ],
    );
  }
}
