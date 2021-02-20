import L from 'leaflet'
L.drawLocal.draw.handlers.polyline.tooltip = {
  start: `点击地图开始画线`,
  cont: `点击地图继续画线`,
  end: `点击最后一个点完成画线`
}

L.drawLocal.draw.handlers.polygon.tooltip = {
  start: `点击地图开始画面`,
  cont: `点击地图继续画面`,
  end: `点击第一个点完成画面`
}

L.drawLocal.draw.handlers.circle = {
  tooltip: {
    start: '点击地图并拖动开始画圆'
  },
  radius: '半径'
}

L.drawLocal.draw.handlers.marker = {
  tooltip: {
    start: '点击地图开始标绘'
  }
}

L.drawLocal.draw.handlers.circlemarker.tooltip = {
  start: `点击地图画圆点`,
}