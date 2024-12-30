import 'package:dio/dio.dart';
import 'package:golocal/src/dio_client.dart';

class EventsDataSource {
  final DioClient _dio = DioClient();

  Future<List<Response>> getEvents(Map<String, dynamic> data) {
    return Future.value([Response(requestOptions: RequestOptions(path: ''))]);
  }

  Future<Response> getEvent(Map<String, dynamic> data) {
    return Future.value(Response(requestOptions: RequestOptions(path: '')));
  }

  Future<Response> createEvent(Map<String, dynamic> data) {
    return Future.value(Response(requestOptions: RequestOptions(path: '')));
  }

  Future<Response> updateEvent(Map<String, dynamic> data) {
    return Future.value(Response(requestOptions: RequestOptions(path: '')));
  }

  Future<void> deleteEvent(Map<String, dynamic> data) {
    return Future.delayed(Duration(seconds: 1));
  }
}
