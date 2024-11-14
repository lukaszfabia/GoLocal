part of 'auth_bloc.dart';

@immutable
sealed class AuthEvent {}

class SignIn extends AuthEvent {}

final class SignInWithEmail extends SignIn {
  final String email;
  final String password;
  SignInWithEmail({required this.email, required this.password});
}

class SignUp extends AuthEvent {}

final class SignUpWithEmail extends SignUp {
  final String email;
  final String password;
  final DateTime birthDate;
  final String firstName;
  final String lastName;
  SignUpWithEmail(
      {required this.email,
      required this.password,
      required this.birthDate,
      required this.firstName,
      required this.lastName});
}

final class Logout extends AuthEvent {}
