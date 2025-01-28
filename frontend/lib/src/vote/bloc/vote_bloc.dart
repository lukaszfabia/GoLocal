import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:golocal/src/vote/data/ivotes_repository.dart';
import 'package:golocal/src/vote/domain/vote.dart';
import 'package:meta/meta.dart';

part 'vote_event.dart';
part 'vote_state.dart';

/// Bloc responsible for handling vote-related events and states.
///
/// This Bloc manages the following events:
/// - [LoadVotes]: Loads the list of votes from the repository.
/// - [CreateVote]: Creates a new vote and reloads the list of votes.
/// - [UpdateVote]: Updates an existing vote and reloads the list of votes.
/// - [DeleteVote]: Deletes a vote and reloads the list of votes.
/// - [VoteOnOption]: Registers a vote on a specific option and reloads the list of votes.
///
/// The Bloc interacts with an [IVotesRepository] to perform the necessary
/// operations and updates the state accordingly.
class VoteBloc extends Bloc<VoteEvent, VoteState> {
  final IVotesRepository votesRepository;

  VoteBloc(this.votesRepository) : super(const VoteState(votes: [])) {
    on<LoadVotes>(_onLoadVotes);
    on<CreateVote>(_onCreateVote);
    on<UpdateVote>(_onUpdateVote);
    on<DeleteVote>(_onDeleteVote);
    on<VoteOnOption>(_onVoteOnOption);
  }

  Future<void> _onLoadVotes(LoadVotes event, Emitter<VoteState> emit) async {
    emit(state.copyWith(status: VoteStatus.loading));
    try {
      final votes = await votesRepository.getVotes();
      emit(state.copyWith(votes: votes, status: VoteStatus.loaded));
    } catch (e) {
      emit(
          state.copyWith(status: VoteStatus.error, errorMessage: e.toString()));
    }
  }

  Future<void> _onCreateVote(CreateVote event, Emitter<VoteState> emit) async {
    try {
      await votesRepository.createVote(event.vote);
      add(LoadVotes());
    } catch (e) {
      emit(
          state.copyWith(status: VoteStatus.error, errorMessage: e.toString()));
    }
  }

  Future<void> _onUpdateVote(UpdateVote event, Emitter<VoteState> emit) async {
    try {
      await votesRepository.updateVote(event.vote);
      add(LoadVotes());
    } catch (e) {
      emit(
          state.copyWith(status: VoteStatus.error, errorMessage: e.toString()));
    }
  }

  Future<void> _onDeleteVote(DeleteVote event, Emitter<VoteState> emit) async {
    try {
      await votesRepository.deleteVote(event.id);
      add(LoadVotes());
    } catch (e) {
      emit(
          state.copyWith(status: VoteStatus.error, errorMessage: e.toString()));
    }
  }

  Future<void> _onVoteOnOption(
      VoteOnOption event, Emitter<VoteState> emit) async {
    try {
      await votesRepository.voteOnOption(event.voteId, event.optionId);

      Future.delayed(Duration(seconds: 1), () {
        add(LoadVotes());
      });
    } catch (e) {
      emit(
          state.copyWith(status: VoteStatus.error, errorMessage: e.toString()));
    }
  }
}
