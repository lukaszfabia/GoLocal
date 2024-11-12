import 'package:golocal/src/event/location/address.dart';
import 'package:golocal/src/event/location/coords.dart';
import 'package:golocal/src/shared/model_base.dart';

class LocationDTO extends Model {
  String city;
  String country;
  String zip;

  CoordsDTO coords;
  AddressDTO address;
  LocationDTO({
    required super.id,
    required this.city,
    required this.country,
    required this.zip,
    required this.coords,
    required this.address,
  });
}
