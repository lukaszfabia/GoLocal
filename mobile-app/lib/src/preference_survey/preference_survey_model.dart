import 'package:golocal/src/shared/model_base.dart';

class PreferenceSurvey extends Model {
  List<PreferenceSurveyQuestion> questions;
  PreferenceSurvey({
    required super.id,
    required this.questions,
  });

  factory PreferenceSurvey.fromJson(Map<String, dynamic> json) {
    return PreferenceSurvey(
      id: json['id'],
      questions: (json['questions'] as List)
          .map((question) => PreferenceSurveyQuestion.fromJson(question))
          .toList(),
    );
  }

  @override
  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'questions': questions.map((question) => question.toJson()).toList(),
    };
  }
}

class PreferenceSurveyQuestion {
  final int id;
  final String text;
  final String type;
  final List<String>? options;
  final bool? toggle;

  PreferenceSurveyQuestion({
    required this.id,
    required this.text,
    required this.type,
    this.options,
    this.toggle,
  });

  factory PreferenceSurveyQuestion.fromJson(Map<String, dynamic> json) {
    return PreferenceSurveyQuestion(
      id: json['ID'],
      text: json['Text'],
      type: json['Type'],
      options:
          json['Options'] != null ? List<String>.from(json['Options']) : null,
      toggle: json['Toggle'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'ID': id,
      'Text': text,
      'Type': type,
      'Options': options,
      'Toggle': toggle,
    };
  }
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
