import 'package:dio/dio.dart';
import 'package:jwt_decoder/jwt_decoder.dart';
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
        receiveTimeout: Duration(seconds: 25),
      ),
    );

    dio.interceptors.add(
      InterceptorsWrapper(
        onRequest: (options, handler) async {
          final accessToken = await _tokenStorage.getAccessToken();
          if (accessToken != null) {
            options.headers['Authorization'] = 'Bearer $accessToken';
            Map<String, dynamic> decodedToken = JwtDecoder.decode(accessToken);
            options.headers['User-Id'] = decodedToken['sub'];
          }
          print("Request: [${options.method}] ${options.uri}");
          print("Headers: ${options.headers}");
          print("Body: ${options.data}");
          handler.next(options);
        },
        onResponse: (response, handler) {
          print(
              "Response: [${response.statusCode}] ${response.requestOptions.uri}");
          print("Headers: ${response.headers}");
          print("Body: ${response.data}");
          handler.next(response);
        },
        onError: (error, handler) async {
          if (error.response?.statusCode == 401 &&
              error.response?.statusMessage == 'Unauthorized') {
            final refreshToken = await _tokenStorage.getRefreshToken();
            if (refreshToken != null) {
              try {
                final response = await dio.post('/refresh-token/', data: {
                  'refresh': refreshToken,
                });
                print('Refresh token response: $response');
                final newAccessToken = response.data['data']['access'];
                await _tokenStorage.saveAccessToken(newAccessToken);
                final newRefreshToken = response.data['data']['refresh'];
                await _tokenStorage.saveRefreshToken(newRefreshToken);

                final options = error.response!.requestOptions;
                options.headers['Authorization'] = 'Bearer $newAccessToken';
                Map<String, dynamic> decodedToken =
                    JwtDecoder.decode(newAccessToken);
                options.headers['User-Id'] = decodedToken['sub'];
                if (options.data is FormData) {
                  options.data = FormData.fromMap(options.data.fields);
                }

                final retryResponse = await dio.fetch(options);
                return handler.resolve(retryResponse);
              } catch (e) {
                // await _tokenStorage.clearTokens();
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

  Future<int?> getLoggedUserId() async {
    final accessToken = await _tokenStorage.getAccessToken();
    if (accessToken != null) {
      Map<String, dynamic> decodedToken = JwtDecoder.decode(accessToken);
      return decodedToken['sub'];
    }
    return null;
  }
}
