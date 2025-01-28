# Getting started

First compilation tips:

```sh
flutter doctor
```

we need only what we want in green

```sh
flutter upgrade
```

if needed

```sh
flutter upgrade --force
```

```sh
flutter pub get
flutter run
```

and everything should work

## Generating documentation

```sh
dart doc .
```

if you want to document something from lib/src/, you should export it (best in a new file, with a corresponding name, in the same directory level that main.dart is)

## Tips

- best to run at least once

```sh
flutter run --no-enable-impeller
```

- if you have a ton of system messages and you are using VSCode, run in debug mode (f5) and in debug console you can filter by "I/flutter"
- out of emulators, the project is the most stable on a flutter emulator (in VSCode, right bottom - select run device - Create Android emulator)
