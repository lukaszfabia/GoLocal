part of 'events_bloc.dart';

sealed class EventsEvent extends Equatable {
  const EventsEvent();

  @override
  List<Object> get props => [];
}

final class FetchEvents extends EventsEvent {
  const FetchEvents({this.refresh = false});
  final bool refresh;

  @override
  List<Object> get props => [refresh];
}

final class SwitchRepository extends EventsEvent {
  const SwitchRepository(this.newRepository);
  final IEventsRepository newRepository;

  @override
  List<Object> get props => [newRepository];
}
