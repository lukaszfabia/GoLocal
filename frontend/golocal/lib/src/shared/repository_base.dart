import 'package:golocal/src/shared/data_source_base.dart';
import 'package:golocal/src/shared/model_base.dart';

abstract class Repository<T extends Model, V extends RemoteDataSource,
    K extends LocalDataSource> {
  final V remoteDataSource;
  final K localDataSource;
  Repository({
    required this.remoteDataSource,
    required this.localDataSource,
  });

  Future<bool> create(T model);
  Future<bool> update(T model);
  Future<bool> delete(T model);
  Future<T> getById(int id);
}
