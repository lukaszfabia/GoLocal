import 'package:flutter_test/flutter_test.dart';
import 'package:golocal/src/user/domain/user.dart';

void main() {
  group('Test user model methods: valid values', () {
    test('User.fromJson()', () {
      var userJson = {
        'id': 1,
        'firstName': 'John',
        'lastName': 'Doe',
        'email': 'johndoe@example.com',
        'birthDate': '1990-01-01',
        'isVerified': true,
        'isPremium': false,
        'avatarUrl': 'avatar/url',
      };
      var user = User.fromJson(userJson);
      expect(user.id, userJson['id']);
      expect(user.firstName, userJson['firstName']);
      expect(user.lastName, userJson['lastName']);
      expect(user.email, userJson['email']);
      expect(user.birthDate, DateTime.parse(userJson['birthDate'].toString()));
      expect(user.isVerified, userJson['isVerified']);
      expect(user.isPremium, userJson['isPremium']);
      expect(user.avatarUrl, userJson['avatarUrl']);
    });
    test('User.toJson()', () {
      var user = User(
        id: 1,
        firstName: 'John',
        lastName: 'Doe',
        email: 'johndoe@example.com',
        birthDate: DateTime(1990, 1, 1),
        isVerified: true,
        isPremium: false,
        avatarUrl: 'avatar/url',
      );
      var userJson = user.toJson();
      expect(userJson['id'], user.id);
      expect(userJson['firstName'], user.firstName);
      expect(userJson['lastName'], user.lastName);
      expect(userJson['email'], user.email);
      expect(userJson['birthDate'], user.birthDate.toString());
      expect(userJson['isVerified'], user.isVerified);
      expect(userJson['isPremium'], user.isPremium);
      expect(userJson['avatarUrl'], user.avatarUrl);
    });
  });
}
