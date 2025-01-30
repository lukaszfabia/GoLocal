import 'package:flutter/material.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/shared/badge_widget.dart';
import 'package:golocal/src/shared/extensions.dart';

class EventCard extends StatelessWidget {
  final Event event;
  const EventCard({required this.event, super.key});

  @override
  Widget build(BuildContext context) {
    return Card(
      clipBehavior: Clip.antiAlias,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(16.0),
      ),
      elevation: 6,
      child: Stack(
        children: [
          Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(16),
              image: DecorationImage(
                image: event.hasImage
                    ? NetworkImage(event.imageUrl!) as ImageProvider
                    : const AssetImage("assets/images/image_not_found.png"),
                fit: BoxFit.cover,
              ),
            ),
          ),
          Positioned.fill(
            child: Container(
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(16),
                gradient: LinearGradient(
                  begin: Alignment.topCenter,
                  end: Alignment.bottomCenter,
                  colors: [
                    Colors.black.withValues(alpha: .3),
                    Colors.black.withValues(alpha: 0.8),
                  ],
                ),
              ),
            ),
          ),
          Positioned(
            top: 12, // Position it at the top
            left: 12, // Align it to the left
            child: Row(
              children: [
                if (event.isPromoted)
                  BadgeWidget(
                      text: "🔥Promoted",
                      backgroundColor: Colors.orange,
                      fontSize: 14),
                if (event.isAdultOnly)
                  BadgeWidget(
                      text: "🔞18+", backgroundColor: Colors.red, fontSize: 14),
              ],
            ),
          ),
          Positioned(
            bottom: 12,
            left: 12,
            right: 12,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  event.eventType.name.toUpperCase(),
                  style: TextStyle(
                    color: Colors.orangeAccent,
                    fontSize: 12,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                Text(
                  event.title,
                  style: const TextStyle(
                    color: Colors.white,
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                  ),
                  maxLines: 1,
                  overflow: TextOverflow.ellipsis,
                ),
                const SizedBox(height: 4),
                Row(
                  children: [
                    Icon(Icons.event, size: 16, color: Colors.white70),
                    const SizedBox(width: 6),
                    Text(
                      "${(event.startDate.formatReadableDate())}${event.endDate != null ? " - ${event.endDate!.formatReadableDate()}" : ""}",
                      style:
                          const TextStyle(color: Colors.white70, fontSize: 12),
                    ),
                  ],
                ),
                const SizedBox(height: 4),
                if (event.location != null)
                  Row(
                    children: [
                      Icon(Icons.location_on, size: 16, color: Colors.white70),
                      const SizedBox(width: 6),
                      Expanded(
                        child: Text(
                          "${event.location!.city}, ${event.location!.address?.street}",
                          style: const TextStyle(
                              color: Colors.white70, fontSize: 12),
                          maxLines: 1,
                          overflow: TextOverflow.ellipsis,
                        ),
                      ),
                    ],
                  ),
                const SizedBox(height: 6),
              ],
            ),
          )
        ],
      ),
    );
  }
}
