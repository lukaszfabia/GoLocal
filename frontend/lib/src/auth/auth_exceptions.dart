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
