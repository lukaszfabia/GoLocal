import 'package:dio/dio.dart';
import 'package:golocal/src/auth/auth_exceptions.dart';
import 'package:golocal/src/auth/auth_service.dart';
import 'package:golocal/src/jwt_token_storage.dart';

class AuthRepository {
  final AuthService _authService = AuthService();
  AuthRepository();

  Future<void> signUpWithEmail(
      String email, String firstName, String lastName, String password) async {
    final data = {
      'email': email,
      'firstName': firstName,
      'lastName': lastName,
      'password': password,
    };
    try {
      print("starting response");
      Response response = await _authService.signUpWithEmail(data);
      print(response.data);
      await TokenStorage().saveAccessToken(response.data['data']['access']);
      await TokenStorage().saveRefreshToken(response.data['data']['refresh']);
    } on DioException catch (e) {
      print(e.toString());
      throw AuthException(e.response?.data['data']);
    } catch (e) {
      print(e.toString());
      throw AuthException(e.toString());
    }
  }

  Future<void> signInWithEmail(String email, String password) async {
    final data = {
      'email': email,
      'password': password,
    };
    try {
      Response response = await _authService.signInWithEmail(data);
      await TokenStorage().saveAccessToken(response.data['data']['access']);
      await TokenStorage().saveRefreshToken(response.data['data']['refresh']);
    } on DioException catch (e) {
      throw AuthException(e.response?.data['data']);
    } catch (e) {
      throw AuthException(e.toString());
    }
  }

  Future<void> logout() async {
    await _authService.logout();
    await TokenStorage().clearTokens();
  }

  Future<bool> initialCheck() async {
    return await TokenStorage().hasTokens();
  }
}
