part of 'promote_bloc.dart';

@immutable
sealed class PromoteEvent extends Equatable {
  const PromoteEvent();

  @override
  List<Object> get props => [];
}

final class PromotePackChosenEvent extends PromoteEvent {
  final PromotePack pack;
  const PromotePackChosenEvent(this.pack);

  @override
  List<Object> get props => [pack];
}

final class PromoteRequestedEvent extends PromoteEvent {
  const PromoteRequestedEvent();
}
