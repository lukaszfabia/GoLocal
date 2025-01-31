import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/event/domain/participationstatus_enum.dart';

// TODO: ten vote jest zduplikowany i przestarzały
/// Represents a vote for an event by a user.
///
/// Each vote is associated with a specific user and event, and indicates the user's participation status.
///
/// Attributes:
/// - `userId`: The ID of the user who cast the vote.
/// - `eventId`: The ID of the event the vote is associated with.
/// - `state`: The participation status of the user for the event.
class Vote extends Model {
  int userId;
  int eventId;
  ParticipationStatus state;
  Vote({
    required super.id,
    required this.userId,
    required this.eventId,
    required this.state,
  });

  Vote.fromJson(super.json)
      : userId = json['userId'],
        eventId = json['eventId'],
        state = json['state'],
        super.fromJson();
}
