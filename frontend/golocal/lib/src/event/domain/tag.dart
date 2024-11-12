import 'package:golocal/src/shared/model_base.dart';

class Tag extends Model {
  String name;

  Tag({
    required super.id,
    required this.name,
  });
}
