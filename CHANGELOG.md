## [1.11.0](https://gitlab.com/kilianpaquier/go-builder-generator/compare/v1.10.5...v1.11.0) (2026-02-19)

### Chores

* **cobra:** show usage only on arguments errors ([f3adb31](https://gitlab.com/kilianpaquier/go-builder-generator/commit/f3adb312005cce852420d8e2a497d4c89114f5d6))
* **deps:** update go dependencies ([1cfd110](https://gitlab.com/kilianpaquier/go-builder-generator/commit/1cfd1106bedfb088c3a7408fd5d9dbb2fb0198b4))
* **deps:** update go dependencies ([feca6d8](https://gitlab.com/kilianpaquier/go-builder-generator/commit/feca6d88c9946dfae76ce8865dfdbcb6b69ac2d8))

### Code Refactoring

* **cobra:** reduce global variables through cobra commands definitions ([c10dc88](https://gitlab.com/kilianpaquier/go-builder-generator/commit/c10dc881aaae24fe2d0d1fbefec2e4d59531519b))
* **cobra:** use os.Args to rebuild generate command instead of flag parsing deactivation / activation ([b603618](https://gitlab.com/kilianpaquier/go-builder-generator/commit/b60361816b08143dfc3f4325b0b07eaecd5e71f2))

### Continuous Integration

* **deps:** update dependency go to v1.25.6 ([8aeb6fc](https://gitlab.com/kilianpaquier/go-builder-generator/commit/8aeb6fc35bc85abbaadeb7386f809e8226fd2be4))
* **deps:** update dependency go to v1.26.0 ([dceb219](https://gitlab.com/kilianpaquier/go-builder-generator/commit/dceb219ad63061ac573994b1bb853ebb041385de))
* **layout:** regenerate kickr layout ([c035c6a](https://gitlab.com/kilianpaquier/go-builder-generator/commit/c035c6afcb5ceecc87cfe6f1c88c3961355768c9))
* **layout:** regenerate kickr layout ([05993a4](https://gitlab.com/kilianpaquier/go-builder-generator/commit/05993a460c62c75d08e578ed9dd4ef404fbbc6e1))
* migrate to GitLab CICD ([dafa050](https://gitlab.com/kilianpaquier/go-builder-generator/commit/dafa0506e6f6d7aac59585554311f79fffe3ea0c))
* update codeowners ([770dce8](https://gitlab.com/kilianpaquier/go-builder-generator/commit/770dce86915c64aaff45c0f39cf687d6a6cc4741))

### Tests

* add Error override for required assertions ([dacad18](https://gitlab.com/kilianpaquier/go-builder-generator/commit/dacad1866b962e4913e12b291aeb660bab871e5f))
* add missing t.Helper() on internal testing func ([e5a4937](https://gitlab.com/kilianpaquier/go-builder-generator/commit/e5a493723568bfc582eadaf11343926152c215bd))
* setup internal simple testing functions ([eead35c](https://gitlab.com/kilianpaquier/go-builder-generator/commit/eead35c735938e43c323d74b462830562b835c38))

## [1.10.5](https://github.com/kilianpaquier/go-builder-generator/compare/v1.10.4...v1.10.5) (2026-01-09)

### Bug Fixes

* **module:::** make it work with local replace ([e32cd71](https://github.com/kilianpaquier/go-builder-generator/commit/e32cd71e87220f01adcd4494dfad473509a5f423))
* **require:** handle submodules cases where the parent and one of its submodules is imported ([e5c9b6c](https://github.com/kilianpaquier/go-builder-generator/commit/e5c9b6cb991daa619e77c6a2597aa593193e0980))
* **tmpl:** invalid filerelpath when using std:: ([adebdc6](https://github.com/kilianpaquier/go-builder-generator/commit/adebdc60a644181e36e9cb36f0e33cb898054cb0))

### Documentation

* **readme:** remove french typo ([0aa503c](https://github.com/kilianpaquier/go-builder-generator/commit/0aa503c79e23ea65781a73722f88d79bee9849c4))

### Continuous Integration

* **layout:** regenerate kickr layout ([db2e45c](https://github.com/kilianpaquier/go-builder-generator/commit/db2e45ce98222bcf57d537b14fbbfb3a3d42225f))
* **lint:** fix/remove lint issues ([ee44ab1](https://github.com/kilianpaquier/go-builder-generator/commit/ee44ab1fd69ccd154e62a8d076ad0faa51fa445d))

## [1.10.4](https://github.com/kilianpaquier/go-builder-generator/compare/v1.10.3...v1.10.4) (2025-12-26)

### Chores

* **deps:** update go dependencies ([ef77cf7](https://github.com/kilianpaquier/go-builder-generator/commit/ef77cf7e7235e3c3ac919b930a6f4dfd0543ddbe))

### Continuous Integration

* **deps:** update dependency go to v1.25.5 ([c762a5b](https://github.com/kilianpaquier/go-builder-generator/commit/c762a5bb1d89ace26ca5e33d9713f20ea59304c3))
* **layout:** regenerate kickr layout ([9fb99f5](https://github.com/kilianpaquier/go-builder-generator/commit/9fb99f534659812c9abc882fc40ff8b6910e0360))
* **layout:** regenerate kickr layout ([efe48bc](https://github.com/kilianpaquier/go-builder-generator/commit/efe48bcd045e89893850009ce24f1e6158a50100))

## [1.10.3](https://github.com/kilianpaquier/go-builder-generator/compare/v1.10.2...v1.10.3) (2025-11-22)

### Chores

* **deps:** update go dependencies ([2a05d68](https://github.com/kilianpaquier/go-builder-generator/commit/2a05d683c88d0e13c9805fcf6251161ed44f3c09))

### Continuous Integration

* **deps:** update dependency go to v1.25.2 ([af9518f](https://github.com/kilianpaquier/go-builder-generator/commit/af9518fc73c25d218463961083b842900648842b))
* **deps:** update dependency go to v1.25.4 ([4dbbf02](https://github.com/kilianpaquier/go-builder-generator/commit/4dbbf029fc7f351cfec61b98c6b4bf9dcd30b479))
* **deps:** update github actions dependencies ([245c11b](https://github.com/kilianpaquier/go-builder-generator/commit/245c11bfebeed2dc10285eeb491e1d198ba24ac7))
* **layout:** regenerate kickr layout ([e0c55ff](https://github.com/kilianpaquier/go-builder-generator/commit/e0c55ff794ca86c1c679ab7cd8cff9bb7ce02718))
* **layout:** regenerate kickr layout ([5f41a64](https://github.com/kilianpaquier/go-builder-generator/commit/5f41a649f57e7a11ba3e0e06c387bf5ecc28871d))
* **layout:** regenerate kickr layout ([0824257](https://github.com/kilianpaquier/go-builder-generator/commit/0824257fccfb8798e72947789c191c3fd380b36e))
* **renovate:** ignore specific testdata version update ([857f333](https://github.com/kilianpaquier/go-builder-generator/commit/857f333094625b97413dddab0ae600a9b144e807))

## [1.10.2](https://github.com/kilianpaquier/go-builder-generator/compare/v1.10.1...v1.10.2) (2025-10-15)

### Chores

* **deps:** bump the minor-patch group with 4 updates ([076ef0c](https://github.com/kilianpaquier/go-builder-generator/commit/076ef0cfb8046f33abab1472b0d2c5c9f73160a2))

## [1.10.1](https://github.com/kilianpaquier/go-builder-generator/compare/v1.10.0...v1.10.1) (2025-10-07)

### Bug Fixes

* **imports:** handle correctly imports having the same name as base package ([b893ca9](https://github.com/kilianpaquier/go-builder-generator/commit/b893ca9609f5b4db4607b814ac269f229bd5f461))

### Chores

* **deps:** remove testify and sprig in favor of local implementation for less module imports ([5c22bb9](https://github.com/kilianpaquier/go-builder-generator/commit/5c22bb92da031d16eca75e84c2a4cf5c4ae2dce4))
* **tmpl:** change `not eq` to `ne` ([55055d5](https://github.com/kilianpaquier/go-builder-generator/commit/55055d5d2b6a0a3e1469d2126131e56bfa99d6a1))

## [1.10.0](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.11...v1.10.0) (2025-09-14)

### Features

* **ast:** handle variadic fields (i.e `func(...string)`) ([d85e329](https://github.com/kilianpaquier/go-builder-generator/commit/d85e329a7a36d96a6a3db086a80482f8e0dedfc5))
* **std:** handle stdlib builders with "std::" prefix like "module::" (i.e. `-f std::go/build/build.go -s Context`) ([5354982](https://github.com/kilianpaquier/go-builder-generator/commit/5354982cc45444fee981e3e2ece4f15394b8d3fe))
* **struct:** handle inline type definitions (i.e. `struct{ Start, End time.Time }` or even `func(in, in2 int64) (out, out2 string)`) ([49a728b](https://github.com/kilianpaquier/go-builder-generator/commit/49a728b1422bee1e4921dbbb837f78c5e8010ee9))

### Bug Fixes

* **builtin:** handle correctly inputs names with direct builtin types ([ac83682](https://github.com/kilianpaquier/go-builder-generator/commit/ac83682a022327169669661d21c69bd208cd8b6d))
* **cmd:** add missing --no-tool arg in cmd in generated files ([4904a38](https://github.com/kilianpaquier/go-builder-generator/commit/4904a38d35602bffb39649008245fa6505a8ea22))
* **options:** return an error on unknown option ([1200f0a](https://github.com/kilianpaquier/go-builder-generator/commit/1200f0a98a440ca50228f9b8c1254c48d7827471))

### Documentation

* **std:** add example and README path ([3a83898](https://github.com/kilianpaquier/go-builder-generator/commit/3a83898e9642dfb69edc7bf0b5940eb400d73773))

### Chores

* **deps:** remove dependency for charmbracelet/log to avoid adding too many dependencies with go tool imports ([9a43808](https://github.com/kilianpaquier/go-builder-generator/commit/9a43808a13a0004b6997be9f8b6cbce6a636a5d3))
* **examples:** change examples names ([6560c90](https://github.com/kilianpaquier/go-builder-generator/commit/6560c909f78882ede8297e7bccb2c20e9b657d6b))
* **regexp:** invert anti-spaces char to its right char ([14a3e20](https://github.com/kilianpaquier/go-builder-generator/commit/14a3e208edbfa7d3dd047d36adb882eaee52e0ce))
* **slash:** change strings replace to filepath toslash ([ece1170](https://github.com/kilianpaquier/go-builder-generator/commit/ece1170f2a2696734ff487237fe48a0d208edb39))

### Code Refactoring

* **testdata:** rework testdata to avoid generation differences when running `make testdata` with `make test` ([0f71415](https://github.com/kilianpaquier/go-builder-generator/commit/0f71415b30317abe69169cda3a664e371cb8111f))

## [1.9.11](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.10...v1.9.11) (2025-08-27)

### Documentation

* **cli:** add precision on how to provide multiple structs in the same call ([3900617](https://github.com/kilianpaquier/go-builder-generator/commit/390061756e86938939ae83ee008d1c357dfc9aba))

### Chores

* **deps:** upgrade dependencies ([e93504a](https://github.com/kilianpaquier/go-builder-generator/commit/e93504a4981e9db607aa98c4e566ae50f1051090))

## [1.9.10](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.9...v1.9.10) (2025-08-10)

### Chores

* **deps:** bump github.com/spf13/pflag in the minor-patch group ([653c1b9](https://github.com/kilianpaquier/go-builder-generator/commit/653c1b97b49a4d9362b6cd38d6236c8e3aba0176))

## [1.9.9](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.8...v1.9.9) (2025-07-17)

### Chores

* **deps:** bump the minor-patch group across 1 directory with 2 updates ([905a35d](https://github.com/kilianpaquier/go-builder-generator/commit/905a35d2e50a4b6101b5d4852d47729bcc24c3b9))

## [1.9.8](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.7...v1.9.8) (2025-06-16)

### Chores

* **deps:** bump the minor-patch group across 1 directory with 4 updates ([633af0a](https://github.com/kilianpaquier/go-builder-generator/commit/633af0afcbae8e3ebbccdad12fd2074273613393))

## [1.9.7](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.6...v1.9.7) (2025-05-27)

### Chores

* **deps:** bump the minor-patch group across 1 directory with 2 updates ([17ccf16](https://github.com/kilianpaquier/go-builder-generator/commit/17ccf16408455a475fed6f1f9cd0d445e4cb785b))

## [1.9.6](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.5...v1.9.6) (2025-04-29)

### Chores

* **deps:** bump github.com/kilianpaquier/compare ([269091d](https://github.com/kilianpaquier/go-builder-generator/commit/269091da26305fdbeedb83d350ea0848360fab49))

## [1.9.5](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.4...v1.9.5) (2025-04-29)

### Chores

* **deps:** upgrade dependencies ([c7cb05e](https://github.com/kilianpaquier/go-builder-generator/commit/c7cb05e470118c21464ef5a741bd8253950fbd1d))

## [1.9.4](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.3...v1.9.4) (2025-04-28)

### Bug Fixes

* **module:::** use GOMODCACHE when possible then fallback on GOPATH or /home/go (like before) ([6220d88](https://github.com/kilianpaquier/go-builder-generator/commit/6220d887bb54d27c3635500821ee3dbaedb9f025))

## [1.9.3](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.2...v1.9.3) (2025-04-09)

### Chores

* define go.mod files for examples and testdata to reduce go-builder-generator dependencies and let using go tool like expected for tests ([199eb1d](https://github.com/kilianpaquier/go-builder-generator/commit/199eb1d71e77585dd48d07f661ef57108fc061bc))

## [1.9.2](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.1...v1.9.2) (2025-04-09)

### Chores

* **deps:** upgrade dependencies ([a548729](https://github.com/kilianpaquier/go-builder-generator/commit/a5487296742090f21b80418554da3be182b5e89e))

## [1.9.1](https://github.com/kilianpaquier/go-builder-generator/compare/v1.9.0...v1.9.1) (2025-04-04)

### Bug Fixes

* **gomod:** comment replace directive to allow go install usage ([5574f86](https://github.com/kilianpaquier/go-builder-generator/commit/5574f86f8e201d0d69df89b719b7e7e74dae4bd5))

## [1.9.0](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.16...v1.9.0) (2025-03-30)

### Features

* **go:generate:** use `go tool` instead of `go run ...` in generated files when possible. See [#66](https://github.com/kilianpaquier/go-builder-generator/issues/66) for full feature. ([689a943](https://github.com/kilianpaquier/go-builder-generator/commit/689a943c2cb20f0a56dbf7a10eae891468621307))

### Chores

* **deps:** upgrade dependencies ([f06fd4d](https://github.com/kilianpaquier/go-builder-generator/commit/f06fd4d9e1d192e5c5c451c9ecd13da2b6d27b47))

## [1.8.16](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.15...v1.8.16) (2025-02-24)

### Chores

* add pre-commit configuration ([3ea5fe0](https://github.com/kilianpaquier/go-builder-generator/commit/3ea5fe00e0a747e40c671eacf0e890ce12f2555d))
* **deps:** upgrade dependencies ([aeab39c](https://github.com/kilianpaquier/go-builder-generator/commit/aeab39c822c89ba53940bfc5dd652cc44e8b5cc1))

## [1.8.15](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.14...v1.8.15) (2025-02-16)

### Chores

* **deps:** bump github.com/kilianpaquier/compare ([7569e42](https://github.com/kilianpaquier/go-builder-generator/commit/7569e42d7a7781fc851124190fc799f3ff726c4c))
* **deps:** upgrade dependencies ([1848b4c](https://github.com/kilianpaquier/go-builder-generator/commit/1848b4c3c139a1f0c2411191b4033e0e43557d21))

## [1.8.14](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.13...v1.8.14) (2025-02-09)

### Chores

* **deps:** bump the minor-patch group across 1 directory with 3 updates ([0174158](https://github.com/kilianpaquier/go-builder-generator/commit/0174158d74bbb92028460da5c85828282693b922))

## [1.8.13](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.12...v1.8.13) (2025-01-27)

### Chores

* **deps:** bump github.com/samber/lo in the minor-patch group ([fee785e](https://github.com/kilianpaquier/go-builder-generator/commit/fee785e1538f849aca33a90d1bb35fc7005c894c))

## [1.8.12](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.11...v1.8.12) (2025-01-13)

### Chores

* **deps:** bump github.com/go-playground/validator/v10 ([019a1ee](https://github.com/kilianpaquier/go-builder-generator/commit/019a1eea11f6013a9123e0c47a96b42641d0d3f6))

## [1.8.11](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.10...v1.8.11) (2025-01-12)

### Chores

* **deps:** migrate dependency cli-sdk to compare ([a07d910](https://github.com/kilianpaquier/go-builder-generator/commit/a07d91082d8586becb28dd52db2b4e96231f6727))

## [1.8.10](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.9...v1.8.10) (2024-12-25)

### Chores

* **log:** migrate to charmbracelet ([c0e1745](https://github.com/kilianpaquier/go-builder-generator/commit/c0e17450cae921ae705810c61a8799c3b0c5722f))

## [1.8.9](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.8...v1.8.9) (2024-12-11)

### Chores

* **deps:** upgrade various dependencies ([08603dc](https://github.com/kilianpaquier/go-builder-generator/commit/08603dc48c100aedc30da9e7ddd65e4989621442))

## [1.8.8](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.7...v1.8.8) (2024-12-08)

### ⚙️ Chores

* **deps:** bump golang.org/x/tools in the minor-patch group ([0d43edf](https://github.com/kilianpaquier/go-builder-generator/commit/0d43edf4fcc3eb0932ef99895a864ff2bd7f7efa))

## [1.8.7](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.6...v1.8.7) (2024-11-25)

### ⚙️ Chores

* **deps:** bump github.com/go-playground/validator/v10 ([1f38e87](https://github.com/kilianpaquier/go-builder-generator/commit/1f38e8702bb72dc0255c5a26d8e04736179a3f36))
* **deps:** bump github.com/stretchr/testify in the minor-patch group ([2d5b37c](https://github.com/kilianpaquier/go-builder-generator/commit/2d5b37c09815c7b13cf9cb711e7ba104bfb15a55))
* **deps:** bump the minor-patch group across 1 directory with 2 updates ([7650d19](https://github.com/kilianpaquier/go-builder-generator/commit/7650d1990633be292766f7f4e04d1e8e2108b24c))

## [1.8.6](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.5...v1.8.6) (2024-10-05)

### Chores

* **deps:** bump golang.org/x/tools in the minor-patch group ([f10408b](https://github.com/kilianpaquier/go-builder-generator/commit/f10408bb6b4af765cbe25abfd21a25f2d2c8fc89))

## [1.8.5](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.4...v1.8.5) (2024-09-10)


### Chores

* **deps:** bump golang.org/x/tools in the minor-patch group ([0ec8dd3](https://github.com/kilianpaquier/go-builder-generator/commit/0ec8dd356215e7340407df651a7078f983605624))

## [1.8.4](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.3...v1.8.4) (2024-09-09)


### Chores

* **deps:** bump github.com/go-playground/validator/v10 ([19bfb8a](https://github.com/kilianpaquier/go-builder-generator/commit/19bfb8a45149542086fe1621bc34e9eee1216000))

## [1.8.3](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.2...v1.8.3) (2024-09-06)


### Chores

* **deps:** bump golang.org/x/mod in the minor-patch group ([2bde82a](https://github.com/kilianpaquier/go-builder-generator/commit/2bde82ad5469f467874e37d1dc4e80625e36452b))

## [1.8.2](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.1...v1.8.2) (2024-08-30)


### Chores

* **deps:** bump github.com/Masterminds/sprig/v3 ([1c6fb38](https://github.com/kilianpaquier/go-builder-generator/commit/1c6fb385a935a27c48e1638cdea8bbcf7b921b1f))

## [1.8.1](https://github.com/kilianpaquier/go-builder-generator/compare/v1.8.0...v1.8.1) (2024-08-17)


### Documentation

* **upgrade:** add default installation destination path in help ([eb47cf3](https://github.com/kilianpaquier/go-builder-generator/commit/eb47cf336fde5fb67128b6900f9a2336fbcc7b51))

## [1.8.0](https://github.com/kilianpaquier/go-builder-generator/compare/v1.7.2...v1.8.0) (2024-08-17)


### Features

* **upgrade:** add new command for easier upgrade of go-builder-generator ([438e8d0](https://github.com/kilianpaquier/go-builder-generator/commit/438e8d05703801f81e86d8337908288af7064064))


### Bug Fixes

* **generate:** show help when flag is provided ([e5bc4ff](https://github.com/kilianpaquier/go-builder-generator/commit/e5bc4ff1839a2d2528b8ee207ce0dccaf35c8204))

## [1.7.2](https://github.com/kilianpaquier/go-builder-generator/compare/v1.7.1...v1.7.2) (2024-08-14)


### Bug Fixes

* **comments:** add missing --package-name when given ([3adc1a4](https://github.com/kilianpaquier/go-builder-generator/commit/3adc1a48d72f7d4b46503de1907228c7226e5e2b))
* **comments:** correctly parse args with cobra and strengthen regexp search ([471a3f6](https://github.com/kilianpaquier/go-builder-generator/commit/471a3f69384419dc0684de5f8cb8e5df6cafccef))

## [1.7.1](https://github.com/kilianpaquier/go-builder-generator/compare/v1.7.0...v1.7.1) (2024-08-11)


### Bug Fixes

* incomplete regular expression to check whether go:generate is present in input file ([c9e6089](https://github.com/kilianpaquier/go-builder-generator/commit/c9e6089f35ddfcae6a3c5c9ff50c597b06178c67))

## [1.7.0](https://github.com/kilianpaquier/go-builder-generator/compare/v1.6.3...v1.7.0) (2024-08-11)


### Features

* provide filepath to where the structs comes from in generated files and don't add go:generate command when it already exists in input file - [#30](https://github.com/kilianpaquier/go-builder-generator/issues/30) ([ebff42c](https://github.com/kilianpaquier/go-builder-generator/commit/ebff42c6aa1c1f295d7679b36e192db71f29cf82))

## [1.6.3](https://github.com/kilianpaquier/go-builder-generator/compare/v1.6.2...v1.6.3) (2024-08-09)


### Chores

* **deps:** upgrade go toolchain to go1.22.6 ([ebdd36d](https://github.com/kilianpaquier/go-builder-generator/commit/ebdd36dea18b1067cf1b3699f1631724f66a39a1))

## [1.6.2](https://github.com/kilianpaquier/go-builder-generator/compare/v1.6.1...v1.6.2) (2024-08-07)


### Chores

* **deps:** bump golang.org/x/tools in the minor-patch group ([9849ce6](https://github.com/kilianpaquier/go-builder-generator/commit/9849ce6cc348e52b89db201c0b15aa195f128170))

## [1.6.1](https://github.com/kilianpaquier/go-builder-generator/compare/v1.6.0...v1.6.1) (2024-08-05)


### Chores

* **deps:** bump golang.org/x/mod in the minor-patch group ([715dcec](https://github.com/kilianpaquier/go-builder-generator/commit/715dcecff5396088d6a074c0229f7ca079940d42))

## [1.6.0](https://github.com/kilianpaquier/go-builder-generator/compare/v1.5.2...v1.6.0) (2024-07-20)


### Features

* **generate:** add --package-name CLI arg ([c884191](https://github.com/kilianpaquier/go-builder-generator/commit/c884191ee2a9db4f253c67c6de40d2e19b3e7d55))
* **generate:** add top comment 'go:generate' in generated files to easily regenerate files alongside disabling option --no-cmd to remove this comment - [#28](https://github.com/kilianpaquier/go-builder-generator/issues/28) ([cd6c56d](https://github.com/kilianpaquier/go-builder-generator/commit/cd6c56da74ad75439f458e6dbac3a5ef4bff789d))


### Bug Fixes

* **generate:** set the right package name when generating in the same package but with a package name difference from directory name ([8f10fdd](https://github.com/kilianpaquier/go-builder-generator/commit/8f10fddc7268ded524cb777b0c9ac31c55b546ef))

## [1.5.2](https://github.com/kilianpaquier/go-builder-generator/compare/v1.5.1...v1.5.2) (2024-07-16)


### Chores

* **deps:** bump github.com/samber/lo ([f96f627](https://github.com/kilianpaquier/go-builder-generator/commit/f96f6271172a8607cab838e3963f216ba332a8ce))

## [1.5.1](https://github.com/kilianpaquier/go-builder-generator/compare/v1.5.0...v1.5.1) (2024-07-08)


### Chores

* **deps:** bump golang.org/x/tools in the minor-patch group ([ba3ba18](https://github.com/kilianpaquier/go-builder-generator/commit/ba3ba18e492272875bab56f02566b82759ba33e0))

## [1.5.0](https://github.com/kilianpaquier/go-builder-generator/compare/v1.4.2...v1.5.0) (2024-07-05)


### Code Refactoring

* remove dependency to filesystem and logrus ([2f8d537](https://github.com/kilianpaquier/go-builder-generator/commit/2f8d537d8c246b7a344dbdedfeabe45b4190fbf9))

## [1.4.2](https://github.com/kilianpaquier/go-builder-generator/compare/v1.4.1...v1.4.2) (2024-07-01)


### Chores

* **deps:** bump github.com/samber/lo in the minor-patch group ([7571304](https://github.com/kilianpaquier/go-builder-generator/commit/7571304b78773d7f4aee536fd27403d6f427ecd7))

## [1.4.1](https://github.com/kilianpaquier/go-builder-generator/compare/v1.4.0...v1.4.1) (2024-06-28)


### Chores

* **deps:** bump github.com/samber/lo in the minor-patch group ([836804f](https://github.com/kilianpaquier/go-builder-generator/commit/836804f064905ab23fe3a91c416152d924a286cb))

## [1.4.0](https://github.com/kilianpaquier/go-builder-generator/compare/v1.3.3...v1.4.0) (2024-06-16)


### Features

* **deps:** upgrade xstrings to v1.5.0 which changes ToCamelCase to respect what the function should initially do ([495434d](https://github.com/kilianpaquier/go-builder-generator/commit/495434db99547427622d4ad9269a21c31a5efff8))

## [1.3.3](https://github.com/kilianpaquier/go-builder-generator/compare/v1.3.2...v1.3.3) (2024-06-01)


### Chores

* **deps:** upgrade filesystem and validator dependencies ([91f0e40](https://github.com/kilianpaquier/go-builder-generator/commit/91f0e4019f728d714131e4f0237b24cbb9fc410c))

## [1.3.2](https://github.com/kilianpaquier/go-builder-generator/compare/v1.3.1...v1.3.2) (2024-05-08)


### Chores

* **deps:** bump golang.org/x/tools in the minor-patch group ([ab743c1](https://github.com/kilianpaquier/go-builder-generator/commit/ab743c1152d4f3880b6cda7ec387e39909738fab))
* **deps:** upgrade toolchain to go1.22.3 ([51a1716](https://github.com/kilianpaquier/go-builder-generator/commit/51a17169ab7758c99b92486f626f47911cfd4249))

## [1.3.1](https://github.com/kilianpaquier/go-builder-generator/compare/v1.3.0...v1.3.1) (2024-05-01)


### Chores

* **deps:** upgrade both filesystem and go-validator dependencies ([b5e8374](https://github.com/kilianpaquier/go-builder-generator/commit/b5e8374394ce1d10fc4d7eff58610de6e30c68e8))

## [1.3.0](https://github.com/kilianpaquier/go-builder-generator/compare/v1.2.2...v1.3.0) (2024-04-26)


### Features

* **options:** add export function to export specific field's builder function (only work in certain cases, please check README.md) ([c39e935](https://github.com/kilianpaquier/go-builder-generator/commit/c39e935e6338b8ede596e85141f6bf2cca0a96a7))

## [1.2.2](https://github.com/kilianpaquier/go-builder-generator/compare/v1.2.1...v1.2.2) (2024-04-25)


### Bug Fixes

* prefix not taken into account anymore in generation ([827b842](https://github.com/kilianpaquier/go-builder-generator/commit/827b8425a87fc9ecdbd61a60db3b7c43c5566f82))

## [1.2.1](https://github.com/kilianpaquier/go-builder-generator/compare/v1.2.0...v1.2.1) (2024-04-25)


### Bug Fixes

* **generic:** add unary (~) implementation for generic builders ([08baa16](https://github.com/kilianpaquier/go-builder-generator/commit/08baa16f225b4f1fddac6362d733bc6274caabb5))
* **generic:** handle acronyms and lowercase generic names ([a642ea0](https://github.com/kilianpaquier/go-builder-generator/commit/a642ea02b21df56721fb70120f983b3336db9c67))


### Documentation

* **readme:** add some cases examples ([37d1bd8](https://github.com/kilianpaquier/go-builder-generator/commit/37d1bd8e1207f8ae49cfb0a45e5a3f9685bbde7d))

## [1.2.0](https://github.com/kilianpaquier/go-builder-generator/compare/v1.1.2...v1.2.0) (2024-04-24)


### Features

* add generic struct builders generation ([db5a8f1](https://github.com/kilianpaquier/go-builder-generator/commit/db5a8f147bd9ab8525937824e2dfb42f8ba318b6))

## [1.1.2](https://github.com/kilianpaquier/go-builder-generator/compare/v1.1.1...v1.1.2) (2024-04-23)


### Bug Fixes

* **struct:** add tags to builders around anonymous struct ([712a4d9](https://github.com/kilianpaquier/go-builder-generator/commit/712a4d97af0485b32df9c81bfd7815d83b3737b1))

## [1.1.1](https://github.com/kilianpaquier/go-builder-generator/compare/v1.1.0...v1.1.1) (2024-04-23)


### Bug Fixes

* error when using generic field ([5b322a7](https://github.com/kilianpaquier/go-builder-generator/commit/5b322a7257f240a645758dfd107d5484087b2c61))

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
