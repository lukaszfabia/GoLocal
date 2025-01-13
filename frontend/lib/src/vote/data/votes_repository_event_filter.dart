import 'package:golocal/src/vote/data/ivotes_repository.dart';
import 'package:golocal/src/vote/domain/vote.dart';
import 'package:golocal/src/event/domain/event.dart';

class VotesRepositoryEventFilter implements IVotesRepository {
  final Event event;
  final IVotesRepository _votesRepository;

  VotesRepositoryEventFilter(this.event, this._votesRepository);

  @override
  Future<void> createVote(Vote vote) async {
    _votesRepository.createVote(vote);
  }

  @override
  Future<void> deleteVote(String id) async {
    _votesRepository.deleteVote(id);
  }

  @override
  Future<Vote> getVote(String id) async {
    return _votesRepository.getVote(id);
  }

  @override
  Future<List<Vote>> getVotes() async {
    return _votesRepository.getVotes().then(
        (votes) => votes.where((vote) => vote.event.id == event.id).toList());
  }

  @override
  Future<void> updateVote(Vote vote) async {
    _votesRepository.updateVote(vote);
  }

  @override
  Future<void> voteOnOption(int voteId, int optionId) async {
    _votesRepository.voteOnOption(voteId, optionId);
  }
}
