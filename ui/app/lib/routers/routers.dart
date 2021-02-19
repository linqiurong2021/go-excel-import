
import 'package:flutter/material.dart';
import 'package:app/pages/home/home.dart';
import 'package:app/pages/user/user.dart';
import 'package:app/pages/map/map.dart';
// 路由
final routes = 
  {
    '/home': (BuildContext context)=> HomePage(),
    '/map': (BuildContext context)=> MapPage(),
    '/user': (BuildContext context)=> UserPage(),
  };