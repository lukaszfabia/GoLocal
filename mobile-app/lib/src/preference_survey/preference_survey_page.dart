import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'bloc/preference_survey_bloc.dart';

class PreferenceSurveyPage extends StatelessWidget {
  const PreferenceSurveyPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () => Navigator.pop(context),
        ),
        title: const Text(
          'Preference survey',
          style: TextStyle(fontWeight: FontWeight.bold, fontSize: 24),
        ),
        centerTitle: true,
      ),
      body: BlocProvider(
        create: (context) =>
            PreferenceSurveyBloc()..add(LoadPreferenceSurvey()),
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
  final Map<int, dynamic> _answers = {};

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<PreferenceSurveyBloc, PreferenceSurveyState>(
      builder: (context, state) {
        if (state is PreferenceSurveyLoading) {
          return const Center(child: CircularProgressIndicator());
        } else if (state is PreferenceSurveyLoaded) {
          return SingleChildScrollView(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                ...state.questions.asMap().entries.map((entry) {
                  int index = entry.key;
                  SurveyQuestion question = entry.value;
                  return _buildQuestion(index, question);
                }),
                Center(
                  child: ElevatedButton(
                    onPressed: () {
                      // Handle survey submission
                    },
                    style: ElevatedButton.styleFrom(
                      backgroundColor: Colors.blue,
                      padding: const EdgeInsets.symmetric(
                          horizontal: 50, vertical: 12),
                    ),
                    child: const Text('Submit',
                        style: TextStyle(color: Colors.white)),
                  ),
                )
              ],
            ),
          );
        } else {
          return const Center(child: Text('Failed to load survey'));
        }
      },
    );
  }

  Widget _buildQuestion(int index, SurveyQuestion question) {
    switch (question.type) {
      case QuestionType.toggle:
        return _buildToggleQuestion(
            index, question.question, question.initialValue!);
      case QuestionType.singleChoice:
        return _buildButtonQuestion(
            index, question.question, question.options ?? [], 0);
      case QuestionType.multiSelect:
        return _buildMultiSelectQuestion(
            index, question.question, question.options!);
    }
  }

  Widget _buildToggleQuestion(int index, String question, bool initialValue) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(question, style: const TextStyle(fontSize: 16)),
        const SizedBox(height: 8),
        Switch(
          value: _answers[index] ?? initialValue,
          onChanged: (bool value) {
            setState(() {
              _answers[index] = value;
            });
          },
        ),
      ],
    );
  }

  Widget _buildButtonQuestion(
      int index, String question, List<String> options, int selectedIndex) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(question, style: const TextStyle(fontSize: 16)),
        const SizedBox(height: 8),
        Wrap(
          spacing: 8.0,
          children: options.asMap().entries.map((entry) {
            int idx = entry.key;
            String text = entry.value;
            return ChoiceChip(
              label: Text(text),
              selected: _answers[index] == idx,
              onSelected: (bool selected) {
                setState(() {
                  _answers[index] = idx;
                });
              },
              selectedColor: Colors.blue,
              labelStyle: TextStyle(
                color: _answers[index] == idx ? Colors.white : Colors.black,
              ),
            );
          }).toList(),
        )
      ],
    );
  }

  Widget _buildMultiSelectQuestion(
      int index, String question, List<String> options) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(question, style: const TextStyle(fontSize: 16)),
        const SizedBox(height: 8),
        Wrap(
          spacing: 8.0,
          children: options.asMap().entries.map((entry) {
            int idx = entry.key;
            String text = entry.value;
            return FilterChip(
              label: Text(text),
              selected: _answers[index]?.contains(idx) ?? false,
              onSelected: (bool selected) {
                setState(() {
                  if (selected) {
                    if (_answers[index] == null) {
                      _answers[index] = [];
                    }
                    _answers[index].add(idx);
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
