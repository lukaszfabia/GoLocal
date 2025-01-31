import 'package:dio/dio.dart';
import 'package:golocal/src/dio_client.dart';

/// A service class that handles authentication-related API calls.
class AuthService {
  /// An instance of Dio for making HTTP requests.
  final Dio _dio = DioClient().dio;

  /// Sends a POST request to the '/login/' endpoint with the provided [data].
  ///
  /// [data] is a map containing the email and password for signing in.
  ///
  /// Returns a [Future] that resolves to a [Response] from the server.
  Future<Response> signInWithEmail(Map<String, dynamic> data) {
    print("logging");
    return _dio.post('/login/', data: data);
  }

  /// Sends a POST request to the '/sign-up/' endpoint with the provided [data].
  ///
  /// [data] is a map containing the necessary information for signing up.
  ///
  /// Returns a [Future] that resolves to a [Response] from the server.
  Future<Response> signUpWithEmail(Map<String, dynamic> data) {
    return _dio.post('/sign-up/', data: data);
  }

  /// Logs out the current user.
  ///
  /// Returns a [Future] that resolves to an empty [Response].
  Future<Response> logout() {
    return Future<Response>.value(
        Response(requestOptions: RequestOptions(path: '')));
  }

  /// Retrieves the current user's information.
  ///
  /// Sends a GET request to the '/user/' endpoint.
  ///
  /// Returns a [Future] that resolves to a [Response] containing the user's information.
  Future<Response> getUser() {
    return _dio.get('/user/');
  }
}
