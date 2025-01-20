import 'package:faker/faker.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/domain/eventtype_enum.dart';
import 'package:golocal/src/event/domain/tag.dart';
import 'package:golocal/src/event/location/location.dart';
import 'package:golocal/src/user/domain/user.dart';
import 'package:golocal/src/dio_client.dart';
import 'package:golocal/src/jwt_token_storage.dart';
import 'package:jwt_decoder/jwt_decoder.dart';
import 'package:dio/dio.dart';

class EventsRepositoryImpl implements IEventsRepository {
  final DioClient _dioClient = DioClient();

  @override
  Future<List<Event>> getEvents() async {
    final response = await _dioClient.dio.get('api/events');
    return (response.data as List).map((json) => Event.fromJson(json)).toList();
  }

  @override
  Future<Event> getEvent(String id) async {
    final response = await _dioClient.dio.get('/events/$id');
    return Event.fromJson(response.data);
  }

  @override
  Future<Event> createEvent(Event event) async {
    final response = await _dioClient.dio.post('/events', data: event.toJson());
    return Event.fromJson(response.data);
  }

  @override
  Future<Event> updateEvent(Event event) async {
    var id = event.id;
    final response =
        await _dioClient.dio.put('/events/$id', data: event.toJson());
    return Event.fromJson(response.data);
  }

  @override
  Future<void> deleteEvent(String id) async {
    await _dioClient.dio.delete('/events/$id');
  }

  @override
  Future<void> reportEvent(int id, String category, String description) async {
    await _dioClient.dio.post('/events/$id/report', data: {
      'category': category,
      'description': description,
    });
  }
}
