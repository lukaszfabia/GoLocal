import 'package:golocal/src/shared/model_base.dart';

class Survey extends Model {
  List<SurveyQuestion> questions;
  Survey({
    required super.id,
    required this.questions,
  });
}

class SurveyQuestion extends Model {
  String question;
  SurveyQuestion({
    required super.id,
    required this.question,
  });
}
