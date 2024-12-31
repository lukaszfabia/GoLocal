import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/event/bloc/events_bloc.dart';

import 'package:golocal/src/event/ui/event_card.dart';

class EventsViewPage extends StatelessWidget {
  const EventsViewPage({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocConsumer<EventsBloc, EventsState>(
      listener: (context, state) {
        if (state.status == EventsStatus.error) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(content: Text(state.errorMessage!)),
          );
        }
      },
      builder: (context, state) {
        double width = MediaQuery.of(context).size.width;
        int crossAxisCount = min(max(width ~/ 300, 1), 3);
        return GridView.builder(
          gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
              crossAxisCount: crossAxisCount, childAspectRatio: 5 / 3),
          itemCount: state.events.length + 1,
          itemBuilder: (context, i) {
            if (i < state.events.length) {
              final event = state.events[i];
              return EventCard(
                event: event,
              );
            } else {
              if (state.hasNext) {
                BlocProvider.of<EventsBloc>(context)
                    .add(FetchEvents(refresh: false));
              }
              if (state.status == EventsStatus.loading) {
                return const Center(child: CircularProgressIndicator());
              }
            }
            return null;
          },
        );
      },
    );
  }
}
