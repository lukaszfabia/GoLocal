import 'package:golocal/src/shared/model_base.dart';

/// A class representing a voting option in a poll or survey.
///
/// The [VoteOption] class extends the [Model] class and includes properties
/// for the option's text, the number of votes it has received, and whether
/// it is currently selected.
///
/// The class provides methods for JSON serialization and deserialization,
/// as well as a method for creating a copy of an instance with modified
/// properties.
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

  VoteOption.fromJson(super.json)
      : text = json['text'],
        votesCount = json["votesCount"] != null ? json['votesCount'] : 0,
        isSelected = json['isSelected'] ?? false,
        super.fromJson();

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
