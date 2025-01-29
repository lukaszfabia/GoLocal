part of 'report_event_bloc.dart';

/// Base class for all report event events, extending Equatable for value comparison.
sealed class ReportEventEvent extends Equatable {
  const ReportEventEvent();

  @override
  List<Object> get props => [];
}

/// Event to send a report.
final class SendReport extends ReportEventEvent {
  const SendReport();

  @override
  List<Object> get props => [];
}

/// Event to update the description of a report.
///
/// Takes a required [description] parameter.
final class UpdateDescription extends ReportEventEvent {
  const UpdateDescription({required this.description});
  final String description;

  @override
  List<Object> get props => [description];
}

/// Event to update the category of a report.
///
/// Takes a required [category] parameter.
final class UpdateCategory extends ReportEventEvent {
  const UpdateCategory({required this.category});
  final String category;

  @override
  List<Object> get props => [category];
}
