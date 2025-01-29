part of 'promote_bloc.dart';

enum PromoteStatus { initial, success, error, loading, packChosen }

@immutable
class PromoteState extends Equatable {
  final PromoteStatus status;
  final String? message;
  final PromotePack? pack;
  const PromoteState({
    this.status = PromoteStatus.initial,
    this.message,
    this.pack,
  });

  @override
  List<Object?> get props => [status, message, pack];

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
