import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_option.dart';

/// A class representing a preference survey question.
///
/// This class extends the [Model] class and contains the following properties:
/// - `text`: The text of the survey question.
/// - `type`: The type of the survey question, represented by the [QuestionType] enum.
/// - `options`: An optional list of [Option] objects representing the possible answers to the question.
///
/// The class provides a factory constructor [PreferenceSurveyQuestion.fromJson] to create an instance
/// from a JSON object, and a private static method [_questionTypeFromString] to convert a string
/// representation of a question type to a [QuestionType] enum value.
///
/// The [QuestionType] enum defines the possible types of survey questions:
/// - `toggle`: A question with a toggle (yes/no) answer.
/// - `singleChoice`: A question with a single choice answer.
/// - `multiSelect`: A question with multiple selectable answers.
class PreferenceSurveyQuestion extends Model {
  final String text;
  final QuestionType type;
  final List<Option>? options;

  PreferenceSurveyQuestion({
    required super.id,
    required this.text,
    required this.type,
    this.options,
  });

  factory PreferenceSurveyQuestion.fromJson(Map<String, dynamic> json) {
    return PreferenceSurveyQuestion(
      id: json['id'],
      text: json['text'],
      type: _questionTypeFromString(json['type']),
      options: json['options'] != null
          ? (json['options'] as List)
              .map((option) => Option.fromJson(option))
              .toList()
          : null,
    );
  }

  static QuestionType _questionTypeFromString(String type) {
    switch (type) {
      case 'TOGGLE':
        return QuestionType.toggle;
      case 'SINGLE_CHOICE':
        return QuestionType.singleChoice;
      case 'MULTIPLE_CHOICE':
        return QuestionType.multiSelect;
      default:
        throw Exception('Unknown question type: $type');
    }
  }
}

enum QuestionType { toggle, singleChoice, multiSelect }
