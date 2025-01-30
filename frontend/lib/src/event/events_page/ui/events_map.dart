import 'package:flutter/material.dart';
import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';

class EventsMap extends StatefulWidget {
  const EventsMap({super.key});

  @override
  State<EventsMap> createState() => _EventsMapState();
}

class _EventsMapState extends State<EventsMap> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: OSMViewer(
      controller: SimpleMapController(
        initPosition: GeoPoint(
          latitude: 47.4358055,
          longitude: 8.4737324,
        ),
        markerHome: const MarkerIcon(
          icon: Icon(Icons.home),
        ),
      ),
      zoomOption: const ZoomOption(
        initZoom: 16,
        minZoomLevel: 11,
      ),
    ));
  }
}
