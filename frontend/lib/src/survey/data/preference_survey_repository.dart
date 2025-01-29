import 'package:golocal/src/dio_client.dart';
import 'package:golocal/src/survey/domain/preference_survey_answer.dart';
import 'package:golocal/src/survey/domain/preference_survey.dart';
import 'package:golocal/src/jwt_token_storage.dart';
import 'package:jwt_decoder/jwt_decoder.dart';
import 'package:golocal/src/survey/data/ipreference_survey_repository.dart';

/// A repository class for handling preference survey operations.
///
/// This class extends [IPreferenceSurveyRepository] and provides methods to
/// submit survey answers and fetch survey data from a remote server.
class PreferenceSurveyRepository extends IPreferenceSurveyRepository {
  final DioClient _dioClient = DioClient();

  /// Submits the survey answers to the server.
  ///
  /// Takes a [preferenceSurveyId] and a map of [answers] where the key is the
  /// question ID and the value is a list of selected answer IDs. The method
  /// retrieves the user's access token, decodes it to get the user ID, and
  /// constructs a list of [PreferenceSurveyAnswer] objects. It then sends a
  /// POST request to the server with the answers.
  ///
  /// Throws an error if the submission fails.
  @override
  Future<void> submitSurvey(
      int preferenceSurveyId, Map<int, List<int>> answers) async {
    try {
      final accessToken = await TokenStorage().getAccessToken();
      final decodedToken = JwtDecoder.decode(accessToken!);
      final userId = int.parse(decodedToken['sub'].toString());
      final List<PreferenceSurveyAnswer> answerList =
          answers.entries.map((entry) {
        final questionId = entry.key;
        final values = entry.value;
        return PreferenceSurveyAnswer.factory(
            preferenceSurveyId, questionId, userId, values);
      }).toList();

      final data = {
        'answers': answerList.map((answer) => answer.toJson()).toList(),
      };

      await _dioClient.dio
          .post('/preference/preference-survey/answer', data: data);
    } catch (e) {
      print('Error submitting survey: $e');
    }
  }

  /// Fetches the preference survey data from the server.
  ///
  /// Sends a GET request to the server to retrieve the survey data. The method
  /// parses the response and returns a [PreferenceSurvey] object.
  ///
  /// Throws an error if the fetch operation fails.
  @override
  Future<PreferenceSurvey> fetchSurvey() async {
    try {
      final response =
          await _dioClient.dio.get('/preference/preference-survey');
      final Map<String, dynamic> data = response.data['data'];
      return PreferenceSurvey.fromJson(data);
    } catch (e) {
      print('Error fetching survey: $e');
      rethrow;
    }
  }
}
