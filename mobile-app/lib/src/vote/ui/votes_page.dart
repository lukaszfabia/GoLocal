import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/vote/bloc/vote_bloc.dart';
import 'package:golocal/src/vote/data/votes_repository_dummy.dart';

class VotesPage extends StatelessWidget {
  const VotesPage({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => VoteBloc(DummyVotesRepository())..add(LoadVotes()),
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Votes'),
        ),
        body: BlocBuilder<VoteBloc, VoteState>(
          builder: (context, state) {
            if (state.status == VoteStatus.loading) {
              return const Center(child: CircularProgressIndicator());
            } else if (state.status == VoteStatus.error) {
              return Center(child: Text('Error: ${state.errorMessage}'));
            } else if (state.status == VoteStatus.loaded) {
              return ListView.builder(
                itemCount: state.votes.length,
                itemBuilder: (context, index) {
                  final vote = state.votes[index];
                  return ListTile(
                    title: Text(vote.text),
                    subtitle: Text(vote.event.title),
                    trailing: Text('${vote.options.length} options'),
                  );
                },
              );
            }
            return const SizedBox.shrink();
          },
        ),
      ),
    );
  }
}
