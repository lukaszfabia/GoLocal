enum ReportCategory {
  inappropriate("Inappropriate content"),
  spam("Spam"),
  illegal("Illegal activity"),
  harmful("Harmful activity"),
  vandalism("Vandalism"),
  other("Other");

  const ReportCategory(this.value);
  final String value;
}
