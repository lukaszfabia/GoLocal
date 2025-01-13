abstract class Model {
  final int id;

  Model({
    required this.id,
  });

  Model.fromJson(Map<String, dynamic> json) : id = json['id'] ?? 0;

  Map<String, dynamic> toJson() {
    return {'id': id};
  }
}
