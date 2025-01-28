import 'package:dio/dio.dart';
import 'package:golocal/src/vote/data/ivotes_repository.dart';
import 'package:golocal/src/vote/domain/vote.dart';
import 'package:golocal/src/dio_client.dart';

/// Implementation of the [IVotesRepository] interface using Dio for HTTP requests.
class VotesRepositoryImpl implements IVotesRepository {
  /// Dio client for making HTTP requests.
  final DioClient _dioClient = DioClient();

  /// Fetches a single vote by its [id].
  ///
  /// Makes a GET request to the `/auth/vote/?=all` endpoint and filters the
  /// results to find the vote with the matching [id].
  ///
  /// Returns a [Vote] object if found.
  @override
  Future<Vote> getVote(String id) async {
    final response = await _dioClient.dio.get('/auth/vote/?=all');

    final data = response.data['data'] as List<dynamic>;

    return data
        .map((json) => Vote.fromJson(json))
        .toList()
        .firstWhere((vote) => vote.id.toString() == id);
  }

  /// Fetches a list of votes.
  ///
  /// Makes a GET request to the `/auth/vote/10` endpoint and returns a list
  /// of [Vote] objects.
  @override
  Future<List<Vote>> getVotes() async {
    final response = await _dioClient.dio.get('/auth/vote/10');

    final data = response.data['data'] as List<dynamic>;

    return data.map((json) => Vote.fromJson(json)).toList();
  }

  /// Creates a new vote.
  ///
  /// This method is not yet implemented.
  @override
  Future<void> createVote(Vote vote) async {
    // TODO: implement createVote
    throw UnimplementedError();
  }

  /// Deletes a vote by its [id].
  ///
  /// This method is not yet implemented.
  @override
  Future<void> deleteVote(String id) async {
    // TODO: implement deleteVote
    throw UnimplementedError();
  }

  /// Fetches a list of votes for a specific event by its [eventId].
  ///
  /// Makes a GET request to the `/auth/vote/?eventID=$eventId` endpoint and
  /// returns a list of [Vote] objects.
  @override
  Future<List<Vote>> getVotesForEvent(String eventId) async {
    final response = await _dioClient.dio.get('/auth/vote/?eventID=$eventId');

    final data = response.data['data'] != null
        ? response.data['data'] as List<dynamic>
        : [];

    return data.map((json) => Vote.fromJson(json)).toList();
  }

  /// Updates an existing vote.
  ///
  /// This method is not yet implemented.
  @override
  Future<void> updateVote(Vote vote) async {
    // TODO: implement updateVote
    throw UnimplementedError();
  }

  /// Casts a vote on a specific option.
  ///
  /// Makes a POST request to the `/auth/vote/` endpoint with the [voteId] and
  /// [optionId]. If the vote cannot be changed, throws an exception with a
  /// specific message.
  @override
  Future<void> voteOnOption(int voteId, int optionId) async {
    try {
      await _dioClient.dio.post('/auth/vote/', data: {
        'voteID': voteId,
        'voteOptionID': optionId,
      });
    } on DioException catch (e) {
      if (e.response?.statusCode == 400 &&
          e.response?.data['message'] ==
              'You tried to change vote on a vote that doesn\'t allow changing votes') {
        throw Exception('you can\'t change vote');
      }
      rethrow;
    }
  }
}
