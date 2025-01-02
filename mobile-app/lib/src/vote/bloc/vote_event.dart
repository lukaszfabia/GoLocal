part of 'vote_bloc.dart';

@immutable
sealed class VoteEvent {}

class LoadVotes extends VoteEvent {}

class CreateVote extends VoteEvent {
  final Vote vote;
  CreateVote(this.vote);
}

class UpdateVote extends VoteEvent {
  final Vote vote;
  UpdateVote(this.vote);
}

class DeleteVote extends VoteEvent {
  final String id;
  DeleteVote(this.id);
}
