import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/comment/domain/participationstatus_enum.dart';

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
