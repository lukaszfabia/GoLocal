/// Enum representing different types of events.
///
/// This enum is used to categorize events into different types such as workshops, cultural events, sports, etc.
///
/// Values:
/// - `workshop`: Workshop.
/// - `cultural`: Cultural event.
/// - `sports`: Sports event.
/// - `social`: Social event.
/// - `community`: Community event.
/// - `charity`: Charity event.
/// - `party`: Party.
/// - `other`: Other.
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
