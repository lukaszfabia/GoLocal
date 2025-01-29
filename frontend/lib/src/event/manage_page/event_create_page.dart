import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:geolocator/geolocator.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/domain/eventtype_enum.dart';
import 'package:golocal/src/event/manage_page/bloc/manage_event_bloc.dart';
import 'package:golocal/src/shared/dialog.dart';
import 'package:golocal/src/shared/position.dart';
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
              padding: const EdgeInsets.all(16.0),
              child: BlocConsumer<ManageEventBloc, ManageEventState>(
                listenWhen: (previous, current) =>
                    current.status != previous.status,
                listener: (context, state) {
                  if (state.status == ManageEventStatus.success) {
                    showMyDialog(context,
                        title: "Event created!",
                        message:
                            "Your event will be soon available to other users",
                        doublePop: true);
                  } else if (state.status == ManageEventStatus.error) {
                    showMyDialog(context,
                        title: "Error",
                        message: state.message ?? "An unknown error occurred");
                  }
                },
                builder: (context, state) {
                  return Column(
                    children: [
                      imageSection(state, context),
                      SizedBox(height: 16),
                      titleSection(state, context),
                      SizedBox(height: 16),
                      descriptionSection(state, context),
                      SizedBox(height: 16),
                      datesSection(state, context),
                      SizedBox(height: 16),
                      eventTypeSection(state, context),
                      SizedBox(height: 16),
                      tagsSection(state, context),
                      SizedBox(height: 16),
                      positionSection(state, context),
                      SizedBox(height: 16),
                      adultsOnlySection(state, context),
                      SizedBox(height: 16),
                      state.status == ManageEventStatus.loading
                          ? Center(child: CircularProgressIndicator())
                          : ElevatedButton(
                              onPressed: () {
                                if (_formKey.currentState!.validate()) {
                                  context
                                      .read<ManageEventBloc>()
                                      .add(SaveEvent());
                                }
                              },
                              style: ElevatedButton.styleFrom(
                                padding: EdgeInsets.symmetric(vertical: 16),
                                shape: RoundedRectangleBorder(
                                  borderRadius: BorderRadius.circular(8),
                                ),
                              ),
                              child:
                                  Text("Save", style: TextStyle(fontSize: 18)),
                            )
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
          style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
        ),
        SizedBox(height: 8),
        Text(
          "Add an image to represent your event. This will help attendees easily identify your event.",
          style: TextStyle(fontSize: 14, color: Colors.grey[600]),
        ),
        SizedBox(height: 12),
        GestureDetector(
          onTap: () async {
            context.read<ManageEventBloc>().add(UpdateImage());
          },
          child: Container(
            height: 200,
            width: double.infinity,
            decoration: BoxDecoration(
              border: Border.all(color: Colors.grey),
              borderRadius: BorderRadius.circular(12),
            ),
            child: state.image == null
                ? Center(child: Text("Tap to select an image"))
                : ClipRRect(
                    borderRadius: BorderRadius.circular(12),
                    child: Image.file(
                      state.image!,
                      fit: BoxFit.cover,
                    ),
                  ),
          ),
        ),
      ],
    );
  }

  TextFormField titleSection(ManageEventState state, BuildContext context) {
    return TextFormField(
      validator: (value) {
        if (value == null || value.isEmpty)
          return "Please enter a title for your event";
        if (value.length < 3) return "Title must be at least 3 characters long";
        return null;
      },
      maxLines: 1,
      initialValue: state.title,
      decoration: const InputDecoration(
        border: OutlineInputBorder(),
        labelText: 'Event Title',
        hintText: 'Enter the title of your event',
      ),
      onChanged: (value) {
        context.read<ManageEventBloc>().add(UpdateTitle(value));
      },
    );
  }

  TextFormField descriptionSection(
      ManageEventState state, BuildContext context) {
    return TextFormField(
      validator: (value) {
        if (value == null || value.isEmpty)
          return "Please provide a description of your event";
        return null;
      },
      maxLines: 3,
      initialValue: state.description,
      decoration: const InputDecoration(
        border: OutlineInputBorder(),
        labelText: 'Event Description',
        hintText: 'Describe your event in detail',
      ),
      onChanged: (value) {
        context.read<ManageEventBloc>().add(UpdateDescription(value));
      },
    );
  }

  Widget datesSection(ManageEventState state, BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          "Event Dates",
          style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
        ),
        SizedBox(height: 8),
        Text(
          "Select the start and end dates for your event. The start date represents when the event begins, and the end date marks when the event finishes.",
          style: TextStyle(fontSize: 14, color: Colors.grey[600]),
        ),
        SizedBox(height: 12),
        Row(
          children: [
            Expanded(
              child: FormField<DateTime?>(
                initialValue: state.startDate,
                autovalidateMode: AutovalidateMode.onUserInteraction,
                validator: (value) {
                  if (value == null) {
                    return "Please provide a start date for the event";
                  }
                  return null;
                },
                builder: (field) => ListTile(
                  title: Text('Starts on'),
                  subtitle: field.errorText != null
                      ? Text(field.errorText!,
                          style: TextStyle(color: Colors.red))
                      : Text(field.value != null
                          ? (field.value!.toFormattedString())
                          : 'Select start date'),
                  trailing: const Icon(Icons.calendar_today),
                  onTap: () async {
                    DateTime? date = await selectDateTime(
                      context: context,
                      initialDate: state.startDate ?? DateTime.now(),
                    );

                    field.didChange(date);

                    if (date != null && context.mounted) {
                      context
                          .read<ManageEventBloc>()
                          .add(UpdateStartDate(date));
                    }
                  },
                ),
              ),
            ),
            Expanded(
              child: FormField<DateTime?>(
                validator: (value) {
                  if (value == null) return "Please provide an end date";
                  if (state.startDate == null) return null;
                  if (value.isBefore(state.startDate!)) {
                    return "End date must be after the start date";
                  }
                  return null;
                },
                initialValue: state.endDate,
                autovalidateMode: AutovalidateMode.onUserInteraction,
                builder: (field) => ListTile(
                  title: Text('Ends on'),
                  subtitle: field.errorText != null
                      ? Text(field.errorText!,
                          style: TextStyle(color: Colors.red))
                      : Text(state.endDate != null
                          ? (state.endDate!.toFormattedString())
                          : 'Select end date'),
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
        ),
      ],
    );
  }

  Widget eventTypeSection(ManageEventState state, BuildContext context) {
    return FormField<EventType?>(
      validator: (value) {
        if (value == null) {
          return "Please select an event type to categorize your event";
        }
        return null;
      },
      autovalidateMode: AutovalidateMode.onUserInteraction,
      builder: (field) => Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            "Event Type",
            style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
          ),
          SizedBox(height: 8),
          Text(
            "Choose an event type to categorize your event. This helps users find similar events easily.",
            style: TextStyle(fontSize: 14, color: Colors.grey[600]),
          ),
          SizedBox(height: 12),
          Wrap(
            spacing: 8.0,
            runSpacing: 4.0,
            children: EventType.values.map((eventType) {
              bool isSelected = state.type == eventType;
              return ChoiceChip(
                label: Text(eventType.name),
                selected: state.type == eventType,
                onSelected: (isSelected) {
                  field.didChange(isSelected ? eventType : null);
                  if (isSelected) {
                    context.read<ManageEventBloc>().add(UpdateType(eventType));
                  }
                },
                selectedColor: Theme.of(context).primaryColor,
                backgroundColor: Colors.grey[200],
                labelStyle: TextStyle(
                  color: isSelected
                      ? Colors.white
                      : Theme.of(context).primaryColor,
                ),
              );
            }).toList(),
          ),
        ],
      ),
    );
  }

  Column tagsSection(ManageEventState state, BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          "Help others find your event easily",
          style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
        ),
        SizedBox(height: 4),
        Text(
          'Create tags that describe your event to make it more discoverable.',
          style: TextStyle(fontSize: 14, color: Colors.grey[600]),
        ),
        SizedBox(height: 12),
        Row(
          children: [
            Expanded(
              child: TextField(
                controller: _tagsFieldController,
                decoration: InputDecoration(
                  labelText: 'Tag',
                  hintText: 'Enter a tag',
                  border: OutlineInputBorder(),
                  contentPadding:
                      EdgeInsets.symmetric(vertical: 10, horizontal: 12),
                ),
                onSubmitted: (value) {
                  _addTag(value, context);
                },
              ),
            ),
            SizedBox(width: 8),
            OutlinedButton(
              onPressed: () {
                _addTag(_tagsFieldController.text.trim(), context);
              },
              style: OutlinedButton.styleFrom(
                shape: CircleBorder(),
                padding: EdgeInsets.all(14),
              ),
              child: Icon(Icons.add, color: Theme.of(context).primaryColor),
            ),
          ],
        ),
        SizedBox(height: 12),
        if (state.tags.isEmpty)
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 8.0),
            child: Text(
              'No tags added yet. Add some tags to make your event more discoverable.',
              style: TextStyle(color: Colors.grey[600]),
            ),
          ),
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
                deleteIconColor: Colors.red,
                backgroundColor: const Color.fromARGB(255, 255, 255, 255),
                deleteIcon: Icon(Icons.remove_circle_outline),
              ),
          ],
        ),
      ],
    );
  }

  Widget positionSection(ManageEventState state, BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          "Event Location",
          style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
        ),
        SizedBox(height: 8),
        Text(
          "Select the location for your event. This will help attendees know where to go.",
          style: TextStyle(fontSize: 14, color: Colors.grey[600]),
        ),
        SizedBox(height: 12),
        Row(
          children: [
            ElevatedButton(
              onPressed: () {
                context.read<ManageEventBloc>().add(UpdateLocation());
              },
              style: ElevatedButton.styleFrom(
                padding: EdgeInsets.all(16),
                shape: CircleBorder(),
              ),
              child: Icon(Icons.gps_fixed),
            ),
            SizedBox(width: 8),
            Text(
              state.lat == null || state.lon == null
                  ? "No location selected"
                  : "Lat: ${state.lat}, Lon: ${state.lon}",
              style: TextStyle(fontSize: 14, color: Colors.grey[600]),
            ),
          ],
        ),
      ],
    );
  }

  Widget adultsOnlySection(ManageEventState state, BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 12.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Row(
            children: [
              Icon(Icons.accessibility_new,
                  color: Theme.of(context).primaryColor),
              SizedBox(width: 8),
              Text(
                "Adults Only",
                style: TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.w500,
                  color: Colors.black,
                ),
              ),
            ],
          ),
          Switch(
            value: state.isAdultsOnly,
            onChanged: (value) {
              context.read<ManageEventBloc>().add(UpdateIsAdultsOnly(value));
            },
            activeColor: Theme.of(context).primaryColor,
            inactiveTrackColor: Colors.grey[300],
            inactiveThumbColor: Colors.grey,
          ),
        ],
      ),
    );
  }

  void _addTag(String tag, BuildContext context) {
    if (tag.isNotEmpty) {
      context.read<ManageEventBloc>().add(UpdateTags(tag));
      _tagsFieldController.clear();
    } else {
      ScaffoldMessenger.of(context)
          .showSnackBar(SnackBar(content: Text('Please enter a tag')));
    }
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
