import 'dart:io';

import 'package:dio/dio.dart';
import 'package:golocal/src/event/domain/eventtype_enum.dart';
import 'package:golocal/src/event/location/location.dart';
import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/event/domain/tag.dart';
import 'package:golocal/src/user/domain/user.dart';

class Event extends Model {
  List<User> eventOrganizers;
  String title;
  String description;
  String? imageUrl;
  bool isAdultOnly;
  EventType eventType;
  List<Tag> tags;

  DateTime startDate;
  DateTime? endDate;

  Location? location;

  bool get hasImage => imageUrl != null && imageUrl!.isNotEmpty;

  Event({
    required super.id,
    required this.title,
    required this.description,
    required this.tags,
    required this.startDate,
    required this.eventOrganizers,
    this.imageUrl,
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
        imageUrl = json['imageUrl'] ?? '',
        isAdultOnly = json['isAdultOnly'],
        eventType = EventType.fromString(json['event_type']),
        tags = json["event_tags"] != null
            ? (json['event_tags'] as List).map((e) => Tag.fromJson(e)).toList()
            : [],
        startDate = DateTime.parse(json['startDate']),
        endDate = json['finishDate'] != null
            ? DateTime.parse(json['finishDate'])
            : null,
        location = Location.fromJson(json['location']),
        eventOrganizers = json['eventOrganizers'] != null
            ? (json['eventOrganizers'] as List)
                .map((e) => User.fromJson(e))
                .toList()
            : [],
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

class EventDTO {
  List<int> organizers;
  String title;
  String description;
  File image;
  DateTime startDate;
  DateTime endDate;
  bool isAdultOnly;
  String eventType;
  List<String> tags;
  var location;

  EventDTO({
    required this.organizers,
    required this.title,
    required this.description,
    required this.image,
    required this.startDate,
    required this.endDate,
    required this.isAdultOnly,
    required this.eventType,
    required this.tags,
    required this.location,
  });

  Future<FormData> toFormData() async {
    final data = FormData.fromMap({
      'organizers': organizers,
      'title': title,
      'description': description,
      'image': await MultipartFile.fromFile(image.path),
      'startDate': startDate,
      'endDate': endDate,
      'isAdultOnly': isAdultOnly,
      'eventType': eventType,
      'tags': tags,
      'location': location,
    });
    return data;
  }
}
