import 'package:bloc/bloc.dart';
import 'package:golocal/src/login/auth/auth_service.dart';
import 'package:golocal/src/user/domain/user.dart';
import 'package:meta/meta.dart';
part 'auth_event.dart';
part 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  //TODO: create auth servivce
  final AuthService authSerivce;
  late User signedInUser;
  AuthBloc({required this.authSerivce}) : super(AuthInitial()) {
    on<SignInWithEmail>(
      (event, emit) async {
        emit(AuthLoading());
        final user = await authSerivce.signInWithEmail(
          event.email,
          event.password,
        );
        if (user != null) {
          signedInUser = user;
          emit(AuthSuccess(user: user));
        } else {
          emit(AuthError(message: 'Something went wrong'));
        }
      },
    );
  }

  Future<User?> signUpWithEmail(String email, String password,
      DateTime birthDate, String firstName, String lastName) async {
    return authSerivce.signUpWithEmail(
        email, password, birthDate, firstName, lastName);
  }

  @override
  void onEvent(AuthEvent event) {
    print("event: $event called");
    super.onEvent(event);
  }
}
