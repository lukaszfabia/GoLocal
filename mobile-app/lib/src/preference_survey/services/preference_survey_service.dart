import 'package:golocal/src/dio_client.dart';
import 'package:golocal/src/preference_survey/preference_survey_model.dart';

class PreferenceSurveyService {
  final DioClient _dioClient = DioClient();

  Future<void> submitSurvey(List<String> answers) async {
    try {
      final response = await _dioClient.dio.post('/survey', data: {
        'answers': answers,
      });
      print('Survey submitted: ${response.data}');
    } catch (e) {
      print('Error submitting survey: $e');
    }
  }

  Future<List<PreferenceSurveyQuestion>> fetchSurvey() async {
    try {
      final response = await _dioClient.dio.get('/preference-survey');
      final List<dynamic> data = response.data['data']['Questions'];
      return data
          .map((json) => PreferenceSurveyQuestion.fromJson(json))
          .toList();
    } catch (e) {
      print('Error fetching survey: $e');
      rethrow;
    }
  }
}
