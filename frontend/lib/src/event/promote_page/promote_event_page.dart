import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/event/data/ievents_repository.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/event/promote_page/bloc/promote_bloc.dart';
import 'package:golocal/src/event/promote_page/promote_pack.dart';
import 'package:golocal/src/shared/dialog.dart';

class PromoteEventPage extends StatelessWidget {
  const PromoteEventPage(this.event, {super.key});
  final Event event;

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => PromoteBloc(
        event: event,
        repository: context.read<IEventsRepository>(),
      ),
      child: BlocListener<PromoteBloc, PromoteState>(
        listenWhen: (previous, current) => current.status != previous.status,
        listener: (context, state) {
          if (state.status == PromoteStatus.success) {
            showMyDialog(context,
                title: "Success!",
                message: "Your event will be soon promoted!",
                doublePop: true);
          } else if (state.status == PromoteStatus.error) {
            showMyDialog(context,
                title: "Error", message: state.message ?? "An error occurred");
          }
        },
        child: Scaffold(
          appBar: AppBar(
            backgroundColor: const Color.fromARGB(149, 255, 255, 255),
            elevation: 0,
            title: Text(
              "Promote your event!",
              style: const TextStyle(
                fontWeight: FontWeight.bold,
                overflow: TextOverflow.ellipsis,
                color: Color.fromARGB(255, 0, 0, 0),
              ),
            ),
            centerTitle: true,
          ),
          body: SingleChildScrollView(
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                children: [
                  const SizedBox(height: 16),

                  Card(
                    elevation: 4.0,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(12),
                    ),
                    child: Padding(
                      padding: const EdgeInsets.all(16.0),
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.stretch,
                        children: [
                          const Text(
                            "How does the promotion work?",
                            style: TextStyle(
                              fontSize: 18,
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                          const SizedBox(height: 8),
                          Text(
                            "When you select a promotion pack, your event will be promoted for a period of time based on your selected pack. This will help you reach more people and get more attention for your event. Here are the details of each pack:",
                            style: TextStyle(fontSize: 14),
                          ),
                          const SizedBox(height: 16),
                        ],
                      ),
                    ),
                  ),

                  BlocBuilder<PromoteBloc, PromoteState>(
                    builder: (context, state) {
                      return GridView.builder(
                        shrinkWrap: true,
                        physics:
                            NeverScrollableScrollPhysics(), // Prevent scrolling inside the GridView
                        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                          crossAxisCount: 2,
                          crossAxisSpacing: 16.0,
                          mainAxisSpacing: 16.0,
                          childAspectRatio: 1, // Keeps the cards square-shaped
                        ),
                        itemCount: PromotePack.values.length,
                        itemBuilder: (context, index) {
                          final pack = PromotePack.values[index];
                          final isSelected = state.pack == pack;
                          return GestureDetector(
                            onTap: () {
                              if (isSelected) return;
                              context
                                  .read<PromoteBloc>()
                                  .add(PromotePackChosenEvent(pack));
                            },
                            child: Card(
                              elevation: 4.0,
                              shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(12),
                              ),
                              child: Container(
                                decoration: BoxDecoration(
                                  border: Border.all(
                                    color: isSelected
                                        ? Colors.amber
                                        : Colors.transparent,
                                    width: 2.0,
                                  ),
                                  borderRadius: BorderRadius.circular(12),
                                ),
                                child: Column(
                                  children: [
                                    Expanded(
                                      child: Center(
                                        child: Text(
                                          pack.name,
                                          style: const TextStyle(
                                            fontSize: 18,
                                            fontWeight: FontWeight.bold,
                                          ),
                                          textAlign: TextAlign.center,
                                        ),
                                      ),
                                    ),
                                    Text(
                                      "Duration: ${pack.duration} days",
                                      style: const TextStyle(fontSize: 14),
                                    ),
                                    const Divider(),
                                    Padding(
                                      padding:
                                          const EdgeInsets.only(bottom: 8.0),
                                      child: Text(
                                        "Cost: \$${pack.cost.toStringAsFixed(2)}",
                                        style: const TextStyle(
                                          fontSize: 14,
                                          fontWeight: FontWeight.bold,
                                        ),
                                      ),
                                    ),
                                  ],
                                ),
                              ),
                            ),
                          );
                        },
                      );
                    },
                  ), // Info Section

                  Card(
                    child: Padding(
                      padding: EdgeInsets.all(16),
                      child: BlocBuilder<PromoteBloc, PromoteState>(
                        builder: (context, state) {
                          if (state.pack == null) {
                            return const SizedBox.shrink();
                          }
                          final selectedPack = state.pack!;
                          return Column(
                            crossAxisAlignment: CrossAxisAlignment.stretch,
                            children: [
                              Text(
                                "You have selected: ${selectedPack.name}",
                                style: const TextStyle(
                                  fontSize: 16,
                                  fontWeight: FontWeight.bold,
                                ),
                              ),
                              const SizedBox(height: 8),
                              Text(
                                "Duration: ${selectedPack.duration} days",
                                style: const TextStyle(fontSize: 14),
                              ),
                              const SizedBox(height: 8),
                              Text(
                                "Price: \$${selectedPack.cost}",
                                style: const TextStyle(fontSize: 14),
                              ),
                              const SizedBox(height: 16),
                              ElevatedButton(
                                onPressed: () {
                                  context
                                      .read<PromoteBloc>()
                                      .add(PromoteRequestedEvent());
                                },
                                child: const Text("Promote"),
                              ),
                            ],
                          );
                        },
                      ),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
