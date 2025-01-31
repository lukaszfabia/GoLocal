import 'package:dio/dio.dart';
import 'package:golocal/src/event/data/impl/events_datasource.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/dio_client.dart';
import 'package:golocal/src/event/promote_page/promote_pack.dart';

/// Implementation of the IEventsRepository interface.
/// This class handles the data operations related to events,
/// including fetching, creating, updating, deleting, reporting, and promoting events.
class EventsRepositoryImpl implements IEventsRepository {
  final DioClient _dioClient = DioClient();
  final EventsDataSource _eventsDataSource = EventsDataSource();

  /// Fetches a list of events from the server.
  ///
  /// Returns a [Future] that resolves to a list of [Event] objects.
  @override
  Future<List<Event>> getEvents() async {
    final response = await _dioClient.dio.get('/auth/event/10');

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
  /// Returns a [Future] that resolves to the created [Event] object, or throws an exception if the creation fails.
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
  /// Returns a [Future] that completes when the event is reported, or throws an exception if the report fails.
  @override
  Future<void> reportEvent(int id, String category, String description) async {
    var reason = "${category.toString()}: $description";
    var data = {'id': id, 'reason': reason};
    var result = await _eventsDataSource.reportEvent(data);
    if (result.statusCode != 201) {
      throw Exception('Failed to report event');
    }
  }

  /// Promotes an event.
  ///
  /// [id] The ID of the event to promote.
  /// [pack] The promotion package details.
  /// Returns a [Future] that resolves to a success message.
  @override
  Future<String> promoteEvent(int id, PromotePack pack) async {
    var data = {'id': id};
    await _eventsDataSource.promoteEvent(data);
    return "Success";
  }

  @override
  Future<bool> hasAccessToRecommendedEvents() async {
    final response =
        await _dioClient.dio.get('/auth/preference/was-survey-filled');

    final data = response.data['data'] as bool;

    return data;
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
