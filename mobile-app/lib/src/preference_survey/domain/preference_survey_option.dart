import 'package:golocal/src/shared/model_base.dart';

class Option extends Model {
  final String text;
  final bool isSelected;

  Option({
    required super.id,
    required this.text,
    required this.isSelected,
  });

  factory Option.fromJson(Map<String, dynamic> json) {
    return Option(
      id: json['ID'],
      text: json['text'],
      isSelected: json['isSelected'],
    );
  }
}
