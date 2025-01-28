part of 'vote_bloc.dart';

/// Represents the different statuses a vote can have.
enum VoteStatus {
  /// Initial status before any action is taken.
  initial,

  /// Status when a vote is in the process of being loaded.
  loading,

  /// Status when a vote has been successfully loaded.
  loaded,

  /// Status when there is an error loading the vote.
  error
}

/// A state class that holds the current state of votes.
class VoteState extends Equatable {
  /// Creates a new instance of [VoteState].
  ///
  /// The [votes] parameter must not be null.
  /// The [status] parameter defaults to [VoteStatus.initial].
  /// The [errorMessage] parameter is optional.
  const VoteState({
    required this.votes,
    this.status = VoteStatus.initial,
    this.errorMessage,
  });

  /// A list of votes.
  final List<Vote> votes;

  /// The current status of the vote.
  final VoteStatus status;

  /// An optional error message if there is an error.
  final String? errorMessage;

  /// Creates a copy of the current [VoteState] with updated values.
  ///
  /// The [votes], [status], and [errorMessage] parameters are optional.
  /// If not provided, the current values will be used.
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
