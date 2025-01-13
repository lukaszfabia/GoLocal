import 'package:golocal/src/vote/domain/vote.dart';
import 'package:golocal/src/vote/data/ivotes_repository.dart';
import 'package:golocal/src/vote/data/votes_repository_dummy.dart';

class VoteService implements IVotesRepository {
  @override
  Future<void> createVote(Vote vote) async {
    // Dummy implementation
  }

  @override
  Future<void> deleteVote(String id) async {
    // Dummy implementation
  }

  @override
  Future<Vote> getVote(String id) async {
    return VotesRepositoryDummy().getVote(id);
  }

  @override
  Future<List<Vote>> getVotes() async {
    return VotesRepositoryDummy().getVotes();
  }

  @override
  Future<void> updateVote(Vote vote) async {
    print("Updating vote: $vote, vote service");
    return VotesRepositoryDummy().updateVote(vote);
  }

  @override
  Future<void> voteOnOption(int voteId, int optionId) async {
    print("Voting on option: $optionId for vote: $voteId");
    return VotesRepositoryDummy().voteOnOption(voteId, optionId);
  }
}
