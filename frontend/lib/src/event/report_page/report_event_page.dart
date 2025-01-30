import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/domain/report_category_enum.dart';
import 'package:golocal/src/event/report_page/bloc/report_event_bloc.dart';
import 'package:golocal/src/shared/dialog.dart';
import 'dart:async';

class ReportEventPage extends StatefulWidget {
  final Event event;

  const ReportEventPage({super.key, required this.event});

  @override
  ReportEventPageState createState() => ReportEventPageState();
}

class ReportEventPageState extends State<ReportEventPage> {
  final _formKey = GlobalKey<FormState>();
  Timer? _debounce;

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => ReportEventBloc(
        context.read<IEventsRepository>(),
        eventId: widget.event.id,
      ),
      child: Scaffold(
        appBar: AppBar(
          centerTitle: true,
          title: Text('Report "${widget.event.title}"'),
        ),
        body: BlocConsumer<ReportEventBloc, ReportEventState>(
          listener: (context, state) {
            if (state.status == ReportEventStatus.success) {
              showMyDialog(
                context,
                doublePop: true,
                title: 'Report Sent',
                message:
                    'Your report has been submitted. Thank you for helping us maintain a safe environment.',
              );
            } else if (state.status == ReportEventStatus.error) {
              showMyDialog(
                context,
                title: 'Error',
                message: state.message ?? 'An unknown error occurred.',
              );
            }
          },
          builder: (context, state) {
            return SingleChildScrollView(
              padding: const EdgeInsets.all(16.0),
              child: Form(
                key: _formKey,
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.stretch,
                  children: [
                    _buildCategoryChips(context, state),
                    const SizedBox(height: 16),
                    _buildDescriptionField(context, state),
                    const SizedBox(height: 24),
                    _buildSubmitButton(context, state),
                  ],
                ),
              ),
            );
          },
        ),
      ),
    );
  }

  //TODO: Make this not expand when a chip is selected
  Widget _buildCategoryChips(BuildContext context, ReportEventState state) {
    return Card(
      elevation: 3,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12.0)),
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            Text('Select a reason',
                style: Theme.of(context).textTheme.labelLarge),
            const SizedBox(height: 8),
            Wrap(
              alignment: WrapAlignment.center,
              spacing: 8.0,
              children: ReportCategory.values.map((category) {
                bool isSelected = state.category == category.value;
                return ChoiceChip(
                  label: Text(category.value),
                  selected: isSelected,
                  selectedColor: Colors.redAccent.withValues(alpha: 0.8),
                  backgroundColor: Colors.grey.shade300,
                  labelStyle: TextStyle(
                    color: isSelected ? Colors.white : Colors.black87,
                    fontWeight:
                        isSelected ? FontWeight.bold : FontWeight.normal,
                  ),
                  onSelected: (selected) {
                    if (selected) {
                      context
                          .read<ReportEventBloc>()
                          .add(UpdateCategory(category: category.value));
                    }
                  },
                );
              }).toList(),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildDescriptionField(BuildContext context, ReportEventState state) {
    return Card(
      elevation: 3,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12.0)),
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text('Description', style: Theme.of(context).textTheme.labelLarge),
            const SizedBox(height: 8),
            TextFormField(
              decoration: const InputDecoration(
                border: OutlineInputBorder(),
                hintText: "Provide additional details...",
              ),
              maxLines: 5,
              initialValue: state.description,
              onChanged: (value) {
                if (_debounce?.isActive ?? false) _debounce!.cancel();
                _debounce = Timer(const Duration(milliseconds: 500), () {
                  context
                      .read<ReportEventBloc>()
                      .add(UpdateDescription(description: value));
                });
              },
              validator: (value) {
                if (value == null || value.trim().isEmpty) {
                  return 'Please enter some details';
                }
                return null;
              },
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildSubmitButton(BuildContext context, ReportEventState state) {
    if (state.description.isEmpty || state.category.isEmpty) {
      return const SizedBox.shrink();
    }
    return ElevatedButton(
      onPressed: state.status == ReportEventStatus.sending
          ? null
          : () {
              if (_formKey.currentState!.validate()) {
                _formKey.currentState!.save();
                context.read<ReportEventBloc>().add(SendReport());
              }
            },
      style: ElevatedButton.styleFrom(
        padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 15),
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(10)),
        elevation: 5,
      ),
      child: state.status == ReportEventStatus.sending
          ? const SizedBox(
              width: 24,
              height: 24,
              child: CircularProgressIndicator(
                  strokeWidth: 2, color: Colors.white))
          : const Text('Submit Report', style: TextStyle(fontSize: 16)),
    );
  }
}
