import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/promote_page/promote_pack.dart';
import 'package:meta/meta.dart';

part 'promote_event.dart';
part 'promote_state.dart';

/// Bloc responsible for handling promotion-related events and states.
///
/// This Bloc manages the promotion process of an event by interacting with
/// the [IEventsRepository] to promote the event and updating the state
/// accordingly.
///
/// The [PromoteBloc] requires an [Event] and an [IEventsRepository] to be
/// provided upon initialization.
///
/// Events handled:
/// - [PromotePackChosenEvent]: Updates the state with the chosen promotion pack.
/// - [PromoteRequestedEvent]: Initiates the promotion process and updates the
///   state based on the success or failure of the promotion.
///
/// States emitted:
/// - [PromoteState]: The current state of the promotion process, including
///   the selected pack, status, and any error messages.
///
/// Example usage:
/// ```dart
/// final promoteBloc = PromoteBloc(event: myEvent, repository: myRepository);
/// promoteBloc.add(PromotePackChosenEvent(pack: myPack));
/// promoteBloc.add(PromoteRequestedEvent());
/// ```
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
