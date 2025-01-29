import 'package:dio/dio.dart';
import 'package:golocal/src/dio_client.dart';

class EventsDataSource {
  Dio _dio = DioClient().dio;
  Future<List<Response>> getEvents(Map<String, dynamic> data) {
    return Future.value([Response(requestOptions: RequestOptions(path: ''))]);
  }

  Future<Response> getEvent(Map<String, dynamic> data) {
    return Future.value(Response(requestOptions: RequestOptions(path: '')));
  }

  Future<Response> createEvent(FormData formData) {
    return _dio.post('/auth/event/',
        data: formData, options: Options(contentType: 'multipart/form-data'));
  }

  Future<Response> updateEvent(Map<String, dynamic> data) {
    return Future.value(Response(requestOptions: RequestOptions(path: '')));
  }

  Future<void> deleteEvent(Map<String, dynamic> data) {
    return Future.delayed(Duration(seconds: 1));
  }

  Future<Response> reportEvent(Map<String, dynamic> data) {
    return _dio.post('/auth/event/report', data: data);
  }
}
