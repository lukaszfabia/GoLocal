import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';

part 'events_event.dart';
part 'events_state.dart';

class EventsBloc extends Bloc<EventsEvent, EventsState> {
  final IEventsRepository _repository;
  EventsBloc(this._repository) : super(EventsState(events: [])) {
    on<FetchEvents>((event, emit) async {
      if (event.refresh) {
        emit(state
            .copyWith(events: [], status: EventsStatus.loading, nextPage: 1));
      } else {
        emit(state.copyWith(status: EventsStatus.loading));
      }
      try {
        final events = await _repository.getEvents();
        emit(state.copyWith(
          events: state.events + events,
          status: EventsStatus.loaded,
          nextPage: state.nextPage + 1,
          hasNext: events.isNotEmpty,
        ));
      } catch (e) {
        emit(state.copyWith(
          status: EventsStatus.error,
          errorMessage: e.toString(),
        ));
      }
    });
  }
}
