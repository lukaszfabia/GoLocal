import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_option.dart';

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
