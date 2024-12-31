import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:meta/meta.dart';
import 'package:golocal/src/preference_survey/services/preference_survey_service.dart';
import 'package:golocal/src/preference_survey/preference_survey_model.dart';

part 'preference_survey_event.dart';
part 'preference_survey_state.dart';

class PreferenceSurveyBloc
    extends Bloc<PreferenceSurveyEvent, PreferenceSurveyState> {
  final PreferenceSurveyService _recommendationService;

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
  }
}
