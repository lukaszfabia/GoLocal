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
