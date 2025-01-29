import 'dart:ffi';

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
