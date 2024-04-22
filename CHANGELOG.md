## [1.1.0](https://github.com/kilianpaquier/go-builder-generator/compare/v1.0.2...v1.1.0) (2024-04-22)


### Features

* force function name ([cfbbda7](https://github.com/kilianpaquier/go-builder-generator/commit/cfbbda749e5071420ab77930e407f66a0cdae4bf))
* option to return a copy of the builder each time a setter function is called ([273030a](https://github.com/kilianpaquier/go-builder-generator/commit/273030a178b9dde650434d92242b4b8ca2652e3f))


### Chores

* **deps:** bump github.com/hashicorp/go-getter in the all group ([1143c00](https://github.com/kilianpaquier/go-builder-generator/commit/1143c0052405170c80aff45765cf3c9ea050b45f))
* **go:** update go to 1.22.2 ([3b79e8e](https://github.com/kilianpaquier/go-builder-generator/commit/3b79e8e9165a3d2dab2a1c1010542780ac94a768))


### Code Refactoring

* **module:** remove ~ and git:: in favor of module:: to handle current module imported modules types generations - [#9](https://github.com/kilianpaquier/go-builder-generator/issues/9) ([9b49445](https://github.com/kilianpaquier/go-builder-generator/commit/9b494459b95e0b7b03445892286f436b4374ef0d))

## [1.0.2](https://github.com/kilianpaquier/go-builder-generator/compare/v1.0.1...v1.0.2) (2024-04-06)


### Bug Fixes

* **deps:** upgrade to go1.22.2 ([3dc05d7](https://github.com/kilianpaquier/go-builder-generator/commit/3dc05d7d93ec4d2d01513983d5dd3a62d532c326))
* **generate:** add unexported property to builder when generation is done in the same package and handle properly function name ([cdf58da](https://github.com/kilianpaquier/go-builder-generator/commit/cdf58da363374f91b8fa30996a3182e1f1da37a6))
* **generate:** throw an error in case an unexported struct is given in generation in another package ([4f60bb3](https://github.com/kilianpaquier/go-builder-generator/commit/4f60bb3e52b1bdb2629a5b063b11c2ef7d4f4cd6))


### Documentation

* **readme:** remove code section language to avoid weird colors ([9c05df7](https://github.com/kilianpaquier/go-builder-generator/commit/9c05df7f12cf3778b06c5342ad05b109829babac))


### Chores

* **golangci:** remove govet deleted option ([6c690b9](https://github.com/kilianpaquier/go-builder-generator/commit/6c690b9dbc4ca7513e5d1dedafe46cd00073b785))

## [1.0.1](https://github.com/kilianpaquier/go-builder-generator/compare/v1.0.0...v1.0.1) (2024-03-19)


### Bug Fixes

* **generate:** ensure Build function returns a copy of built type ([4c29c0d](https://github.com/kilianpaquier/go-builder-generator/commit/4c29c0d51a3058fd5c685f4f8cfcedc221c3b8f5))

## 1.0.0 (2024-03-18)


### ⚠ BREAKING CHANGES

* **generate:** no prefix is applied by default on builder functions, while before this commit "Set" would be applied
* **generate:** use-validator has been removed, as such there's no integration of validate anymore. See updated examples.

### Features

* **ci:** add release branches handling ([b45e0cd](https://github.com/kilianpaquier/go-builder-generator/commit/b45e0cd3b4b893eaa4ec8a91d3e143c99dad37f5))
* **generate:** add prefix option defaulting to empty meaning a builder function will just be the property name ([cd790b9](https://github.com/kilianpaquier/go-builder-generator/commit/cd790b943946d44dfc8ab4a3843eb672e5231d76))
* **generate:** change use-validator to validate-func to allow one to use any library to validate its structs ([8ebd755](https://github.com/kilianpaquier/go-builder-generator/commit/8ebd755fda796c4f3a1ded1ae33629f86b14753f))
* **gen:** return shallow copy at the end of Build function (just like Copy) ([a7c8050](https://github.com/kilianpaquier/go-builder-generator/commit/a7c80501ef0717b3bec148915a65255064ef6941))
* handle remote git structs for builder generation ([19ff472](https://github.com/kilianpaquier/go-builder-generator/commit/19ff472fc5b4fe8c9c55ab16aac5196f7ab329ac))
* handle tild go files path (to use with go/pkg) ([6ebf990](https://github.com/kilianpaquier/go-builder-generator/commit/6ebf99083c2fec6ec48afb0840484807032172c8))
* import project from gitlab ([0d432f9](https://github.com/kilianpaquier/go-builder-generator/commit/0d432f9fd0fa32ace26ce564858f9abad1acef87))


### Bug Fixes

* bad slices import ([296233b](https://github.com/kilianpaquier/go-builder-generator/commit/296233b1c6c684b56bd6642f27ab1afd3336a301))
* **ci:** bad codecov configuration ([2a4d14b](https://github.com/kilianpaquier/go-builder-generator/commit/2a4d14be95b663a06d1127ac6a955606b1d340ee))
* **ci:** bad coverage exclusions for codecov ([d663392](https://github.com/kilianpaquier/go-builder-generator/commit/d6633920e31cb017e6cc0548eee8f54d74db51e1))
* **ci:** codecov config in subdir really doesn't work ([2bd8eeb](https://github.com/kilianpaquier/go-builder-generator/commit/2bd8eebc303084ef65342847ded7ba698281b782))
* **ci:** codecov in subdir .... ([635e25d](https://github.com/kilianpaquier/go-builder-generator/commit/635e25d53225242f64c4e28cc31a9e5b801e3948))
* **ci:** handle correctly dependabot codecov ignore ([057df1c](https://github.com/kilianpaquier/go-builder-generator/commit/057df1c920bfcbdd084cab510012f7e676b561cf))
* **deps:** missing sprig dependency ([3fdceb6](https://github.com/kilianpaquier/go-builder-generator/commit/3fdceb61be1f1e008a268a17931187582394f9a9))
* handle go files at root go.mod ([0d11e09](https://github.com/kilianpaquier/go-builder-generator/commit/0d11e09bf10eba1bc81a7cc7780ba7cd02da7491))
* handle keywords in SetXXX functions parameter name ([59d8e76](https://github.com/kilianpaquier/go-builder-generator/commit/59d8e76620dccc78314ceaa9f82bd25558fbbe56))
* **imports:** aliased imports not taken into account ([8e19bd9](https://github.com/kilianpaquier/go-builder-generator/commit/8e19bd961750210bfb0c17660cedcf8091311cee))
* linting issues ([e7e6130](https://github.com/kilianpaquier/go-builder-generator/commit/e7e6130364296adf12e01ab355d9444dac5c30a1))


### Documentation

* **readme:** update sections with just CC of binary commands ([a894eb0](https://github.com/kilianpaquier/go-builder-generator/commit/a894eb0f587638bb45c5c6736332bac0a34f330b))
* update indents on README ([517b238](https://github.com/kilianpaquier/go-builder-generator/commit/517b23878f7dc304fda06e1b8615dd4411e4241f))
* update README indents ([75c8c5b](https://github.com/kilianpaquier/go-builder-generator/commit/75c8c5b092b0ce0310b27171e041df5d327a6f0b))


### Chores

* **ci:** add strategy execution for tests ([20e53ef](https://github.com/kilianpaquier/go-builder-generator/commit/20e53efe1b98911a18d9cc797c4ac022cf71c1f9))
* **ci:** regenerate files ([4a72690](https://github.com/kilianpaquier/go-builder-generator/commit/4a72690892aca324269d9b5e60cdcc6112aebb2e))
* **ci:** regenerate layout ([0394d4a](https://github.com/kilianpaquier/go-builder-generator/commit/0394d4acc6e45de667ab966f2540f423351509a6))
* **ci:** remove build/ci directory ([9009ebf](https://github.com/kilianpaquier/go-builder-generator/commit/9009ebfb8ae6b190a704f9ed0a212ca6680467cc))
* **ci:** tune codecov configuration ([e8db7ce](https://github.com/kilianpaquier/go-builder-generator/commit/e8db7cec9965b3a3f5cda4e3a5f864b1dbcc4e81))
* **deps:** add dependabot ([12748cf](https://github.com/kilianpaquier/go-builder-generator/commit/12748cf778a5f26f623e0ee7285d9ee6301794b1))
* **deps:** update dependencies ([ba16c36](https://github.com/kilianpaquier/go-builder-generator/commit/ba16c36643d00044505a96613128d9ff944142b4))
* **deps:** update dependencies ([1a76018](https://github.com/kilianpaquier/go-builder-generator/commit/1a7601810aee6c5fe61892aac9b5c6ffd73e8243))
* **examples:** add remote builder generation example ([db519d7](https://github.com/kilianpaquier/go-builder-generator/commit/db519d72f2d8c9063cf83286a833e33acd3b52c3))
* **layout:** regenerate project layout with craft ([b28ce0c](https://github.com/kilianpaquier/go-builder-generator/commit/b28ce0cf74e8e464747e37e5b0442894a03ee388))
* **testing:** use latest version of filesystem to automatically ignore windows / linux diffs ([543617c](https://github.com/kilianpaquier/go-builder-generator/commit/543617cb8ab3d06b31d0a3a5358263ed0ef9420d))
