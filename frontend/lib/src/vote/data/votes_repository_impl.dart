import 'package:golocal/src/vote/data/ivotes_repository.dart';
import 'package:golocal/src/vote/domain/vote.dart';
import 'package:golocal/src/dio_client.dart';

class VotesRepositoryImpl implements IVotesRepository {
  final DioClient _dioClient = DioClient();

  @override
  Future<Vote> getVote(String id) async {
    final response = await _dioClient.dio.get('/auth/vote/?=all');

    final data = response.data['data'] as List<dynamic>;

    return data
        .map((json) => Vote.fromJson(json))
        .toList()
        .firstWhere((vote) => vote.id.toString() == id);
  }

  @override
  Future<List<Vote>> getVotes() async {
    final response = await _dioClient.dio.get('/auth/vote/10');

    final data = response.data['data'] as List<dynamic>;

    return data.map((json) => Vote.fromJson(json)).toList();
  }

  @override
  Future<void> createVote(Vote vote) async {
    // TODO: implement createVote
    throw UnimplementedError();
  }

  @override
  Future<void> deleteVote(String id) async {
    // TODO: implement deleteVote
    throw UnimplementedError();
  }

  @override
  Future<List<Vote>> getVotesForEvent(String eventId) async {
    final response = await _dioClient.dio.get('/auth/vote/?eventID=$eventId');

    final data = response.data['data'] != null
        ? response.data['data'] as List<dynamic>
        : [];

    return data.map((json) => Vote.fromJson(json)).toList();
  }

  @override
  Future<void> updateVote(Vote vote) async {
    // TODO: implement updateVote
    throw UnimplementedError();
  }

  @override
  Future<void> voteOnOption(int voteId, int optionId) async {
    final response = await _dioClient.dio.post('/auth/vote/', data: {
      'voteID': voteId,
      'voteOptionID': optionId,
    });

    if (response.statusCode != 200) {
      throw Exception('Failed to vote on option');
    }
  }
}
