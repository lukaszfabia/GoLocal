part of 'report_event_bloc.dart';

/// Enum representing the status of a report event.
enum ReportEventStatus {
  /// Initial status of the report event.
  initial,

  /// Status when the report event is being sent.
  sending,

  /// Status when the report event has been successfully sent.
  success,

  /// Status when there was an error sending the report event.
  error
}

/// State class for the report event, containing the description, category, status, and an optional message.
class ReportEventState extends Equatable {
  /// Creates a new instance of [ReportEventState].
  ///
  /// The [description] and [category] default to empty strings,
  /// [status] defaults to [ReportEventStatus.initial], and [message] is optional.
  const ReportEventState({
    this.description = '',
    this.category = '',
    this.status = ReportEventStatus.initial,
    this.message,
  });

  /// The current status of the report event.
  final ReportEventStatus status;

  /// The description of the report event.
  final String description;

  /// The category of the report event.
  final String category;

  /// An optional message associated with the report event.
  final String? message;

  /// Returns a list of properties to compare for equality.
  @override
  List<Object?> get props => [
        description,
        category,
        status,
        message,
      ];

  /// Creates a copy of the current [ReportEventState] with the given fields replaced with new values.
  ///
  /// If a field is not provided, the current value of that field is used.
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
