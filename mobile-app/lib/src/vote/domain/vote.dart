import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/vote/domain/vote_option.dart';
import 'package:golocal/src/event/domain/event.dart';

class Vote extends Model {
  final Event event;
  final String text;
  final List<VoteOption> options;

  Vote({
    required super.id,
    this.text = '',
    required this.options,
    required this.event,
  });

  factory Vote.fromJson(Map<String, dynamic> json) {
    return Vote(
      id: json['ID'],
      options: (json['options'] as List)
          .map((option) => VoteOption.fromJson(option))
          .toList(),
      text: json['text'],
      event: Event.fromJson(json['event']),
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
  }) {
    return Vote(
      id: id,
      text: text ?? this.text,
      options: options ?? this.options,
      event: event ?? this.event,
    );
  }
}
