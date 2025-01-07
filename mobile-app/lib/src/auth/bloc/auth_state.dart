part of 'auth_bloc.dart';

sealed class AuthState extends Equatable {
  const AuthState();

  @override
  List<Object> get props => [];
}

final class AuthInitial extends AuthState {}

final class AuthLoading extends AuthState {}

final class AuthError extends AuthState {
  final String message;

  const AuthError(this.message);

  @override
  List<Object> get props => [message];
}

final class Authenticated extends AuthState {
  final String message;

  const Authenticated(this.message);

  @override
  List<Object> get props => [message];
}

final class Unathenticated extends AuthState {
  const Unathenticated();

  @override
  List<Object> get props => [];
}
