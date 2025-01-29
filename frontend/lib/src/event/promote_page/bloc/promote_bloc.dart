import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/promote_page/promote_pack.dart';
import 'package:meta/meta.dart';

part 'promote_event.dart';
part 'promote_state.dart';

class PromoteBloc extends Bloc<PromoteEvent, PromoteState> {
  final Event event;
  final IEventsRepository repository;
  PromoteBloc({required this.event, required this.repository})
      : super(PromoteState()) {
    on<PromotePackChosenEvent>((event, emit) {
      emit(state.copyWith(pack: event.pack));
    });
    on<PromoteRequestedEvent>((event, emit) async {
      emit(state.copyWith(status: PromoteStatus.loading));
      if (state.pack == null) {
        emit(state.copyWith(
            status: PromoteStatus.error, message: "No pack selected"));
        return;
      }
      try {
        await repository.promoteEvent(this.event.id, state.pack!);
        emit(state.copyWith(status: PromoteStatus.success));
      } catch (e) {
        emit(
            state.copyWith(status: PromoteStatus.error, message: e.toString()));
      }
    });
  }
}
