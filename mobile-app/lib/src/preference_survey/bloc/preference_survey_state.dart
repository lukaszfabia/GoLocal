part of 'preference_survey_bloc.dart';

@immutable
sealed class PreferenceSurveyState {}

class PreferenceSurveyLoading extends PreferenceSurveyState {}

class PreferenceSurveyLoaded extends PreferenceSurveyState {
  final List<SurveyQuestion> questions;
  PreferenceSurveyLoaded(this.questions);
}

class SurveyQuestion {
  final String question;
  final QuestionType type;
  final List<String>? options;
  final bool? initialValue;

  SurveyQuestion({
    required this.question,
    required this.type,
    this.options,
    this.initialValue,
  });
}

enum QuestionType { toggle, singleChoice, multiSelect }
