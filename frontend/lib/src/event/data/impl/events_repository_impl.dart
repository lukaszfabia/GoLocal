import 'package:dio/dio.dart';
import 'package:golocal/src/event/data/impl/events_datasource.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/dio_client.dart';
import 'package:golocal/src/event/promote_page/promote_pack.dart';

class EventsRepositoryImpl implements IEventsRepository {
  final DioClient _dioClient = DioClient();
  final EventsDataSource _eventsDataSource = EventsDataSource();

  @override
  Future<List<Event>> getEvents() async {
    final response = await _dioClient.dio.get('/auth/event/10');

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
  Future<Event?> createEvent(EventDTO event) async {
    final FormData formData = await event.toFormData();

    final response = await _eventsDataSource.createEvent(formData);
    if (response.statusCode == 201) {
      return Event.fromJson(response.data['data']);
    } else {
      throw Exception('Failed to create event');
    }
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
    var reason = "${category.toString()}: $description";
    var data = {'id': id, 'reason': reason};
    var result = await _eventsDataSource.reportEvent(data);
    if (result.statusCode != 201) {
      throw Exception('Failed to report event');
    }
  }

  @override
  Future<String> promoteEvent(int id, PromotePack pack) async {
    var data = {'id': id};
    await _eventsDataSource.promoteEvent(data);
    return "Success";
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
