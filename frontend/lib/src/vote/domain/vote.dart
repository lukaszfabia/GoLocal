import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/vote/domain/vote_option.dart';
import 'package:golocal/src/event/domain/event.dart';

enum VoteType { canChangeVote, cannotChangeVote }

/// Represents a vote in the system.
///
/// A vote can be associated with an event and contains multiple vote options.
/// The vote can have a type indicating whether the vote can be changed or not,
/// and an optional end date.
///
/// The [Vote] class extends the [Model] class.
///
/// Properties:
/// - `event`: The event associated with the vote.
/// - `text`: The text description of the vote.
/// - `options`: The list of vote options.
/// - `type`: The type of the vote, indicating if it can be changed or not.
/// - `endsAt`: The end date of the vote.
///
/// Constructors:
/// - `Vote`: Creates a new vote instance.
/// - `Vote.fromJson`: Creates a new vote instance from a JSON object.
///
/// Methods:
/// - `parseVoteType`: Parses a string to a [VoteType].
/// - `toJson`: Converts the vote instance to a JSON object.
/// - `copyWith`: Creates a copy of the vote instance with optional new values.
class Vote extends Model {
  final Event? event;
  final String text;
  final List<VoteOption> options;
  final VoteType type;
  final DateTime? endsAt;

  Vote({
    required super.id,
    required this.options,
    required this.type,
    required this.endsAt,
    this.text = '',
    this.event,
  });

  Vote.fromJson(super.json)
      : options = json["options"] != null
            ? (json['options'] as List)
                .map((option) => VoteOption.fromJson(option))
                .toList()
            : []
          ..sort((a, b) => a.text.compareTo(b.text)),
        text = json['text'],
        event = json["event"] != null ? Event.fromJson(json['event']) : null,
        type = parseVoteType(json['voteType']),
        endsAt =
            json['endDate'] != null ? DateTime.parse(json['endDate']) : null,
        super.fromJson();

  static VoteType parseVoteType(String type) {
    switch (type) {
      case 'CAN_CHANGE_VOTE':
        return VoteType.canChangeVote;
      case 'CANNOT_CHANGE_VOTE':
        return VoteType.cannotChangeVote;
      default:
        print('Unknown vote type: $type');
        return VoteType.canChangeVote;
    }
  }

  @override
  Map<String, dynamic> toJson() {
    return {
      'ID': id,
      'text': text,
      'options': options.map((option) => option.toJson()).toList(),
      'event': event?.toJson(),
    };
  }

  Vote copyWith({
    String? text,
    List<VoteOption>? options,
    Event? event,
    VoteType? type,
    DateTime? endsAt,
  }) {
    return Vote(
      id: id,
      text: text ?? this.text,
      options: options ?? this.options,
      event: event ?? this.event,
      type: type ?? this.type,
      endsAt: endsAt ?? this.endsAt,
    );
  }
}
