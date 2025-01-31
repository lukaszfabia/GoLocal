part of 'promote_bloc.dart';

/// Base class for all promote events, which are immutable and extend Equatable.
@immutable
sealed class PromoteEvent extends Equatable {
  const PromoteEvent();

  @override
  List<Object> get props => [];
}

/// Event triggered when a promote pack is chosen.
///
/// Contains the selected [PromotePack].
final class PromotePackChosenEvent extends PromoteEvent {
  final PromotePack pack;

  /// Creates a PromotePackChosenEvent with the given [pack].
  const PromotePackChosenEvent(this.pack);

  @override
  List<Object> get props => [pack];
}

/// Event triggered when a promote request is made.
final class PromoteRequestedEvent extends PromoteEvent {
  const PromoteRequestedEvent();
}
