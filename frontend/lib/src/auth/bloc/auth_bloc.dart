import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:golocal/src/auth/auth_exceptions.dart';
import 'package:golocal/src/auth/data/iauth_repository.dart';
import 'package:golocal/src/notifications_service/notification_service.dart';

part 'auth_event.dart';
part 'auth_state.dart';

/// AuthBloc handles authentication-related events and states.
///
/// This Bloc listens to the following events:
/// - [AuthSignIn]: Handles user sign-in with email and password.
/// - [AuthSignUp]: Handles user sign-up with email, first name, last name, and password.
/// - [AuthInitialCheck]: Checks if the user has valid tokens and is already authenticated.
/// - [AuthLogout]: Handles user logout.
///
/// The Bloc emits the following states:
/// - [AuthInitial]: Initial state of the authentication.
/// - [Authenticated]: State when the user is successfully authenticated.
/// - [AuthError]: State when there is an authentication error.
/// - [Unathenticated]: State when the user is not authenticated.
class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final IAuthRepository _authRepository;
  AuthBloc(this._authRepository) : super(AuthInitial()) {
    on<AuthSignIn>((event, emit) async {
      try {
        await _authRepository.signInWithEmail(event.email, event.password);

        await NotificationService.instance.registerDevice();

        emit(Authenticated("You are logged in"));
      } on AuthException catch (e) {
        emit(AuthError(e.toString()));
      }
    });
    on<AuthSignUp>((event, emit) async {
      try {
        await _authRepository.signUpWithEmail(
            event.email, event.firstName, event.lastName, event.password);

        await NotificationService.instance.registerDevice();

        emit(Authenticated("Sign up success"));
      } on AuthException catch (e) {
        emit(AuthError(e.toString()));
      }
    });
    on<AuthInitialCheck>((event, emit) async {
      bool hasTokens = await _authRepository.initialCheck();
      if (hasTokens) {
        await NotificationService.instance.registerDevice();

        emit(Authenticated("You are logged in"));
      } else {
        emit(Unathenticated());
      }
    });
    on<AuthLogout>((event, emit) async {
      await _authRepository.logout();
      emit(Unathenticated());
    });
  }

  @override
  void onEvent(AuthEvent event) {
    super.onEvent(event);
  }

  @override
  void onTransition(Transition<AuthEvent, AuthState> transition) {
    super.onTransition(transition);
  }
}
