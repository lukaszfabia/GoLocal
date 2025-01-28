/// A class representing an answer to a preference survey.
///
/// This class holds information about a user's response to a specific question
/// in a preference survey. It includes the survey ID, question ID, user ID,
/// and the user's selected options or toggle state.
///
/// The class provides a factory constructor for creating an instance from a
/// list of option IDs, and a method for converting the instance to a JSON
/// representation.
class PreferenceSurveyAnswer {
  /// The ID of the survey.
  final int surveyId;

  /// The ID of the question.
  final int questionId;

  /// The ID of the user.
  final int userId;

  /// The toggle state of the answer, if applicable.
  final bool? toggle;

  /// The ID of the selected option, if applicable.
  final int? optionId;

  /// The list of selected option IDs, if applicable.
  final List<int>? options;

  /// Creates a new instance of [PreferenceSurveyAnswer].
  ///
  /// The [surveyId], [questionId], and [userId] parameters are required.
  /// The [toggle], [optionId], and [options] parameters are optional.
  PreferenceSurveyAnswer({
    required this.surveyId,
    required this.questionId,
    required this.userId,
    this.toggle,
    this.optionId,
    this.options,
  });

  /// Factory constructor for creating a [PreferenceSurveyAnswer] from a list of option IDs.
  ///
  /// The [preferenceSurveyId], [questionId], and [userId] parameters are required.
  /// The [value] parameter is a list of option IDs.
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

  /// Converts the [PreferenceSurveyAnswer] instance to a JSON representation.
  ///
  /// Returns a map containing the survey ID, question ID, user ID, toggle state,
  /// option ID, and list of option IDs.
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
