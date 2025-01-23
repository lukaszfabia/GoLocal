import 'package:golocal/src/vote/data/ivotes_repository.dart';
import 'package:golocal/src/vote/domain/vote.dart';
import 'package:golocal/src/vote/domain/vote_option.dart';
import 'package:golocal/src/event/data/events_repository_dummy.dart';

class VotesRepositoryDummy implements IVotesRepository {
  final List<Vote> _votes = dummyVotes();

  static List<Vote> dummyVotes() {
    final mockEvents = EventsRepositoryDummy().dummyEvents();

    return [
      Vote(
        id: 1,
        text: 'Dummy Vote',
        options: [
          VoteOption(id: 1, text: 'Option 1', votesCount: 4, isSelected: false),
          VoteOption(id: 2, text: 'Option 2', votesCount: 2, isSelected: false),
        ],
        event: mockEvents[0],
        type: VoteType.single,
        endsAt: DateTime.now().add(Duration(days: 1)),
      ),
      Vote(
        id: 2,
        text: 'Dummy Vote 2',
        options: [
          VoteOption(
              id: 3, text: 'Option 2.1', votesCount: 3421, isSelected: false),
          VoteOption(
              id: 4, text: 'Option 2.2', votesCount: 123, isSelected: true),
        ],
        type: VoteType.multiple,
        event: mockEvents[1],
        endsAt: DateTime.now().add(Duration(days: 2)),
      )
    ];
  }

  @override
  Future<void> createVote(Vote vote) async {
    _votes.add(vote);
  }

  @override
  Future<void> deleteVote(String id) async {
    _votes.removeWhere((vote) => vote.id.toString() == id);
  }

  @override
  Future<Vote> getVote(String id) async {
    return _votes.firstWhere((vote) => vote.id.toString() == id);
  }

  @override
  Future<List<Vote>> getVotes() async {
    return _votes;
  }

  @override
  Future<void> updateVote(Vote vote) async {
    print("Updating vote: $vote");
    final index = _votes.indexWhere((v) => v.id == vote.id);
    if (index != -1) {
      _votes[index] = vote;
    }
  }

  @override
  Future<void> voteOnOption(int voteId, int optionId) async {
    print("Voting on option: $optionId for vote: $voteId");
    final vote = await getVote(voteId.toString());
    final updatedOptions = vote.options.map((option) {
      if (option.id == optionId) {
        return option.copyWith(votesCount: option.votesCount + 1);
      }
      return option;
    }).toList();
    final updatedVote = vote.copyWith(options: updatedOptions);
    await updateVote(updatedVote);
  }
}
