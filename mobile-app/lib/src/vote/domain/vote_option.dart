import 'package:golocal/src/shared/model_base.dart';

class VoteOption extends Model {
  final String text;
  final int votesCount;

  VoteOption({
    required super.id,
    required this.text,
    required this.votesCount,
  });

  factory VoteOption.fromJson(Map<String, dynamic> json) {
    return VoteOption(
      id: json['ID'],
      text: json['text'],
      votesCount: json['votesCount'],
    );
  }

  @override
  Map<String, dynamic> toJson() {
    return {
      'ID': id,
      'Text': text,
    };
  }
}
