import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

Future<void> showMyDialog(
  BuildContext context, {
  required String title,
  required String message,
  bool isWarning = false,
  bool isInfo = false,
  bool doublePop = false,
  Future<void> Function()? onOk,
  bool showCancel = false,
}) async {
  IconData dialogIcon = Icons.info;
  Color iconColor = Colors.blue;

  if (isWarning) {
    dialogIcon = Icons.warning_amber_rounded;
    iconColor = Colors.orange;
  } else if (isInfo) {
    dialogIcon = Icons.info_outline;
    iconColor = Colors.blue;
  }

  return showDialog(
    context: context,
    barrierDismissible: false,
    builder: (BuildContext context) {
      return AlertDialog(
        shape:
            RoundedRectangleBorder(borderRadius: BorderRadius.circular(12.0)),
        title: Row(
          children: [
            Icon(dialogIcon, color: iconColor),
            const SizedBox(width: 10),
            Expanded(
                child:
                    Text(title, style: TextStyle(fontWeight: FontWeight.bold))),
          ],
        ),
        content: Text(message, textAlign: TextAlign.center),
        actionsAlignment: MainAxisAlignment.spaceEvenly,
        actions: <Widget>[
          if (showCancel)
            TextButton(
              onPressed: () => context.pop(),
              child: const Text('Cancel', style: TextStyle(color: Colors.grey)),
            ),
          TextButton(
            onPressed: () async {
              context.pop();
              if (doublePop && context.canPop()) {
                context.pop();
              }
              if (onOk != null) {
                await onOk();
              }
            },
            child:
                const Text('OK', style: TextStyle(fontWeight: FontWeight.bold)),
          ),
        ],
      );
    },
  );
}
