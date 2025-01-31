/// Enum representing different promotion packages.
///
/// Each promotion package has a cost, a name, and a duration in days.
///
/// - `oneDay`: A promotion package for one day with a cost of 5.0.
/// - `threeDays`: A promotion package for three days with a cost of 12.0.
/// - `week`: A promotion package for one week with a cost of 25.0.
/// - `month`: A promotion package for one month with a cost of 90.0.
///
/// Properties:
/// - `cost` (`double`): The cost of the promotion package.
/// - `name` (`String`): The name of the promotion package.
/// - `duration` (`int`): The duration of the promotion package in days.
enum PromotePack {
  oneDay(5.0, "One Day", 1),
  threeDays(12.0, "Three Days", 3),
  week(25.0, "One Week", 7),
  month(90.0, "One Month", 30);

  const PromotePack(this.cost, this.name, this.duration);
  final double cost;
  final String name;
  final int duration;
}
