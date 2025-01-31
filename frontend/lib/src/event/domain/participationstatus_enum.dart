/// Enum representing the participation status of a user for an event.
///
/// This enum is used to indicate whether a user is interested, will participate, or is not interested in an event.
///
/// Values:
/// - `interested`: The user is interested in the event.
/// - `willParticipate`: The user will participate in the event.
/// - `notInterested`: The user is not interested in the event.
enum ParticipationStatus {
  interested('INTERESTED'),
  willParticipate('WILL_PARTICIPATE'),
  notInterested('NOT_INTERESTED');

  final String name;
  const ParticipationStatus(this.name);

  @override
  String toString() {
    return name;
  }

  factory ParticipationStatus.fromString(String name) {
    for (ParticipationStatus type in ParticipationStatus.values) {
      if (type.name.toLowerCase() == name.toLowerCase()) {
        return type;
      }
    }
    return ParticipationStatus.notInterested;
  }
}
