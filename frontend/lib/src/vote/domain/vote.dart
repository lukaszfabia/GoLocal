import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/vote/domain/vote_option.dart';
import 'package:golocal/src/event/domain/event.dart';

enum VoteType { CAN_CHANGE_VOTE, CANNOT_CHANGE_VOTE }

class Vote extends Model {
  final Event event;
  final String text;
  final List<VoteOption> options;
  final VoteType type;
  final DateTime? endsAt;

  Vote({
    required super.id,
    this.text = '',
    required this.options,
    required this.event,
    required this.type,
    required this.endsAt,
  });

  Vote.fromJson(super.json)
      : options = json["options"] != null
            ? (json['options'] as List)
                .map((option) => VoteOption.fromJson(option))
                .toList()
            : []
          ..sort((a, b) => a.text.compareTo(b.text)),
        text = json['text'],
        event = Event.fromJson(json['event']),
        type = VoteType.values
            .firstWhere((e) => e.toString() == 'VoteType.${json['voteType']}'),
        endsAt =
            json['endDate'] != null ? DateTime.parse(json['endDate']) : null,
        super.fromJson();

  @override
  Map<String, dynamic> toJson() {
    return {
      'ID': id,
      'text': text,
      'options': options.map((option) => option.toJson()).toList(),
      'event': event.toJson(),
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
