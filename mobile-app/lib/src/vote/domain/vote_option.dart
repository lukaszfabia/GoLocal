import 'package:golocal/src/shared/model_base.dart';

class VoteOption extends Model {
  final String text;

  VoteOption({
    required super.id,
    required this.text,
  });

  factory VoteOption.fromJson(Map<String, dynamic> json) {
    return VoteOption(
      id: json['ID'],
      text: json['Text'],
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
