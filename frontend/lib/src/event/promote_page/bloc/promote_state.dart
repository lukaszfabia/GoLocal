part of 'promote_bloc.dart';

/// Represents the different statuses that the promotion process can have.
enum PromoteStatus {
  initial, // Initial state before any action is taken.
  success, // State when the promotion is successful.
  error, // State when there is an error in the promotion process.
  loading, // State when the promotion process is in progress.
  packChosen // State when a promotion pack is chosen.
}

/// Immutable state class for the promotion process.
@immutable
class PromoteState extends Equatable {
  /// The current status of the promotion process.
  final PromoteStatus status;

  /// An optional message that provides additional information about the current state.
  final String? message;

  /// An optional promotion pack that has been chosen.
  final PromotePack? pack;

  /// Creates a new instance of [PromoteState].
  ///
  /// The [status] defaults to [PromoteStatus.initial] if not provided.
  const PromoteState({
    this.status = PromoteStatus.initial,
    this.message,
    this.pack,
  });

  @override
  List<Object?> get props => [status, message, pack];

  /// Creates a copy of the current [PromoteState] with updated values.
  ///
  /// If a value is not provided, the current value is used.
  PromoteState copyWith({
    PromoteStatus? status,
    String? message,
    PromotePack? pack,
  }) {
    return PromoteState(
      status: status ?? this.status,
      message: message ?? this.message,
      pack: pack ?? this.pack,
    );
  }
}
