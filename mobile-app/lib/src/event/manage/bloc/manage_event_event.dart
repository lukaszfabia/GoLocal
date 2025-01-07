part of 'manage_event_bloc.dart';

sealed class ManageEventEvent extends Equatable {
  const ManageEventEvent();

  @override
  List<Object> get props => [];
}

final class SaveEvent extends ManageEventEvent {
  const SaveEvent();
}

final class UpdateTitle extends ManageEventEvent {
  const UpdateTitle(this.title);
  final String title;

  @override
  List<Object> get props => [title];
}

final class UpdateDescription extends ManageEventEvent {
  const UpdateDescription(this.description);
  final String description;

  @override
  List<Object> get props => [description];
}

final class UpdateType extends ManageEventEvent {
  const UpdateType(this.type);
  final EventType type;

  @override
  List<Object> get props => [type];
}

final class UpdateTags extends ManageEventEvent {
  const UpdateTags(this.tag, {this.remove = false});
  final String tag;
  final bool remove;

  @override
  List<Object> get props => [tag, remove];
}

final class UpdateStartDate extends ManageEventEvent {
  const UpdateStartDate(this.startDate);
  final DateTime startDate;

  @override
  List<Object> get props => [startDate];
}

final class UpdateEndDate extends ManageEventEvent {
  const UpdateEndDate(this.endDate);
  final DateTime endDate;

  @override
  List<Object> get props => [endDate];
}

final class UpdateIsAdultsOnly extends ManageEventEvent {
  const UpdateIsAdultsOnly(this.isAdultsOnly);
  final bool isAdultsOnly;

  @override
  List<Object> get props => [isAdultsOnly];
}

final class UpdateOrganizers extends ManageEventEvent {
  const UpdateOrganizers(this.organizers);
  final List<User> organizers;

  @override
  List<Object> get props => [organizers];
}

final class UpdateLocation extends ManageEventEvent {
  const UpdateLocation(this.location);
  final Location location;

  @override
  List<Object> get props => [location];
}
