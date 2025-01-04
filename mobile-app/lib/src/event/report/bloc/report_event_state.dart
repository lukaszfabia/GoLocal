part of 'report_event_bloc.dart';

enum ReportEventStatus { initial, sending, success, error }

class ReportEventState extends Equatable {
  const ReportEventState({
    this.description = '',
    this.category = '',
    this.status = ReportEventStatus.initial,
    this.message,
  });
  final ReportEventStatus status;
  final String description;
  final String category;
  final String? message;

  @override
  List<Object?> get props => [
        description,
        category,
        status,
        message,
      ];

  ReportEventState copyWith({
    ReportEventStatus? status,
    String? description,
    String? category,
    String? message,
  }) {
    return ReportEventState(
      status: status ?? this.status,
      description: description ?? this.description,
      category: category ?? this.category,
      message: message ?? this.message,
    );
  }
}
