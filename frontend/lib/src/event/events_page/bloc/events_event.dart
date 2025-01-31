part of 'events_bloc.dart';

/// Base class for all events related to events.
sealed class EventsEvent extends Equatable {
  const EventsEvent();

  @override
  List<Object> get props => [];
}

/// Event to fetch events.
final class FetchEvents extends EventsEvent {
  /// Creates an instance of [FetchEvents].
  const FetchEvents({this.refresh = false});

  /// Indicates if the events should be refreshed.
  final bool refresh;

  @override
  List<Object> get props => [refresh];
}

/// Event to switch the repository.
final class SwitchRepository extends EventsEvent {
  /// Creates an instance of [SwitchRepository].
  const SwitchRepository(this.newRepository);

  /// The new repository to switch to.
  final IEventsRepository newRepository;

  @override
  List<Object> get props => [newRepository];
}
