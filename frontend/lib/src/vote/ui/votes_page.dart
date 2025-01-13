import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/event/location/location.dart';
import 'package:golocal/src/routing/router.dart';
import 'package:golocal/src/vote/bloc/vote_bloc.dart';
import 'package:golocal/src/vote/domain/vote.dart';
import 'package:golocal/src/vote/domain/vote_option.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/vote/data/ivotes_repository.dart';

class VotesPage extends StatelessWidget {
  const VotesPage({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) =>
          VoteBloc(context.read<IVotesRepository>())..add(LoadVotes()),
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
}
