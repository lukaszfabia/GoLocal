import 'package:golocal/src/user/domain/user.dart';

// TODO: implement real auth service
class AuthService {
  Future<User?> signUpWithEmail(String email, String password,
      DateTime birthDate, String firstName, String lastName) async {
    return Future.delayed(Duration(seconds: 1), () {
      return User(
          id: 1,
          email: email,
          firstName: firstName,
          lastName: lastName,
          birthDate: birthDate,
          isVerified: true,
          isPremium: false);
    });
  }

  Future<User?> signInWithEmail(String email, String password) {
    return Future.delayed(Duration(seconds: 1), () {
      return User(
          id: 1,
          email: email,
          firstName: 'Peter',
          lastName: 'Fun',
          birthDate: DateTime(1965, 7, 27),
          isVerified: true,
          isPremium: false);
    });
  }

  Future<void> signOut() {
    return Future.delayed(Duration(milliseconds: 500), () {});
  }
}
