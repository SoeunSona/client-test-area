### For Go

1. go.mod
```
replace github.com/chartmogul/chartmogul-go/v4 => /Users/soeun/workspace/chartmogul-go <-- Replace a repository you are developing
```

### For Node
1. Please copy your developing chartmogul-node and pase inside of node_modules
2. please check .package-lock.json
```
{
  "lockfileVersion": 3,
  "requires": true,
  "packages": {
    "../../chartmogul-node": {
      "version": "3.2.1",
      "license": "MIT",
      "dependencies": {
        "nyc": "^15.1.0",
        "request": "^2.88.2",
        "retry": "^0.13.1",
        "uri-templates": "^0.2.0"
      },
      "devDependencies": {
        "@babel/traverse": ">=7.23.2",
        "chai": "",
        "dependency-lint": "^7.1.0",
        "eslint": "^8.54.0",
        "eslint-config-standard": "",
        "eslint-plugin-import": "",
        "eslint-plugin-n": "",
        "eslint-plugin-promise": "",
        "eslint-plugin-standard": "",
        "mocha": "",
        "nock": "",
        "pre-commit": ""
      },
      "engines": {
        "node": ">=16"
      }
    },
    "node_modules/chartmogul-node": {
      "resolved": "../../chartmogul-node",
      "link": true
    }
  }
}

```

### For PHP
1. please copy your developing chartmogul-php and paste in side of vender folder(vender/chartmogul/chartmogul-php)
2. run local server and test it

### For Python
create a virtual environment, and then installing the package in development mode.

```
$ virtualenv dev_env
$ source dev_env/bin/activate
$ cd ~/project_folder
$ pip install -e .
```
