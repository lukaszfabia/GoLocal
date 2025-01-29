import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:meta/meta.dart';
import 'package:golocal/src/survey/data/ipreference_survey_repository.dart';
import 'package:golocal/src/survey/domain/preference_survey.dart';

part 'preference_survey_event.dart';
part 'preference_survey_state.dart';

/// Bloc for managing the state of the preference survey.
class PreferenceSurveyBloc
    extends Bloc<PreferenceSurveyEvent, PreferenceSurveyState> {
  /// The repository for fetching and submitting the preference survey.
  final IPreferenceSurveyRepository _recommendationService;

  /// Creates a [PreferenceSurveyBloc] with the given [IPreferenceSurveyRepository].
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
