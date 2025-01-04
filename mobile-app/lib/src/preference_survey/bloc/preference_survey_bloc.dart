import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:meta/meta.dart';
import 'package:golocal/src/preference_survey/data/preference_survey_repository.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey.dart';

part 'preference_survey_event.dart';
part 'preference_survey_state.dart';

class PreferenceSurveyBloc
    extends Bloc<PreferenceSurveyEvent, PreferenceSurveyState> {
  final PreferenceSurveyRepository _recommendationService;

  PreferenceSurveyBloc(this._recommendationService)
      : super(PreferenceSurveyLoading()) {
    on<LoadPreferenceSurvey>((event, emit) async {
      try {
        final questions = await _recommendationService.fetchSurvey();
        emit(PreferenceSurveyLoaded(questions));
      } catch (e) {
        emit(PreferenceSurveyError('Failed to load survey'));
      }
    });

    on<SubmitPreferenceSurvey>((event, emit) async {
      try {
        await _recommendationService.submitSurvey(
            event.surveyId, event.answers);
        emit(PreferenceSurveySubmitted());
      } catch (e) {
        emit(PreferenceSurveyError('Failed to submit survey'));
      }
    });
  }
}
