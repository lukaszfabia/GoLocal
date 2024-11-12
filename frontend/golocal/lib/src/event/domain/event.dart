import 'package:golocal/src/comment/domain/comment.dart';
import 'package:golocal/src/event/domain/eventtype_enum.dart';
import 'package:golocal/src/event/location/location.dart';
import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/event/domain/tag.dart';
import 'package:golocal/src/user/domain/user.dart';
import 'package:golocal/src/event/domain/vote.dart';

class Event extends Model {
  List<User> eventOrganizers;
  String description;
  String imageUrl;
  bool isAdultOnly;
  EventType eventType;
  List<Tag> tags;

  DateTime startDate;
  DateTime endDate;

  LocationDTO location;

  List<Comment> comments;
  List<Vote> votes;

  Event({
    required super.id,
    required this.eventOrganizers,
    required this.description,
    required this.imageUrl,
    required this.isAdultOnly,
    required this.tags,
    required this.startDate,
    required this.endDate,
    required this.location,
    required this.comments,
    required this.votes,
    required this.eventType,
  });
}
