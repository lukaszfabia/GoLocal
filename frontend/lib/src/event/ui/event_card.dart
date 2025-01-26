import 'package:flutter/material.dart';
import 'package:golocal/src/event/domain/event.dart';

class EventCard extends StatelessWidget {
  final Event event;
  const EventCard({required this.event, super.key});

  @override
  Widget build(BuildContext context) {
    return Card(
      clipBehavior: Clip.antiAlias,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(14.0),
      ),
      child: Container(
        decoration: BoxDecoration(
          image: DecorationImage(
            image: event.hasImage
                ? NetworkImage(event.imageUrl!)
                : AssetImage("assets/images/image_not_found.png"),
            fit: BoxFit.cover,
          ),
        ),
        child: Stack(
          children: [
            Positioned(
              bottom: 0,
              left: 0,
              right: 0,
              child: Container(
                  decoration: BoxDecoration(
                    color: Colors.grey.withValues(alpha: 0.8),
                  ),
                  child: Padding(
                    padding: const EdgeInsets.all(6.0),
                    child: Column(
                      children: [
                        Text(event.title),
                        Text(event.description,
                            maxLines: 2, overflow: TextOverflow.ellipsis),
                      ],
                    ),
                  )),
            ),
          ],
        ),
      ),
    );
  }
}
