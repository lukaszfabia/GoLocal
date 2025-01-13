import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/routing/router.dart';
import 'package:golocal/src/vote/bloc/vote_bloc.dart';
import 'package:golocal/src/vote/domain/vote.dart';
import 'package:golocal/src/vote/domain/vote_option.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/vote/data/ivotes_repository.dart';
import 'package:golocal/src/vote/data/votes_repository_event_filter.dart';
import 'package:golocal/src/event/domain/event.dart';
import 'package:golocal/src/vote/ui/vote_card.dart';

class VotesForEventPage extends StatelessWidget {
  final Event event;
  const VotesForEventPage(this.event, {super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => VoteBloc(
          VotesRepositoryEventFilter(event, context.read<IVotesRepository>()))
        ..add(LoadVotes()),
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Votes'),
          actions: [
            IconButton(
              icon: const Icon(Icons.filter_list),
              onPressed: () {},
            ),
          ],
        ),
        body: BlocBuilder<VoteBloc, VoteState>(
          builder: (context, state) {
            if (state.status == VoteStatus.loading) {
              return const Center(child: CircularProgressIndicator());
            } else if (state.status == VoteStatus.error) {
              return Center(child: Text('Error: ${state.errorMessage}'));
            } else if (state.status == VoteStatus.loaded) {
              return _buildVoteList(state.votes);
            }
            return const SizedBox.shrink();
          },
        ),
      ),
    );
  }

  Widget _buildVoteList(List<Vote> votes) {
    return ListView.builder(
      itemCount: votes.length,
      itemBuilder: (context, index) {
        final vote = votes[index];
        return VoteCard(vote: vote);
      },
    );
  }
}
