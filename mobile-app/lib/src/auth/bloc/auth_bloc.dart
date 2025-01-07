import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:golocal/src/auth/auth_exceptions.dart';
import 'package:golocal/src/auth/data/iauth_repository.dart';

part 'auth_event.dart';
part 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final IAuthRepository _authRepository;
  AuthBloc(this._authRepository) : super(AuthInitial()) {
    on<AuthSignIn>((event, emit) async {
      try {
        await _authRepository.signInWithEmail(event.email, event.password);
        emit(Authenticated("You are logged in"));
      } on AuthException catch (e) {
        emit(AuthError(e.toString()));
      }
    });
    on<AuthSignUp>((event, emit) async {
      try {
        await _authRepository.signUpWithEmail(
            event.email, event.firstName, event.lastName, event.password);
        emit(Authenticated("Sign up success"));
      } on AuthException catch (e) {
        emit(AuthError(e.toString()));
      }
    });
    on<AuthInitialCheck>((event, emit) async {
      bool hasTokens = await _authRepository.initialCheck();
      if (hasTokens) {
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
    print(event);
  }

  @override
  void onTransition(Transition<AuthEvent, AuthState> transition) {
    super.onTransition(transition);
    print(transition);
  }
}
