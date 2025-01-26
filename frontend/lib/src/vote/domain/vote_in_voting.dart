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
