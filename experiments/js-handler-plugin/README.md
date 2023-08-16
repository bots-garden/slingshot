
## Install Parcel

```bash
npm install esbuild --save-dev
```

> esbuild.js
```javascript
const esbuild = require('esbuild');

esbuild
    .build({
        entryPoints: ['src/index.js'],
        outdir: 'dist',
        bundle: true,
        sourcemap: true,
        minify: false, // might want to use true for production build
        format: 'cjs', // needs to be CJS for now
        target: ['es2020'] // don't go over es2020 because quickjs doesn't support it
    })
```


> package.json
```json
{
  "name": "js-handler-plugin",
  "version": "1.0.0",
  "main": "index.js",
  "scripts": {
    "build": "node esbuild.js && extism-js dist/index.js -o ./handler-js.wasm"
  },
  "author": "@k33g_org",
  "license": "MIT",
  "devDependencies": {
    "esbuild": "^0.18.17"
  }
}

```

=> `npm run build`
