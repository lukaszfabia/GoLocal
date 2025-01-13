import 'package:golocal/src/shared/model_base.dart';

class VoteOption extends Model {
  final String text;
  final int votesCount;
  bool isSelected;

  VoteOption({
    required super.id,
    required this.text,
    required this.votesCount,
    required this.isSelected,
  });

  factory VoteOption.fromJson(Map<String, dynamic> json) {
    return VoteOption(
      id: json['ID'],
      text: json['text'],
      votesCount: json['votesCount'],
      isSelected: json['isSelected'] ?? false,
    );
  }

  @override
  Map<String, dynamic> toJson() {
    return {
      'ID': id,
      'text': text,
      'votesCount': votesCount,
      'isSelected': isSelected,
    };
  }

  VoteOption copyWith({
    int? id,
    String? text,
    int? votesCount,
    bool? isSelected,
  }) {
    return VoteOption(
      id: id ?? this.id,
      text: text ?? this.text,
      votesCount: votesCount ?? this.votesCount,
      isSelected: isSelected ?? this.isSelected,
    );
  }
}
