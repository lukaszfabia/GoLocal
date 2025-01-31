part of 'events_bloc.dart';

/// Enum representing the status of events.
enum EventsStatus { initial, loading, loaded, error }

/// State class for managing events.
class EventsState extends Equatable {
  /// Creates an instance of [EventsState].
  const EventsState({
    required this.events,
    this.status = EventsStatus.initial,
    this.hasNext = true,
    this.nextPage = 1,
    this.errorMessage,
  });

  /// List of events.
  final List<Event> events;

  /// Indicates if there are more events to load.
  final bool hasNext;

  /// The next page to load.
  final int nextPage;

  /// The current status of events.
  final EventsStatus status;

  /// Error message in case of failure.
  final String? errorMessage;

  /// Creates a copy of the current state with updated values.
  EventsState copyWith({
    List<Event>? events,
    bool? hasNext,
    int? nextPage,
    EventsStatus? status,
    String? errorMessage,
  }) {
    return EventsState(
      events: events ?? this.events,
      hasNext: hasNext ?? this.hasNext,
      nextPage: nextPage ?? this.nextPage,
      status: status ?? this.status,
      errorMessage: errorMessage ?? this.errorMessage,
    );
  }

  @override
  List<Object?> get props => [
        events,
        hasNext,
        nextPage,
        status,
        errorMessage,
      ];
}
