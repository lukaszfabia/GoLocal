import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';

part 'events_event.dart';
part 'events_state.dart';

class EventsBloc extends Bloc<EventsEvent, EventsState> {
  final IEventsRepository _repository;
  final List<Event> events = [];
  EventsBloc(this._repository) : super(EventsInitial()) {
    on<FetchEvents>((event, emit) async {
      emit(EventsLoading());
      await _repository.getEvents().then((events) {
        emit(EventsLoaded(events));
      }).catchError((e) {
        emit(EventsError(message: e.toString()));
      });
    });
  }
}
