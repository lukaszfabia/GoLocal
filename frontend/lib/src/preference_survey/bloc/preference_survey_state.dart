part of 'preference_survey_bloc.dart';

/// Base state for the preference survey.
@immutable
sealed class PreferenceSurveyState {}

/// State indicating that the preference survey is loading.
class PreferenceSurveyLoading extends PreferenceSurveyState {}

/// State indicating that an error occurred while loading or submitting the survey.
class PreferenceSurveyError extends PreferenceSurveyState {
  /// The error message.
  final String message;

  /// Creates a [PreferenceSurveyError] with the given error [message].
  PreferenceSurveyError(this.message);
}

/// State indicating that the preference survey has been loaded successfully.
class PreferenceSurveyLoaded extends PreferenceSurveyState {
  /// The loaded preference survey.
  final PreferenceSurvey survey;

  /// Creates a [PreferenceSurveyLoaded] with the given [survey].
  PreferenceSurveyLoaded(this.survey);
}

/// State indicating that the preference survey has been submitted successfully.
class PreferenceSurveySubmitted extends PreferenceSurveyState {}
