module.exports = {
  devServer: { host: '0.0.0.0', disableHostCheck: true },
  chainWebpack: config => {
    config.performance
      .maxEntrypointSize(40000000)
      .maxAssetSize(40000000)
  },
  runtimeCompiler: true,
  productionSourceMap: false
}
