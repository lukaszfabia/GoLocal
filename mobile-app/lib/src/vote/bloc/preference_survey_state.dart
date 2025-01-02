part of 'preference_survey_bloc.dart';

@immutable
sealed class PreferenceSurveyState {}

class PreferenceSurveyLoading extends PreferenceSurveyState {}

class PreferenceSurveyError extends PreferenceSurveyState {
  final String message;
  PreferenceSurveyError(this.message);
}

class PreferenceSurveyLoaded extends PreferenceSurveyState {
  final PreferenceSurvey survey;
  PreferenceSurveyLoaded(this.survey);
}

class PreferenceSurveySubmitted extends PreferenceSurveyState {}
