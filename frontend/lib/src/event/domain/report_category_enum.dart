/// Enum representing different categories of reports.
///
/// Each category represents a type of issue that can be reported, such as inappropriate content or spam.
///
/// Values:
/// - `inappropriate`: Inappropriate content.
/// - `spam`: Spam.
/// - `illegal`: Illegal activity.
/// - `harmful`: Harmful activity.
/// - `vandalism`: Vandalism.
/// - `other`: Other.
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
