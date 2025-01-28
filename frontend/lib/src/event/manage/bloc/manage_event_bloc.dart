import 'dart:io';
import 'dart:math';

import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/domain/eventtype_enum.dart';
import 'package:golocal/src/event/domain/tag.dart';
import 'package:golocal/src/event/location/location.dart';
import 'package:golocal/src/user/domain/user.dart';
import 'package:image_picker/image_picker.dart';

part 'manage_event_event.dart';
part 'manage_event_state.dart';

class ManageEventBloc extends Bloc<ManageEventEvent, ManageEventState> {
  final IEventsRepository _repository;
  final Event? event;
  ManageEventBloc(this._repository, {this.event})
      : super(event == null
            ? ManageEventState(organizers: [], tags: [])
            : ManageEventState.copyFromEvent(event)) {
    on<UpdateTitle>((event, emit) {
      emit(state.copyWith(title: event.title));
    });
    on<UpdateDescription>((event, emit) {
      emit(state.copyWith(description: event.description));
    });
    on<UpdateType>((event, emit) {
      emit(state.copyWith(type: event.type));
    });
    on<UpdateTags>((event, emit) {
      var tags = List.of(state.tags);
      if (event.remove) {
        tags.removeWhere((element) => element == event.tag);
      } else {
        if (event.tag.isEmpty) return;
        if (tags.contains(event.tag)) return;
        tags.add(event.tag);
      }
      emit(state.copyWith(tags: tags));
    });
    on<UpdateStartDate>((event, emit) {
      if (state.endDate != null && event.startDate.isAfter(state.endDate!)) {
        emit(state.copyWith(
          startDate: event.startDate,
          endDate: event.startDate.add(const Duration(hours: 1)),
        ));
      } else {
        emit(state.copyWith(startDate: event.startDate));
      }
    });
    on<UpdateEndDate>((event, emit) {
      emit(state.copyWith(endDate: event.endDate));
    });
    on<UpdateIsAdultsOnly>((event, emit) {
      emit(state.copyWith(isAdultsOnly: event.isAdultsOnly));
    });
    on<UpdateOrganizers>((event, emit) {
      var organizers = List.of(state.organizers);
      organizers.add(event.organizers);
      emit(state.copyWith(organizers: organizers));
    });
    on<UpdateLocation>((event, emit) {
      emit(state.copyWith(location: event.location));
    });
    on<UpdateImage>((event, emit) async {
      final image = await ImagePicker().pickImage(source: ImageSource.gallery);
      if (image == null) return;
      emit(state.copyWith(image: File(image.path)));
    });
    on<SaveEvent>((event, emit) async {
      emit(state.copyWith(status: ManageEventStatus.loading));
      final dto = EventDTO(
        title: state.title,
        description: state.description,
        startDate: state.startDate!,
        endDate: state.endDate!,
        isAdultOnly: state.isAdultsOnly,
        organizers: state.organizers.map((e) => 1).toList(),
        tags: state.tags,
        image: state.image!,
        eventType: state.type!.name,
        lat: state.lat.toString(),
        lon: state.lon.toString(),
      );
      try {
        final created = await _repository.createEvent(dto);
        emit(state.copyWith(
          status: ManageEventStatus.success,
          message: 'Event saved successfully',
        ));
      } catch (e) {
        print(e);
        emit(state.copyWith(
          status: ManageEventStatus.error,
          message: e.toString(),
        ));
      }
    });
  }

  @override
  void onTransition(Transition<ManageEventEvent, ManageEventState> transition) {
    super.onTransition(transition);
  }
}
