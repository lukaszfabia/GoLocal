import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/report/bloc/report_event_bloc.dart';
import 'package:golocal/src/shared/dialog.dart';

enum ReportCategory {
  inappropriate("Inappropriate content"),
  spam("Spam"),
  illegal("Illegal content/activity"),
  harmful("Harmful content/activity"),
  vandalism("Vandalism"),
  other("Other");

  const ReportCategory(this.value);
  final String value;
}

class ReportEventPage extends StatelessWidget {
  ReportEventPage({super.key, required this.event});
  final Event event;
  final _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => ReportEventBloc(
        context.read<IEventsRepository>(),
        eventId: event.id,
      ),
      child: Scaffold(
        appBar: AppBar(
          centerTitle: true,
          title: Text('Report ${event.title}'),
        ),
        body: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              BlocConsumer<ReportEventBloc, ReportEventState>(
                listener: (context, state) {
                  if (state.status == ReportEventStatus.success) {
                    showMyDialog(context,
                        doublePop: true,
                        title: 'Report sent',
                        message:
                            'Your report will be revieved by our staff as soon as possible. We will take appropriate actions if needed. Thank you for contributing to a safe environment :)');
                  } else if (state.status == ReportEventStatus.error) {
                    showMyDialog(context,
                        message: state.message ?? 'An unknown error occured',
                        title: 'Error');
                  }
                },
                builder: (context, state) {
                  return Form(
                    key: _formKey,
                    child: Column(
                      children: [
                        Card(
                          elevation: 4,
                          child: Padding(
                            padding: const EdgeInsets.all(12.0),
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                Text(
                                  'Select a category',
                                  style: Theme.of(context).textTheme.labelLarge,
                                ),
                                const SizedBox(height: 8),
                                for (var category in ReportCategory.values)
                                  RadioListTile<String>(
                                    title: Text(category.value),
                                    value: category.value,
                                    groupValue: state.category,
                                    onChanged: (value) {
                                      context.read<ReportEventBloc>().add(
                                          UpdateCategory(category: value!));
                                    },
                                  ),
                              ],
                            ),
                          ),
                        ),
                        const SizedBox(height: 16),
                        Card(
                          elevation: 4,
                          child: Padding(
                            padding: const EdgeInsets.all(12.0),
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                Text(
                                  'Description',
                                  style: Theme.of(context).textTheme.labelLarge,
                                ),
                                const SizedBox(height: 8),
                                TextFormField(
                                  decoration: InputDecoration(
                                    border: OutlineInputBorder(),
                                  ),
                                  maxLines: 5,
                                  onChanged: (value) {
                                    /* 
                                    TODO: it may not be the best way, dont have time to change it now. it would be good  to have something like 
                                    debouncing while the user is typing because now the add method is called with every letter change
                                    */
                                    context.read<ReportEventBloc>().add(
                                        UpdateDescription(description: value));
                                  },
                                  validator: (value) {
                                    if (value == null || value.isEmpty) {
                                      return 'Please enter some information';
                                    }
                                    return null;
                                  },
                                ),
                              ],
                            ),
                          ),
                        ),
                        const SizedBox(height: 16),
                        Center(
                          child: ElevatedButton(
                            onPressed: () {
                              if (state.status == ReportEventStatus.sending) {
                                return;
                              }
                              if (_formKey.currentState!.validate()) {
                                _formKey.currentState!.save();
                                context
                                    .read<ReportEventBloc>()
                                    .add(SendReport());
                              }
                            },
                            style: ElevatedButton.styleFrom(
                              padding: const EdgeInsets.symmetric(
                                  horizontal: 30, vertical: 15),
                              shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(10),
                              ),
                              elevation: 5,
                            ),
                            child: Text('Submit Report'),
                          ),
                        ),
                      ],
                    ),
                  );
                },
              ),
            ],
          ),
        ),
      ),
    );
  }
}
