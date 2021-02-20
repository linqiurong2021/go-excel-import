```
 // css相关配置
css: {
  // 是否使用css分离插件 ExtractTextPlugin
  extract: true,
  // 开启 CSS source maps?
  sourceMap: false,
  // css预设器配置项
  loaderOptions: {
    scss: {
      prependData: `@import "./src/styles/main.scss";`
    }
  }
},
```

::: tip
请注意 App.vue的文件中

<style lang="scss">
#app {}
</style>

其中 <font color="#ff000">lang="scss" </font> 且 <font color="#ff0000" >#app</font>需要存在
:::

