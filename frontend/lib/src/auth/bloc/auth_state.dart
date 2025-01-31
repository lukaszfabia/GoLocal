part of 'auth_bloc.dart';

/// The base class for all authentication states.
///
/// This class extends [Equatable] to allow for easy comparison of states.
sealed class AuthState extends Equatable {
  const AuthState();

  @override
  List<Object> get props => [];
}

/// Represents the initial state of authentication.
final class AuthInitial extends AuthState {}

/// Represents the loading state of authentication.
final class AuthLoading extends AuthState {}

/// Represents an error state in authentication.
///
/// Contains an error message describing the issue.
final class AuthError extends AuthState {
  final String message;

  const AuthError(this.message);

  @override
  List<Object> get props => [message];
}

/// Represents a successful authentication state.
///
/// Contains a message, typically a success message or token.
final class Authenticated extends AuthState {
  final String message;

  const Authenticated(this.message);

  @override
  List<Object> get props => [message];
}

/// Represents an unauthenticated state.
final class Unathenticated extends AuthState {
  const Unathenticated();

  @override
  List<Object> get props => [];
}
