import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/comment/domain/participationstatus_enum.dart';

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
}
