part of 'vote_bloc.dart';

@immutable
sealed class PreferenceSurveyEvent {}

class LoadPreferenceSurvey extends PreferenceSurveyEvent {}

class SubmitPreferenceSurvey extends PreferenceSurveyEvent {
  final int surveyId;
  final Map<int, String> answers;
  SubmitPreferenceSurvey({required this.surveyId, required this.answers});
}
