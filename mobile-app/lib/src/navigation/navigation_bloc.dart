import 'package:flutter_bloc/flutter_bloc.dart';

// States
abstract class NavigationState {}

class EventsPageState extends NavigationState {}

class ProfilePageState extends NavigationState {}

class PreferenceSurveyState extends NavigationState {}

// Events
abstract class NavigationEvent {}

class GoToEventsPage extends NavigationEvent {}

class GoToProfilePage extends NavigationEvent {}

class GoToPreferenceSurvey extends NavigationEvent {}

// Bloc
class NavigationBloc extends Bloc<NavigationEvent, NavigationState> {
  NavigationBloc() : super(EventsPageState()) {
    on<GoToEventsPage>((event, emit) => emit(EventsPageState()));
    on<GoToProfilePage>((event, emit) => emit(ProfilePageState()));
    on<GoToPreferenceSurvey>((event, emit) => emit(PreferenceSurveyState()));
  }
}
