import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/domain/eventtype_enum.dart';
import 'package:golocal/src/event/manage/bloc/manage_event_bloc.dart';
import 'package:image_picker/image_picker.dart';

class EventCreatePage extends StatefulWidget {
  const EventCreatePage({super.key, this.event});
  final Event? event;

  @override
  State<EventCreatePage> createState() => _EventCreatePageState();
}

class _EventCreatePageState extends State<EventCreatePage> {
  final _formKey = GlobalKey<FormState>();
  final _tagsFieldController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => ManageEventBloc(context.read<IEventsRepository>(),
          event: widget.event),
      child: Scaffold(
        appBar: AppBar(
          centerTitle: true,
          title: Text(widget.event == null ? 'Create Event' : 'Edit Event'),
        ),
        body: SingleChildScrollView(
          child: Form(
            key: _formKey,
            child: Padding(
              padding: const EdgeInsets.all(8.0),
              child: BlocBuilder<ManageEventBloc, ManageEventState>(
                builder: (context, state) {
                  return Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      imageSection(state, context),
                      SizedBox(height: 8),
                      titleSection(state, context),
                      SizedBox(height: 8),
                      descriptionSection(state, context),
                      SizedBox(height: 8),
                      datesSection(state, context),
                      SizedBox(height: 8),
                      eventTypeSection(state, context),
                      SizedBox(height: 8),
                      tagsSection(state, context),
                      SizedBox(height: 8),
                      adultsOnlySection(state, context),
                      SizedBox(height: 8),
                      OutlinedButton(
                          onPressed: () {
                            if (_formKey.currentState!.validate()) {
                              context.read<ManageEventBloc>().add(SaveEvent());
                            }
                          },
                          child: Text("Save"))
                    ],
                  );
                },
              ),
            ),
          ),
        ),
      ),
    );
  }

  Widget imageSection(ManageEventState state, BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          "Event Image",
          style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
        ),
        SizedBox(height: 8),
        GestureDetector(
          onTap: () async {
            context.read<ManageEventBloc>().add(UpdateImage());
          },
          child: Container(
            height: 150,
            width: double.infinity,
            decoration: BoxDecoration(
              border: Border.all(color: Colors.grey),
              borderRadius: BorderRadius.circular(8),
            ),
            child: state.image == null
                ? Center(child: Text("Tap to select an image"))
                : Image.file(
                    state.image!,
                    fit: BoxFit.cover,
                  ),
          ),
        ),
      ],
    );
  }

  Column tagsSection(ManageEventState state, BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          "Help other people find your event easier",
          style: TextStyle(fontSize: 16),
        ),
        Text('Create tags which describe your event'),
        SizedBox(height: 8),
        Row(
          children: [
            Expanded(
              child: TextField(
                  controller: _tagsFieldController,
                  decoration: InputDecoration(
                    labelText: 'Tag',
                    border: OutlineInputBorder(),
                  ),
                  onSubmitted: (value) {
                    context.read<ManageEventBloc>().add(UpdateTags(value));
                    _tagsFieldController.clear();
                  }),
            ),
            SizedBox(width: 8),
            ElevatedButton(
              onPressed: () {
                context
                    .read<ManageEventBloc>()
                    .add(UpdateTags(_tagsFieldController.text));
                _tagsFieldController.clear();
              },
              style: ElevatedButton.styleFrom(
                padding: EdgeInsets.all(16),
                shape: CircleBorder(),
              ),
              child: Icon(Icons.add),
            ),
          ],
        ),
        SizedBox(height: 8),
        Wrap(
          spacing: 8,
          children: [
            for (var tag in state.tags)
              Chip(
                label: Text(tag),
                onDeleted: () {
                  context
                      .read<ManageEventBloc>()
                      .add(UpdateTags(tag, remove: true));
                },
              ),
          ],
        ),
      ],
    );
  }

  Column organizersSection(ManageEventState state, BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          "Tell us more about the organizers",
          style: TextStyle(fontSize: 16),
        ),
        SizedBox(height: 8),
        Row(
          children: [
            Expanded(
              child: TextField(
                  controller: _tagsFieldController,
                  decoration: InputDecoration(
                    labelText: 'Organizer',
                    border: OutlineInputBorder(),
                  ),
                  onSubmitted: (value) {
                    context
                        .read<ManageEventBloc>()
                        .add(UpdateOrganizers(value));
                    _tagsFieldController.clear();
                  }),
            ),
            SizedBox(width: 8),
            ElevatedButton(
              onPressed: () {
                context
                    .read<ManageEventBloc>()
                    .add(UpdateTags(_tagsFieldController.text));
                _tagsFieldController.clear();
              },
              style: ElevatedButton.styleFrom(
                padding: EdgeInsets.all(16),
                shape: CircleBorder(),
              ),
              child: Icon(Icons.add),
            ),
          ],
        ),
        SizedBox(height: 8),
        Wrap(
          spacing: 8,
          children: [
            for (var organizer in state.organizers)
              Chip(
                label: Text(organizer),
                onDeleted: () {
                  context
                      .read<ManageEventBloc>()
                      .add(UpdateOrganizers(organizer));
                },
              ),
          ],
        ),
      ],
    );
  }

  TextFormField descriptionSection(
      ManageEventState state, BuildContext context) {
    return TextFormField(
      validator: (value) {
        if (value == null) return "Please enter description";
        if (value.isEmpty) return "Please enter description";
        return null;
      },
      maxLines: 3,
      initialValue: state.description,
      decoration: const InputDecoration(
        border: OutlineInputBorder(),
        labelText: 'Description',
      ),
      onChanged: (value) {
        context.read<ManageEventBloc>().add(UpdateDescription(value));
      },
    );
  }

  TextFormField titleSection(ManageEventState state, BuildContext context) {
    return TextFormField(
      validator: (value) {
        if (value == null) return "Please enter title";
        if (value.isEmpty) return "Please enter title";
        if (value.length < 3) {
          return "Title must be at least 3 characters long";
        }
        return null;
      },
      maxLines: 1,
      initialValue: state.title,
      decoration: const InputDecoration(
        border: OutlineInputBorder(),
        labelText: 'Title',
      ),
      onChanged: (value) {
        context.read<ManageEventBloc>().add(UpdateTitle(value));
      },
    );
  }

  Widget adultsOnlySection(ManageEventState state, BuildContext context) {
    return Row(
      children: [
        Text("Adults only"),
        Switch(
          value: state.isAdultsOnly,
          onChanged: (value) {
            context.read<ManageEventBloc>().add(UpdateIsAdultsOnly(value));
          },
        ),
      ],
    );
  }

  Widget datesSection(ManageEventState state, BuildContext context) {
    return Row(
      children: [
        Expanded(
          child: FormField<DateTime?>(
            initialValue: state.startDate,
            autovalidateMode: AutovalidateMode.onUserInteraction,
            validator: (value) {
              if (value == null) {
                return "Provide date";
              }
              return null;
            },
            builder: (field) => ListTile(
              title: Text('Starts at'),
              subtitle: field.errorText != null
                  ? Text(field.errorText!, style: TextStyle(color: Colors.red))
                  : Text(
                      field.value != null
                          ? (field.value!.toFormattedString())
                          : 'Select start date',
                    ),
              trailing: const Icon(Icons.calendar_today),
              onTap: () async {
                DateTime? date = await selectDateTime(
                  context: context,
                  initialDate: state.startDate ?? DateTime.now(),
                );

                field.didChange(date);

                if (date != null && context.mounted) {
                  context.read<ManageEventBloc>().add(UpdateStartDate(date));
                }
              },
            ),
          ),
        ),
        Expanded(
          child: FormField<DateTime?>(
            validator: (value) {
              if (value == null) return null;
              if (state.startDate == null) return null;
              if (value.isBefore(state.startDate!)) {
                return "Must be after start date";
              }
              return null;
            },
            initialValue: state.endDate,
            autovalidateMode: AutovalidateMode.onUserInteraction,
            builder: (field) => ListTile(
              title: Text('Ends at'),
              subtitle: field.errorText != null
                  ? Text(field.errorText!)
                  : Text(
                      state.endDate != null
                          ? (state.endDate!.toFormattedString())
                          : 'Select end date',
                    ),
              trailing: const Icon(Icons.calendar_today),
              onTap: () async {
                DateTime? date = await selectDateTime(
                  context: context,
                  initialDate:
                      state.endDate ?? state.startDate ?? DateTime.now(),
                  firstDate: state.startDate ?? DateTime.now(),
                );
                if (date != null && context.mounted) {
                  context.read<ManageEventBloc>().add(UpdateEndDate(date));
                }
              },
            ),
          ),
        ),
      ],
    );
  }

  Widget eventTypeSection(ManageEventState state, BuildContext context) {
    return FormField<EventType?>(
      validator: (value) {
        if (value == null) {
          return "Choose event type";
        }
        return null;
      },
      autovalidateMode: AutovalidateMode.onUserInteraction,
      builder: (field) => Column(
        children: [
          DropdownMenu<EventType>(
            initialSelection: state.type,
            dropdownMenuEntries: EventType.values
                .map((e) => DropdownMenuEntry(label: e.name, value: e))
                .toList(),
            inputDecorationTheme: InputDecorationTheme(
              border: OutlineInputBorder(),
            ),
            label: field.errorText == null
                ? Text("Event type")
                : Text(
                    field.errorText!,
                    style: TextStyle(color: Colors.red),
                  ),
            onSelected: (value) {
              field.didChange(value);
              if (value != null) {
                context.read<ManageEventBloc>().add(UpdateType(value));
              }
            },
            expandedInsets: EdgeInsets.all(0),
          ),
        ],
      ),
    );
  }

  Future<DateTime?> selectDateTime({
    required BuildContext context,
    required DateTime? initialDate,
    DateTime? firstDate,
  }) async {
    firstDate ??= DateTime.now();
    initialDate ??= firstDate;
    DateTime lastDate = DateTime.now().add(Duration(days: 365 * 10));
    DateTime? date = await showDatePicker(
      context: context,
      initialDate: initialDate,
      firstDate: firstDate,
      lastDate: lastDate,
    );
    if (date == null || !context.mounted) return null;
    TimeOfDay? time = await showTimePicker(
      context: context,
      initialTime: TimeOfDay.fromDateTime(initialDate),
    );
    if (time == null || !context.mounted) return null;
    return DateTime.utc(
      date.year,
      date.month,
      date.day,
      time.hour,
      time.minute,
    );
  }
}

extension on DateTime {
  String toFormattedString() {
    return '$day/$month/$year $hour:$minute';
  }
}
