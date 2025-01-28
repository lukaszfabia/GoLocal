part of 'vote_bloc.dart';

/// Represents the base class for all vote-related events.
@immutable
sealed class VoteEvent {}

/// Event to load all votes.
class LoadVotes extends VoteEvent {}

/// Event to create a new vote.
///
/// Takes a [Vote] object as a parameter.
class CreateVote extends VoteEvent {
  final Vote vote;
  CreateVote(this.vote);
}

/// Event to update an existing vote.
///
/// Takes a [Vote] object as a parameter.
class UpdateVote extends VoteEvent {
  final Vote vote;
  UpdateVote(this.vote);
}

/// Event to delete a vote.
///
/// Takes the ID of the vote to be deleted as a parameter.
class DeleteVote extends VoteEvent {
  final String id;
  DeleteVote(this.id);
}

/// Event to vote on an option.
///
/// Takes the ID of the vote, the ID of the option, and the new value as parameters.
class VoteOnOption extends VoteEvent {
  final int voteId;
  final int optionId;
  final bool newValue;
  VoteOnOption(this.voteId, this.optionId, this.newValue);
}
