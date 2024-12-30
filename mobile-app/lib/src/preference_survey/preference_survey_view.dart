import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'bloc/preference_survey_bloc.dart';

class PreferenceSurveyView extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Preference Survey'),
      ),
      body: BlocProvider(
        create: (context) => PreferenceSurveyBloc(),
        child: PreferenceSurveyForm(),
      ),
    );
  }
}

class PreferenceSurveyForm extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return BlocBuilder<PreferenceSurveyBloc, PreferenceSurveyState>(
      builder: (context, state) {
        if (state is PreferenceSurveyLoading) {
          return Center(child: CircularProgressIndicator());
        } else if (state is PreferenceSurveyLoaded) {
          return ListView.builder(
            itemCount: state.questions.length,
            itemBuilder: (context, index) {
              final question = state.questions[index];
              return ListTile(
                title: Text(question),
              );
            },
          );
        } else {
          return Center(child: Text('Failed to load survey'));
        }
      },
    );
  }
}
