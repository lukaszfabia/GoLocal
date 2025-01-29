import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:intl/intl.dart';

class EventDetailPage extends StatelessWidget {
  final Event event;

  const EventDetailPage({super.key, required this.event});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      extendBodyBehindAppBar: true,
      appBar: AppBar(
        backgroundColor: const Color.fromARGB(134, 255, 255, 255),
        elevation: 0,
        title: Text(
          event.title,
          style: const TextStyle(
            fontWeight: FontWeight.bold,
            overflow: TextOverflow.ellipsis,
            color: Color.fromARGB(255, 0, 0, 0),
          ),
        ),
        centerTitle: true,
        actions: [
          IconButton(
            icon: const Icon(Icons.info, color: Colors.white),
            onPressed: () {
              context.push('/events/${event.id}/info');
            },
          ),
          IconButton(
            icon: const Icon(Icons.report, color: Colors.red),
            onPressed: () {
              context.push('/events/${event.id}/report', extra: event);
            },
          ),
        ],
      ),
      body: SingleChildScrollView(
        child: Column(
          children: [
            _imageWithOverlay(),
            const SizedBox(height: 16),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 16.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  _title(),
                  const SizedBox(height: 8),
                  _buildTags(),
                  const SizedBox(height: 12),
                  _buildDetailCard(
                      "📅 Starts at", [_formatDate(event.startDate)]),
                  if (event.endDate != null)
                    _buildDetailCard(
                        "⏳ Ends at", [_formatDate(event.endDate!)]),
                  _buildDetailCard("📍 Location",
                      event.location != null ? [_formatLocation(event)] : null),
                  _buildDetailCard(
                      "👥 Organized by",
                      event.eventOrganizers
                          .map((o) => "${o.firstName} ${o.lastName}")
                          .toList()),
                  _buildDetailCard("📝 Description", [event.description]),
                  const SizedBox(height: 12),
                  _buildVotesCard(context),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _imageWithOverlay() {
    return Stack(
      children: [
        Container(
          height: 280,
          decoration: BoxDecoration(
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
              gradient: LinearGradient(
                begin: Alignment.topCenter,
                end: Alignment.bottomCenter,
                colors: [
                  Colors.black.withOpacity(0.5),
                  Colors.transparent,
                  Colors.black.withOpacity(0.8),
                ],
              ),
            ),
          ),
        ),
        Positioned(
          bottom: 16,
          left: 16,
          child: Row(
            children: [
              if (event.isAdultOnly) _buildBadge("18+", Colors.redAccent),
              if (event.isPromoted) _buildBadge("Promoted", Colors.orange),
            ],
          ),
        ),
      ],
    );
  }

  Widget _title() {
    return Text(
      event.title,
      style: const TextStyle(
        fontSize: 24.0,
        fontWeight: FontWeight.bold,
        color: Colors.black87,
      ),
    );
  }

  Widget _buildDetailCard(String label, List<String>? values) {
    if (values == null || values.isEmpty) {
      return const SizedBox.shrink();
    }
    return Card(
      elevation: 3.0,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12.0)),
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            Text(
              label,
              style:
                  const TextStyle(fontWeight: FontWeight.bold, fontSize: 16.0),
            ),
            const SizedBox(height: 6),
            Text(
              values.join(', '),
              style: const TextStyle(fontSize: 14.0, color: Colors.black87),
              maxLines: 2,
              overflow: TextOverflow.ellipsis,
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildTags() {
    return Wrap(
      spacing: 6,
      children: event.tags
          .map(
            (tag) => Chip(
              label: Text(tag.name),
              backgroundColor: Colors.blueGrey.withOpacity(0.2),
              labelStyle: const TextStyle(color: Colors.black87),
            ),
          )
          .toList(),
    );
  }

  Widget _buildVotesCard(BuildContext context) {
    return Card(
      elevation: 3.0,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12.0)),
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            const Text(
              "🗳 Related Votes",
              style: TextStyle(fontWeight: FontWeight.bold, fontSize: 16.0),
            ),
            const SizedBox(height: 8),
            ElevatedButton(
              style: ElevatedButton.styleFrom(
                backgroundColor: Colors.blueAccent,
                padding: const EdgeInsets.symmetric(vertical: 12),
              ),
              onPressed: () {
                context.push('/events/${event.id}/votes', extra: event);
              },
              child: const Text("View Votes", style: TextStyle(fontSize: 16)),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildBadge(String text, Color color) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
      margin: const EdgeInsets.only(right: 6),
      decoration: BoxDecoration(
        color: color,
        borderRadius: BorderRadius.circular(12),
      ),
      child: Text(
        text,
        style: const TextStyle(
          color: Colors.white,
          fontWeight: FontWeight.bold,
          fontSize: 12,
        ),
      ),
    );
  }

  String _formatDate(DateTime date) {
    return DateFormat("MMM d, yyyy • HH:mm").format(date);
  }

  String _formatLocation(Event event) {
    return "${event.location!.city}, ${event.location!.country}";
  }
}
