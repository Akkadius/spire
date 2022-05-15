const path = require('path')

module.exports = {
  devServer: {
    host: '0.0.0.0',
    disableHostCheck: true,
    watchOptions: {
      ignored: ['/node_modules/', '/public/'],
    }
  },
  // pluginOptions: {
  //   webpackBundleAnalyzer: {
  //     openAnalyzer: false
  //   }
  // },

  chainWebpack: config => {
    config.performance
      .maxEntrypointSize(40000000)
      .maxAssetSize(40000000)

    // ignore asset preview during development to keep build times down
    if (process.env.NODE_ENV !== 'production') {
      config.plugin('copy').tap(([options]) => {
        options[0].ignore.push('eq-asset-preview-master/**/*')

        return [options]
      })
    }

    //
    config.output
      .filename('[name].[hash].js')
      .path(path.resolve(__dirname, 'dist'))
      .clean(true)

    // config.optimization.moduleIds    = 'deterministic'
    // config.optimization.runtimeChunk = 'single'
    config.optimization.splitChunks = {
      cacheGroups: {
        vendor: {
          test: /[\\/]node_modules[\\/]/,
          name: 'vendors',
          chunks: 'all',
        },
      },
    }
    // console.log(config)
  },
  runtimeCompiler: true,
  productionSourceMap: false
}
