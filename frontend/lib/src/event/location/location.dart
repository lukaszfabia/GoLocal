import 'package:golocal/src/event/location/address.dart';
import 'package:golocal/src/event/location/coords.dart';
import 'package:golocal/src/shared/model_base.dart';

class Location extends Model {
  String city;
  String country;
  String? zip;

  Coords? coords;
  Address? address;
  Location({
    required super.id,
    required this.city,
    required this.country,
    this.zip,
    this.coords,
    this.address,
  });

  // super.json doesn't work for some reason
  Location.fromJson(super.json)
      : city = json['city'],
        country = json['country'],
        zip = json['zip'],
        coords =
            json['coords'] != null ? Coords.fromJson(json['coords']) : null,
        address =
            json['address'] != null ? Address.fromJson(json['address']) : null,
        super.fromJson();

  @override
  Map<String, dynamic> toJson() {
    final data = super.toJson();
    data['city'] = city;
    data['country'] = country;
    data['zip'] = zip ?? '';
    data['coords'] = coords?.toJson() ?? {};
    data['address'] = address?.toJson() ?? {};
    return data;
  }

  @override
  String toString() {
    return 'Location{city: $city, country: $country, zip: $zip, coords: $coords, address: $address}';
  }
}
