import 'package:golocal/src/shared/model_base.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_option.dart';

class PreferenceSurveyQuestion extends Model {
  final String text;
  final QuestionType type;
  final List<Option>? options;
  final bool? toggle;

  PreferenceSurveyQuestion({
    required super.id,
    required this.text,
    required this.type,
    this.options,
    this.toggle,
  });

  factory PreferenceSurveyQuestion.fromJson(Map<String, dynamic> json) {
    return PreferenceSurveyQuestion(
      id: json['ID'],
      text: json['Text'],
      type: _questionTypeFromString(json['Type']),
      options: json['Options'] != null
          ? (json['Options'] as List)
              .map((option) => Option.fromJson(option))
              .toList()
          : null,
      toggle: json['Toggle'],
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
