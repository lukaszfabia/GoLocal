import 'package:golocal/src/vote/domain/vote.dart';
import 'package:golocal/src/vote/domain/vote_option.dart';
import 'package:golocal/src/vote/data/ivotes_repository.dart';
import 'package:golocal/src/event/data/events_repository_dummy.dart';

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
    // Dummy implementation
    return Vote(
      id: 1,
      text: 'Dummy Vote',
      options: [
        VoteOption(id: 1, text: 'Option 1'),
        VoteOption(id: 2, text: 'Option 2'),
      ],
      event: (await EventsRepositoryDummy().getEvents())[0],
    );
  }

  @override
  Future<List<Vote>> getVotes() async {
    // Dummy implementation
    return [
      Vote(
        id: 1,
        text: 'Dummy Vote',
        options: [
          VoteOption(id: 1, text: 'Option 1'),
          VoteOption(id: 2, text: 'Option 2'),
        ],
        event: (await EventsRepositoryDummy().getEvents())[0],
      ),
      Vote(
        id: 2,
        text: 'Dummy Vote 2',
        options: [
          VoteOption(id: 3, text: 'Option 2.1'),
          VoteOption(id: 4, text: 'Option 2.2'),
        ],
        event: (await EventsRepositoryDummy().getEvents())[1],
      )
    ];
  }

  @override
  Future<void> updateVote(Vote vote) async {
    // Dummy implementation
  }
}
