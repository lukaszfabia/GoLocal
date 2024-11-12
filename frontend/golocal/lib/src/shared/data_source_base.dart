abstract class DataSource {
  // For now it returns either true or false,
  // TODO: change to Either<Success,Failure>
  Future<bool> create(Map<String, dynamic> data);
  Future<bool> update(Map<String, dynamic> data);
  Future<bool> delete(Map<String, dynamic> data);
  Future<List<Map<String, dynamic>>> getAll();
  Future<Map<String, dynamic>> getById(int id);
}

abstract class RemoteDataSource extends DataSource {}

abstract class LocalDataSource extends DataSource {}
