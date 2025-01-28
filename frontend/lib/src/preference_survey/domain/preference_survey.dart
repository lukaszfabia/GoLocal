import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_question.dart';

/// A model representing a preference survey.
///
/// This class extends the [Model] class and contains information about a
/// preference survey, including its title, description, and a list of
/// questions.
///
/// The [PreferenceSurvey] class includes a factory constructor for creating
/// an instance from a JSON object.
///
/// Properties:
/// - `title` (String): The title of the survey.
/// - `description` (String): A brief description of the survey.
/// - `questions` (List<PreferenceSurveyQuestion>): A list of questions included in the survey.
///
/// Factory Constructors:
/// - `factory PreferenceSurvey.fromJson(Map<String, dynamic> json)`:
///   Creates a new instance of [PreferenceSurvey] from a JSON object.
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
      id: json['id'],
      questions: (json['questions'] as List)
          .map((question) => PreferenceSurveyQuestion.fromJson(question))
          .toList(),
      title: json['title'],
      description: json['description'],
    );
  }
}
