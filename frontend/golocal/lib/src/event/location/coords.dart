import 'package:golocal/src/shared/model_base.dart';

class Coords extends Model {
  double latitude;
  double longitude;

  Coords({
    required super.id,
    required this.latitude,
    required this.longitude,
  });

  Coords.fromJson(super.json)
      : latitude = json['latitude'],
        longitude = json['longitude'],
        super.fromJson();

  @override
  Map<String, dynamic> toJson() {
    final data = super.toJson();
    data['latitude'] = latitude;
    data['longitude'] = longitude;
    return data;
  }
}
