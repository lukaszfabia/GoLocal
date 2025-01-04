import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/event/domain/event.dart';

class EventDetailPage extends StatelessWidget {
  final Event event;

  const EventDetailPage({super.key, required this.event});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.transparent,
        title: Text(
          event.title,
          style: TextStyle(
            fontWeight: FontWeight.bold,
            overflow: TextOverflow.ellipsis,
          ),
        ),
        centerTitle: true,
        actions: [
          IconButton(
            icon: Icon(Icons.info, color: Colors.blueAccent),
            onPressed: () {
              context.push('/events/${event.id}/info');
            },
          ),
          IconButton(
            icon: Icon(Icons.report, color: Colors.red),
            onPressed: () {
              context.push('/events/${event.id}/report', extra: event);
            },
          ),
        ],
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Card(
                child: Column(
                  children: [
                    _image(),
                    _title(),
                  ],
                ),
              ),
              _buildDetailCard("Description", [event.description]),
              _buildDetailCard("Starts at", [_formatDate(event.startDate)]),
              if (event.endDate != null)
                _buildDetailCard("Ends at", [_formatDate(event.endDate!)]),
              _buildDetailCard("Tags", [for (var tag in event.tags) tag.name]),
              _buildDetailCard("By", [
                for (var organizer in event.eventOrganizers)
                  "${organizer.firstName} ${organizer.lastName}"
              ]),
              _buildDetailCard(
                  "Location",
                  event.location != null
                      ? [
                          event.location!.city,
                          event.location!.address != null
                              ? "${event.location!.address!.street}"
                              : "",
                          event.location!.country,
                        ]
                      : null),
            ],
          ),
        ),
      ),
    );
  }

  Widget _image() {
    return ClipRRect(
      borderRadius: BorderRadius.circular(8.0),
      child: Image.network(
        event.imageUrl ??
            'https://via.placeholder.com/150', // Placeholder if no image
        fit: BoxFit.fitWidth,
        width: double.infinity,
      ),
    );
  }

  Widget _title() {
    return Text(
      event.title, // Assuming `title` is a property of `Event`
      style: TextStyle(
        fontSize: 24.0,
        fontWeight: FontWeight.bold,
        color: Colors.black87,
      ),
    );
  }

  Widget _buildDetailCard(String label, List<String>? values) {
    if (values == null) {
      return const SizedBox.shrink();
    }
    return Card(
      elevation: 2.0,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8.0)),
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            Text(
              label,
              style: TextStyle(fontWeight: FontWeight.bold, fontSize: 16.0),
            ),
            Text(
              values.join(', '),
              style: TextStyle(fontSize: 16.0, color: Colors.black87),
              overflow: TextOverflow.ellipsis,
              maxLines: 2,
            ),
          ],
        ),
      ),
    );
  }

  String _formatDate(DateTime date) {
    return '${date.hour}:${date.minute} ${date.day}/${date.month}/${date.year} ';
  }
}
