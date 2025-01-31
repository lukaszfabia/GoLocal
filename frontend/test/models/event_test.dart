import 'package:flutter_test/flutter_test.dart';
import 'package:golocal/src/event/domain/address.dart';
import 'package:golocal/src/event/domain/coords.dart';
import 'package:golocal/src/event/domain/location.dart';

void main() {
  var coordsJson = {
    'id': 1,
    'latitude': 10.0,
    'longitude': 10.0,
  };
  var coords = Coords(id: 1, latitude: 10.0, longitude: 10.0);
  group("Test coords model methods: valid values", () {
    test("Coords.fromJson()", () {
      var coords = Coords.fromJson(coordsJson);
      expect(coords.id, coordsJson['id']);
      expect(coords.latitude, coordsJson['latitude']);
      expect(coords.longitude, coordsJson['longitude']);
    });
    test("Coords.toJson()", () {
      var coordsJson = coords.toJson();
      expect(coordsJson['id'], coords.id);
      expect(coordsJson['latitude'], coords.latitude);
      expect(coordsJson['longitude'], coords.longitude);
    });
  });

  var addressJson = {
    'id': 1,
    'street': 'Plac Grunwaldzki',
    'streetNumber': "1",
    'additionalInfo': 'additionalInfo',
  };
  var address = Address(
    id: 1,
    street: 'Plac Grunwaldzki',
    streetNumber: 1,
    additionalInfo: 'additionalInfo',
  );
  group("Test address model methods: valid values", () {
    test("Address.fromJson()", () {
      var address = Address.fromJson(addressJson);
      expect(address.id, addressJson['id']);
      expect(address.street, addressJson['street']);
      expect(address.streetNumber,
          int.parse(addressJson['streetNumber'] as String));
      expect(address.additionalInfo, addressJson['additionalInfo']);
    });
    test("Address.toJson()", () {
      var addressJson = address.toJson();
      expect(addressJson['id'], address.id);
      expect(addressJson['street'], address.street);
      expect(addressJson['streetNumber'], address.streetNumber);
      expect(addressJson['additionalInfo'], address.additionalInfo);
    });
  });

  var locationJson = {
    'id': 1,
    'city': 'Wroclaw',
    'country': 'Poland',
    'zip': '50-000',
    'coords': coordsJson,
    'address': addressJson,
  };
  var location = Location(
      id: 1,
      city: 'Wroclaw',
      country: 'Poland',
      zip: '50-000',
      coords: coords,
      address: address);

  group("Test location model methods", () {
    test("Location.fromJson()", () {
      var location = Location.fromJson(locationJson);
      expect(location.id, locationJson['id']);
      expect(location.city, locationJson['city']);
      expect(location.country, locationJson['country']);
      expect(location.zip, locationJson['zip']);
      var coords = locationJson['coords'] as Map<String, dynamic>;
      expect(location.coords!.latitude, coords['latitude']);
      expect(location.coords!.longitude, coords['longitude']);
      var address = locationJson['address'] as Map<String, dynamic>;
      expect(location.address!.street, address['street']);
      expect(
          location.address!.streetNumber, int.parse(address['streetNumber']));
      expect(location.address!.additionalInfo, address['additionalInfo']);
    });
    test("Location.toJson()", () {
      var locationJson = location.toJson();
      expect(locationJson['id'], location.id);
      expect(locationJson['city'], location.city);
      expect(locationJson['country'], location.country);
      expect(locationJson['zip'], location.zip);
      expect(locationJson['coords']['latitude'], location.coords!.latitude);
      expect(locationJson['coords']['longitude'], location.coords!.longitude);
      expect(locationJson['address']['street'], location.address!.street);
      expect(locationJson['address']['streetNumber'],
          location.address!.streetNumber);
      expect(locationJson['address']['additionalInfo'],
          location.address!.additionalInfo);
    });
  });
  group("Test event model methods: valid values", () {
    test("Event.fromJson()", () {});
  });
}
