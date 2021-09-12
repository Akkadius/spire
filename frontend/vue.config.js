module.exports = {
  devServer: { host: '0.0.0.0', disableHostCheck: true },
  chainWebpack: config => {
    config.performance
      .maxEntrypointSize(400000)
      .maxAssetSize(400000)
  },
  runtimeCompiler: true,
  productionSourceMap: false
}
