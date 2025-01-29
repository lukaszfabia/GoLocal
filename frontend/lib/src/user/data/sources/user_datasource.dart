import 'package:golocal/src/dio_client.dart';

// TODO: Implement when account routes are finished
class UserDataSource {
  Future<int?> getLoggedUserId() async {
    final int? userId = await DioClient().getLoggedUserId();
    return userId;
  }
}
