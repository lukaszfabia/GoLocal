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
