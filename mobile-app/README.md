# Getting started

First compilation tips:

- flutter doctor # we need only what we want in green
- flutter upgrade (--force) # if needed
- flutter pub get
- flutter run
and it should work

## Tips

- impeller error something something -> flutter run --no-enable-impeller
- if you have a ton of system messages, if using VSCode run in debug mode (f5) and in debug console you can filter by "I/flutter". Another approach is to run

```cmd
flutter run | sed '/^\(V\|I\|W\|E\)\/flutter/!d'
```
