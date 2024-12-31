import 'package:dio/dio.dart';

import 'jwt_token_storage.dart';
import 'dart:io' show Platform;

class DioClient {
  static final DioClient _instance = DioClient._internal();
  late final Dio dio;
  final TokenStorage _tokenStorage = TokenStorage();

  factory DioClient() => _instance;

  DioClient._internal() {
    dio = Dio(
      BaseOptions(
        baseUrl:
            'http://${Platform.isAndroid ? "10.0.2.2" : "127.0.0.1"}:8080/api',
        headers: {
          'Content-Type': 'application/json',
        },
        connectTimeout: Duration(seconds: 5),
        receiveTimeout: Duration(seconds: 10),
      ),
    );

    dio.interceptors.add(
      InterceptorsWrapper(
        onRequest: (options, handler) async {
          final accessToken = await _tokenStorage.getAccessToken();
          if (accessToken != null) {
            options.headers['Authorization'] = 'Bearer $accessToken';
          }
          return handler.next(options);
        },
        onError: (error, handler) async {
          // Unauthorized, try to refresh token if exists
          if (error.response?.statusCode == 401) {
            final refreshToken = await _tokenStorage.getRefreshToken();
            if (refreshToken != null) {
              try {
                final response = await dio.post('refresh-token/', data: {
                  'refresh': refreshToken,
                });
                final newAccessToken = response.data['data']['access'];
                await _tokenStorage.saveAccessToken(newAccessToken);
                final newRefreshToken = response.data['data']['refresh'];
                await _tokenStorage.saveRefreshToken(newRefreshToken);

                final options = error.response!.requestOptions;
                options.headers['Authorization'] = 'Bearer $newAccessToken';
                final retryResponse = await dio.fetch(options);
                return handler.resolve(retryResponse);
              } catch (e) {
                await _tokenStorage.clearTokens();
                return handler.reject(
                  DioException(
                      requestOptions: error.requestOptions,
                      error: "Session expired"),
                );
              }
            }
          }
          return handler.next(error);
        },
      ),
    );
  }
}
