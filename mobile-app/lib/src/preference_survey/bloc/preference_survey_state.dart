part of 'preference_survey_bloc.dart';

abstract class PreferenceSurveyState {}

class PreferenceSurveyLoading extends PreferenceSurveyState {}

class PreferenceSurveyLoaded extends PreferenceSurveyState {
  final List<String> questions;

  PreferenceSurveyLoaded(this.questions);
}
