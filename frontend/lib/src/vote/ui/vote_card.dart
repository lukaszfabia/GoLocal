import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:golocal/src/routing/router.dart';
import 'package:golocal/src/vote/bloc/vote_bloc.dart';
import 'package:golocal/src/vote/domain/vote.dart';
import 'package:go_router/go_router.dart';
import 'package:golocal/src/event/location/location.dart';
import 'package:golocal/src/vote/domain/vote_option.dart';

class VoteCard extends StatelessWidget {
  final Vote vote;
  const VoteCard({required this.vote, super.key});

  String __formatLocation(Location? location) {
    return location != null
        ? [
            location.city,
            location.address != null ? location.address!.street : "",
            location.country,
          ].where((element) => element.isNotEmpty).join(", ")
        : "No location available";
  }

  @override
  Widget build(BuildContext context) {
    var optionCounts = vote.options.map((option) => option.votesCount).toList();
    var totalVotes = optionCounts.reduce((a, b) => a + b);

    bool canVote = !(vote.type == VoteType.cannotChangeVote &&
        vote.options.any((option) => option.isSelected));

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
            subtitle: Text(__formatLocation(vote.event.location)),
            trailing: Icon(Icons.star_border),
            onTap: () {
              context.push('${AppRoute.events.path}/${vote.event.id}',
                  extra: vote.event);
            },
          ),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(vote.text, style: TextStyle(fontWeight: FontWeight.bold)),
                const SizedBox(height: 8),
                ...vote.options.map((option) => _buildOptionCard(
                    context, option, totalVotes, vote, canVote)),
              ],
            ),
          ),
          const SizedBox(height: 12),
        ],
      ),
    );
  }

  Widget _buildOptionCard(BuildContext context, VoteOption option,
      int totalVotes, Vote vote, bool canVote) {
    totalVotes = totalVotes == 0 ? 1 : totalVotes;

    return Card(
      margin: const EdgeInsets.symmetric(vertical: 4),
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
      child: ListTile(
        leading: CircleAvatar(
          child: IconButton(
            icon: option.isSelected
                ? Icon(Icons.circle)
                : Icon(Icons.circle_outlined),
            onPressed: canVote
                ? () {
                    context.read<VoteBloc>().add(
                        VoteOnOption(vote.id, option.id, !option.isSelected));
                  }
                : null,
          ),
        ),
        title: Text(option.text),
        subtitle: LinearProgressIndicator(
          value: option.votesCount / totalVotes,
          minHeight: 6,
          backgroundColor: Colors.grey[200],
          color: Colors.blue,
        ),
        trailing: Text(
            '${(option.votesCount / totalVotes * 100).toStringAsFixed(2)}%'),
        tileColor: canVote ? null : Colors.grey[300], // Grey out if cannot vote
      ),
    );
  }
}
