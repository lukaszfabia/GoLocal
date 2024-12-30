import 'package:bloc/bloc.dart';
import 'package:golocal/src/auth/auth_service.dart';
import 'package:golocal/src/user/domain/user.dart';
part 'auth_event.dart';
part 'auth_state.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  final AuthService authService;
  late User signedInUser;
  AuthBloc({required this.authService}) : super(AuthInitial()) {
    on<SignInWithEmail>(
      (event, emit) async {
        emit(AuthLoading());
        final user = await signInWithEmail(
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
    return authService.signUpWithEmail(
        email, password, birthDate, firstName, lastName);
  }

  Future<User?> signInWithEmail(String email, String password) async {
    return authService.signInWithEmail(email, password);
  }
}
