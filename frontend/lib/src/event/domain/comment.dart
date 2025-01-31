import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/event/domain/participationstatus_enum.dart';

/// Represents a comment made by a user on an event.
///
/// Attributes:
/// - `userId`: The ID of the user who made the comment.
/// - `eventId`: The ID of the event the comment is associated with.
/// - `content`: The content of the comment.
/// - `state`: The participation status of the user for the event.
class Comment extends Model {
  int userId;
  int eventId;
  String content;
  ParticipationStatus state;
  Comment({
    required super.id,
    required this.userId,
    required this.eventId,
    required this.state,
    required this.content,
  });
}
