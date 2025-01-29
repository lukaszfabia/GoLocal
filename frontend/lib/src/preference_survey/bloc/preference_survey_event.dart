part of 'preference_survey_bloc.dart';

/// Base event for the preference survey.
@immutable
sealed class PreferenceSurveyEvent {}

/// Event to load the preference survey.
class LoadPreferenceSurvey extends PreferenceSurveyEvent {}

/// Event to submit the preference survey.
class SubmitPreferenceSurvey extends PreferenceSurveyEvent {
  /// The ID of the survey to be submitted.
  final int surveyId;

  /// The answers to the survey questions.
  final Map<int, List<int>> answers;

  /// Creates a [SubmitPreferenceSurvey] with the given [surveyId] and [answers].
  SubmitPreferenceSurvey({required this.surveyId, required this.answers});
}
