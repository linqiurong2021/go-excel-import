
import 'package:flutter/material.dart';
import 'package:flutter_map/plugin_api.dart';
import 'package:latlong/latlong.dart';

class MapPage extends StatelessWidget {
  
  @override
  Widget build(BuildContext context) {
    return FlutterMap(options: MapOptions(
      center: LatLng(24.4646,118.0404),
      zoom: 13.0,
    ),
    layers: [
      TileLayerOptions(
        urlTemplate: "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
        subdomains: ['a', 'b', 'c']
      ),
      MarkerLayerOptions(
        markers: [
          Marker(
            width: 80.0,
            height: 80.0,
            point: LatLng(24.4646,118.0404),
            builder: (ctx) =>
            Container(
              child: FlutterLogo(),
            ),
          ),
        ],
      ),
    ],);
  }
}