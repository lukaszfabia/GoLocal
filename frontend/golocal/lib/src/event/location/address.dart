import 'package:golocal/src/shared/model_base.dart';

class AddressDTO extends Model {
  String street;
  int streetNumber;
  String sdditionalInfo;
  AddressDTO({
    required super.id,
    required this.street,
    required this.streetNumber,
    required this.sdditionalInfo,
  });
}
