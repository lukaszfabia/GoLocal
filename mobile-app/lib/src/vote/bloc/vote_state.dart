part of 'vote_bloc.dart';

enum VoteStatus { initial, loading, loaded, error }

class VoteState extends Equatable {
  const VoteState({
    required this.votes,
    this.status = VoteStatus.initial,
    this.errorMessage,
  });

  final List<Vote> votes;
  final VoteStatus status;
  final String? errorMessage;

  VoteState copyWith({
    List<Vote>? votes,
    VoteStatus? status,
    String? errorMessage,
  }) {
    return VoteState(
      votes: votes ?? this.votes,
      status: status ?? this.status,
      errorMessage: errorMessage ?? this.errorMessage,
    );
  }

  @override
  List<Object?> get props => [votes, status, errorMessage];
}
