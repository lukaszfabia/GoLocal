import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/vote/bloc/vote_bloc.dart';
import 'package:golocal/src/vote/data/votes_repository_dummy.dart';
import 'package:golocal/src/vote/domain/vote.dart';
import 'package:golocal/src/vote/domain/vote_option.dart';

class VotesPage extends StatelessWidget {
  const VotesPage({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => VoteBloc(DummyVotesRepository())..add(LoadVotes()),
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
        return _buildVoteCard(
          vote,
        );
      },
    );
  }

  Widget _buildVoteCard(Vote vote) {
    var optionCounts = vote.options.map((option) => option.votesCount).toList();
    var totalVotes = optionCounts.reduce((a, b) => a + b);

    return Card(
      margin: const EdgeInsets.symmetric(horizontal: 12, vertical: 8),
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          ListTile(
            leading: CircleAvatar(
              backgroundImage: vote.event.imageUrl == null
                  ? AssetImage('assets/image_not_found.png')
                  : NetworkImage(vote.event.imageUrl!),
            ),
            title: Text(vote.event.title,
                style: TextStyle(fontWeight: FontWeight.bold)),
            subtitle: Text(vote.event.location?.address?.toString() ??
                'No address available'),
            trailing: Icon(Icons.star_border),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(vote.text, style: TextStyle(fontWeight: FontWeight.bold)),
                const SizedBox(height: 8),
                ...vote.options
                    .map((option) => _buildOptionCard(option, totalVotes)),
              ],
            ),
          ),
          const SizedBox(height: 12),
        ],
      ),
    );
  }

  Widget _buildOptionCard(VoteOption option, int totalVotes) {
    return Card(
      margin: const EdgeInsets.symmetric(vertical: 4),
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
      child: ListTile(
        title: Text(option.text),
        subtitle: LinearProgressIndicator(
          value: option.votesCount / totalVotes,
          minHeight: 6,
          backgroundColor: Colors.grey[200],
          color: Colors.blue,
        ),
        trailing: Text(
            '${(option.votesCount / totalVotes * 100).toStringAsFixed(2)}%'),
      ),
    );
  }
}
