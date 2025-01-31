/// A custom exception class for authentication errors.
///
/// This class extends the [Error] class and provides a way to handle
/// authentication-related errors with a custom message.
///
/// Example usage:
/// ```dart
/// throw AuthException('Invalid credentials');
/// ```
///
/// If no message is provided, a default message 'An unknown error occurred' will be used.
class AuthException extends Error {
  late final String message;
  AuthException(String? message) {
    this.message = message ?? 'An unknown error occurred';
  }
  @override
  String toString() {
    return message;
  }
}
