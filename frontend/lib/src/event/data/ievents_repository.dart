import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/promote_page/promote_pack.dart';

abstract class IEventsRepository {
  Future<List<Event>> getEvents();
  Future<Event> getEvent(String id);
  Future<Event?> createEvent(EventDTO event);
  Future<Event> updateEvent(Event event);
  Future<void> deleteEvent(String id);
  Future<void> reportEvent(int id, String category, String description);
  Future<String> promoteEvent(int id, PromotePack pack);
  Future<bool> hasAccessToRecommendedEvents();
}
