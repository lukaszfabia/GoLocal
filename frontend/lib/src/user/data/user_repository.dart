import 'package:golocal/src/user/data/sources/user_datasource.dart';

class UserRepository {
  final UserDataSource _userDataSource = UserDataSource();

  Future<int?> getLoggedUserId() async {
    final response = await _userDataSource.getLoggedUserId();
    return response;
  }
}
