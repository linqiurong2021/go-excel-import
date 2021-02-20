import resolve from 'rollup-plugin-node-resolve'; // node_modules 打包问题
import css from 'rollup-plugin-css-only' // css 编译问题
import json from '@rollup/plugin-json'
import commonjs from '@rollup/plugin-commonjs'; // 解决 import {} 或 require 等问题
import postcss from 'rollup-plugin-postcss'
// import { uglify } from 'rollup-plugin-uglify' // 压缩代码的插件
import image from 'rollup-plugin-image' // 图片资源
export default {
  input: './Entry.js',
  output: {
    file: './dist/main.min.js',
    format: 'umd',
    name: 'EginMap'
  },
  plugins: [
    resolve({
      // 将自定义选项传递给解析插件
      customResolveOptions: {
        moduleDirectory: 'node_modules'
      }
    }),
    json(),
    commonjs({
      dynamicRequireTargets: [
        'node_modules/esri-leaf/*.js'
      ]
    }),
    // uglify(),
    image(),
    postcss({
      minimize: true,
      extensions: ['.css']
    }),
    css({ output: './dist/eginmap.css' })
  ],
}