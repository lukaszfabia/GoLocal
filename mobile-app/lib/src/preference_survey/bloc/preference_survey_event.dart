part of 'preference_survey_bloc.dart';

@immutable
sealed class PreferenceSurveyEvent {}

class LoadPreferenceSurvey extends PreferenceSurveyEvent {}

class SubmitPreferenceSurvey extends PreferenceSurveyEvent {
  final List<String> answers;
  SubmitPreferenceSurvey({required this.answers});
}
