typedef Validator = String? Function(String? value);

Validator emailValidator() {
  return (String? value) {
    if (value == null || value.isEmpty) {
      return 'Please enter email';
    }
    if (!RegExp(r'^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$').hasMatch(value)) {
      return 'Please enter a valid email';
    }
    return null;
  };
}

Validator passwordValidator({int minLen = 6}) {
  return (String? value) {
    if (value == null || value.isEmpty) {
      return 'Please enter password';
    }
    if (value.length < minLen) {
      return 'Password must be at least $minLen characters';
    }
    return null;
  };
}

Validator requiredValidator({String? fieldName}) {
  return (String? value) {
    if (value == null || value.isEmpty) {
      return "${fieldName ?? 'This field'} is required";
    }
    return null;
  };
}
