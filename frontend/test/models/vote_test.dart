import 'package:flutter_test/flutter_test.dart';
import 'package:golocal/src/vote/domain/vote.dart';
import 'package:golocal/src/vote/domain/vote_in_voting.dart';

void main() {
  group('Voting', () {
    test('Parses VoteInVoting to json', () {
      final voteInVoting = VoteInVoting(voteId: 1, optionId: 1);

      final json = voteInVoting.toJson();

      expect(json['voteID'], 1);
      expect(json['optionID'], 1);
    });

    test('Parses JSON to Vote', () {
      final json = {
        'id': 1,
        'text': 'Vote title',
        'options': [
          {
            'id': 1,
            'text': 'Option 1',
          },
          {
            'id': 2,
            'text': 'Option 2',
          }
        ],
        'voteType': "CAN_CHANGE_VOTE",
      };

      final vote = Vote.fromJson(json);

      expect(vote.id, 1);
      expect(vote.text, 'Vote title');
      expect(vote.options.length, 2);

      var id = 1;
      for (final option in vote.options) {
        expect(option.id, id);
        expect(option.text, 'Option $id');
        id++;
      }
    });
  });
}
