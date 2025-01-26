import 'dart:async';

import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';

part 'report_event_event.dart';
part 'report_event_state.dart';

class ReportEventBloc extends Bloc<ReportEventEvent, ReportEventState> {
  final IEventsRepository _repository;
  final int eventId;
  ReportEventBloc(this._repository, {required this.eventId})
      : super(ReportEventState()) {
    on<UpdateCategory>(_onUpdateCategory);
    on<UpdateDescription>(_onUpdateDescription);
    on<SendReport>(_onSendReport);
  }

  FutureOr<void> _onUpdateCategory(
      UpdateCategory event, Emitter<ReportEventState> emit) {
    emit(state.copyWith(category: event.category));
  }

  FutureOr<void> _onUpdateDescription(
      UpdateDescription event, Emitter<ReportEventState> emit) {
    emit(state.copyWith(description: event.description));
  }

  FutureOr<void> _onSendReport(
      SendReport event, Emitter<ReportEventState> emit) async {
    emit(state.copyWith(status: ReportEventStatus.sending));
    try {
      await _repository.reportEvent(eventId, state.category, state.description);
      emit(state.copyWith(status: ReportEventStatus.success));
    } on Exception catch (e) {
      emit(state.copyWith(
        status: ReportEventStatus.error,
        message: e.toString(),
      ));
    }
  }

  @override
  void onTransition(Transition<ReportEventEvent, ReportEventState> transition) {
    super.onTransition(transition);
  }
}
