class PreferenceSurveyQuestion {
  final int id;
  final String text;
  final QuestionType type;
  final List<Option>? options;
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
      type: _questionTypeFromString(json['Type']),
      options: json['Options'] != null
          ? (json['Options'] as List)
              .map((option) => Option.fromJson(option))
              .toList()
          : null,
      toggle: json['Toggle'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'ID': id,
      'Text': text,
      'Type': _questionTypeToString(type),
      'Options': options?.map((option) => option.toJson()).toList(),
      'Toggle': toggle,
    };
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

  static String _questionTypeToString(QuestionType type) {
    switch (type) {
      case QuestionType.toggle:
        return 'TOGGLE';
      case QuestionType.singleChoice:
        return 'SINGLE_CHOICE';
      case QuestionType.multiSelect:
        return 'MULTIPLE_CHOICE';
    }
  }
}

enum QuestionType { toggle, singleChoice, multiSelect }

class Option {
  final int id;
  final String text;
  final bool isSelected;

  Option({
    required this.id,
    required this.text,
    required this.isSelected,
  });

  factory Option.fromJson(Map<String, dynamic> json) {
    return Option(
      id: json['ID'],
      text: json['Text'],
      isSelected: json['IsSelected'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'ID': id,
      'Text': text,
      'IsSelected': isSelected,
    };
  }
}
