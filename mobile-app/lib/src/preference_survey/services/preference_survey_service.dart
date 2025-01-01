import 'package:golocal/src/dio_client.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_answer.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey.dart';
import 'package:golocal/src/jwt_token_storage.dart';
import 'package:jwt_decoder/jwt_decoder.dart';

class PreferenceSurveyService {
  final DioClient _dioClient = DioClient();

  Future<void> submitSurvey(
      int preferenceSurveyId, Map<int, String> answers) async {
    try {
      print(answers);
      final accessToken = await TokenStorage().getAccessToken();
      final userId = int.parse(JwtDecoder.decode(accessToken!)['sub']);
      final List<PreferenceSurveyAnswer> answerList =
          answers.entries.map((entry) {
        final questionId = entry.key;
        final value = entry.value;

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
      }).toList();

      final response =
          await _dioClient.dio.post('/preference-survey/answer', data: {
        'answers': answerList.map((answer) => answer.toJson()).toList(),
      });
      print('Survey submitted: ${response.data}');
    } catch (e) {
      print('Error submitting survey: $e');
    }
  }

  Future<PreferenceSurvey> fetchSurvey() async {
    try {
      final response = await _dioClient.dio.get('/preference-survey');
      final Map<String, dynamic> data = response.data['data'];
      return PreferenceSurvey.fromJson(data);
    } catch (e) {
      print('Error fetching survey: $e');
      rethrow;
    }
  }
}
