import 'package:faker/faker.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/domain/eventtype_enum.dart';
import 'package:golocal/src/event/domain/tag.dart';
import 'package:golocal/src/event/location/location.dart';
import 'package:golocal/src/user/domain/user.dart';

class EventsRepositoryDummy implements IEventsRepository {
  var faker = Faker();

  List<Event> dummyEvents() {
    var events = List.generate(
      5,
      (i) {
        return Event(
          id: i,
          title: faker.lorem.words(3).join(' '),
          description: faker.lorem.sentence(),
          imageUrl: faker.image
              .loremPicsum(random: faker.randomGenerator.integer(1000)),
          tags: [Tag(id: i, name: faker.lorem.word())],
          startDate: DateTime.now(),
          eventOrganizers: [
            User(
              id: i,
              firstName: faker.person.firstName(),
              lastName: faker.person.lastName(),
              email: faker.internet.safeEmail(),
              birthDate: faker.date.dateTime(),
              isVerified: faker.randomGenerator.boolean(),
              isPremium: faker.randomGenerator.boolean(),
            )
          ],
          eventType: faker.randomGenerator.element(EventType.values),
          location: Location(
            id: i,
            city: faker.address.city(),
            country: faker.address.country(),
          ),
        );
      },
    );

    return events;
  }

  @override
  Future<Event> createEvent(Event event) {
    // TODO: implement createEvent
    throw UnimplementedError();
  }

  @override
  Future<void> deleteEvent(String id) {
    // TODO: implement deleteEvent
    throw UnimplementedError();
  }

  @override
  Future<Event> getEvent(String id) {
    // TODO: implement getEvent
    throw UnimplementedError();
  }

  @override
  Future<List<Event>> getEvents() {
    var events = List.generate(
      5,
      (i) {
        return Event(
          id: i,
          title: faker.lorem.words(3).join(' '),
          description: faker.lorem.sentence(),
          imageUrl: faker.image
              .loremPicsum(random: faker.randomGenerator.integer(1000)),
          tags: [Tag(id: i, name: faker.lorem.word())],
          startDate: DateTime.now(),
          eventOrganizers: [
            User(
              id: i,
              firstName: faker.person.firstName(),
              lastName: faker.person.lastName(),
              email: faker.internet.safeEmail(),
              birthDate: faker.date.dateTime(),
              isVerified: faker.randomGenerator.boolean(),
              isPremium: faker.randomGenerator.boolean(),
            )
          ],
          eventType: faker.randomGenerator.element(EventType.values),
          location: Location(
            id: i,
            city: faker.address.city(),
            country: faker.address.country(),
          ),
        );
      },
    );
    return Future.delayed(Duration(seconds: 1), () => events);
  }

  @override
  Future<Event> updateEvent(Event event) {
    // TODO: implement updateEvent
    throw UnimplementedError();
  }

  @override
  Future<void> reportEvent(int id, String category, String description) async {
    await Future.delayed(Duration(seconds: 1));
    return;
  }
}
