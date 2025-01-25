import 'package:golocal/src/shared/model_base.dart';

class User extends Model {
  String firstName;
  String lastName;
  String email;
  DateTime? birthDate;

  bool isVerified;
  bool isPremium;
  String? avatarUrl;

  // TODO: comments, votes??? wg. mnie nie

  User({
    required super.id,
    required this.firstName,
    required this.lastName,
    required this.email,
    this.birthDate,
    required this.isVerified,
    required this.isPremium,
    this.avatarUrl,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'] ?? 0,
      firstName: json['firstName'] ?? '',
      lastName: json['lastName'] ?? '',
      email: json['email'] ?? '',
      birthDate: json['birthDate'] != null && json['birthDate'].isNotEmpty
          ? DateTime.parse(json['birthDate'])
          : null,
      isVerified: json['isVerified'] ?? false,
      isPremium: json['isPremium'] ?? false,
      avatarUrl: json['avatarUrl'] ?? '',
    );
  }
  @override
  Map<String, dynamic> toJson() {
    return {
      'firstName': firstName,
      'lastName': lastName,
      'email': email,
      'birthDate': birthDate.toString(),
      'isVerified': isVerified,
      'isPremium': isPremium,
      'avatarUrl': avatarUrl,
    }..addAll(super.toJson());
  }
}
