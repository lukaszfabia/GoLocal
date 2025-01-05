class PreferenceSurveyAnswer {
  final int surveyId;
  final int questionId;
  final int userId;
  final bool? toggle;
  final int? optionId;
  final List<Map<String, int>>? options;

  PreferenceSurveyAnswer({
    required this.surveyId,
    required this.questionId,
    required this.userId,
    this.toggle,
    this.optionId,
    this.options,
  });

  factory PreferenceSurveyAnswer.factory(
      int preferenceSurveyId, int questionId, int userId, String value) {
    if (value == 'true' || value == 'false') {
      return PreferenceSurveyAnswer(
        surveyId: preferenceSurveyId,
        questionId: questionId,
        userId: userId,
        toggle: value == 'true',
      );
    } else if (value.contains(',')) {
      final options = value
          .split(',')
          .map((e) => {'OptionID': int.parse(e.trim())})
          .toList();
      return PreferenceSurveyAnswer(
        surveyId: preferenceSurveyId,
        questionId: questionId,
        userId: userId,
        options: options,
      );
    } else {
      return PreferenceSurveyAnswer(
        surveyId: preferenceSurveyId,
        questionId: questionId,
        userId: userId,
        optionId: int.tryParse(value),
      );
    }
  }

  Map<String, dynamic> toJson() {
    return {
      'SurveyID': surveyId,
      'QuestionID': questionId,
      'UserID': userId,
      'Toggle': toggle,
      'OptionID': optionId,
      'Options': options,
    };
  }
}
