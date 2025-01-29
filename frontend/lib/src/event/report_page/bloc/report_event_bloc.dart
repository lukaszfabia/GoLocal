/// Bloc for handling report events.
///
/// This Bloc manages the state and events related to reporting an event.
/// It interacts with an `IEventsRepository` to send the report data.
///
/// Events:
/// - `UpdateCategory`: Updates the category of the report.
/// - `UpdateDescription`: Updates the description of the report.
/// - `SendReport`: Sends the report to the repository.
///
/// States:
/// - `ReportEventState`: Holds the current state of the report, including category, description, status, and message.
///
/// Methods:
/// - `_onUpdateCategory`: Handles the `UpdateCategory` event and updates the category in the state.
/// - `_onUpdateDescription`: Handles the `UpdateDescription` event and updates the description in the state.
/// - `_onSendReport`: Handles the `SendReport` event, sends the report to the repository, and updates the status in the state.
///
/// Parameters:
/// - `_repository`: The repository used to send the report.
/// - `eventId`: The ID of the event being reported.
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
