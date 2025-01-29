part of 'manage_event_bloc.dart';

/// Base class for all events related to managing an event.
sealed class ManageEventEvent extends Equatable {
  const ManageEventEvent();

  @override
  List<Object> get props => [];
}

/// Event to save the current event.
final class SaveEvent extends ManageEventEvent {
  const SaveEvent();
}

/// Event to update the title of the event.
final class UpdateTitle extends ManageEventEvent {
  /// The new title of the event.
  final String title;

  const UpdateTitle(this.title);

  @override
  List<Object> get props => [title];
}

/// Event to update the description of the event.
final class UpdateDescription extends ManageEventEvent {
  /// The new description of the event.
  final String description;

  const UpdateDescription(this.description);

  @override
  List<Object> get props => [description];
}

/// Event to update the type of the event.
final class UpdateType extends ManageEventEvent {
  /// The new type of the event.
  final EventType type;

  const UpdateType(this.type);

  @override
  List<Object> get props => [type];
}

/// Event to update the tags of the event.
final class UpdateTags extends ManageEventEvent {
  /// The tag to be added or removed.
  final String tag;

  /// Whether the tag should be removed.
  final bool remove;

  const UpdateTags(this.tag, {this.remove = false});

  @override
  List<Object> get props => [tag, remove];
}

/// Event to update the start date of the event.
final class UpdateStartDate extends ManageEventEvent {
  /// The new start date of the event.
  final DateTime startDate;

  const UpdateStartDate(this.startDate);

  @override
  List<Object> get props => [startDate];
}

/// Event to update the end date of the event.
final class UpdateEndDate extends ManageEventEvent {
  /// The new end date of the event.
  final DateTime endDate;

  const UpdateEndDate(this.endDate);

  @override
  List<Object> get props => [endDate];
}

/// Event to update whether the event is for adults only.
final class UpdateIsAdultsOnly extends ManageEventEvent {
  /// Whether the event is for adults only.
  final bool isAdultsOnly;

  const UpdateIsAdultsOnly(this.isAdultsOnly);

  @override
  List<Object> get props => [isAdultsOnly];
}

/// Event to update the organizers of the event.
final class UpdateOrganizers extends ManageEventEvent {
  /// The new organizers of the event.
  final String organizers;

  const UpdateOrganizers(this.organizers);

  @override
  List<Object> get props => [organizers];
}

/// Event to update the location of the event.
final class UpdateLocation extends ManageEventEvent {
  /// The new location of the event.

  const UpdateLocation();

  @override
  List<Object> get props => [];
}

/// Event to update the image of the event.
final class UpdateImage extends ManageEventEvent {
  const UpdateImage();

  @override
  List<Object> get props => [];
}
