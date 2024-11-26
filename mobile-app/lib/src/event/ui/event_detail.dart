import 'package:flutter/widgets.dart';
import 'package:golocal/src/event/domain/event.dart';

class EventDetail extends StatelessWidget {
  final Event event;
  const EventDetail({super.key, required this.event});

  @override
  Widget build(BuildContext context) {
    return SizedBox.expand(
      child: Column(
        children: [
          Text(event.title),
          Text(event.description),
          Text(event.startDate.toString()),
          Text(event.eventType.toString()),
        ],
      ),
    );
  }
}
