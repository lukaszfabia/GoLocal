import 'package:golocal/src/user/data/sources/user_local.dart';
import 'package:golocal/src/user/data/sources/user_remote.dart';
import 'package:golocal/src/user/data/user_repository.dart';

// void main() async {
//   runApp(const GoLocalApp());
// }

void main() async {
  var userrepository = UserRepository(
      remoteDataSource: UserRemoteDataSource(),
      localDataSource: UserLocalDataSource());

  var user = await userrepository.getById(1);
  await userrepository.create(user);
  print(user.toJson());
}
