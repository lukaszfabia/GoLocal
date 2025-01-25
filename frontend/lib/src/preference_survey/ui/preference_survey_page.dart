import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_option.dart';
import '../bloc/preference_survey_bloc.dart';
import 'package:golocal/src/preference_survey/data/preference_survey_repository.dart';
import 'package:golocal/src/preference_survey/domain/preference_survey_question.dart';

class PreferenceSurveyPage extends StatelessWidget {
  const PreferenceSurveyPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text(
          'Preference survey',
          style: TextStyle(fontWeight: FontWeight.bold, fontSize: 24),
        ),
        centerTitle: true,
      ),
      body: BlocProvider(
        create: (context) => PreferenceSurveyBloc(PreferenceSurveyRepository())
          ..add(LoadPreferenceSurvey()),
        child: const PreferenceSurveyForm(),
      ),
    );
  }
}

class PreferenceSurveyForm extends StatefulWidget {
  const PreferenceSurveyForm({super.key});

  @override
  _PreferenceSurveyFormState createState() => _PreferenceSurveyFormState();
}

class _PreferenceSurveyFormState extends State<PreferenceSurveyForm> {
  final Map<int, List<int>> _answers = {};

  @override
  void initState() {
    super.initState();
  }

  void _submitSurvey(int surveyId) {
    context
        .read<PreferenceSurveyBloc>()
        .add(SubmitPreferenceSurvey(surveyId: surveyId, answers: _answers));
  }

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<PreferenceSurveyBloc, PreferenceSurveyState>(
      builder: (context, state) {
        if (state is PreferenceSurveyLoading) {
          return const Center(child: CircularProgressIndicator());
        } else if (state is PreferenceSurveyLoaded) {
          return SingleChildScrollView(
              scrollDirection: Axis.vertical,
              child: Padding(
                padding: const EdgeInsets.all(16.0),
                child: Column(
                  children: [
                    Text(
                      state.survey.description,
                      style: const TextStyle(fontSize: 18),
                    ),
                    ...state.survey.questions.asMap().entries.map((entry) {
                      int index = entry.value.id;
                      PreferenceSurveyQuestion question = entry.value;
                      return _buildQuestion(index, question);
                    }),
                    const SizedBox(height: 20),
                    ElevatedButton(
                      onPressed: () => _submitSurvey(state.survey.id),
                      child: const Text('Submit Survey'),
                    ),
                  ],
                ),
              ));
        } else if (state is PreferenceSurveySubmitted) {
          return const Center(child: Text('Survey submitted successfully!'));
        } else if (state is PreferenceSurveyError) {
          return Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                Text('Error: ${state.message}'),
                ElevatedButton(
                  onPressed: () => context
                      .read<PreferenceSurveyBloc>()
                      .add(LoadPreferenceSurvey()),
                  child: const Text('Retry'),
                )
              ],
            ),
          );
        } else {
          return const Center(child: Text('Unknown state'));
        }
      },
    );
  }

  Widget _buildQuestion(int index, PreferenceSurveyQuestion question) {
    switch (question.type) {
      case QuestionType.toggle:
        return _buildToggleQuestion(
            index, question.text, question.options![0].isSelected);
      case QuestionType.singleChoice:
        return _buildSingleSelectQuestion(
            index, question.text, question.options ?? [], 0);
      case QuestionType.multiSelect:
        return _buildMultiSelectQuestion(
            index, question.text, question.options ?? []);
    }
  }

  Widget _buildToggleQuestion(int index, String question, bool initialValue) {
    return const SizedBox.shrink();
    // return Column(
    //   crossAxisAlignment: CrossAxisAlignment.center,
    //   children: [
    //     Text(question, style: const TextStyle(fontSize: 16)),
    //     const SizedBox(height: 8),
    //     Switch(
    //       value: _answers[index] ?? initialValue,
    //       onChanged: (bool value) {
    //         setState(() {
    //           _answers[index] = value;
    //         });
    //       },
    //     ),
    //   ],
    // );
  }

  Widget _buildSingleSelectQuestion(
      int index, String question, List<Option> options, int selectedIndex) {
    if (_answers[index] == null) {
      _answers[index] = [selectedIndex];
    }

    return Column(
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        Text(question, style: const TextStyle(fontSize: 16)),
        const SizedBox(height: 8),
        Wrap(
          spacing: 8.0,
          children: options.asMap().entries.map((entry) {
            int idx = entry.value.id;
            String text = entry.value.text;
            return ChoiceChip(
              label: Text(text),
              selected: _answers[index]?[0] == idx,
              onSelected: (bool selected) {
                setState(() {
                  _answers[index] = [idx];
                });
              },
              selectedColor: Colors.blue,
              labelStyle: TextStyle(
                color: (_answers.isNotEmpty && _answers[index]![0] == idx)
                    ? Colors.white
                    : Colors.black,
              ),
            );
          }).toList(),
        )
      ],
    );
  }

  Widget _buildMultiSelectQuestion(
      int index, String question, List<Option> options) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        Text(question, style: const TextStyle(fontSize: 16)),
        const SizedBox(height: 8),
        Wrap(
          spacing: 8.0,
          children: options.asMap().entries.map((entry) {
            int idx = entry.value.id;
            String text = entry.value.text;
            return FilterChip(
              label: Text(text),
              selected: _answers[index]?.contains(idx) ?? false,
              onSelected: (bool selected) {
                setState(() {
                  if (selected) {
                    if (_answers[index] == null) {
                      _answers[index] = [];
                    }
                    _answers[index]!.add(idx);
                  } else {
                    _answers[index]?.remove(idx);
                  }
                });
              },
              selectedColor: Colors.blue,
              labelStyle: TextStyle(
                color: _answers[index]?.contains(idx) ?? false
                    ? Colors.white
                    : Colors.black,
              ),
            );
          }).toList(),
        )
      ],
    );
  }
}
