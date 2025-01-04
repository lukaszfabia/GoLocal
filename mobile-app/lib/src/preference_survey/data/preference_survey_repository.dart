import 'package:golocal/src/dio_client.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_answer.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey.dart';
import 'package:golocal/src/jwt_token_storage.dart';
import 'package:jwt_decoder/jwt_decoder.dart';
import 'package:golocal/src/preference_survey/data/ipreference_survey_repository.dart';

class PreferenceSurveyRepository extends IPreferenceSurveyRepository {
  final DioClient _dioClient = DioClient();

  @override
  Future<void> submitSurvey(
      int preferenceSurveyId, Map<int, String> answers) async {
    try {
      final accessToken = await TokenStorage().getAccessToken();
      final decodedToken = JwtDecoder.decode(accessToken!);
      final userId = int.parse(decodedToken['sub'].toString());
      final List<PreferenceSurveyAnswer> answerList =
          answers.entries.map((entry) {
        final questionId = entry.key;
        final value = entry.value;
        return PreferenceSurveyAnswer.fromJson(
            preferenceSurveyId, questionId, userId, value);
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

  @override
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
