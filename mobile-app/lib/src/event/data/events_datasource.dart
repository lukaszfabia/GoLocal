import 'package:dio/dio.dart';

class EventsDataSource {
  final Dio _dio = Dio();

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
