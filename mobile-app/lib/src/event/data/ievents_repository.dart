import 'package:golocal/src/event/domain/event.dart';

abstract class IEventsRepository {
  Future<List<Event>> getEvents();
  Future<Event> getEvent(String id);
  Future<void> createEvent(Event event);
  Future<void> updateEvent(Event event);
  Future<void> deleteEvent(String id);
}
