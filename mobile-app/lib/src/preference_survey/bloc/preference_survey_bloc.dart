import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:meta/meta.dart';
part 'preference_survey_event.dart';
part 'preference_survey_state.dart';

class PreferenceSurveyBloc
    extends Bloc<PreferenceSurveyEvent, PreferenceSurveyState> {
  PreferenceSurveyBloc() : super(PreferenceSurveyLoading()) {
    on<LoadPreferenceSurvey>((event, emit) {
      // Mock questions
      final questions = [
        'What is your favorite color?',
        'Do you prefer cats or dogs?',
        'What is your favorite season?',
      ];
      emit(PreferenceSurveyLoaded(questions));
    });
  }
}
