import 'package:golocal/src/preference_survey/domain/preference_survey.dart';

abstract class IPreferenceSurveyRepository {
  Future<PreferenceSurvey> fetchSurvey();
  Future<void> submitSurvey(
      int preferenceSurveyId, Map<int, List<int>> answers);
}
