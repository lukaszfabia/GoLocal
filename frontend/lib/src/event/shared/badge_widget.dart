import 'package:flutter/material.dart';

class BadgeWidget extends StatelessWidget {
  const BadgeWidget({
    super.key,
    required this.backgroundColor,
    this.text,
    this.textColor = Colors.white,
    this.child,
    this.fontSize = 12,
  }) : assert(text != null || child != null);

  final String? text;
  final Color backgroundColor;
  final double fontSize;

  final Color textColor;

  final Widget? child;

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
      margin: const EdgeInsets.only(right: 6),
      decoration: BoxDecoration(
        color: backgroundColor,
        borderRadius: BorderRadius.circular(12),
      ),
      child: child ??
          Text(
            text!,
            style: TextStyle(
              color: textColor,
              fontSize: fontSize,
              fontWeight: FontWeight.bold,
            ),
          ),
    );
  }
}
