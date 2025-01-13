enum EventType {
  workshop('WORKSHOP'),
  cultural('CULTURAL'),
  sports('SPORTS'),
  social('SOCIAL'),
  community('COMMUNITY'),
  charity('CHARITY'),
  party('PARTY'),
  other('OTHER');

  final String name;
  const EventType(this.name);

  @override
  String toString() {
    return name;
  }

  factory EventType.fromString(String name) {
    for (EventType type in EventType.values) {
      if (type.name.toLowerCase() == name.toLowerCase()) {
        return type;
      }
    }
    return EventType.other;
  }
}
