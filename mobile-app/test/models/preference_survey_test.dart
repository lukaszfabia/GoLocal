import 'package:flutter_test/flutter_test.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_question.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_option.dart';

void main() {
  group('PreferenceSurvey', () {
    test('fromJson creates a valid PreferenceSurvey object', () {
      final json = {
        'ID': 1,
        'Title': 'Survey Title',
        'Description': 'Survey Description',
        'Questions': [
          {
            'ID': 1,
            'Text': 'Question 1',
            'Type': 'SINGLE_CHOICE',
            'Options': [
              {'ID': 1, 'Text': 'Option 1', 'IsSelected': false},
              {'ID': 2, 'Text': 'Option 2', 'IsSelected': true},
            ],
            'Toggle': null,
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
        'Title': 'Survey Title 2',
        'Description': 'Survey Description 2',
        'Questions': [],
      };

      final survey = PreferenceSurvey.fromJson(json);

      expect(survey.id, 2);
      expect(survey.title, 'Survey Title 2');
      expect(survey.description, 'Survey Description 2');
      expect(survey.questions.isEmpty, true);
    });
  });
}
