import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_question.dart';

class PreferenceSurvey extends Model {
  final String title;
  final String description;
  final List<PreferenceSurveyQuestion> questions;

  PreferenceSurvey({
    required super.id,
    this.title = '',
    this.description = '',
    required this.questions,
  });

  factory PreferenceSurvey.fromJson(Map<String, dynamic> json) {
    return PreferenceSurvey(
      id: json['ID'],
      questions: (json['questions'] as List)
          .map((question) => PreferenceSurveyQuestion.fromJson(question))
          .toList(),
      title: json['title'],
      description: json['description'],
    );
  }
}
