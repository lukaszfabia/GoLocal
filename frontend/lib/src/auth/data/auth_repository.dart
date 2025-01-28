import 'package:dio/dio.dart';
import 'package:golocal/src/auth/auth_exceptions.dart';
import 'package:golocal/src/auth/auth_service.dart';
import 'package:golocal/src/jwt_token_storage.dart';

import 'iauth_repository.dart';

class AuthRepository implements IAuthRepository {
  final AuthService _authService = AuthService();
  AuthRepository();

  @override
  Future<void> signUpWithEmail(
      String email, String firstName, String lastName, String password) async {
    final data = {
      'email': email,
      'firstName': firstName,
      'lastName': lastName,
      'password': password,
    };
    try {
      Response response = await _authService.signUpWithEmail(data);
      await TokenStorage().saveAccessToken(response.data['data']['access']);
      await TokenStorage().saveRefreshToken(response.data['data']['refresh']);
    } on DioException catch (e) {
      throw AuthException(e.response?.data['data']);
    } catch (e) {
      throw AuthException(e.toString());
    }
  }

  @override
  Future<void> signInWithEmail(String email, String password) async {
    final data = {
      'email': email,
      'password': password,
    };
    try {
      Response response = await _authService.signInWithEmail(data);
      print(response);

      await TokenStorage().saveAccessToken(response.data['data']['access']);
      await TokenStorage().saveRefreshToken(response.data['data']['refresh']);
    } on DioException catch (e) {
      print("dio exception");

      throw AuthException(e.response?.data['data']);
    } catch (e) {
      print(e.toString());
      throw AuthException(e.toString());
    }
  }

  @override
  Future<void> logout() async {
    await _authService.logout();
    await TokenStorage().clearTokens();
  }

  @override
  Future<bool> initialCheck() async {
    return await TokenStorage().hasTokens();
  }
}
