import 'package:golocal/src/shared/model_base.dart';

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
