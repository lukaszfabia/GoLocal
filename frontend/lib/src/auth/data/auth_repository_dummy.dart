import 'package:golocal/src/auth/auth_exceptions.dart';

import 'iauth_repository.dart';

class AuthRepositoryDummy implements IAuthRepository {
  bool _isAuthenticated = false;
  Map<String, String> dummyCredentials = {
    'golocal@gmail.com': "Golocal1!",
    'a@a.a': "AAAaaa1!",
  };
  @override
  Future<bool> initialCheck() {
    return Future.delayed(Duration(seconds: 1), () => _isAuthenticated);
  }

  @override
  Future<void> logout() {
    return Future.delayed(Duration(milliseconds: 250), () {
      _isAuthenticated = false;
    });
  }

  @override
  Future<void> signInWithEmail(String email, String password) {
    return Future.delayed(Duration(seconds: 1), () {
      if (dummyCredentials.containsKey(email) &&
          dummyCredentials[email] == password) {
        _isAuthenticated = true;
      } else if (dummyCredentials.containsKey(email)) {
        throw AuthException("Invalid password");
      } else {
        throw AuthException("User with this email does not exist");
      }
    });
  }

  @override
  Future<void> signUpWithEmail(
      String email, String firstName, String lastName, String password) {
    return Future.delayed(Duration(seconds: 1), () {
      if (dummyCredentials.containsKey(email)) {
        throw AuthException("User with this email already exists");
      } else {
        dummyCredentials[email] = password;
        _isAuthenticated = true;
      }
    });
  }
}
