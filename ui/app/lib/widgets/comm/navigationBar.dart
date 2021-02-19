import 'package:flutter/material.dart';
import 'package:app/pages/map/map.dart';
import 'package:app/pages/home/home.dart';
import 'package:app/pages/user/user.dart';

class NavigationBar extends StatefulWidget {
  @override
  _NavigationBarState createState() => _NavigationBarState();
}

class _NavigationBarState extends State<NavigationBar> {
  // 当前选中的项
  int currentIndex = 0;
  // 底部菜单
  final homeBarItem =BottomNavigationBarItem(
    icon: Icon(Icons.home),
    label: "主页",
  );

  // 底部菜单
  final mapBarItem =BottomNavigationBarItem(
    icon: Icon(Icons.map),
    label: "地图",
  );

   // 底部菜单
  final userBarItem =BottomNavigationBarItem(
    icon: Icon(Icons.verified_user),
    label: "用户",
  );

  // 点击事件
  _tap (int index) {
    if (index != this.currentIndex) {
      setState(() {
        this.currentIndex = index;
      });
    }
  }
  // 跳转的页面
  final pages = [ HomePage(), MapPage(), UserPage()];
  // 路由
  final routes = ['/home','/map', '/user'];
  // 标题
  final titles = ['首页','地图','用户中心'];
  // 菜单
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text(titles[this.currentIndex])),
      body: pages[this.currentIndex],
      bottomNavigationBar:  BottomNavigationBar( items:[homeBarItem,mapBarItem,userBarItem] , onTap: (index){_tap(index);},currentIndex: currentIndex,iconSize: 24.0),
    );
  }
}