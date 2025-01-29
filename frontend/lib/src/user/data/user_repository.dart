import 'package:golocal/src/shared/repository_base.dart';
import 'package:golocal/src/user/data/sources/user_datasource.dart';
import 'package:golocal/src/user/domain/user.dart';

class UserRepository {
  final UserDataSource _userDataSource = UserDataSource();

  Future<int?> getLoggedUserId() async {
    final response = await _userDataSource.getLoggedUserId();
    return response;
  }
}
