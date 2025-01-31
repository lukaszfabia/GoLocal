import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/dio_client.dart';
import 'package:golocal/src/event/promote_page/promote_pack.dart';

/// Implementation of the IEventsRepository interface for recommended events.
/// This class handles the data operations related to recommended events,
/// including fetching, creating, updating, deleting, and reporting events.
class RecommendedRepositoryImpl implements IEventsRepository {
  final DioClient _dioClient = DioClient();

  /// Fetches a list of recommended events from the server.
  ///
  /// Returns a [Future] that resolves to a list of [Event] objects.
  @override
  Future<List<Event>> getEvents() async {
    final response = await _dioClient.dio.get('/auth/recommendations');

    final data = response.data['data'] as List<dynamic>;

    return data.map((json) => Event.fromJson(json)).toList();
  }

  /// Fetches a single event by its ID.
  ///
  /// [id] The ID of the event to fetch.
  /// Returns a [Future] that resolves to an [Event] object.
  @override
  Future<Event> getEvent(String id) async {
    final response = await _dioClient.dio.get('/auth/event/id=$id');

    final data = response.data['data'] as dynamic;

    return Event.fromJson(data);
  }

  /// Creates a new event.
  ///
  /// [event] The event data transfer object containing the event details.
  /// Returns a [Future] that resolves to the created [Event] object.
  @override
  Future<Event> createEvent(EventDTO event) async {
    // Return null for now
    return Event(
        id: 1,
        title: "1",
        description: "",
        tags: [],
        startDate: DateTime(2024),
        eventOrganizers: []);
  }

  /// Updates an existing event.
  ///
  /// [event] The event object containing the updated details.
  /// Returns a [Future] that resolves to the updated [Event] object.
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

  /// Deletes an event by its ID.
  ///
  /// [id] The ID of the event to delete.
  /// Returns a [Future] that completes when the event is deleted.
  @override
  Future<void> deleteEvent(String id) async {
    // Do nothing for now
  }

  /// Reports an event.
  ///
  /// [id] The ID of the event to report.
  /// [category] The category of the report.
  /// [description] The description of the report.
  /// Returns a [Future] that completes when the event is reported.
  @override
  Future<void> reportEvent(int id, String category, String description) async {
    // Do nothing for now
  }

  /// Promotes an event.
  ///
  /// [id] The ID of the event to promote.
  /// [pack] The promotion package details.
  /// Throws an [UnimplementedError] as this method is not implemented.
  @override
  Future<String> promoteEvent(int id, PromotePack pack) {
    throw UnimplementedError();
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
