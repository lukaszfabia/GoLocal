import 'package:intl/intl.dart';

extension DateTimeExtensions on DateTime {
  String formatFullDateTime({String divider = '/'}) {
    return '$day$divider$month$divider$year $hour:$minute';
  }

  String formatDateNumeric({String divider = '/'}) {
    return '$day$divider$month$divider$year';
  }

  String formatTimeOnly() {
    return '$hour:$minute';
  }

  String formatDateOnly({String divider = '/'}) {
    return '$day$divider$month$divider$year';
  }

  String formatReadableDate({bool includeTime = false}) {
    if (includeTime) {
      return DateFormat("MMM d, yyyy • HH:mm").format(this);
    }
    return DateFormat("MMM d, yyyy").format(this);
  }
}
