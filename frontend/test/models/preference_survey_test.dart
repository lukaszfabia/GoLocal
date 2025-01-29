import 'package:flutter_test/flutter_test.dart';
import 'package:golocal/src/survey/domain/preference_survey.dart';
import 'package:golocal/src/survey/domain/preference_survey_question.dart';

void main() {
  group('PreferenceSurvey', () {
    test('parses JSON creates a valid PreferenceSurvey object', () {
      final json = {
        'id': 1,
        'title': 'Survey title',
        'description': 'Survey description',
        'questions': [
          {
            'id': 1,
            'text': 'Question text',
            'type': 'TOGGLE',
            'options': [
              {
                'id': 1,
                'text': 'Option text',
              },
              {
                'id': 2,
                'text': 'Option text',
              }
            ],
          },
        ],
      };

      final survey = PreferenceSurvey.fromJson(json);

      expect(survey.id, 1);
      expect(survey.title, 'Survey title');
      expect(survey.description, 'Survey description');
      expect(survey.questions.length, 1);

      final question = survey.questions.first;
      expect(question.id, 1);
      expect(question.text, 'Question text');
      expect(question.type, QuestionType.toggle);
      expect(question.options!.length, 2);

      var id = 1;
      for (final option in question.options!) {
        expect(option.id, id++);
        expect(option.text, 'Option text');
      }
    });

    test('from JSON handles missing required fields', () {
      final json = {
        'id': 1,
        'title': null,
        'questions': [
          {
            'id': 1,
            'text': 'Question text',
            'type': 'TOGGLE',
            'options': [
              {
                'id': 1,
                'text': 'Option text',
              },
            ],
          },
        ],
      };

      expect(() => PreferenceSurvey.fromJson(json), throwsA(isA<TypeError>()));

      final json2 = {
        'id': 1,
        'title': 'Survey title',
        'questions': null,
      };

      expect(() => PreferenceSurvey.fromJson(json2), throwsA(isA<TypeError>()));

      final json3 = {
        'id': 1,
        'title': 'Survey title',
        'questions': [],
      };

      expect(() => PreferenceSurvey.fromJson(json3), throwsA(isA<TypeError>()));

      final json4 = {
        'id': 1,
        'title': 'Survey title',
        'questions': [
          {
            'id': 1,
            'text': 'Question text',
            'type': 'TOGGLE',
            'options': null,
          },
        ],
      };

      expect(() => PreferenceSurvey.fromJson(json4), throwsA(isA<TypeError>()));

      final json5 = {
        'id': 1,
        'title': 'Survey title',
        'questions': [
          {
            'id': 1,
            'text': 'Question text',
            'type': 'TOGGLE',
            'options': [
              {
                'id': 1,
                'text': null,
              },
            ],
          },
        ],
      };

      expect(() => PreferenceSurvey.fromJson(json5), throwsA(isA<TypeError>()));
    });
  });
}
