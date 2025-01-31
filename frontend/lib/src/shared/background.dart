import 'package:flutter/material.dart';
import 'dart:ui';

class BackgroundWidget extends StatelessWidget {
  const BackgroundWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        ColorCircle(
          color: const Color(0xE0E99CFF),
          opacity: 0.8,
          bottom: -0.3,
          right: -0.3,
          size: 1,
        ),
        ColorCircle(
          color: const Color(0x6851FFE4),
          opacity: 0.8,
          top: -0.3,
          left: -0.3,
          size: 1,
        ),
        BackdropFilter(
          filter: ImageFilter.blur(sigmaX: 100, sigmaY: 100),
          child: Container(
            color: Colors.transparent,
          ),
        ),
      ],
    );
  }
}

class ColorCircle extends StatelessWidget {
  const ColorCircle({
    required this.color,
    this.opacity = 0.5,
    this.left,
    this.top,
    this.bottom,
    this.right,
    this.size,
    super.key,
  });

  final double opacity;
  final Color color;
  final double? top;
  final double? bottom;
  final double? right;
  final double? left;
  final double? size;

  @override
  Widget build(BuildContext context) {
    final double screenWidth = MediaQuery.of(context).size.width;
    final double screenHeight = MediaQuery.of(context).size.height;
    return Positioned(
      top: top != null ? screenHeight * top! : null,
      left: left != null ? screenWidth * left! : null,
      bottom: bottom != null ? screenHeight * bottom! : null,
      right: right != null ? screenWidth * right! : null,
      child: Container(
        width: size != null ? screenWidth * size! : 100,
        height: size != null ? screenWidth * size! : 100,
        decoration: BoxDecoration(
          color: color.withValues(alpha: opacity),
          shape: BoxShape.circle,
        ),
      ),
    );
  }
}
