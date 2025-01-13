abstract class IAuthRepository {
  Future<void> signUpWithEmail(
      String email, String firstName, String lastName, String password);

  Future<void> signInWithEmail(String email, String password);

  Future<void> logout();

  Future<bool> initialCheck();
}
