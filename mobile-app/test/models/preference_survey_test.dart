import 'package:flutter_test/flutter_test.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_question.dart';

void main() {
  group('PreferenceSurvey', () {
    test('fromJson creates a valid PreferenceSurvey object', () {
      final json = {
        'ID': 1,
        'title': 'Survey Title',
        'description': 'Survey Description',
        'questions': [
          {
            'ID': 1,
            'text': 'Question 1',
            'type': 'SINGLE_CHOICE',
            'options': [
              {'ID': 1, 'text': 'Option 1', 'isSelected': false},
              {'ID': 2, 'text': 'Option 2', 'isSelected': true},
            ],
            'toggle': null,
          },
        ],
      };

      final survey = PreferenceSurvey.fromJson(json);

      expect(survey.id, 1);
      expect(survey.title, 'Survey Title');
      expect(survey.description, 'Survey Description');
      expect(survey.questions.length, 1);

      final question = survey.questions.first;
      expect(question.id, 1);
      expect(question.text, 'Question 1');
      expect(question.type, QuestionType.singleChoice);
      expect(question.options?.length, 2);

      final option1 = question.options?.first;
      expect(option1?.id, 1);
      expect(option1?.text, 'Option 1');
      expect(option1?.isSelected, false);

      final option2 = question.options?.last;
      expect(option2?.id, 2);
      expect(option2?.text, 'Option 2');
      expect(option2?.isSelected, true);
    });

    test('fromJson handles missing optional fields', () {
      final json = {
        'ID': 2,
        'title': 'Survey Title 2',
        'description': 'Survey Description 2',
        'questions': [],
      };

      final survey = PreferenceSurvey.fromJson(json);

      expect(survey.id, 2);
      expect(survey.title, 'Survey Title 2');
      expect(survey.description, 'Survey Description 2');
      expect(survey.questions.isEmpty, true);
    });
  });
}
