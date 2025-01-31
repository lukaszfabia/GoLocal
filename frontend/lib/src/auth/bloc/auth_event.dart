part of 'auth_bloc.dart';

/// The base class for all authentication events, extending [Equatable] to
/// support value comparison.
sealed class AuthEvent extends Equatable {
  const AuthEvent();

  @override
  List<Object> get props => [];
}

/// Event triggered when a user attempts to sign in.
///
/// Contains the user's [email] and [password].
final class AuthSignIn extends AuthEvent {
  final String email;
  final String password;

  const AuthSignIn(this.email, this.password);

  @override
  List<Object> get props => [email, password];
}

/// Event triggered when a user attempts to sign up.
///
/// Contains the user's [email], [firstName], [lastName], and [password].
final class AuthSignUp extends AuthEvent {
  final String email;
  final String firstName;
  final String lastName;
  final String password;

  const AuthSignUp(this.email, this.firstName, this.lastName, this.password);

  @override
  List<Object> get props => [email, firstName, lastName, password];
}

/// Event triggered to perform an initial authentication check.
final class AuthInitialCheck extends AuthEvent {
  const AuthInitialCheck();
}

/// Event triggered when a user attempts to log out.
final class AuthLogout extends AuthEvent {
  const AuthLogout();
}
