part of 'preference_survey_bloc.dart';

@immutable
sealed class PreferenceSurveyEvent {}

class LoadPreferenceSurvey extends PreferenceSurveyEvent {}

class SubmitPreferenceSurvey extends PreferenceSurveyEvent {
  final int surveyId;
  final Map<int, List<int>> answers;
  SubmitPreferenceSurvey({required this.surveyId, required this.answers});
}
