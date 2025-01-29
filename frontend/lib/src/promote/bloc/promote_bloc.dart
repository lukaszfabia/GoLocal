import 'package:bloc/bloc.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:meta/meta.dart';

part 'promote_event.dart';
part 'promote_state.dart';

class PromoteBloc extends Bloc<PromoteEvent, PromoteState> {
  final Event event;
  final IEventsRepository repository;
  PromoteBloc({required this.event, required this.repository})
      : super(PromoteInitial()) {
    on<PromoteEvent>((event, emit) {
      // TODO: implement event handler
    });
  }
}
