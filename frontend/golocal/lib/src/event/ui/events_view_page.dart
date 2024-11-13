import 'package:flutter/material.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/domain/eventtype_enum.dart';
import 'package:golocal/src/event/ui/event_card.dart';
import 'package:golocal/src/user/domain/user.dart';

class EventsViewPage extends StatelessWidget {
  EventsViewPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Events'),
      ),
      // ignore: prefer_const_constructors
      body: ListView.builder(
        itemBuilder: (context, index) => EventCard(
          event: Event(
            id: 1,
            title: "Najazd na 16 tarnogaj",
            eventOrganizers: [
              User(
                  id: 0,
                  firstName: "Anna",
                  lastName: "Robak",
                  email: "AnnaRobak123",
                  birthDate: DateTime(2000, 1, 11),
                  isVerified: true,
                  isPremium: false)
            ],
            description:
                "A very nice eventA very nice eventA very nice eventA very nice eventA very nice eventA very nice eventA very nice eventA very nice eventA very nice eventA very nice eventA very A very nice eventA very nice eventA very nice eventA very nice eventnice eventA very nice eventA very nice event",
            imageUrl:
                "https://www.contentviewspro.com/wp-content/uploads/2017/07/default_image.png",
            tags: [],
            startDate: DateTime(2024, 11, 15),
            eventType: EventType.party,
          ),
        ),
      ),
    );
  }
}
