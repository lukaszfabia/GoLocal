import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/shared/data_source_base.dart';
import 'package:golocal/src/shared/repository_base.dart';

class EventRepository
    implements Repository<Event, EventRemoteDataSource, EventLocalDataSource> {
  @override
  final EventRemoteDataSource remoteDataSource;
  @override
  final EventLocalDataSource localDataSource;

  EventRepository({
    required this.remoteDataSource,
    required this.localDataSource,
  });

  @override
  Future<bool> create(Event model) {
    // TODO
    return localDataSource.create(model.toJson());
  }

  @override
  Future<bool> delete(Event model) {
    // TODO
    localDataSource.delete(model.toJson());
    remoteDataSource.delete(model.toJson());
    return Future.value(true);
  }

  @override
  Future<Event> getById(int id) {
    return localDataSource.getById(id).then((value) => Event.fromJson(value));
  }

  @override
  Future<bool> update(Event model) {
    // TODO: implement update
    throw UnimplementedError();
  }
}

class EventLocalDataSource implements LocalDataSource {
  @override
  Future<bool> create(Map<String, dynamic> data) {
    // TODO: implement create
    throw UnimplementedError();
  }

  @override
  Future<bool> delete(Map<String, dynamic> data) {
    // TODO: implement delete
    throw UnimplementedError();
  }

  @override
  Future<List<Map<String, dynamic>>> getAll() {
    // TODO: implement getAll
    throw UnimplementedError();
  }

  @override
  Future<Map<String, dynamic>> getById(int id) {
    // TODO: implement getById
    throw UnimplementedError();
  }

  @override
  Future<bool> update(Map<String, dynamic> data) {
    // TODO: implement update
    throw UnimplementedError();
  }
}

class EventRemoteDataSource implements RemoteDataSource {
  @override
  Future<bool> create(Map<String, dynamic> data) {
    // TODO: implement create
    throw UnimplementedError();
  }

  @override
  Future<bool> delete(Map<String, dynamic> data) {
    // TODO: implement delete
    throw UnimplementedError();
  }

  @override
  Future<List<Map<String, dynamic>>> getAll() {
    // TODO: implement getAll
    throw UnimplementedError();
  }

  @override
  Future<Map<String, dynamic>> getById(int id) {
    // TODO: implement getById
    throw UnimplementedError();
  }

  @override
  Future<bool> update(Map<String, dynamic> data) {
    // TODO: implement update
    throw UnimplementedError();
  }
}
