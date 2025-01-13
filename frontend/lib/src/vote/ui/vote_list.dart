import 'package:flutter/material.dart';
import 'package:golocal/src/vote/domain/vote.dart';
import 'package:golocal/src/vote/ui/vote_card.dart';

class VoteList extends StatelessWidget {
  final List<Vote> votes;
  const VoteList({required this.votes, super.key});

  Widget build(BuildContext context) {
    return ListView.builder(
      itemCount: votes.length,
      itemBuilder: (context, index) {
        final vote = votes[index];
        return VoteCard(vote: vote);
      },
    );
  }
}
