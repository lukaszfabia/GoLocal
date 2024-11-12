import 'package:golocal/src/shared/model_base.dart';

class CoordsDTO extends Model {
  double latitude;
  double longitude;

  CoordsDTO({
    required super.id,
    required this.latitude,
    required this.longitude,
  });
}
