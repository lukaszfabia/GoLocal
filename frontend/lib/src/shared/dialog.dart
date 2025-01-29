import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

Future<dynamic> showMyDialog(BuildContext context,
    {required String title,
    required String message,
    bool isWarning = false,
    bool isInfo = false,
    bool doublePop = false}) {
  return showDialog(
    context: context,
    builder: (BuildContext context) {
      return AlertDialog(
        title: Text(title),
        actionsAlignment: MainAxisAlignment.center,
        content: Text(message),
        actions: <Widget>[
          TextButton(
            onPressed: () {
              context.pop();
              if (doublePop && context.canPop()) {
                context.pop();
              }
            },
            child: Text('OK'),
          ),
        ],
      );
    },
  );
}
