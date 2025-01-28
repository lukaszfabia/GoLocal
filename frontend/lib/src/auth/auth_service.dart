import 'package:dio/dio.dart';
import 'package:golocal/src/dio_client.dart';

class AuthService {
  final Dio _dio = DioClient().dio;

  Future<Response> signInWithEmail(Map<String, dynamic> data) {
    print("logging");
    return _dio.post('/login/', data: data);
  }

  Future<Response> signUpWithEmail(Map<String, dynamic> data) {
    return _dio.post('/sign-up/', data: data);
  }

  Future<Response> logout() {
    return Future<Response>.value(
        Response(requestOptions: RequestOptions(path: '')));
  }
}
