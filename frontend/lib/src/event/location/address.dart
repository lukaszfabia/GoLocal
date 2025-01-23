import 'package:golocal/src/shared/model_base.dart';

class Address extends Model {
  String street;
  int streetNumber;
  String? additionalInfo;
  Address({
    required super.id,
    required this.street,
    required this.streetNumber,
    this.additionalInfo,
  });

  Address.fromJson(super.json)
      : street = json['street'],
        streetNumber = int.tryParse(json['streetNumber']) ?? 0,
        additionalInfo = json['additionalInfo'],
        super.fromJson();

  @override
  Map<String, dynamic> toJson() {
    final data = super.toJson();
    data['street'] = street;
    data['streetNumber'] = streetNumber;
    data['additionalInfo'] = additionalInfo ?? '';
    return data;
  }
}
