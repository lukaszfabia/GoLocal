part of 'auth_bloc.dart';

sealed class AuthEvent extends Equatable {
  const AuthEvent();

  @override
  List<Object> get props => [];
}

final class AuthSignIn extends AuthEvent {
  final String email;
  final String password;

  const AuthSignIn(this.email, this.password);

  @override
  List<Object> get props => [email, password];
}

final class AuthSignUp extends AuthEvent {
  final String email;
  final String firstName;
  final String lastName;
  final String password;

  const AuthSignUp(this.email, this.firstName, this.lastName, this.password);

  @override
  List<Object> get props => [email, firstName, lastName, password];
}

final class AuthInitialCheck extends AuthEvent {
  const AuthInitialCheck();
}

final class AuthLogout extends AuthEvent {
  const AuthLogout();
}
