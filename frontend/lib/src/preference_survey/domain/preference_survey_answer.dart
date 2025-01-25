class PreferenceSurveyAnswer {
  final int surveyId;
  final int questionId;
  final int userId;
  final bool? toggle;
  final int? optionId;
  final List<int>? options;

  PreferenceSurveyAnswer({
    required this.surveyId,
    required this.questionId,
    required this.userId,
    this.toggle,
    this.optionId,
    this.options,
  });

  factory PreferenceSurveyAnswer.factory(
      int preferenceSurveyId, int questionId, int userId, List<int> value) {
    final options = value;
    return PreferenceSurveyAnswer(
      surveyId: preferenceSurveyId,
      questionId: questionId,
      userId: userId,
      options: options,
    );
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
