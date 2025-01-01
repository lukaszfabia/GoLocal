import 'package:golocal/src/dio_client.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey.dart';

class PreferenceSurveyService {
  final DioClient _dioClient = DioClient();

  Future<void> submitSurvey(List<String> answers) async {
    try {
      final response =
          await _dioClient.dio.post('/preference-survey/answer', data: {
        'answers': answers,
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
