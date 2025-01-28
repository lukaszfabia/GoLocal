///@nodoc
library;

import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/widgets.dart';
import 'package:golocal/src/app.dart';
import 'package:golocal/src/notifications_service/notification_service.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp();
  await NotificationService.instance.init();
  runApp(GoLocalApp());
}
