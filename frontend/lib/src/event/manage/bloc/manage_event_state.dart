part of 'manage_event_bloc.dart';

enum ManageEventStatus { initial, success, error, loading }

class ManageEventState extends Equatable {
  const ManageEventState({
    this.status = ManageEventStatus.initial,
    this.title = '',
    this.description = '',
    this.type,
    required this.tags,
    this.startDate,
    this.endDate,
    this.isAdultsOnly = false,
    required this.organizers,
    this.image,
    this.message,
    this.id,
    this.lat,
    this.lon,
  });
  final ManageEventStatus status;
  final int? id;
  final String title;
  final String description;
  final EventType? type;
  final List<String> tags;
  final DateTime? startDate;
  final DateTime? endDate;
  final bool isAdultsOnly;
  final List<String> organizers;
  final File? image;
  final double? lat;
  final double? lon;

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
        lat,
        lon,
        message,
        image,
      ];

  ManageEventState copyWith({
    ManageEventStatus? status,
    String? title,
    String? description,
    EventType? type,
    List<String>? tags,
    DateTime? startDate,
    DateTime? endDate,
    bool? isAdultsOnly,
    List<String>? organizers,
    Location? location,
    String? message,
    File? image,
    double? lat,
    double? lon,
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
      lat: lat ?? this.lat,
      lon: lon ?? this.lon,
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
          tags: event.tags.map((e) => e.name).toList(),
          startDate: event.startDate,
          endDate: event.endDate,
          isAdultsOnly: event.isAdultOnly,
          organizers: event.eventOrganizers.map((e) => e.email).toList(),
          lat: event.location?.coords?.latitude,
          lon: event.location?.coords?.longitude,
          image: File(event.imageUrl ?? ''),
        );
}
