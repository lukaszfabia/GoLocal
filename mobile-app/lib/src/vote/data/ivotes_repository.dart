import 'package:golocal/src/vote/domain/vote.dart';

abstract class IVotesRepository {
  Future<List<Vote>> getVotes();
  Future<Vote> getVote(String id);
  Future<void> createVote(Vote vote);
  Future<void> updateVote(Vote vote);
  Future<void> deleteVote(String id);
}
