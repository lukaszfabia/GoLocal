import 'package:golocal/src/shared/data_source_base.dart';

// String firstName;
// String lastName;
// String email;
// DateTime birthDate;

// bool isVerified;
// bool isPremium;
// String? avatarUrl;

class UserLocalDataSource implements LocalDataSource {
  List<Map<String, dynamic>> users = <Map<String, dynamic>>[
    {
      'id': 1,
      'firstName': 'John',
      'lastName': 'Doe',
      'email': 'johndoe@gmail.com',
      'isVerified': true,
      'isPremium': false,
      'avatarUrl': null,
      'birthDate': '1990-01-01',
    },
    {
      'id': 2,
      'firstName': 'Bella',
      'lastName': 'Sprrow',
      'email': 'bellasparrow@gmail.com',
      'isVerified': true,
      'isPremium': true,
      'avatarUrl': null,
      'birthDate': '1990-01-01',
    },
    {
      'id': 3,
      'firstName': 'John',
      'lastName': 'Doe',
      'email': 'johndoe@gmail.com',
      'isVerified': false,
      'isPremium': false,
      'avatarUrl': null,
      'birthDate': '1990-01-01',
    },
  ];
  @override
  Future<bool> create(Map<String, dynamic> data) {
    users.add(data);
    return Future.value(true);
  }

  @override
  Future<bool> delete(
    Map<String, dynamic> data,
  ) {
    return Future.value(users.remove(data));
  }

  @override
  Future<List<Map<String, dynamic>>> getAll() {
    return Future.value(users);
  }

  @override
  Future<Map<String, dynamic>> getById(int id) {
    return Future.value(users.firstWhere((element) => element['id'] == id));
  }

  @override
  Future<bool> update(Map<String, dynamic> data) {
    users[users.indexWhere((element) => element['id'] == data['id'])] = data;
    return Future.value(true);
  }
}
