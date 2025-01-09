import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/vote/domain/vote_option.dart';
import 'package:golocal/src/event/domain/event.dart';

enum VoteType { single, multiple }

class Vote extends Model {
  final Event event;
  final String text;
  final List<VoteOption> options;
  final VoteType type;
  final DateTime endsAt;

  Vote({
    required super.id,
    this.text = '',
    required this.options,
    required this.event,
    required this.type,
    required this.endsAt,
  });

  factory Vote.fromJson(Map<String, dynamic> json) {
    return Vote(
      id: json['ID'],
      options: (json['options'] as List)
          .map((option) => VoteOption.fromJson(option))
          .toList(),
      text: json['text'],
      event: Event.fromJson(json['event']),
      type: VoteType.values
          .firstWhere((e) => e.toString() == 'VoteType.${json['type']}'),
      endsAt: DateTime.parse(json['endsAt']),
    );
  }

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
