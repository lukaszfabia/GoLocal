import 'package:golocal/src/shared/repository_base.dart';
import 'package:golocal/src/user/data/sources/user_local.dart';
import 'package:golocal/src/user/data/sources/user_remote.dart';
import 'package:golocal/src/user/domain/user.dart';

class UserRepository
    implements Repository<User, UserRemoteDataSource, UserLocalDataSource> {
  @override
  final UserRemoteDataSource remoteDataSource;
  @override
  final UserLocalDataSource localDataSource;
  UserRepository({
    required this.remoteDataSource,
    required this.localDataSource,
  });

  @override
  Future<bool> create(User model) {
    return localDataSource.create(model.toJson());
  }

  @override
  Future<bool> delete(User model) {
    return localDataSource.delete(model.toJson());
  }

  @override
  Future<User> getById(int id) {
    Future<Map<String, dynamic>> user = localDataSource.getById(id);
    return user.then((value) => User.fromJson(value));
  }

  @override
  Future<bool> update(User model) {
    return localDataSource.update(model.toJson());
  }
}
