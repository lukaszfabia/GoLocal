import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/auth/bloc/auth_bloc.dart';
import 'package:golocal/src/event/bloc/events_bloc.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/domain/eventtype_enum.dart';
import 'package:golocal/src/event/ui/event_card.dart';
import 'package:golocal/src/event/ui/event_detail.dart';
import 'package:golocal/src/user/domain/user.dart';

class EventsViewPage extends StatelessWidget {
  EventsViewPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        actions: [
          IconButton(
            icon: const Icon(Icons.exit_to_app),
            onPressed: () {
              BlocProvider.of<AuthBloc>(context).add(AuthLogout());
            },
          ),
        ],
        title: const Text('Events'),
      ),
      body: BlocBuilder<EventsBloc, EventsState>(
        builder: (context, state) {
          if (state is EventsInitial) {
            BlocProvider.of<EventsBloc>(context).add(const FetchEvents());
          }
          if (state is EventsLoading) {
            return const Center(child: CircularProgressIndicator());
          }
          if (state is EventsLoaded) {
            return ListView.builder(
              itemCount: state.events.length,
              itemBuilder: (context, index) {
                return EventCard(
                  event: state.events[index],
                );
              },
            );
          }
          return const Center(child: Text('No events found'));
        },
      ),
    );
  }
}
