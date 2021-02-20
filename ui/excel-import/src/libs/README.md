
# 地图的主要文件

## ClipPolygon 图形切割类工具

### 方法说明

 unionPolygon  合并多边形

 polygonClipByLine 线切割多边形

 _singlePolygonClip 单个多边形切割

 _multiPolygonClip  多个多边形切割

 connectLine  连线

 isOnLine  是否在线上

 getIntersectPoints  获得交点

 multiPolygon2polygons multiPolygon转polygons,不涉及属性

 polygons2MultiPolygon polygons转multiPolygon,不涉及属性，只输出属性为{}



## EginLayer 图层类

### 方法说明

 setOpacity  设置图层透明度

 layerGroup 图形分组

 getLayerGroup 通过ID获取group

 removeGroup  通过ID删除group

 addGeoJsonLayer  添加geoJson

 addEsriFeatureLayer  添加arcgis图层

 addEsriFeatureLayerToMap  添加到地图上

 addTileLayer 添加瓦片图层

 addClusterLayer 添加聚合图层

 setBaseLayer  设置底图

 setRetinaLayer 设置RetinalLayer

 setVectorLayer 矢量图层



## EginMap 实例化地图类

### 属性说明

 L  Leaflet类

 Esri esriforleaflet 类

 map 地图

 basemapLayers  底图图层

 layers  图层数据

 options  其它属性

 printTool  打印工具

 Tools 工具类

### 方法说明

 setDefaultIcon  设置默认ICON

 setView 设置视图中心

 setZoom 设置缩放级别

 flyTo  定位有动画

 panTo  定位无动画

 setZoomPosition  设置缩放的位置

 widgetActive  控件激活转换状态

 loadBasempWidget 加载底图控件

 loadLayersWidget 加载图层控件

 loadToolsWidget 加载工具箱控件

 loadStaticWidget 加载统计控件

 loadScaleWidget 加载比例尺

 loadPrintWidget 添加打印

 exportMap 导出地图

 printMap 打印地图

 splitSync 分屏比对

 wrap 卷帘





## EginMapTool 与地图相关的工具

### 方法说明

 newFeatureGroup  新增图形分组

 newMakerIcon 设置ICON

 initOptions 初始化参数属性

 initDrawPolyline  初始化画线工具

 initDrawPolygon  初始化画面工具

 initTextMarker  初始化文本工具

 initMarker  初始化标点工具

 initCircle  初始化画圆工具

 initRectangle  初始化画矩形工具

 initPane  初始化pane

 getArea 获取面积

 query  Esri.query查询 

 getPopupContent  获取点击图形弹窗

 editCircle 编辑圆

 newDivIcon 新增DIVICON

 newMarker 新增点

 initDelete 初始化删除



## EginTool 与地图无关的工具

### 方法说明

 latLngToLngLat   y,x  转 x,y

 geoJsonToArcgis geoJson转arcgis

 arcgisToGeoJSON  arcgis转geoJSON

 guid  生成GUID

 _getGeometry  获取空间数据
