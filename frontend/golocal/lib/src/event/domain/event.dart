import 'package:golocal/src/event/domain/eventtype_enum.dart';
import 'package:golocal/src/event/location/location.dart';
import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/event/domain/tag.dart';
import 'package:golocal/src/user/domain/user.dart';

class Event extends Model {
  List<User> eventOrganizers;
  String title;
  String description;
  String imageUrl;
  bool isAdultOnly;
  EventType eventType;
  List<Tag> tags;

  DateTime startDate;
  DateTime? endDate;

  Location? location;

  Event({
    required super.id,
    required this.title,
    required this.description,
    required this.imageUrl,
    required this.tags,
    required this.startDate,
    required this.eventOrganizers,
    this.eventType = EventType.other,
    this.isAdultOnly = false,
    this.location,
    this.endDate,
  }) : assert(
          endDate == null || endDate.isAfter(startDate),
          eventOrganizers.isNotEmpty,
        );

  Event.fromJson(super.json)
      : title = json['title'],
        description = json['description'],
        imageUrl = json['imageUrl'],
        isAdultOnly = json['isAdultOnly'],
        eventType = json['eventType'],
        tags = json['tags'],
        startDate = DateTime.parse(json['startDate']),
        endDate =
            json['endDate'] != null ? DateTime.parse(json['endDate']) : null,
        location = Location.fromJson(json['location']),
        eventOrganizers = (json['eventOrganizers'] as List)
            .map((e) => User.fromJson(e))
            .toList(),
        super.fromJson();

  @override
  Map<String, dynamic> toJson() {
    final data = super.toJson();
    data['title'] = title;
    data['description'] = description;
    data['imageUrl'] = imageUrl;
    data['isAdultOnly'] = isAdultOnly;
    data['eventType'] = eventType;
    data['tags'] = tags.map((e) => e.toJson()).toList();
    data['startDate'] = startDate.toString();
    data['endDate'] = endDate?.toString();
    data['location'] = location?.toJson() ?? {};
    data['eventOrganizers'] = eventOrganizers.map((e) => e.toJson()).toList();
    return data;
  }
}
