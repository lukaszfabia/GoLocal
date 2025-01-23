part of 'manage_event_bloc.dart';

enum ManageEventStatus { initial, success, error, loading }

class ManageEventState extends Equatable {
  const ManageEventState(
      {this.status = ManageEventStatus.initial,
      this.title = '',
      this.description = '',
      this.type,
      required this.tags,
      this.startDate,
      this.endDate,
      this.isAdultsOnly = false,
      required this.organizers,
      this.location,
      this.image,
      this.message});
  final ManageEventStatus status;
  final String title;
  final String description;
  final EventType? type;
  final List<Tag> tags;
  final DateTime? startDate;
  final DateTime? endDate;
  final bool isAdultsOnly;
  final List<User> organizers;
  final Location? location;
  final XFile? image;

  final String? message;

  @override
  List<Object?> get props => [
        status,
        title,
        description,
        type,
        tags,
        startDate,
        endDate,
        isAdultsOnly,
        organizers,
        location,
        message,
        image,
      ];

  ManageEventState copyWith({
    ManageEventStatus? status,
    String? title,
    String? description,
    EventType? type,
    List<Tag>? tags,
    DateTime? startDate,
    DateTime? endDate,
    bool? isAdultsOnly,
    List<User>? organizers,
    Location? location,
    String? message,
    XFile? image,
  }) {
    return ManageEventState(
      status: status ?? this.status,
      title: title ?? this.title,
      description: description ?? this.description,
      type: type ?? this.type,
      tags: tags ?? this.tags,
      startDate: startDate ?? this.startDate,
      endDate: endDate ?? this.endDate,
      isAdultsOnly: isAdultsOnly ?? this.isAdultsOnly,
      organizers: organizers ?? this.organizers,
      location: location ?? this.location,
      message: message ?? this.message,
      image: image ?? this.image,
    );
  }

  ManageEventState.copyFromEvent(Event event)
      : this(
          status: ManageEventStatus.initial,
          title: event.title,
          description: event.description,
          type: event.eventType,
          tags: event.tags,
          startDate: event.startDate,
          endDate: event.endDate,
          isAdultsOnly: event.isAdultOnly,
          organizers: event.eventOrganizers,
          location: event.location,
        );
}
