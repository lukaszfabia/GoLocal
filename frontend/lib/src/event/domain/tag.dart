import 'package:golocal/src/shared/model_base.dart';

/// Represents a tag associated with an event.
///
/// Tags are used to categorize events and make them easier to search and filter.
///
/// Attributes:
/// - `name`: The name of the tag.
class Tag extends Model {
  String name;

  Tag({
    required super.id,
    required this.name,
  });

  Tag.fromJson(super.json)
      : name = json['Name'],
        super.fromJson();

  @override
  Map<String, dynamic> toJson() {
    final data = super.toJson();
    data['name'] = name;
    return data;
  }
}
