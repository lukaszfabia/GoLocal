import 'package:golocal/src/event/domain/address.dart';
import 'package:golocal/src/event/domain/coords.dart';
import 'package:golocal/src/shared/model_base.dart';

/// Represents the location of an event.
///
/// The location includes details such as the city, country, and optional coordinates and address.
///
/// Attributes:
/// - `city`: The city where the event is located.
/// - `country`: The country where the event is located.
/// - `zip`: The zip code of the event location (optional).
/// - `coords`: The geographical coordinates of the event location (optional).
/// - `address`: The address of the event location (optional).
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
