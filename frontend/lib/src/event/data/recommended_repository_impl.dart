import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/dio_client.dart';

class RecommendedRepositoryImpl implements IEventsRepository {
  final DioClient _dioClient = DioClient();

  @override
  Future<List<Event>> getEvents() async {
    final response = await _dioClient.dio.get('/auth/recommendations');

    final data = response.data['data'] as List<dynamic>;

    return data.map((json) => Event.fromJson(json)).toList();
  }

  @override
  Future<Event> getEvent(String id) async {
    final response = await _dioClient.dio.get('/auth/event/id=$id');

    final data = response.data['data'] as dynamic;

    return Event.fromJson(data);
  }

  @override
  Future<Event> createEvent(Event event) async {
    // Return null for now
    return Event(
        id: 1,
        title: "1",
        description: "",
        tags: [],
        startDate: DateTime(2024),
        eventOrganizers: []);
  }

  @override
  Future<Event> updateEvent(Event event) async {
    // Return null for now
    return Event(
        id: 1,
        title: "1",
        description: "",
        tags: [],
        startDate: DateTime(2024),
        eventOrganizers: []);
  }

  @override
  Future<void> deleteEvent(String id) async {
    // Do nothing for now
  }

  @override
  Future<void> reportEvent(int id, String category, String description) async {
    // Do nothing for now
  }

  // @override
  // Future<Event> getEvent(String id) async {
  //   final response = await _dioClient.dio.get('/auth/event?id=$id');
  //   return Event.fromJson(response.data);
  // }

  // @override
  // Future<Event> createEvent(Event event) async {
  //   final response =
  //       await _dioClient.dio.post('/auth/event', data: event.toJson());
  //   return Event.fromJson(response.data);
  // }

  // @override
  // Future<Event> updateEvent(Event event) async {
  //   var id = event.id;
  //   final response =
  //       await _dioClient.dio.put('/events/$id', data: event.toJson());
  //   return Event.fromJson(response.data);
  // }

  // @override
  // Future<void> deleteEvent(String id) async {
  //   await _dioClient.dio.delete('/events/$id');
  // }

  // @override
  // Future<void> reportEvent(int id, String category, String description) async {
  //   await _dioClient.dio.post('/events/$id/report', data: {
  //     'category': category,
  //     'description': description,
  //   });
  // }
}
