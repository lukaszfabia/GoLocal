part of 'events_bloc.dart';

enum EventsStatus { initial, loading, loaded, error }

class EventsState extends Equatable {
  EventsState({
    required this.events,
    this.status = EventsStatus.initial,
    this.hasNext = true,
    this.nextPage = 1,
    this.errorMessage,
  });
  final List<Event> events;
  final bool hasNext;
  final int nextPage;
  final EventsStatus status;
  final String? errorMessage;

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
