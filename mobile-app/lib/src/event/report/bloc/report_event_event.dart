part of 'report_event_bloc.dart';

sealed class ReportEventEvent extends Equatable {
  const ReportEventEvent();

  @override
  List<Object> get props => [];
}

final class SendReport extends ReportEventEvent {
  const SendReport();
  @override
  List<Object> get props => [];
}

final class UpdateDescription extends ReportEventEvent {
  const UpdateDescription({required this.description});
  final String description;

  @override
  List<Object> get props => [description];
}

final class UpdateCategory extends ReportEventEvent {
  const UpdateCategory({required this.category});
  final String category;

  @override
  List<Object> get props => [category];
}
