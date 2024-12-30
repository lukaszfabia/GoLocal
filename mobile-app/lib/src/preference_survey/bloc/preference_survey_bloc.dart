import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:meta/meta.dart';
part 'preference_survey_event.dart';
part 'preference_survey_state.dart';

class PreferenceSurveyBloc
    extends Bloc<PreferenceSurveyEvent, PreferenceSurveyState> {
  PreferenceSurveyBloc() : super(PreferenceSurveyLoading()) {
    on<LoadPreferenceSurvey>((event, emit) async {
      final questions = [
        SurveyQuestion(
            question: 'Are you interested in adult-only activities?',
            type: QuestionType.toggle,
            initialValue: false),
        SurveyQuestion(
            question: 'Do you prefer to relax, or spend time actively?',
            type: QuestionType.singleChoice,
            options: ['High-energy', 'Relaxation']),
        SurveyQuestion(
            question:
                'What are your age/family constraints for events and activities?',
            type: QuestionType.singleChoice,
            options: ['Family-friendly', 'Couple-friendly', 'Adult-only']),
        SurveyQuestion(
            question:
                'Do you prefer indoors or outdoors events and activities?',
            type: QuestionType.singleChoice,
            options: ['Indoors', 'Outdoors']),
        SurveyQuestion(
            question: 'What more are you interested in?',
            type: QuestionType.multiSelect,
            options: ['Learning', 'Music', 'Sports']),
      ];
      emit(PreferenceSurveyLoaded(questions));
    });
  }
}
