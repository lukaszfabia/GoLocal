/// A class representing a vote in a voting process.
///
/// This class contains the `voteId` and `optionId` which are required to
/// identify a vote and the option it is associated with.
///
/// The class provides a constructor for initializing these fields, a
/// named constructor `fromJson` for creating an instance from JSON data,
/// and a `toJson` method for converting an instance to JSON format.
class VoteInVoting {
  final int voteId;
  final int optionId;

  VoteInVoting({
    required this.voteId,
    required this.optionId,
  });

  VoteInVoting.fromJson(this.voteId, this.optionId);

  Map<String, dynamic> toJson() {
    return {
      'voteID': voteId,
      'optionID': optionId,
    };
  }
}
