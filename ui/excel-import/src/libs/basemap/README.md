# 地图调用



# 引入
```
import Baidu from './Baidu.js'  # 引用路径需要注意  目前是在当前目录下
import GaoDe from './GaoDe.js' # 引用路径需要注意  目前是在当前目录下
import Geoq from './Geoq.js'  # 引用路径需要注意  目前是在当前目录下
import Google from './Google.js'  # 引用路径需要注意  目前是在当前目录下
import TiandiTu from './TiandiTu.js'  # 引用路径需要注意  目前是在当前目录下
```


# 使用说明

::: tip
如果有相关的天地图或其它的token 请更换为自已申请的token
:::


## 天地图参数说明
```
 # 卫星图
 addTianDiTuSatelliteMap()
 # 卫星注记图
 addTianDiTuSatelliteAnnotion()
 # 矢量图
 addTianDiTuNormalMap()
 # 矢量注记图  
 addTianDiTuNormalAnnotion()

```

## 高德地图参数说明

```
# 矢量图与注记图
addGaoDeNormalMap()
# 卫星图
addGaoDeSatelliteMap()
# 卫星注记图
addGaoDeSatelliteAnnotion()

```


## 谷歌地图参数说明
```
# 矢量图
addMap()
# 卫星图
addSatelliteMap()
# 卫星注记图
addSatelliteAnnotion()
```


## 智图参数说明
```
# 矢量图
addGeoqNormalMap()
# 矢量图灰色
addGeoqNormalGray()
# 矢量图
addGeoqNormalHydro()
# 矢量图
addGeoqNormalWarm()
# 矢量图
addGeoqNormalPurplishBlue()
```


## 百度地图参数说明

```
# 矢量图
addBaiDuNormalMap()
# 卫星图
addBaiDuSatelliteMap()
# 卫星图注记图
addBaiDuSatelliteAnnotion()
```


# 代码事例

> 天地图调用
```javascript
// 添加底图
const tiandiTu = new TiandiTu({ key: '4b350b4f343fa22cdb2047e93b4d8712' })
// 卫星图
let tdtSatellite = tiandiTu.addTianDiTuSatelliteMap()
//  卫星注记图
let tdtSatelliteAnnotion = tiandiTu.addTianDiTuSatelliteAnnotion()
// 分组
let tdtSatelliteGroup = this.eginMap.L.layerGroup([
  tdtSatellite,
  tdtSatelliteAnnotion
])
// 添加 title 与 url 属性 为切换底图做准备
tdtSatelliteGroup.name = '卫星图'
tdtSatelliteGroup.url = ''
// 添加到地图
this.eginMap.map.addLayer(tdtSatelliteGroup)
// 矢量图
let tdtNormal = tiandiTu.addTianDiTuNormalMap()
//  矢量注记图
let tdttdtNormalAnnotion = tiandiTu.addTianDiTuNormalAnnotion()
// 分组
let tdtNormalGroup = this.eginMap.L.layerGroup([
  tdtNormal,
  tdttdtNormalAnnotion
])
tdtNormalGroup.name = '矢量图'
tdtNormalGroup.url = ''

// 添加到底图里
this.eginMap.basemapLayers.push(tdtSatelliteGroup)
// 添加到底图里
this.eginMap.basemapLayers.push(tdtNormalGroup)

```


> 高德地图调用
```javascript
const gaode = new GaoDe({ key: 'xxxx' })
// 矢量图
let gaodeNormal = gaode.addGaoDeNormalMap()
// 添加到地图
this.eginMap.map.addLayer(gaodeNormal)

// 卫星图
let gaode2 = gaode.addGaoDeSatelliteMap()
// 卫星注记图
let gaode3 = gaode.addGaoDeSatelliteAnnotion()

let gaodeSatelliteGroup = this.eginMap.L.layerGroup([
  gaode2,
  gaode3
])
// 添加到地图
this.eginMap.map.addLayer(gaodeSatelliteGroup)

```

> 谷歌地图调用
```javascript
const google = new Google({ key: 'xxxx' })
// 矢量图
let googleNormal = gaode.addMap()
// 添加到地图
this.eginMap.map.addLayer(googleNormal)

// 卫星图
let googgle1 = gaode.addSatelliteMap()
// 卫星注记图
let googgle2 = gaode.addSatelliteAnnotion()

let googleSatelliteGroup = this.eginMap.L.layerGroup([
  googgle1,
  googgle2
])
// 添加到地图
this.eginMap.map.addLayer(googleSatelliteGroup)


```

> 智图地图调用
```javascript
const geoq = new Geoq({ key: 'xxxx' })
// 矢量图
let geoqNormal = geoq.addGeoqNormalMap()
// 添加到地图
this.eginMap.map.addLayer(geoqNormal)


// 矢量图
let addGeoqNormalGray = geoq.addGeoqNormalGray()
// 添加到地图
this.eginMap.map.addLayer(addGeoqNormalGray)


// 矢量图
let addGeoqNormalHydro = geoq.addGeoqNormalHydro()
// 添加到地图
this.eginMap.map.addLayer(addGeoqNormalHydro)


// 矢量图
let addGeoqNormalWarm = geoq.addGeoqNormalWarm()
// 添加到地图
this.eginMap.map.addLayer(addGeoqNormalWarm)


// 矢量图
let addGeoqNormalPurplishBlue = geoq.addGeoqNormalPurplishBlue()
// 添加到地图
this.eginMap.map.addLayer(addGeoqNormalPurplishBlue)


```

> 百度地图调用
```javascript

const baidu = new Baidu({ key: 'xxxx' })
// 矢量图
let addBaiDuNormalMap = baidu.addBaiDuNormalMap()
// 添加到地图
this.eginMap.map.addLayer(addBaiDuNormalMap)

// 卫星图
let addBaiDuNormalMap = baidu.addBaiDuSatelliteMap()
// 卫星注记
let addBaiDuSatelliteAnnotion = baidu.addBaiDuSatelliteAnnotion()

let baiduGroup = this.eginMap.L.layerGroup([
  addBaiDuNormalMap,
  addBaiDuSatelliteAnnotion
])

// 添加到地图
this.eginMap.map.addLayer(baiduGroup)




```