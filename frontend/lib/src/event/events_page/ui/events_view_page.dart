import 'dart:math';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/event/events_page/bloc/events_bloc.dart';
import 'package:golocal/src/event/data/impl/events_repository_impl.dart';
import 'package:golocal/src/event/data/impl/recommended_repository_impl.dart';
import 'package:golocal/src/event/events_page/ui/event_card.dart';
import 'package:golocal/src/routing/router.dart';

class EventsViewPage extends StatefulWidget {
  const EventsViewPage({super.key});

  @override
  State<EventsViewPage> createState() => _EventsViewPageState();
}

class _EventsViewPageState extends State<EventsViewPage> {
  final ScrollController _scrollController = ScrollController();
  bool _showSearchBar = false;
  bool _showFilterBar = false;
  bool _isScrollingDown = false;

  @override
  void initState() {
    super.initState();
    _scrollController.addListener(_handleScroll);
    print(_showSearchBar);
  }

  @override
  void dispose() {
    _scrollController.dispose();
    super.dispose();
  }

  void _handleScroll() {
    if (_scrollController.position.userScrollDirection ==
        ScrollDirection.reverse) {
      if (!_isScrollingDown) {
        _isScrollingDown = true;
        _toggleBarsOff();
      }
    } else if (_scrollController.position.userScrollDirection ==
        ScrollDirection.forward) {
      if (_isScrollingDown) {
        _isScrollingDown = false;
        _toggleSearchBar(true);
      }
    }
  }

  void _toggleSearchBar(bool show) {
    _toggleBarsOff();

    setState(() {
      _showSearchBar = show;
    });
  }

  void _toggleFilterBar(bool show) {
    _toggleBarsOff();

    setState(() {
      _showFilterBar = show;
    });
  }

  void _toggleBarsOff() {
    setState(() {
      _showSearchBar = false;
      _showFilterBar = false;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        centerTitle: true,
        title: const Text('Events'),
        leading: IconButton(
          icon: const Icon(Icons.add),
          onPressed: () => context.push('${AppRoute.events.path}/create'),
        ),
        actions: [
          IconButton(
            icon: const Icon(Icons.filter_list),
            onPressed: () => _toggleFilterBar(true),
          ),
          // IconButton(
          //   icon: const Icon(Icons.search),
          //   onPressed: () => _toggleSearchBar(true),
          // ),
        ],
      ),
      body: Column(
        children: [
          // SearchBar(showSearchBar: _showSearchBar),
          FilterBar(showFilterBar: _showFilterBar),
          Expanded(
            child: BlocConsumer<EventsBloc, EventsState>(
              listener: (context, state) {
                if (state.status == EventsStatus.error) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text(state.errorMessage!)),
                  );
                }
              },
              builder: (context, state) {
                double width = MediaQuery.of(context).size.width;
                int crossAxisCount = min(max(width ~/ 300, 1), 3);
                return GridView.builder(
                  controller: _scrollController,
                  gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                      crossAxisCount: crossAxisCount, childAspectRatio: 5 / 3),
                  itemCount: state.events.length + 1,
                  itemBuilder: (context, i) {
                    if (i < state.events.length) {
                      final event = state.events[i];
                      return GestureDetector(
                        child: EventCard(
                          event: event,
                        ),
                        onTap: () => context.push(
                            '${AppRoute.events.path}/${event.id}',
                            extra: event),
                      );
                    } else {
                      if (state.hasNext) {
                        BlocProvider.of<EventsBloc>(context)
                            .add(FetchEvents(refresh: false));
                      }
                      if (state.status == EventsStatus.loading) {
                        return const Center(child: CircularProgressIndicator());
                      }
                    }
                    return null;
                  },
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}

class SearchBar extends StatelessWidget {
  const SearchBar({
    super.key,
    required bool showSearchBar,
  }) : _showSearchBar = showSearchBar;

  final bool _showSearchBar;

  @override
  Widget build(BuildContext context) {
    return AnimatedSize(
      duration: const Duration(milliseconds: 200),
      curve: Curves.easeInOut,
      child: _showSearchBar
          ? Padding(
              padding:
                  const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0),
              child: TextField(
                decoration: InputDecoration(
                  hintText: 'Search events by name or location',
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(8.0),
                  ),
                ),
              ),
            )
          : const SizedBox.shrink(),
    );
  }
}

class FilterBar extends StatelessWidget {
  const FilterBar({
    super.key,
    required bool showFilterBar,
  }) : _showFilterBar = showFilterBar;

  final bool _showFilterBar;

  @override
  Widget build(BuildContext context) {
    return AnimatedSize(
      duration: const Duration(milliseconds: 200),
      curve: Curves.easeInOut,
      child: _showFilterBar
          ? Padding(
              padding:
                  const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: [
                  ElevatedButton(
                    onPressed: () {
                      final eventsRepository = EventsRepositoryImpl();
                      context
                          .read<EventsBloc>()
                          .add(SwitchRepository(eventsRepository));
                    },
                    child: const Text('Recent'),
                  ),
                  ElevatedButton(
                    onPressed: () {
                      final recommendedRepository = RecommendedRepositoryImpl();
                      context
                          .read<EventsBloc>()
                          .add(SwitchRepository(recommendedRepository));
                    },
                    child: const Text('Recommended'),
                  ),
                ],
              ),
            )
          : const SizedBox.shrink(),
    );
  }
}
