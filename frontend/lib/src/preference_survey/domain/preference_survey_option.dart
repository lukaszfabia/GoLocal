import 'package:golocal/src/shared/model_base.dart';

/// Represents an option in a preference survey.
///
/// An `Option` contains the text of the option and whether it is selected or not.
/// It extends the `Model` class which provides an `id` for the option.
class Option extends Model {
  /// The text of the option.
  final String text;

  /// Indicates whether the option is selected.
  final bool isSelected;

  /// Creates an `Option` with the given [id], [text], and [isSelected] status.
  ///
  /// The [id] is inherited from the `Model` class.
  Option({
    required super.id,
    required this.text,
    required this.isSelected,
  });

  /// Creates an `Option` from a JSON object.
  ///
  /// The [json] parameter must contain the keys 'id', 'text', and optionally 'isSelected'.
  /// If 'isSelected' is not provided, it defaults to `false`.
  factory Option.fromJson(Map<String, dynamic> json) {
    return Option(
      id: json['id'],
      text: json['text'],
      isSelected: json['isSelected'] ?? false,
    );
  }
}
