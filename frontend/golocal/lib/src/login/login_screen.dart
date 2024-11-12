import 'dart:ui';

import 'package:flutter/material.dart';

class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: Stack(
        fit: StackFit.loose,
        alignment: Alignment.center,
        clipBehavior: Clip.hardEdge,
        children: [
          ColorCircle(
            color: Colors.purpleAccent,
            opacity: 0.8,
            top: -0.1,
            left: -0.5,
            size: 1.3,
          ),
          ColorCircle(
            color: Colors.deepPurple,
            opacity: 0.9,
            top: -0.5,
            left: -0.6,
            size: 1.5,
          ),
          ColorCircle(
            color: Colors.yellow,
            opacity: 0.8,
            size: 2,
            right: -1.4,
            top: 0.3,
          ),
          BackdropFilter(
            filter: ImageFilter.blur(sigmaX: 100, sigmaY: 100),
            child: Container(
              color: Colors.transparent,
            ),
          ),
          Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Padding(
                padding: const EdgeInsets.all(8.0),
                child: ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    elevation: 0,
                    shadowColor: Colors.transparent,
                    foregroundColor: Colors.black,
                    backgroundColor: Colors.white.withOpacity(0.15),
                    minimumSize: Size(MediaQuery.of(context).size.width / 1.2,
                        MediaQuery.of(context).size.height / 15),
                  ),
                  onPressed: () {},
                  child: Text("Sign In with Google"),
                ),
              ),
              Padding(
                padding: const EdgeInsets.all(8.0),
                child: ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    elevation: 0,
                    shadowColor: Colors.transparent,
                    foregroundColor: Colors.black,
                    backgroundColor: Colors.white.withOpacity(0.15),
                    minimumSize: Size(MediaQuery.of(context).size.width / 1.2,
                        MediaQuery.of(context).size.height / 15),
                  ),
                  onPressed: () {},
                  child: Text("Sign In with Facebook"),
                ),
              ),
            ],
          )
        ],
      ),
    );
  }
}

class ColorCircle extends StatelessWidget {
  ColorCircle({
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
          color: color.withOpacity(opacity),
          shape: BoxShape.circle,
        ),
      ),
    );
  }
}
