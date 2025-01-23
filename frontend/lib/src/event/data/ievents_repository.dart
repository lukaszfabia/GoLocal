import 'package:golocal/src/event/domain/event.dart';

abstract class IEventsRepository {
  Future<List<Event>> getEvents();
  Future<Event> getEvent(String id);
  Future<Event> createEvent(Event event);
  Future<Event> updateEvent(Event event);
  Future<void> deleteEvent(String id);
  Future<void> reportEvent(int id, String category, String description);
}
