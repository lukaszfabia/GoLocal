part of 'events_bloc.dart';

sealed class EventsState extends Equatable {
  const EventsState();

  @override
  List<Object> get props => [];
}

final class EventsInitial extends EventsState {}

final class EventsLoading extends EventsState {}

final class EventsLoaded extends EventsState {
  final List<Event> events;

  const EventsLoaded(this.events);

  @override
  List<Object> get props => [events];
}

final class EventsError extends EventsState {
  final String message;
  final List<Event>? events;

  const EventsError({required this.message, this.events});

  @override
  List<Object> get props => [message];
}
