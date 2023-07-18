# Changelog

# [2.8.0](https://github.com/tj-actions/auto-doc/compare/v2.7.1...v2.8.0) - (2023-07-18)

## <!-- 0 -->🚀 Features

- Increase test coverage and fix bug with inputs ([#498](https://github.com/tj-actions/auto-doc/issues/498)) ([6a2669f](https://github.com/tj-actions/auto-doc/commit/6a2669fed415ebb5749c34f713ea8abdac4407af))  - (Tonye Jack)

## <!-- 1 -->🐛 Bug Fixes

- Fixed deprecation warning and increased fetch-depth
 ([f6b4fa9](https://github.com/tj-actions/auto-doc/commit/f6b4fa925e77586475e50e974aef324da96a64a8))  - (Tonye Jack)
- Usage of deprecated archives.replacements ([4050d60](https://github.com/tj-actions/auto-doc/commit/4050d60894895938b6ee9efaa6f1d2bf5829fca7))  - (Tonye Jack)

## <!-- 17 -->➖ Remove

- Deleted .github/workflows/auto-merge.yml ([5764663](https://github.com/tj-actions/auto-doc/commit/5764663d146290feebabeea0f2400c6194070339))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Updated goreleaser.yaml
 ([22c090b](https://github.com/tj-actions/auto-doc/commit/22c090b88275a2425b3b32fbea856a4c914c97c4))  - (Tonye Jack)
- Updated gorelease config
 ([13d63b3](https://github.com/tj-actions/auto-doc/commit/13d63b327ccdc576035653555f8822a1d1fb9bb0))  - (Tonye Jack)
- Update .goreleaser.yaml ([e928446](https://github.com/tj-actions/auto-doc/commit/e9284464dad169129a4cd1ed3dc261176aaf1faa))  - (Tonye Jack)
- Update .goreleaser.yaml ([91306a8](https://github.com/tj-actions/auto-doc/commit/91306a872cc22f623d7073f7ffbeacb67746a68f))  - (Tonye Jack)
- Updated renovate.json ([03b66c9](https://github.com/tj-actions/auto-doc/commit/03b66c9d85863676eed93e0043949998a0c9bcd1))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#494](https://github.com/tj-actions/auto-doc/pull/494): usage of deprecated archives.replacements ([a312178](https://github.com/tj-actions/auto-doc/commit/a3121780456a6a94fe5ca8ca86c6ac0c0c8d0d4b))  - (repo-ranger[bot])
- PR [#492](https://github.com/tj-actions/auto-doc/pull/492): to v2.7.1 ([1a30ecd](https://github.com/tj-actions/auto-doc/commit/1a30ecdd80268a844f889828dd14eabd2b2d6552))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Don't run setup-bin if bin_path input is provided ([#493](https://github.com/tj-actions/auto-doc/issues/493)) ([7a164ea](https://github.com/tj-actions/auto-doc/commit/7a164ea952a68498642081d3fd8fc804c8b4ac69))  - (Tonye Jack)

# [2.7.1](https://github.com/tj-actions/auto-doc/compare/v2.7.0...v2.7.1) - (2023-06-23)

## <!-- 1 -->🐛 Bug Fixes

- Fixed adding more spaces
 ([4835a03](https://github.com/tj-actions/auto-doc/commit/4835a034a771bb8d60eb494645d9cfb5130999af))  - (Tonye Jack)
- Bug where extra spaces get added to output ([3f13a09](https://github.com/tj-actions/auto-doc/commit/3f13a09c204c0704e1e37543fcdd8d28cd649eb8))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Update action.go ([c722781](https://github.com/tj-actions/auto-doc/commit/c722781eb753c4198db60f690e6f26a212411dbf))  - (Tonye Jack)
- Update action.go ([0ec2c5e](https://github.com/tj-actions/auto-doc/commit/0ec2c5eb12d5c272a16ad883d986aa9797265956))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#490](https://github.com/tj-actions/auto-doc/pull/490): bug where extra spaces get added to output ([9c63fba](https://github.com/tj-actions/auto-doc/commit/9c63fba96d9a5b2744c01b442fc6d0a4735bab0f))  - (repo-ranger[bot])
- Merge branch 'main' into fix/bug-where-extra-spaces-get-added-to-output ([0029ca6](https://github.com/tj-actions/auto-doc/commit/0029ca6f2e0dad8bd9ad90b4a696d8b267d5a795))  - (Tonye Jack)
- PR [#486](https://github.com/tj-actions/auto-doc/pull/486): to v2.7.0 ([f90aa0e](https://github.com/tj-actions/auto-doc/commit/f90aa0e1e928b4264c0cb0c95024eccda5e49336))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Upgraded from v2.7.0 -> v2.7.1 ([d02dccf](https://github.com/tj-actions/auto-doc/commit/d02dccf93632674633d89eb4ec72b3907d04cbb1))  - (github-actions[bot])
- **deps:** Update tj-actions/verify-changed-files action to v16 ([8b78197](https://github.com/tj-actions/auto-doc/commit/8b7819715a14b3e5a2e119f15a92b9ee3711bf60))  - (renovate[bot])
- **deps:** Update peter-evans/create-pull-request action to v5.0.2 ([41ddaa9](https://github.com/tj-actions/auto-doc/commit/41ddaa9f11db6fd2e36857f4d3c1cb30c736397e))  - (renovate[bot])
- **deps:** Update tj-actions/verify-changed-files action to v15 ([a0e14ec](https://github.com/tj-actions/auto-doc/commit/a0e14ec6d518274908236a62817b9f56c8e01dfa))  - (renovate[bot])

# [2.7.0](https://github.com/tj-actions/auto-doc/compare/v2.6.1...v2.7.0) - (2023-06-07)

## <!-- 0 -->🚀 Features

- Add a new input to optionally set names as markdown anchors ([#480](https://github.com/tj-actions/auto-doc/issues/480)) ([b750d7a](https://github.com/tj-actions/auto-doc/commit/b750d7a55150cf63905fd20747da03323c6a41fc))  - (Viacheslav Kudinov)

## <!-- 26 -->🔄 Update

- Updated README.md
 ([c9b3014](https://github.com/tj-actions/auto-doc/commit/c9b30141364dde6527a5ba828d6abea05bc00a5d))  - (Tonye Jack)
- Update README.md ([5a8d1f8](https://github.com/tj-actions/auto-doc/commit/5a8d1f80ef72171af12584dd08ef58efd76e246e))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#484](https://github.com/tj-actions/auto-doc/pull/484): to v2.6.1 ([1b233f2](https://github.com/tj-actions/auto-doc/commit/1b233f2e70a16ad49511ba09e2d7bb2ba18475bc))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Upgraded from v2.6.1 -> v2.7.0 ([2a18124](https://github.com/tj-actions/auto-doc/commit/2a181246010fd682c82ce84e80219ca1e0928b4d))  - (github-actions[bot])
- Improve code coverage ([#485](https://github.com/tj-actions/auto-doc/issues/485)) ([33205a6](https://github.com/tj-actions/auto-doc/commit/33205a6530e0fd158dcd95b4f8b18b6ec62c81fb))  - (Tonye Jack)
- Updated coverage badge. ([6b78036](https://github.com/tj-actions/auto-doc/commit/6b78036deb9beb6937bf15318e8d83049190c5c3))  - (GitHub Action)

# [2.6.1](https://github.com/tj-actions/auto-doc/compare/v2.6.0...v2.6.1) - (2023-06-06)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([d9e68b4](https://github.com/tj-actions/auto-doc/commit/d9e68b40fe5ca7f42477c7064a6c51400374b572))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Updated README.
 ([29b19bd](https://github.com/tj-actions/auto-doc/commit/29b19bdad4586166b73e3122a9b01d4b0d3f8524))  - (github-actions[bot])

## <!-- 30 -->📝 Other

- PR [#483](https://github.com/tj-actions/auto-doc/pull/483): update formatting ([a6a6e0a](https://github.com/tj-actions/auto-doc/commit/a6a6e0a3716a767bd87dee3b664bad766f0a66df))  - (repo-ranger[bot])
- Merge baeca55efb26a45771460c86f83853b301f610b7 into a00777cb8d47091e0bb7f2619bce0de160cf203f
 ([b3167e8](https://github.com/tj-actions/auto-doc/commit/b3167e8d62815eada73e6969d37ea6de457dbb70))  - (Tonye Jack)
- Merge branch 'main' into chore/update-formatting ([baeca55](https://github.com/tj-actions/auto-doc/commit/baeca55efb26a45771460c86f83853b301f610b7))  - (repo-ranger[bot])
- PR [#482](https://github.com/tj-actions/auto-doc/pull/482): to v2.6.0 ([a00777c](https://github.com/tj-actions/auto-doc/commit/a00777cb8d47091e0bb7f2619bce0de160cf203f))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Upgraded from v2.6.0 -> v2.6.1 ([e332e0a](https://github.com/tj-actions/auto-doc/commit/e332e0a5dd5d2243c968292decde099c321bb953))  - (github-actions[bot])
- Update formatting ([757551c](https://github.com/tj-actions/auto-doc/commit/757551c3ea61196541919b81c5ad5b2612a6412c))  - (Tonye Jack)

# [2.6.0](https://github.com/tj-actions/auto-doc/compare/v2.5.4...v2.6.0) - (2023-06-06)

## <!-- 0 -->🚀 Features

- Add support for indicating deprecated inputs ([#481](https://github.com/tj-actions/auto-doc/issues/481)) ([1034c3e](https://github.com/tj-actions/auto-doc/commit/1034c3e0222de0e0f7b07a3cb37c87757bdc4631))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Update README.md ([7aae7c3](https://github.com/tj-actions/auto-doc/commit/7aae7c39efa57b5c5a7c105733c3459838b2e191))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#477](https://github.com/tj-actions/auto-doc/pull/477): to v2.5.4 ([8d294a2](https://github.com/tj-actions/auto-doc/commit/8d294a22ce090ec6fec7ec19a17ef8c6cbb3276a))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Upgraded from v2.5.4 -> v2.6.0 ([24706b6](https://github.com/tj-actions/auto-doc/commit/24706b638b1c2c532054bf7dc9aa0d0b7cbd3898))  - (github-actions[bot])
- **deps:** Update tj-actions/branch-names action to v7 ([fe77e45](https://github.com/tj-actions/auto-doc/commit/fe77e45b4b325230d9ca37002ded009bf3127e7c))  - (renovate[bot])

# [2.5.4](https://github.com/tj-actions/auto-doc/compare/v2.5.3...v2.5.4) - (2023-05-12)

## <!-- 26 -->🔄 Update

- Update action.yml ([9b9d172](https://github.com/tj-actions/auto-doc/commit/9b9d1727f337d5594938fd05288e7d4c9c19efd8))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#476](https://github.com/tj-actions/auto-doc/pull/476): to v2.5.3 ([ea29a7f](https://github.com/tj-actions/auto-doc/commit/ea29a7f1fc2974b5765bcc856e9348a86cd84458))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Upgraded from v2.5.3 -> v2.5.4 ([f5d2fe1](https://github.com/tj-actions/auto-doc/commit/f5d2fe14eeb47c2d25098a7863fe47929be76a3d))  - (github-actions[bot])

# [2.5.3](https://github.com/tj-actions/auto-doc/compare/v2.5.0...v2.5.3) - (2023-05-12)

## <!-- 0 -->🚀 Features

- Switch to use setup bin action ([#475](https://github.com/tj-actions/auto-doc/issues/475)) ([a511420](https://github.com/tj-actions/auto-doc/commit/a5114207122941e5d4b631dddad665b3834202e0))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Updated README.
 ([fa3bf22](https://github.com/tj-actions/auto-doc/commit/fa3bf2294de33cdec5726c66e5ed911b7f388283))  - (github-actions[bot])
- Update README.md ([61fc5c0](https://github.com/tj-actions/auto-doc/commit/61fc5c0ab121b39d324fd968371b8bb227f034f8))  - (Tonye Jack)
- Updated comment
 ([3454e51](https://github.com/tj-actions/auto-doc/commit/3454e5186150933e348991cfae8bc35287567c4e))  - (Tonye Jack)
- Update sync-release-version.yml ([b1c32bb](https://github.com/tj-actions/auto-doc/commit/b1c32bbbbb4b65e6d09e50f56e4afac6f85be5e8))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#471](https://github.com/tj-actions/auto-doc/pull/471): to v2.5.0 ([34c0c4c](https://github.com/tj-actions/auto-doc/commit/34c0c4cd3812e8ba86334e43a1d606b4a71d8825))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Upgraded from v2.5.2 -> v2.5.3 ([cf27187](https://github.com/tj-actions/auto-doc/commit/cf27187403627d870cf3d278aa0723bdff3d5830))  - (github-actions[bot])
- Remove debug mode ([0ead27c](https://github.com/tj-actions/auto-doc/commit/0ead27c29f963014ceddd7b065ac764ba6eff7f8))  - (Tonye Jack)
- Skip fields that are within the limit ([#472](https://github.com/tj-actions/auto-doc/issues/472)) ([45a609f](https://github.com/tj-actions/auto-doc/commit/45a609f057e55686834a3b6ec66296d6a8461312))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded to v2.5.2 ([#474](https://github.com/tj-actions/auto-doc/issues/474))

Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com> ([ca5874c](https://github.com/tj-actions/auto-doc/commit/ca5874cea2ac21e059d32ec33ae4ae8085332853))  - (Tonye Jack)
- Upgraded to v2.5.1 ([#473](https://github.com/tj-actions/auto-doc/issues/473))

Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com>
Co-authored-by: repo-ranger[bot] <39074581+repo-ranger[bot]@users.noreply.github.com> ([44f07c4](https://github.com/tj-actions/auto-doc/commit/44f07c4337a00fa5c79aefd75e220fc88e6a6784))  - (Tonye Jack)

# [2.5.0](https://github.com/tj-actions/auto-doc/compare/v2.4.0...v2.5.0) - (2023-05-04)

## <!-- 0 -->🚀 Features

- Improve handling description data ([#469](https://github.com/tj-actions/auto-doc/issues/469)) ([22dba1f](https://github.com/tj-actions/auto-doc/commit/22dba1f9595c76e9ee78ce08c9db57487102a080))  - (Tonye Jack)

## <!-- 1 -->🐛 Bug Fixes

- Update bytes_utils.go ([#468](https://github.com/tj-actions/auto-doc/issues/468)) ([8dc1ee9](https://github.com/tj-actions/auto-doc/commit/8dc1ee9aca0f48b2dbdc22e62a5d7c08aceeb637))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Update README.md ([e07ced3](https://github.com/tj-actions/auto-doc/commit/e07ced3021659416d40248d93e74d7cd531712cd))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#464](https://github.com/tj-actions/auto-doc/pull/464): to v2.4.0 ([4713314](https://github.com/tj-actions/auto-doc/commit/471331484f5b2644eb48965fb40570ec78da8e7a))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Upgraded from v2.4.0 -> v2.5.0 ([cda3a43](https://github.com/tj-actions/auto-doc/commit/cda3a43e7a76e27a7cc9664d70bbd6c895ceffcf))  - (github-actions[bot])
- Rename variable ([#470](https://github.com/tj-actions/auto-doc/issues/470)) ([68d8e50](https://github.com/tj-actions/auto-doc/commit/68d8e503d59f0084b90b16e6a4909a77031af5d2))  - (Tonye Jack)
- **deps:** Update peter-evans/create-pull-request action to v5.0.1 ([3cc77da](https://github.com/tj-actions/auto-doc/commit/3cc77dacdb189d76c1db9f356938fa9e1203c717))  - (renovate[bot])

# [2.4.0](https://github.com/tj-actions/auto-doc/compare/v2.3.2...v2.4.0) - (2023-04-30)

## <!-- 0 -->🚀 Features

- Move to use shellscript ([#463](https://github.com/tj-actions/auto-doc/issues/463)) ([cb8c54a](https://github.com/tj-actions/auto-doc/commit/cb8c54aa3c25598acd294db66c6c5c1c39fe3de9))  - (Tonye Jack)

## <!-- 1 -->🐛 Bug Fixes

- **deps:** Update module github.com/spf13/cobra to v1.7.0 ([4c02a2e](https://github.com/tj-actions/auto-doc/commit/4c02a2eec7c7c543eb10f62aa28fbf9451a7fa67))  - (renovate[bot])

## <!-- 25 -->🎨 Format

- Formatted code.
 ([3c02ecf](https://github.com/tj-actions/auto-doc/commit/3c02ecf5c4b61994cc3e8dff363ea2873720b9ca))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Update action.yml ([1767975](https://github.com/tj-actions/auto-doc/commit/1767975f451d6f68ececcf3681c4dd0d7dfb99df))  - (Tonye Jack)
- Updated README.
 ([2356861](https://github.com/tj-actions/auto-doc/commit/23568619628238cbda843103c69216384b285493))  - (github-actions[bot])
- Update README.md ([87ea759](https://github.com/tj-actions/auto-doc/commit/87ea7594bec12ec5d7d567ea64041091336735a5))  - (Tonye Jack)
- Update README.md ([f9c60d9](https://github.com/tj-actions/auto-doc/commit/f9c60d9405a5f0b9ae71ded7ce3ad11788f75710))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#460](https://github.com/tj-actions/auto-doc/pull/460): update default col max words ([91f5405](https://github.com/tj-actions/auto-doc/commit/91f5405f392ec786d9b9e94b61aedf8dc232eff2))  - (repo-ranger[bot])
- PR [#458](https://github.com/tj-actions/auto-doc/pull/458): to v2.3.2 ([dc98ce6](https://github.com/tj-actions/auto-doc/commit/dc98ce64cfc74609021a4e031b2e8216e1e2a430))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Upgraded from v2.3.2 -> v2.4.0 ([0e1ab9a](https://github.com/tj-actions/auto-doc/commit/0e1ab9a1ca2a9f532049650adb5a26c9ae74f677))  - (github-actions[bot])
- **deps:** Update peter-evans/create-pull-request action to v5 ([#462](https://github.com/tj-actions/auto-doc/issues/462)) ([1bde7a4](https://github.com/tj-actions/auto-doc/commit/1bde7a45ca20516fa3c7cf12267b73a12692fe6b))  - (renovate[bot])
- Update default col max words ([74826ee](https://github.com/tj-actions/auto-doc/commit/74826ee21d6b243e3c09d8cc1089a29e944fcdbc))  - (Tonye Jack)
- Update description ([#459](https://github.com/tj-actions/auto-doc/issues/459)) ([0c1bad3](https://github.com/tj-actions/auto-doc/commit/0c1bad37088c100af42ece5fc76f84bae3d95613))  - (Tonye Jack)

# [2.3.2](https://github.com/tj-actions/auto-doc/compare/v2.3.1...v2.3.2) - (2023-03-28)

## <!-- 26 -->🔄 Update

- Updated README.
 ([9f4c4c2](https://github.com/tj-actions/auto-doc/commit/9f4c4c23745efbe3ffdf35693c54d4df76e67886))  - (github-actions[bot])
- Update test.yml ([6ca8ec7](https://github.com/tj-actions/auto-doc/commit/6ca8ec75eba1aedbfa62fb1d6ad0551a206a26d0))  - (Tonye Jack)
- Update test.yml ([08979cc](https://github.com/tj-actions/auto-doc/commit/08979ccbf0dcad2f761394d3dc187b40468be365))  - (Tonye Jack)
- Update test.yml ([671c770](https://github.com/tj-actions/auto-doc/commit/671c770588deca61b5111a052d9b0fc5a3c35d24))  - (Tonye Jack)
- Update test.yml ([40a00ff](https://github.com/tj-actions/auto-doc/commit/40a00fff229b87a1374854767ea247b5a29fe0bb))  - (Tonye Jack)
- Update test.yml ([b411daf](https://github.com/tj-actions/auto-doc/commit/b411daf8b5f8816b1df309d274371a6af0c90dc4))  - (Tonye Jack)
- Update action.yml ([f4037b2](https://github.com/tj-actions/auto-doc/commit/f4037b2be39b2581e0626b885f207f42662793da))  - (Tonye Jack)
- Update README.md ([d67b0ee](https://github.com/tj-actions/auto-doc/commit/d67b0ee6bf9078018c83c64a5e3d183ead35d007))  - (Tonye Jack)
- Updated docs
 ([425ba4c](https://github.com/tj-actions/auto-doc/commit/425ba4c189c14edd45864b808eddbb227f1be258))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- Don't enclose non-strings with quotes ([#457](https://github.com/tj-actions/auto-doc/issues/457))

fixes #456

Co-authored-by: Tonye Jack <jtonye@ymail.com> ([866f2d6](https://github.com/tj-actions/auto-doc/commit/866f2d643e0e277c0cdadcbc97eb8c7a790e8411))  - (Andreas Åman)
- PR [#455](https://github.com/tj-actions/auto-doc/pull/455): to v2.3.1 ([402b755](https://github.com/tj-actions/auto-doc/commit/402b755a8684a8080f72cbcea305713a2faa6154))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Upgraded from v2.3.1 -> v2.3.2 ([a2bed47](https://github.com/tj-actions/auto-doc/commit/a2bed47b05b8f6f3c8ec24244d15bba48f1c6fad))  - (github-actions[bot])

# [2.3.1](https://github.com/tj-actions/auto-doc/compare/v1.7.4...v2.3.1) - (2023-03-26)

## <!-- 1 -->🐛 Bug Fixes

- Bug with matching only headings that occur at the start ([9c38951](https://github.com/tj-actions/auto-doc/commit/9c3895129d0712c6a142acc2f3237e21bdce3e15))  - (Tonye Jack)
- Fixed bug with download
 ([5c34e2d](https://github.com/tj-actions/auto-doc/commit/5c34e2d25938b317236a217e8a7010aae49f9e6b))  - (Tonye Jack)
- Fixed bug with download
 ([3881916](https://github.com/tj-actions/auto-doc/commit/3881916d9ec2cfe83c0edbff3a356cbc4fdf5c94))  - (Tonye Jack)
- Fixed bug with download
 ([d2931cc](https://github.com/tj-actions/auto-doc/commit/d2931ccb24d2aede4398b1977c9f2ee6da0e5aed))  - (Tonye Jack)
- Fixed bug with download
 ([c898b33](https://github.com/tj-actions/auto-doc/commit/c898b33db02acdc90a56f6618f792a2398bb1f22))  - (Tonye Jack)
- Fixed the job
 ([bcec606](https://github.com/tj-actions/auto-doc/commit/bcec606948b241a1f9870a283254cfc87fedbb8a))  - (Tonye Jack)
- Fixed the job
 ([b6d7a0d](https://github.com/tj-actions/auto-doc/commit/b6d7a0db8bcc3bfd77367bf3b1ed5acea948bde5))  - (Tonye Jack)
- Fixed error with format action
 ([5bd8cb3](https://github.com/tj-actions/auto-doc/commit/5bd8cb3bea17fe8318f8f2fc3cac15cd14f6db6b))  - (Tonye Jack)

## <!-- 11 -->⏪ Reverts

- Revert "Updated the test" ([#412](https://github.com/tj-actions/auto-doc/issues/412))

 ([df4e152](https://github.com/tj-actions/auto-doc/commit/df4e1525cc4226856c68aef4030bfe0947d13460))  - (Tonye Jack)

## <!-- 16 -->➕ Add

- Create ([3f9cb08](https://github.com/tj-actions/auto-doc/commit/3f9cb081f563081691449906c6d10b276cd7ec24))  - (Tonye Jack)
- Added comment and removed unused code
 ([4bfc0b5](https://github.com/tj-actions/auto-doc/commit/4bfc0b57059f6526c73f677ec56effea2c096337))  - (Tonye Jack)

## <!-- 17 -->➖ Remove

- Delete release.yml ([7b1e32f](https://github.com/tj-actions/auto-doc/commit/7b1e32f0ed0ee2c2dd4ef8b06c475d19eb782194))  - (Tonye Jack)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([4734c07](https://github.com/tj-actions/auto-doc/commit/4734c07d34c523b66c274c515b068b2f2fb61838))  - (github-actions[bot])
- Formatted code.
 ([57a2454](https://github.com/tj-actions/auto-doc/commit/57a2454823ebff84d28e9dc68f3c02fb28c2318e))  - (github-actions[bot])
- Formatted code.
 ([0f003ec](https://github.com/tj-actions/auto-doc/commit/0f003ece9b4659cf1309a5e81f3dc534136152ef))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Update README.md ([7a8dec6](https://github.com/tj-actions/auto-doc/commit/7a8dec64094c8b7cdf1ea55b863341b5387a3783))  - (Tonye Jack)
- Update code-coverage.yml ([d99c585](https://github.com/tj-actions/auto-doc/commit/d99c5852dbe58b31cae03cf0ee37c27994fb5c28))  - (Tonye Jack)
- Update code-coverage.yml ([f137b8b](https://github.com/tj-actions/auto-doc/commit/f137b8b25bfc270d302c2c5fb96f390ab245ddee))  - (Tonye Jack)
- Update word_wrap.go ([cb0cef1](https://github.com/tj-actions/auto-doc/commit/cb0cef17c46d3a497eec60249389c5399c28d0be))  - (Tonye Jack)
- Update format_value.go ([d26c19b](https://github.com/tj-actions/auto-doc/commit/d26c19b2531838bc01d265eff1e5112cbda2c372))  - (Tonye Jack)
- Update word_wrap.go ([589c16c](https://github.com/tj-actions/auto-doc/commit/589c16ca262bebf5f8f97fc3c7a8746407017e72))  - (Tonye Jack)
- Update constants.go ([5a7bc16](https://github.com/tj-actions/auto-doc/commit/5a7bc1651fb2707e9a6ea8e1dd6c42390590aa0b))  - (Tonye Jack)
- Update format_value.go ([bf281a6](https://github.com/tj-actions/auto-doc/commit/bf281a6f1b754729d14be59d722f8b30ff7c5827))  - (Tonye Jack)
- Update README.md ([#452](https://github.com/tj-actions/auto-doc/issues/452))

* Update README.md

* Updated README.

---------

Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com> ([1c258ad](https://github.com/tj-actions/auto-doc/commit/1c258ad15e5cfd1b646035fd05fa0785d635f28e))  - (Tonye Jack)
- Update README.md ([d16ef18](https://github.com/tj-actions/auto-doc/commit/d16ef18086b0864f596fb982b1442f519c629a91))  - (Tonye Jack)
- Update README.md ([#451](https://github.com/tj-actions/auto-doc/issues/451))

* Update README.md

* Update test.yml ([63fdd1f](https://github.com/tj-actions/auto-doc/commit/63fdd1fdf232f949f2c3ac71eeea57c7117e34d6))  - (Tonye Jack)
- Update code-coverage.yml ([#449](https://github.com/tj-actions/auto-doc/issues/449))

 ([ba71d0d](https://github.com/tj-actions/auto-doc/commit/ba71d0d34d7933690f19b966079dfc9b761fe75f))  - (Tonye Jack)
- Update test.yml ([#448](https://github.com/tj-actions/auto-doc/issues/448))

 ([bf98416](https://github.com/tj-actions/auto-doc/commit/bf9841692497ae57ddfc8e0749bc57d213760ec3))  - (Tonye Jack)
- Update test.yml ([d371bee](https://github.com/tj-actions/auto-doc/commit/d371bee86d7a9ae52cc4a253a0aeb3d7fa730b38))  - (Tonye Jack)
- Updated README.
 ([7f2e301](https://github.com/tj-actions/auto-doc/commit/7f2e3016665138efd4eb915516244d265315b0f1))  - (github-actions[bot])
- Update README.md ([9986086](https://github.com/tj-actions/auto-doc/commit/9986086a945e66bb41c811d58f53511b370f525a))  - (Tonye Jack)
- Update format-tidy.yml ([#444](https://github.com/tj-actions/auto-doc/issues/444))

* Update format-tidy.yml

* Update test.yml

* Update test.yml

* Update test.yml

* Update test.yml ([c84d8bc](https://github.com/tj-actions/auto-doc/commit/c84d8bc40886f66308678f947f9a99b2cd5039ea))  - (Tonye Jack)
- Update format-tidy.yml ([bfae19d](https://github.com/tj-actions/auto-doc/commit/bfae19d495eb4f4610e0f864052fca883250dc66))  - (Tonye Jack)
- Update action.yml ([b5500ec](https://github.com/tj-actions/auto-doc/commit/b5500ec1dae1c5ccb48039a69c601b2b44e7e939))  - (Tonye Jack)
- Update action.yml ([36f701a](https://github.com/tj-actions/auto-doc/commit/36f701a13db333a792a702b2694af53649e20417))  - (Tonye Jack)
- Update action.yml ([b3bed18](https://github.com/tj-actions/auto-doc/commit/b3bed18d2e426c40bd382731203a2ea9c3232baa))  - (Tonye Jack)
- Update test.yml ([d98f7bc](https://github.com/tj-actions/auto-doc/commit/d98f7bc52c93aa97c7bed0005f0e8ef16205ca27))  - (Tonye Jack)
- Update README.md ([b8b1720](https://github.com/tj-actions/auto-doc/commit/b8b17200c89cb40debe18e74e5bbadc3f130d107))  - (Tonye Jack)
- Update .gitattributes ([d377a3b](https://github.com/tj-actions/auto-doc/commit/d377a3b9904279d67eebd3b839f7ea6c45e72861))  - (Tonye Jack)
- Update README.md ([26e5b4e](https://github.com/tj-actions/auto-doc/commit/26e5b4e505c47ddf95b8adece88566280542d1a4))  - (Tonye Jack)
- Update test.yml ([6e7f851](https://github.com/tj-actions/auto-doc/commit/6e7f8514577587b1a5a39b1ecd4d68813665d777))  - (Tonye Jack)
- Update README.md ([ca7398f](https://github.com/tj-actions/auto-doc/commit/ca7398fc9390df2c46e3a27c208b2707c149059c))  - (Tonye Jack)
- Update README.md ([14c09c8](https://github.com/tj-actions/auto-doc/commit/14c09c8fb11b276270a470758ec7dfd64f98d4fe))  - (Tonye Jack)
- Update action.go ([8a65e3d](https://github.com/tj-actions/auto-doc/commit/8a65e3d542f50de3e9eb1b35eff67745370ac8b7))  - (Tonye Jack)
- Update action.go ([5de73b8](https://github.com/tj-actions/auto-doc/commit/5de73b82cc15faef5a18f6eb6350009c5bec3526))  - (Tonye Jack)
- Update reusable.go ([ba22ce2](https://github.com/tj-actions/auto-doc/commit/ba22ce2307e89e4c8080d7a69894798568e9f000))  - (Tonye Jack)
- Update action.yml ([b050500](https://github.com/tj-actions/auto-doc/commit/b05050034ae1c71d56e2c70ef9c647d73c863b53))  - (Tonye Jack)
- Update sync-release-version.yml ([9d5d9aa](https://github.com/tj-actions/auto-doc/commit/9d5d9aa96ffbba9e4cb27928a647779b0c4ea125))  - (Tonye Jack)
- Update action.yml ([5940acc](https://github.com/tj-actions/auto-doc/commit/5940accd52fd41c9aa6f6cf91fb0cdd2c53e3d8c))  - (Tonye Jack)
- Update sync-release-version.yml ([25ee792](https://github.com/tj-actions/auto-doc/commit/25ee7929787ecca0d5c0a4c5265847bd7fd5d385))  - (Tonye Jack)
- Update sync-release-version.yml ([3401544](https://github.com/tj-actions/auto-doc/commit/3401544116aa0f848879f2e720b02fe275ae0c73))  - (Tonye Jack)
- Update sync-release-version.yml ([151867a](https://github.com/tj-actions/auto-doc/commit/151867ade53ee5c1d81c774f6c91836fa3105dd1))  - (Tonye Jack)
- Update sync-release-version.yml ([b74c84f](https://github.com/tj-actions/auto-doc/commit/b74c84fc9d3f6e96ce10c27995cbe96ca80f0792))  - (Tonye Jack)
- Update action.yml ([1b4e73f](https://github.com/tj-actions/auto-doc/commit/1b4e73f27ccc4ab9b1ab32cce7fc2bc858d5d613))  - (Tonye Jack)
- Updated README.
 ([7ca3edc](https://github.com/tj-actions/auto-doc/commit/7ca3edcaa4b989f88e4dcd47631ced75f063d180))  - (github-actions[bot])
- Update tj-actions/github-changelog-generator action to v1.18 ([7a0db0d](https://github.com/tj-actions/auto-doc/commit/7a0db0dfed93d4d4cffd3b986b14700d886c864f))  - (renovate[bot])
- Update README.md ([b2e8be8](https://github.com/tj-actions/auto-doc/commit/b2e8be8fc4500b0499197b22c4fddf630e09dc92))  - (Tonye Jack)
- Update README.md ([fe3873f](https://github.com/tj-actions/auto-doc/commit/fe3873f8471ed65031f7a3a0cb71e8fbd099a429))  - (Tonye Jack)
- Update README.md ([d7c9d5f](https://github.com/tj-actions/auto-doc/commit/d7c9d5f8ede9f4feedffcd4e20c4d0b5d8ffd6ba))  - (Tonye Jack)
- Update action.yml ([36e6ebf](https://github.com/tj-actions/auto-doc/commit/36e6ebfbe264506656368772018fe90d3d26ddec))  - (Tonye Jack)
- Update action.yml ([8c2a6e0](https://github.com/tj-actions/auto-doc/commit/8c2a6e012c68bd9f4d2b470a0e34d6823a531bec))  - (Tonye Jack)
- Update .goreleaser.yaml ([bd95368](https://github.com/tj-actions/auto-doc/commit/bd95368d59096ff8a3f3df765e4b377394acec7d))  - (Tonye Jack)
- Update release.yml ([bd6ffa7](https://github.com/tj-actions/auto-doc/commit/bd6ffa777642586fc96f13fd478e4d2ed12ab263))  - (Tonye Jack)
- Update .goreleaser.yaml ([15f623e](https://github.com/tj-actions/auto-doc/commit/15f623ed9c91c6fc4e35eafa980960af7395f3d4))  - (Tonye Jack)
- Update README.md ([d692181](https://github.com/tj-actions/auto-doc/commit/d6921819d3fe168a854ec92926148ab3ad2ada93))  - (Tonye Jack)
- Update action.yml ([d8b8758](https://github.com/tj-actions/auto-doc/commit/d8b87587bf226da9367cf5f26da5d6adb3211f52))  - (Tonye Jack)
- Updated help message
 ([dc750b0](https://github.com/tj-actions/auto-doc/commit/dc750b01ddf583f5638c3014daa6304ad3966c4b))  - (Tonye Jack)
- Update action.yml ([b196819](https://github.com/tj-actions/auto-doc/commit/b1968194255b14540a12eed58ef0c65be5611eda))  - (Tonye Jack)
- Updated help message
 ([3a073c3](https://github.com/tj-actions/auto-doc/commit/3a073c39fd5c6422fae76205080c34bbc19f8e71))  - (Tonye Jack)
- Updated the version
 ([a538967](https://github.com/tj-actions/auto-doc/commit/a538967cdcbebbc6570ed64244515ef3d26159cd))  - (Tonye Jack)
- Update spacing
 ([2ff7de7](https://github.com/tj-actions/auto-doc/commit/2ff7de77ea9dd7252a548445a5c6d95d05dd8c4c))  - (Tonye Jack)
- Updated README.
 ([2452559](https://github.com/tj-actions/auto-doc/commit/2452559e35af778aa4758fda26f5c4e5e36b6023))  - (github-actions[bot])
- Updated test
 ([56de299](https://github.com/tj-actions/auto-doc/commit/56de299f8e555c82fe015611d83b620179f1e9bb))  - (Tonye Jack)
- Updated the test ([#410](https://github.com/tj-actions/auto-doc/issues/410))

 ([675e517](https://github.com/tj-actions/auto-doc/commit/675e51728511e6c13d3b325bd90f030c24cdb0e4))  - (Tonye Jack)
- Updated the test
 ([6c3a68a](https://github.com/tj-actions/auto-doc/commit/6c3a68a2ae1831375f25f6797cb5e7837cbf68f2))  - (Tonye Jack)
- Update README.md ([fd86fa3](https://github.com/tj-actions/auto-doc/commit/fd86fa3c6a7d50fd7cd5e8fa1131e76a9b4ca7d0))  - (Tonye Jack)
- Updated README.
 ([96f6a43](https://github.com/tj-actions/auto-doc/commit/96f6a43e3375ca1bf82f503b2800094a178b0a07))  - (github-actions[bot])
- Update README.md ([cd60cdd](https://github.com/tj-actions/auto-doc/commit/cd60cddc3953728a1d13ddd36d30f414e8666951))  - (Tonye Jack)
- Updated README.
 ([6139819](https://github.com/tj-actions/auto-doc/commit/61398195c2ed552a5432206b7f671c63df81a977))  - (github-actions[bot])
- Update README.md ([21991a8](https://github.com/tj-actions/auto-doc/commit/21991a8a4dcb590eea538d246eabb5c4b40d1032))  - (Tonye Jack)

## <!-- 3 -->📚 Documentation

- Update .all-contributorsrc [skip ci] ([43adee6](https://github.com/tj-actions/auto-doc/commit/43adee6dc0cdeb88f4959eeae95f1860355e8443))  - (allcontributors[bot])
- Update README.md [skip ci] ([91d3869](https://github.com/tj-actions/auto-doc/commit/91d3869fb322a4bca70bc7a0c8a85be0cbe7a402))  - (allcontributors[bot])
- Add andreas-aman as a contributor for code, and doc ([#408](https://github.com/tj-actions/auto-doc/issues/408)) ([6521c53](https://github.com/tj-actions/auto-doc/commit/6521c53eb17d3bc14305d15ba3b0c1c260aff610))  - (allcontributors[bot])

## <!-- 30 -->📝 Other

- PR [#453](https://github.com/tj-actions/auto-doc/pull/453): to v2.3.0 ([0b77703](https://github.com/tj-actions/auto-doc/commit/0b777039f0311ea71d16b4e18db0f674ce030b2e))  - (repo-ranger[bot])
- [Bugfix] Type of inputs are always string fix ([#447](https://github.com/tj-actions/auto-doc/issues/447))

* Fixed input type of reusable workflow

* Fixed typo

* Updated README.

* Clean up

* Without formatting

* Without formatting

* Back yaml

* Formatted code.

* Updated README.

* restored original reame:s

* restored original reame:s

* Update .github/workflows/test.yml

Co-authored-by: Tonye Jack <jtonye@ymail.com>

* Update reusable.go

* Formatted code.

* Update reusable.go

* Formatted code.

---------

Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com>
Co-authored-by: Tonye Jack <jtonye@ymail.com>
Co-authored-by: repo-ranger[bot] <39074581+repo-ranger[bot]@users.noreply.github.com> ([92db81e](https://github.com/tj-actions/auto-doc/commit/92db81e6e0da47e0c3481c360cb92c255b3a961a))  - (Viacheslav Kudinov)
- PR [#450](https://github.com/tj-actions/auto-doc/pull/450): add ViacheslavKudinov as a contributor for code, and test ([703a2b8](https://github.com/tj-actions/auto-doc/commit/703a2b892a7b25c055df239c11c9d5ba1f9e00a3))  - (repo-ranger[bot])
- PR [#445](https://github.com/tj-actions/auto-doc/pull/445): README.md ([a2cf565](https://github.com/tj-actions/auto-doc/commit/a2cf565ab25a7f95b67b3d82b51f601c10045fd5))  - (repo-ranger[bot])
- PR [#442](https://github.com/tj-actions/auto-doc/pull/442): to v2.2.6 ([b4e9b8f](https://github.com/tj-actions/auto-doc/commit/b4e9b8f9474b65692cf60d443862efa3356253dd))  - (repo-ranger[bot])
- PR [#441](https://github.com/tj-actions/auto-doc/pull/441): to v2.2.5 ([77ba9ed](https://github.com/tj-actions/auto-doc/commit/77ba9ed2123d5c31be40988ea8dda87fe955192b))  - (repo-ranger[bot])
- Merge branch 'main' into upgrade-to-v2.2.5 ([7c337a5](https://github.com/tj-actions/auto-doc/commit/7c337a5bbf276e3362f48cc2d452b31ade527cc9))  - (repo-ranger[bot])
- PR [#437](https://github.com/tj-actions/auto-doc/pull/437): switch back to gobinaries ([fc3a280](https://github.com/tj-actions/auto-doc/commit/fc3a280666cce2e6148293b8faebb95e10d9202d))  - (repo-ranger[bot])
- PR [#436](https://github.com/tj-actions/auto-doc/pull/436): to v2.2.4 ([156141c](https://github.com/tj-actions/auto-doc/commit/156141c018163487d0b225ecf7d88fb63d932b68))  - (repo-ranger[bot])
- PR [#432](https://github.com/tj-actions/auto-doc/pull/432): bug with matching only headings that occur at the start ([6874954](https://github.com/tj-actions/auto-doc/commit/6874954a61f4dfdff865a1015ad08e0196b8bf07))  - (repo-ranger[bot])
- Merge branch 'main' into fix/bug-with-matching-only-headings-that-occur-at-the-start ([d7d7fec](https://github.com/tj-actions/auto-doc/commit/d7d7fec895d8e2e4d0e807104ba413095484a207))  - (Tonye Jack)
- Merge 9c3895129d0712c6a142acc2f3237e21bdce3e15 into 1e4f6a0f1df805f378dc2c76aa14e744cc16e8d3
 ([8a22e51](https://github.com/tj-actions/auto-doc/commit/8a22e51077ee3c1347e25902ffe4875e90d4b51b))  - (Tonye Jack)
- PR [#428](https://github.com/tj-actions/auto-doc/pull/428): headers instead of any matching text ([1e4f6a0](https://github.com/tj-actions/auto-doc/commit/1e4f6a0f1df805f378dc2c76aa14e744cc16e8d3))  - (repo-ranger[bot])
- Merge branch 'main' into replace_headers ([8169a71](https://github.com/tj-actions/auto-doc/commit/8169a71343e9cac7e6a7d111730da94b06a547e6))  - (Tonye Jack)
- Use correct input for first replace
 ([5f93bfb](https://github.com/tj-actions/auto-doc/commit/5f93bfbd15a2de1d47f95825bd734afd1759b218))  - (Andreas Åman)
- Only replace headers in beginning of line

fixes #427
 ([ae90733](https://github.com/tj-actions/auto-doc/commit/ae90733450c1767a2ff9a1e5d761da07b80d1dff))  - (Andreas Åman)
- PR [#425](https://github.com/tj-actions/auto-doc/pull/425): to v2.2.3 ([8d0b00e](https://github.com/tj-actions/auto-doc/commit/8d0b00ef3058405a643685e10d8e1c3383959527))  - (repo-ranger[bot])
- Don't add 2 newlines if nothing in table ([#423](https://github.com/tj-actions/auto-doc/issues/423))

Co-authored-by: Tonye Jack <jtonye@ymail.com>
 ([571f879](https://github.com/tj-actions/auto-doc/commit/571f879016160aa495e695a3128c5ec41fc5c40a))  - (Andreas Åman)
- PR [#421](https://github.com/tj-actions/auto-doc/pull/421): to v2.2.0 ([b21bdfa](https://github.com/tj-actions/auto-doc/commit/b21bdfa3ded1cf8e53bf06e8eb8e684042808903))  - (repo-ranger[bot])
- PR [#419](https://github.com/tj-actions/auto-doc/pull/419): to v2.0.3 ([f12956d](https://github.com/tj-actions/auto-doc/commit/f12956d563fefb8ad2c7e9b6aee4a135cc1704d1))  - (repo-ranger[bot])
- [skip ci]: Update version ([517eef5](https://github.com/tj-actions/auto-doc/commit/517eef55df6ea0745ab6fb2583c92d6f290b9af0))  - (Tonye Jack)
- PR [#413](https://github.com/tj-actions/auto-doc/pull/413): the download ([0df34ae](https://github.com/tj-actions/auto-doc/commit/0df34aee50fa7565007545c5a502a2d9c30c53d7))  - (repo-ranger[bot])
- Enable debug mode
 ([04ddf38](https://github.com/tj-actions/auto-doc/commit/04ddf3821ac63eb5e49fa8bd47e583c90d2fbd4a))  - (Tonye Jack)
- Enable debug mode
 ([36b102a](https://github.com/tj-actions/auto-doc/commit/36b102a3b878f19ebfb8eb050de22b9c2613e7a1))  - (Tonye Jack)
- Merge b6d7a0db8bcc3bfd77367bf3b1ed5acea948bde5 into 17f8357497eeb3b70ffeaa37acae64a4a67f17b9
 ([15a0102](https://github.com/tj-actions/auto-doc/commit/15a010267d0e64ab848393446a1cf6d6b1559462))  - (Tonye Jack)
- PR [#409](https://github.com/tj-actions/auto-doc/pull/409): fix add missing comments ([1cd7763](https://github.com/tj-actions/auto-doc/commit/1cd776313d1e3634954a99767b3682a5e6b1df69))  - (repo-ranger[bot])
- Merge 93c34cbb978f72ea4875d832506721a4288a1370 into fd86fa3c6a7d50fd7cd5e8fa1131e76a9b4ca7d0
 ([39668ba](https://github.com/tj-actions/auto-doc/commit/39668ba43bcc979602b701599b7f66332b5c12df))  - (Tonye Jack)
- Merge branch 'main' into chore/fix-add-missing-comments ([93c34cb](https://github.com/tj-actions/auto-doc/commit/93c34cbb978f72ea4875d832506721a4288a1370))  - (Tonye Jack)
- Merge 0f003ece9b4659cf1309a5e81f3dc534136152ef into cd60cddc3953728a1d13ddd36d30f414e8666951
 ([84b9cab](https://github.com/tj-actions/auto-doc/commit/84b9cab8cfc9d08829fc3863c8d2b1a4fb9b031b))  - (Tonye Jack)
- Merge f9f7231fcb22ec1c91b60a86c9afdb3b128ca50e into cd60cddc3953728a1d13ddd36d30f414e8666951
 ([31eb041](https://github.com/tj-actions/auto-doc/commit/31eb041a79bc2bcab26783b35a56e4324cb07d93))  - (Tonye Jack)
- Merge b84d4ecca608ab6123149ad0b62dcb58b01809ee into cd60cddc3953728a1d13ddd36d30f414e8666951
 ([1fa788b](https://github.com/tj-actions/auto-doc/commit/1fa788b417390f6a4e5b4317dce6752761ccf2a7))  - (Tonye Jack)
- Support reusable workflows ([#405](https://github.com/tj-actions/auto-doc/issues/405))

* First draft for supporting reusable workflows

* Check for error

* Update tests

* Update action details

* Refactoring, and actually saving result to README.md

* Update README.md

Co-authored-by: Tonye Jack <jtonye@ymail.com>

* Update action.yml

Co-authored-by: Tonye Jack <jtonye@ymail.com>

* Update action.yml

Co-authored-by: Tonye Jack <jtonye@ymail.com>

* Update action.yml

Co-authored-by: Tonye Jack <jtonye@ymail.com>

* Remove unused variable

* Add support for output in reusable workflows

* Document ReusaböeOutput

* Add error handling for function calls

* Code cleanup and added test

* Add upstream change

* Update the README.md

* Updated the test

* Updated the test

* Updated the test

* Updated the test

* Fixed error with test

* Fixed error with test

* Fixed error with test

* Fixed error with test

* Fixed the test

* Updated line ending

---------

Co-authored-by: Tonye Jack <jtonye@ymail.com> ([1b0b9eb](https://github.com/tj-actions/auto-doc/commit/1b0b9eb26d8d887368f8d48ea886cbb8207b7aef))  - (Andreas Åman)
- PR [#403](https://github.com/tj-actions/auto-doc/pull/403): to v1.7.4 ([d081905](https://github.com/tj-actions/auto-doc/commit/d081905b7bd7423c223412bc4c053c1ea7bca645))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Upgraded from v2.3.0 -> v2.3.1 ([9f9fa19](https://github.com/tj-actions/auto-doc/commit/9f9fa193dbd4e7a4248b32fe57879d80edea7e38))  - (github-actions[bot])
- Upgrade to v2 ([#454](https://github.com/tj-actions/auto-doc/issues/454)) ([d5c6879](https://github.com/tj-actions/auto-doc/commit/d5c6879be8337d44e2658300dfd230ffbc473959))  - (Tonye Jack)
- Upgraded from v2.2.7 -> v2.3.0 ([12d1707](https://github.com/tj-actions/auto-doc/commit/12d170709a097462ce9979c985b0c42650ee5336))  - (github-actions[bot])
- Upgraded from v2.2.6 -> v2.2.7 ([#446](https://github.com/tj-actions/auto-doc/issues/446)) ([7219fcb](https://github.com/tj-actions/auto-doc/commit/7219fcbdac9e3948b964c3484507c4bcc0094239))  - (Tonye Jack)
- Update description of types ([#443](https://github.com/tj-actions/auto-doc/issues/443)) ([d8b74fb](https://github.com/tj-actions/auto-doc/commit/d8b74fbd9fcb7cf12071d12b3f60bed4c2956bd4))  - (Tonye Jack)
- Upgraded from v2.2.5 -> v2.2.6 ([6bb919c](https://github.com/tj-actions/auto-doc/commit/6bb919c7205768c3dee117b8f9b41752888e2ac2))  - (github-actions[bot])
- Update README.md ([#440](https://github.com/tj-actions/auto-doc/issues/440)) ([eaddd66](https://github.com/tj-actions/auto-doc/commit/eaddd6638d3d2ec98f92d0592df8320cd84bc1ff))  - (Tonye Jack)
- Upgraded from v2.2.4 -> v2.2.5 ([35c4850](https://github.com/tj-actions/auto-doc/commit/35c4850f665ce760be1ea33ef3c2f2fbe292f5dd))  - (github-actions[bot])
- Update README.md ([#439](https://github.com/tj-actions/auto-doc/issues/439)) ([ce815cc](https://github.com/tj-actions/auto-doc/commit/ce815cc9483b57410994cacf25e9f34a4111eef0))  - (Tonye Jack)
- Upgraded from v2.2.4 -> v2.2.5 ([#438](https://github.com/tj-actions/auto-doc/issues/438)) ([b042212](https://github.com/tj-actions/auto-doc/commit/b042212cbdddc9e4c8890d8c79996f9149008ce8))  - (Tonye Jack)
- Switch back to gobinaries ([f3dc2c1](https://github.com/tj-actions/auto-doc/commit/f3dc2c1c74269d10ecdc86a6f0c513f4009922c6))  - (Tonye Jack)
- Update README.md ([#434](https://github.com/tj-actions/auto-doc/issues/434)) ([4b8a2dd](https://github.com/tj-actions/auto-doc/commit/4b8a2dd0af7c3f98917711471da935d406df8bf0))  - (Tonye Jack)
- Upgraded from v2.2.3 -> v2.2.4 ([8459b25](https://github.com/tj-actions/auto-doc/commit/8459b25de3932169e8413179409af86ca640be4b))  - (github-actions[bot])
- **deps:** Update tj-actions/verify-changed-files action to v14 ([c47839d](https://github.com/tj-actions/auto-doc/commit/c47839db4e95e2a24b55485c95b2e00b090edd3e))  - (renovate[bot])
- Update README.md ([#433](https://github.com/tj-actions/auto-doc/issues/433)) ([0792a5e](https://github.com/tj-actions/auto-doc/commit/0792a5e8a2f0737a56ad704436cc8143a903bb17))  - (Tonye Jack)
- **deps:** Update actions/setup-go action to v4 ([4a02870](https://github.com/tj-actions/auto-doc/commit/4a02870ff9497c527b949d63a4f16d37bb611bbc))  - (renovate[bot])
- **deps:** Update peter-evans/create-pull-request action to v4.2.4 ([697628d](https://github.com/tj-actions/auto-doc/commit/697628dfc8e0920a0e40301ad19655e17d12ef1f))  - (renovate[bot])
- Upgraded from v2.2.2 -> v2.2.3 ([8108f27](https://github.com/tj-actions/auto-doc/commit/8108f272ddcbe6ddd6ad0df6f2c962d9a3b67172))  - (github-actions[bot])
- Tidy imports ([#424](https://github.com/tj-actions/auto-doc/issues/424)) ([7f48c5f](https://github.com/tj-actions/auto-doc/commit/7f48c5fb829c8ad25f48a8d2254408982ca2e343))  - (Tonye Jack)
- Chore/hard code version ([#417](https://github.com/tj-actions/auto-doc/issues/417))

 ([5365de3](https://github.com/tj-actions/auto-doc/commit/5365de316efca8cb29099fda0e839ccbcd0c23db))  - (Tonye Jack)
- Updated coverage badge. ([f9f7231](https://github.com/tj-actions/auto-doc/commit/f9f7231fcb22ec1c91b60a86c9afdb3b128ca50e))  - (GitHub Action)
- Fix add missing comments ([b84d4ec](https://github.com/tj-actions/auto-doc/commit/b84d4ecca608ab6123149ad0b62dcb58b01809ee))  - (Tonye Jack)
- **deps:** Update codacy/codacy-analysis-cli-action action to v4.3.0 ([9b2569c](https://github.com/tj-actions/auto-doc/commit/9b2569c072851445bb2f3ee6857f0ca451909ed9))  - (renovate[bot])
- **deps:** Update pascalgn/automerge-action action to v0.15.6 ([af0d6b6](https://github.com/tj-actions/auto-doc/commit/af0d6b6fa971c67f864c01594a2ef0c718096d7e))  - (renovate[bot])

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v2.1.0 -> v2.2.0
 ([38dd389](https://github.com/tj-actions/auto-doc/commit/38dd3891dc5f9f9c21fa1ad690d6f3e2fa3984c3))  - (jackton1)
- Upgraded to v2.1.0 ([#420](https://github.com/tj-actions/auto-doc/issues/420))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com> ([5ae001d](https://github.com/tj-actions/auto-doc/commit/5ae001d713d115385ca7f580c029a6c068a0d358))  - (Tonye Jack)
- Upgraded from v2.0.2 -> v2.0.3
 ([3bc4775](https://github.com/tj-actions/auto-doc/commit/3bc47758a74c544c86b7df9323dc409d75bec0ea))  - (jackton1)
- Upgraded to v2.0.2 ([#418](https://github.com/tj-actions/auto-doc/issues/418))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com> ([1509ffc](https://github.com/tj-actions/auto-doc/commit/1509ffc76b868aae4bae3c5b91595ce0b4376888))  - (Tonye Jack)
- Upgraded to v2.0.1 ([#414](https://github.com/tj-actions/auto-doc/issues/414))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com> ([a0e52cd](https://github.com/tj-actions/auto-doc/commit/a0e52cd596a9b1356d2c25fb3c3bef57363dac19))  - (Tonye Jack)
- Upgraded to v2 ([#411](https://github.com/tj-actions/auto-doc/issues/411))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com>
Co-authored-by: repo-ranger[bot] <39074581+repo-ranger[bot]@users.noreply.github.com> ([17f8357](https://github.com/tj-actions/auto-doc/commit/17f8357497eeb3b70ffeaa37acae64a4a67f17b9))  - (Tonye Jack)
- Upgraded from v1.7.3 -> v1.7.4
 ([cef5715](https://github.com/tj-actions/auto-doc/commit/cef5715537b65df237ddb2a3e5a14299d806a1cc))  - (jackton1)

# [1.7.4](https://github.com/tj-actions/auto-doc/compare/v1.7.3...v1.7.4) - (2023-01-22)

## <!-- 26 -->🔄 Update

- Update action.yml ([e7f517d](https://github.com/tj-actions/auto-doc/commit/e7f517da90caf0dc74db319dd7769c7f87613257))  - (Tonye Jack)
- Update README.md ([8007fbf](https://github.com/tj-actions/auto-doc/commit/8007fbf34c39dad5a7bc39a7951d201e0815a91a))  - (Tonye Jack)
- Update code-coverage.yml ([2680b3c](https://github.com/tj-actions/auto-doc/commit/2680b3c32250f06197bdef743354bb59427ef4bf))  - (Tonye Jack)
- Update test.yml ([5498690](https://github.com/tj-actions/auto-doc/commit/5498690582c0fef5de021f49582d00b4a20224e5))  - (Tonye Jack)
- Update test.yml ([ab547f8](https://github.com/tj-actions/auto-doc/commit/ab547f893198ded8c83b3bebae0dad589441da9e))  - (Tonye Jack)
- Updated comment
 ([b1b3685](https://github.com/tj-actions/auto-doc/commit/b1b36858156ea9b2bd91f64270972ed09c9f9495))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#402](https://github.com/tj-actions/auto-doc/pull/402): to v1.7.3 ([b453caf](https://github.com/tj-actions/auto-doc/commit/b453caf25b733d0f2016c041bdf7ed3dde671ebc))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Upload coverage reports to codacy ([#366](https://github.com/tj-actions/auto-doc/issues/366)) ([963c736](https://github.com/tj-actions/auto-doc/commit/963c736451e7dc0664378da0e5cecc185a60662d))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.7.2 -> v1.7.3
 ([3b1c8e3](https://github.com/tj-actions/auto-doc/commit/3b1c8e3784d9c94183f29bc767448ba39b913008))  - (jackton1)

# [1.7.3](https://github.com/tj-actions/auto-doc/compare/v1.7.2...v1.7.3) - (2023-01-03)

## <!-- 26 -->🔄 Update

- Update README.md ([d467978](https://github.com/tj-actions/auto-doc/commit/d467978d41f0860d31beaaf0acc8a19b7ff5a740))  - (Tonye Jack)
- Update action.yml ([9d42253](https://github.com/tj-actions/auto-doc/commit/9d4225399cc1f732fba9770a12a8d28ebc0aaadd))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#400](https://github.com/tj-actions/auto-doc/pull/400): to v1.7.2 ([ac17a9b](https://github.com/tj-actions/auto-doc/commit/ac17a9bf882c6e64cffbcfa0d1eb9055395027fa))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Update readme ([#401](https://github.com/tj-actions/auto-doc/issues/401)) ([987c4ed](https://github.com/tj-actions/auto-doc/commit/987c4ed540caeb4273765dc4504e79ecfd8e6203))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.7.1 -> v1.7.2
 ([27fe883](https://github.com/tj-actions/auto-doc/commit/27fe8834a51667cca80d3977de2b5b341e96f77f))  - (jackton1)

# [1.7.2](https://github.com/tj-actions/auto-doc/compare/v1.7.1...v1.7.2) - (2022-12-31)

## <!-- 30 -->📝 Other

- PR [#398](https://github.com/tj-actions/auto-doc/pull/398): to v1.7.1 ([a231339](https://github.com/tj-actions/auto-doc/commit/a2313397c9280175a8a4980d8fe999f1cda36302))  - (repo-ranger[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Update to remove trailing br tags ([#399](https://github.com/tj-actions/auto-doc/issues/399)) ([07fd015](https://github.com/tj-actions/auto-doc/commit/07fd015d86f1426e81ab57b73a0e04416d04e32b))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.7.0 -> v1.7.1
 ([1d9d667](https://github.com/tj-actions/auto-doc/commit/1d9d6674b2dffd707e011558de4473b368c9181e))  - (jackton1)

# [1.7.1](https://github.com/tj-actions/auto-doc/compare/v1.7.0...v1.7.1) - (2022-12-30)

## <!-- 1 -->🐛 Bug Fixes

- Bug with spacing ([#397](https://github.com/tj-actions/auto-doc/issues/397)) ([1fe62c4](https://github.com/tj-actions/auto-doc/commit/1fe62c41005139bf353f999bbf44e18fb2b1c64b))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#396](https://github.com/tj-actions/auto-doc/pull/396): to v1.7.0 ([51fd40f](https://github.com/tj-actions/auto-doc/commit/51fd40f54709a72280f341f50672ebe56ffed867))  - (repo-ranger[bot])

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.6.0 -> v1.7.0
 ([73aa6af](https://github.com/tj-actions/auto-doc/commit/73aa6af0dbb6afbf22abb0765623a4f4a08ed4d9))  - (jackton1)

# [1.7.0](https://github.com/tj-actions/auto-doc/compare/v1.6.0...v1.7.0) - (2022-12-30)

## <!-- 17 -->➖ Remove

- Delete update-readme.yml ([6c24266](https://github.com/tj-actions/auto-doc/commit/6c242669026f2a9abc2b312e4a890b8d88611647))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Update action.yml ([cb49714](https://github.com/tj-actions/auto-doc/commit/cb497147249dbf049366c11aae4d3942b52e9a9e))  - (Tonye Jack)
- Updated README.
 ([6b40971](https://github.com/tj-actions/auto-doc/commit/6b4097176de4de71c4b75070107c919237086ef7))  - (github-actions[bot])
- Update action.yml ([dde2d06](https://github.com/tj-actions/auto-doc/commit/dde2d067586cb4cb137c882f0a6897a3c05f4028))  - (Tonye Jack)
- Updated README.md ([#388](https://github.com/tj-actions/auto-doc/issues/388))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com>
Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com> ([1c5df2f](https://github.com/tj-actions/auto-doc/commit/1c5df2f422e3f35783e11b7d10d59379930b76ed))  - (Tonye Jack)
- Updated README.md ([#387](https://github.com/tj-actions/auto-doc/issues/387))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com>
Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com> ([bd3389d](https://github.com/tj-actions/auto-doc/commit/bd3389d3154166c338c4ffc85f425b2aa576d9f5))  - (Tonye Jack)
- Updated README.md ([#386](https://github.com/tj-actions/auto-doc/issues/386))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com>
Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com> ([34ebed9](https://github.com/tj-actions/auto-doc/commit/34ebed99e73d133026831fa5ef90aa092ea3361a))  - (Tonye Jack)
- Updated README.md ([#385](https://github.com/tj-actions/auto-doc/issues/385))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com>
Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com> ([e0118be](https://github.com/tj-actions/auto-doc/commit/e0118be485940599a89643ee731298eca0c90781))  - (Tonye Jack)
- Updated README.md ([#384](https://github.com/tj-actions/auto-doc/issues/384))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com>
Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com> ([b41ebfd](https://github.com/tj-actions/auto-doc/commit/b41ebfde335a8226f633fe833a4879170230051b))  - (Tonye Jack)
- Updated README.md ([#382](https://github.com/tj-actions/auto-doc/issues/382))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com> ([02cefcd](https://github.com/tj-actions/auto-doc/commit/02cefcd3e0cde260808df61bd3c804a29c5103c9))  - (Tonye Jack)
- Update tj-actions/verify-changed-files action to v13 ([3ce246c](https://github.com/tj-actions/auto-doc/commit/3ce246c745b7228418712198949d4fb9503d235d))  - (renovate[bot])
- Update tj-actions/auto-doc action to v1.6.0 ([b50a586](https://github.com/tj-actions/auto-doc/commit/b50a586bfca7db90afb2acac92f053dc46cd5dbe))  - (renovate[bot])

## <!-- 30 -->📝 Other

- PR [#390](https://github.com/tj-actions/auto-doc/pull/390): update readme ([73af155](https://github.com/tj-actions/auto-doc/commit/73af15504ea06e9ce19f59b6a888ec021507cfb8))  - (repo-ranger[bot])
- Merge 07abd196b2751839fbee016970ca958f49622fbf into dde2d067586cb4cb137c882f0a6897a3c05f4028
 ([7f53368](https://github.com/tj-actions/auto-doc/commit/7f53368e3e7091f96f34b148b7f98086d706ba1e))  - (Tonye Jack)

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Chore/update default ([#395](https://github.com/tj-actions/auto-doc/issues/395))

Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com> ([c4b295d](https://github.com/tj-actions/auto-doc/commit/c4b295db10d47f723654962a49dd7b568a849c72))  - (Tonye Jack)
- Update the default col_max_words ([#393](https://github.com/tj-actions/auto-doc/issues/393)) ([cd20d96](https://github.com/tj-actions/auto-doc/commit/cd20d96e5e4eab6ce9ffa5f895e3549a160f787a))  - (Tonye Jack)
- Add support to skip adding br tags to links or code ([#394](https://github.com/tj-actions/auto-doc/issues/394)) ([1926ef3](https://github.com/tj-actions/auto-doc/commit/1926ef382ff26bcb286fedf7803c73859a58e333))  - (Tonye Jack)
- Add support to skip adding br tags to links or code ([#392](https://github.com/tj-actions/auto-doc/issues/392)) ([c17c3df](https://github.com/tj-actions/auto-doc/commit/c17c3df854d69a481d9a8ee9f2d1bad4df18c990))  - (Tonye Jack)
- Update readme ([07abd19](https://github.com/tj-actions/auto-doc/commit/07abd196b2751839fbee016970ca958f49622fbf))  - (Tonye Jack)
- **deps:** Update tj-actions/github-changelog-generator action to v1.17 ([d50b5f3](https://github.com/tj-actions/auto-doc/commit/d50b5f39489a3ff23aa4f3a782ee459c7a71b758))  - (renovate[bot])

## <!-- 9 -->⬆️ Upgrades

- Upgraded to v1.6.0 ([#379](https://github.com/tj-actions/auto-doc/issues/379))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com> ([408f344](https://github.com/tj-actions/auto-doc/commit/408f3446ece0aba4eb2ae1b936ecbd988facfdbb))  - (Tonye Jack)

# [1.6.0](https://github.com/tj-actions/auto-doc/compare/v1.5.0...v1.6.0) - (2022-12-15)

## <!-- 26 -->🔄 Update

- Update goreleaser/goreleaser-action action to v4 ([0329f01](https://github.com/tj-actions/auto-doc/commit/0329f01f25582ed544ba264b3dab6bde22ef2e22))  - (renovate[bot])
- Update actions/checkout action to v3.2.0 ([c535e45](https://github.com/tj-actions/auto-doc/commit/c535e4567dee87fd65d8884fad8d2fe6cbf34c57))  - (renovate[bot])
- Update sync-release-version.yml ([0c6e00a](https://github.com/tj-actions/auto-doc/commit/0c6e00a2f0c03d188119b44a6f261a389afd250a))  - (Tonye Jack)
- Update tj-actions/auto-doc action to v1.5.0 ([b4ef1e9](https://github.com/tj-actions/auto-doc/commit/b4ef1e932ed8d5b6506da94946e37dccf15688b6))  - (renovate[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Update the default col max words ([#378](https://github.com/tj-actions/auto-doc/issues/378)) ([4749e10](https://github.com/tj-actions/auto-doc/commit/4749e10abe9b3d2a5b3996c788c35c1652b42d5f))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.4.3 -> v1.5.0 ([#374](https://github.com/tj-actions/auto-doc/issues/374))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com> ([da02d78](https://github.com/tj-actions/auto-doc/commit/da02d78a84d27f7f56a07917b354c506a367cad9))  - (Tonye Jack)

# [1.5.0](https://github.com/tj-actions/auto-doc/compare/v1.4.3...v1.5.0) - (2022-12-03)

## <!-- 0 -->🚀 Features

- Switch to use a temp dir ([#372](https://github.com/tj-actions/auto-doc/issues/372)) ([8bad465](https://github.com/tj-actions/auto-doc/commit/8bad46507680cf653b0690be4f21d9adcd1885e9))  - (Tonye Jack)

## <!-- 13 -->📦 Bumps

- Bump actions/cache from 2 to 3

Bumps [actions/cache](https://github.com/actions/cache) from 2 to 3.
- [Release notes](https://github.com/actions/cache/releases)
- [Changelog](https://github.com/actions/cache/blob/main/RELEASES.md)
- [Commits](https://github.com/actions/cache/compare/v2...v3)

---
updated-dependencies:
- dependency-name: actions/cache
  dependency-type: direct:production
  update-type: version-update:semver-major
...

Signed-off-by: dependabot[bot] <support@github.com> ([818afd6](https://github.com/tj-actions/auto-doc/commit/818afd625465b3b6d269c4cdfb60f0a6926272f2))  - (dependabot[bot])
- Bump actions/setup-go from 2 to 3

Bumps [actions/setup-go](https://github.com/actions/setup-go) from 2 to 3.
- [Release notes](https://github.com/actions/setup-go/releases)
- [Commits](https://github.com/actions/setup-go/compare/v2...v3)

---
updated-dependencies:
- dependency-name: actions/setup-go
  dependency-type: direct:production
  update-type: version-update:semver-major
...

Signed-off-by: dependabot[bot] <support@github.com> ([5e0e3ef](https://github.com/tj-actions/auto-doc/commit/5e0e3efe01aa1468a934919c32dee025c96e1af9))  - (dependabot[bot])
- Bump hmarr/auto-approve-action from 2 to 3

Bumps [hmarr/auto-approve-action](https://github.com/hmarr/auto-approve-action) from 2 to 3.
- [Release notes](https://github.com/hmarr/auto-approve-action/releases)
- [Commits](https://github.com/hmarr/auto-approve-action/compare/v2...v3)

---
updated-dependencies:
- dependency-name: hmarr/auto-approve-action
  dependency-type: direct:production
  update-type: version-update:semver-major
...

Signed-off-by: dependabot[bot] <support@github.com> ([cfaf23b](https://github.com/tj-actions/auto-doc/commit/cfaf23bc38b3419a629bc5374303466556e1dfd1))  - (dependabot[bot])

## <!-- 26 -->🔄 Update

- Update peter-evans/create-pull-request action to v4.2.3 ([364280d](https://github.com/tj-actions/auto-doc/commit/364280d2938067e01fe228a020a28d7e4e9cfc04))  - (renovate[bot])
- Update peter-evans/create-pull-request action to v4.2.2 ([a3c598a](https://github.com/tj-actions/auto-doc/commit/a3c598a78a3bb7834a0ed753fa5188fcb7f2e826))  - (renovate[bot])
- Update actions/checkout action to v3 ([717164d](https://github.com/tj-actions/auto-doc/commit/717164da68c1a1d4ec061f4539c8b781c760262a))  - (renovate[bot])
- Update code-coverage.yml ([e167786](https://github.com/tj-actions/auto-doc/commit/e1677868199724ca08be3e6bf03743e56946b2c9))  - (Tonye Jack)
- Update root.go ([b14395c](https://github.com/tj-actions/auto-doc/commit/b14395c440c923e21aacfd4a0b39add70fdc1c10))  - (Tonye Jack)
- Update peter-evans/create-pull-request action to v4.2.1 ([2ab5b11](https://github.com/tj-actions/auto-doc/commit/2ab5b11350bee62e8da46adaec342a6cc41f1d66))  - (renovate[bot])
- Update tj-actions/auto-doc action to v1.4.3 ([4a5bccb](https://github.com/tj-actions/auto-doc/commit/4a5bccb24219a29024410356a0bfe048a79fc374))  - (renovate[bot])

## <!-- 30 -->📝 Other

- PR [#373](https://github.com/tj-actions/auto-doc/pull/373): update test.yml ([37063b0](https://github.com/tj-actions/auto-doc/commit/37063b05ab6d4fc58fac6c378a367234d15474bb))  - (repo-ranger[bot])
- PR [#365](https://github.com/tj-actions/auto-doc/pull/365): add support for coverage reporting ([bf28c8a](https://github.com/tj-actions/auto-doc/commit/bf28c8a915ce926decf795fe06b0932da1c2ea1a))  - (repo-ranger[bot])
- Merge 56788b3234632f3cb7e8e10a7da5ad649b50ece7 into b14395c440c923e21aacfd4a0b39add70fdc1c10
 ([14a4985](https://github.com/tj-actions/auto-doc/commit/14a4985658dcb5f06001f77cb5cd692027cb854b))  - (Tonye Jack)

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Update test.yml ([2ff714d](https://github.com/tj-actions/auto-doc/commit/2ff714d25a4a0d032fd4f479acb6ac2976faf63e))  - (Tonye Jack)
- Updated coverage badge. ([566e963](https://github.com/tj-actions/auto-doc/commit/566e963ef33fccf7fd4e2515fadd73cb15576549))  - (GitHub Action)
- Add support for coverage reporting ([56788b3](https://github.com/tj-actions/auto-doc/commit/56788b3234632f3cb7e8e10a7da5ad649b50ece7))  - (Tonye Jack)
- Update test ([#361](https://github.com/tj-actions/auto-doc/issues/361)) ([5965f81](https://github.com/tj-actions/auto-doc/commit/5965f817a5a312291cab398076fc9790ed61ef45))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded to v1.4.3 ([#359](https://github.com/tj-actions/auto-doc/issues/359))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com> ([7b1a30e](https://github.com/tj-actions/auto-doc/commit/7b1a30e383315e33dc97220e19348118727dc28c))  - (Tonye Jack)

# [1.4.3](https://github.com/tj-actions/auto-doc/compare/v1.4.2...v1.4.3) - (2022-10-30)

## <!-- 26 -->🔄 Update

- Update tj-actions/auto-doc action to v1.4.2 ([#358](https://github.com/tj-actions/auto-doc/issues/358))

Co-authored-by: renovate[bot] <29139614+renovate[bot]@users.noreply.github.com> ([c9f5aea](https://github.com/tj-actions/auto-doc/commit/c9f5aea1d491372bea97865cc3866a34b1e0994e))  - (renovate[bot])
- Updated README.md ([#356](https://github.com/tj-actions/auto-doc/issues/356))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com>
Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com> ([8c8396a](https://github.com/tj-actions/auto-doc/commit/8c8396a95946834ae7ced5b4d56bb70262b76384))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded to ([#357](https://github.com/tj-actions/auto-doc/issues/357))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com> ([2b97fb7](https://github.com/tj-actions/auto-doc/commit/2b97fb7b2b0012259f6b5207e78ba387d6236ba0))  - (Tonye Jack)

# [1.4.2](https://github.com/tj-actions/auto-doc/compare/v1.4.1...v1.4.2) - (2022-10-29)

## <!-- 16 -->➕ Add

- Create update-readme.yml ([ab29640](https://github.com/tj-actions/auto-doc/commit/ab296403ab80f122a82a87a44d2b3ddc1a01361c))  - (Tonye Jack)
- Create codeql.yml ([3a88e90](https://github.com/tj-actions/auto-doc/commit/3a88e902880de7d23e0039f0dacd9b0f86d32878))  - (Tonye Jack)
- Added .github/workflows/codacy-analysis.yml ([0d0f184](https://github.com/tj-actions/auto-doc/commit/0d0f184260a8373ce9d37d794eacdcc85f4b2d39))  - (Tonye Jack)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([0add7a3](https://github.com/tj-actions/auto-doc/commit/0add7a35d70fe4508e025a03bf5b8d28b39aaf08))  - (github-actions[bot])
- Formatted code.
 ([0dd7312](https://github.com/tj-actions/auto-doc/commit/0dd7312b7f96c4513d79ef07ddb98d6c82e7029c))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Updated README.md ([#355](https://github.com/tj-actions/auto-doc/issues/355))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com>
Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com> ([9426e56](https://github.com/tj-actions/auto-doc/commit/9426e56ee1e2827389ec4cfaaeaf6d09fa3cd905))  - (Tonye Jack)
- Updated README.md ([#354](https://github.com/tj-actions/auto-doc/issues/354))

* Updated README.md

* Updated README.

Co-authored-by: jackton1 <jackton1@users.noreply.github.com>
Co-authored-by: github-actions[bot] <github-actions[bot]@users.noreply.github.com> ([1ab8bd5](https://github.com/tj-actions/auto-doc/commit/1ab8bd57bf58757fc4b15a2d9beab59d09c9d65b))  - (Tonye Jack)
- Update README.md ([7ec0dd2](https://github.com/tj-actions/auto-doc/commit/7ec0dd21786eaa1283fed8cc47067d64e65f8615))  - (Tonye Jack)
- Updated .github/ISSUE_TEMPLATE/bug_report.yaml ([81d228b](https://github.com/tj-actions/auto-doc/commit/81d228b368d5a0d6f900144716f5f81ef66c4d0d))  - (Tonye Jack)
- Update module github.com/spf13/cobra to v1.6.1 ([2964083](https://github.com/tj-actions/auto-doc/commit/2964083cb8086a65a31e4c544eeb389176219d1a))  - (renovate[bot])
- Updated README.md ([#351](https://github.com/tj-actions/auto-doc/issues/351))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com> ([b1e6f79](https://github.com/tj-actions/auto-doc/commit/b1e6f79027639db2296764829198f79b46229fb9))  - (Tonye Jack)
- Update update-readme.yml ([27234d4](https://github.com/tj-actions/auto-doc/commit/27234d49f1a564a201f1647e8280012e79f7e40e))  - (Tonye Jack)
- Update action.yml ([270eb8b](https://github.com/tj-actions/auto-doc/commit/270eb8b39024538853f9a6753f2dc31e9cccb060))  - (Tonye Jack)
- Update peter-evans/create-pull-request action to v4.2.0 ([a3e9a3b](https://github.com/tj-actions/auto-doc/commit/a3e9a3bf4f49e2cafb29e878e7252b6c2d51bd38))  - (renovate[bot])
- Update pascalgn/automerge-action action to v0.15.5 ([cfab5c9](https://github.com/tj-actions/auto-doc/commit/cfab5c992452e3b78c1caf47b1e91e92178e00ff))  - (renovate[bot])
- Update peter-evans/create-pull-request action to v4.1.4 ([869d406](https://github.com/tj-actions/auto-doc/commit/869d40670821f6fa05cea9f15276fe3e543924f2))  - (renovate[bot])
- Update tj-actions/verify-changed-files action to v12 ([5d6e9a2](https://github.com/tj-actions/auto-doc/commit/5d6e9a278311b1d6a1e6eaee1abf7c81f2638218))  - (renovate[bot])
- Updated README.
 ([0f8d40e](https://github.com/tj-actions/auto-doc/commit/0f8d40e1fd073917fc2a5f86247ca537109e0de6))  - (github-actions[bot])
- Update module github.com/spf13/cobra to v1.6.0 ([b3e7e81](https://github.com/tj-actions/auto-doc/commit/b3e7e815274cd150f4be561a8aa050e48d4bcc57))  - (renovate[bot])
- Update action.yml ([90d80f0](https://github.com/tj-actions/auto-doc/commit/90d80f0254c282bfc95337b3afbc0ea68d124db4))  - (Tonye Jack)
- Updated .github/workflows/greetings.yml ([436e6b7](https://github.com/tj-actions/auto-doc/commit/436e6b7d18bbe8bf93f6c6ab6906fe33bca12ae6))  - (Tonye Jack)
- Update actions/first-interaction action to v1.1.1 ([e183d1a](https://github.com/tj-actions/auto-doc/commit/e183d1a714d33ed799e4a9070287bc35f360d826))  - (renovate[bot])
- Update actions/checkout action to v3.1.0 ([f1f2510](https://github.com/tj-actions/auto-doc/commit/f1f2510c4321d0c7904a18530d4ee3dc0adc7ad1))  - (renovate[bot])
- Update peter-evans/create-pull-request action to v4.1.3 ([7500c9f](https://github.com/tj-actions/auto-doc/commit/7500c9fc3a089f92f0c295fd8a715f76c1104546))  - (renovate[bot])
- Update peter-evans/create-pull-request action to v4.1.2 ([1cd60c5](https://github.com/tj-actions/auto-doc/commit/1cd60c5fa20c3cd4144507ee892528c38587b8aa))  - (renovate[bot])
- Update codacy/codacy-analysis-cli-action action to v4.2.0 ([5730e93](https://github.com/tj-actions/auto-doc/commit/5730e93fe9fa86700eac8800bde65074ca004f9e))  - (renovate[bot])
- Updated .github/workflows/sync-release-version.yml ([21e08fe](https://github.com/tj-actions/auto-doc/commit/21e08feb81acb6bdfdb4c703c166d24d369fd279))  - (Tonye Jack)
- Updated .github/workflows/sync-release-version.yml ([7150e9a](https://github.com/tj-actions/auto-doc/commit/7150e9a68c6bb5c23b70fe89280147c0f6592de6))  - (Tonye Jack)
- Updated .github/workflows/codacy-analysis.yml ([d9231ab](https://github.com/tj-actions/auto-doc/commit/d9231ab2a5bbc968e12e3f3664f9c8c3af2aed56))  - (Tonye Jack)
- Updated .github/workflows/codacy-analysis.yml ([1f42fee](https://github.com/tj-actions/auto-doc/commit/1f42feed46830b7a9a2ddc8fdfadedc942257da5))  - (Tonye Jack)
- Updated .github/workflows/greetings.yml ([37424e6](https://github.com/tj-actions/auto-doc/commit/37424e65e65e5742b2e57881ad079833031de169))  - (Tonye Jack)
- Updated .github/ISSUE_TEMPLATE/bug_report.yaml ([04231f4](https://github.com/tj-actions/auto-doc/commit/04231f4aa5795e33b3eb07bcfc1c8dbba08c6a0e))  - (Tonye Jack)
- Updated .github/ISSUE_TEMPLATE/feature_request.yaml ([133d39e](https://github.com/tj-actions/auto-doc/commit/133d39e47ca20c76202c69ab299a445499bed16b))  - (Tonye Jack)
- Updated .github/ISSUE_TEMPLATE/bug_report.yaml ([d09e9b4](https://github.com/tj-actions/auto-doc/commit/d09e9b4d98d0f962b7d26cf7224555072504ff26))  - (Tonye Jack)
- Updated .github/workflows/sync-release-version.yml ([38aaacf](https://github.com/tj-actions/auto-doc/commit/38aaacfb7503cab89cc1f1b16746c5a766c677ec))  - (Tonye Jack)
- Update tj-actions/sync-release-version action to v13 ([7748126](https://github.com/tj-actions/auto-doc/commit/7748126d45c225ec827e0ae0448de231c55ea9e9))  - (renovate[bot])
- Update tj-actions/github-changelog-generator action to v1.15 ([1f13fdc](https://github.com/tj-actions/auto-doc/commit/1f13fdca0bd10930ed48fce309d74b794bf599d7))  - (renovate[bot])
- Update tj-actions/verify-changed-files action to v11 ([a7ed774](https://github.com/tj-actions/auto-doc/commit/a7ed7749d95047ecb0d990c885ada2b68fdc19ce))  - (renovate[bot])

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Update table alignment ([#353](https://github.com/tj-actions/auto-doc/issues/353)) ([1676970](https://github.com/tj-actions/auto-doc/commit/16769700cb07ddaa3140281a3414917f301a85ee))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.4.0 -> v1.4.1 ([#336](https://github.com/tj-actions/auto-doc/issues/336))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com> ([bcdc885](https://github.com/tj-actions/auto-doc/commit/bcdc885577201dd8a810bda7cb55e9e084bf0d98))  - (Tonye Jack)

# [1.4.1](https://github.com/tj-actions/auto-doc/compare/v1.4.0...v1.4.1) - (2022-08-18)

## <!-- 1 -->🐛 Bug Fixes

- Bug handling newline characters ([#335](https://github.com/tj-actions/auto-doc/issues/335)) ([889e7f5](https://github.com/tj-actions/auto-doc/commit/889e7f5cb038db85ed6b969305a8c2152b1ef6fb))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded to v1.4.0 ([#334](https://github.com/tj-actions/auto-doc/issues/334))

Co-authored-by: jackton1 <jackton1@users.noreply.github.com> ([036c06b](https://github.com/tj-actions/auto-doc/commit/036c06b754fcbe75dce30a20a5e4f9d5e57e6524))  - (Tonye Jack)

# [1.4.0](https://github.com/tj-actions/auto-doc/compare/v1.3.1...v1.4.0) - (2022-08-18)

## <!-- 0 -->🚀 Features

- Rename variable and fixed multiline defaults ([#332](https://github.com/tj-actions/auto-doc/issues/332)) ([e0edb1d](https://github.com/tj-actions/auto-doc/commit/e0edb1d099fead46959cdaaaf37a501375b89e6a))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Update to go-version-file
 ([d053ee6](https://github.com/tj-actions/auto-doc/commit/d053ee6aef88a0449ede5da00ec9b343f3a9fb95))  - (Tonye Jack)
- Update module go to 1.19 ([3803ab9](https://github.com/tj-actions/auto-doc/commit/3803ab9aa9a5762a436c439042bb593046522ef2))  - (renovate[bot])

## <!-- 30 -->📝 Other

- PR [#331](https://github.com/tj-actions/auto-doc/pull/331): ([d6c4d56](https://github.com/tj-actions/auto-doc/commit/d6c4d56bbc8f9c487860a1c7188f4c07d8d6c43a))  - (Tonye Jack)
- PR [#329](https://github.com/tj-actions/auto-doc/pull/329): to v1.3.1 ([7cae39b](https://github.com/tj-actions/auto-doc/commit/7cae39b4d5787513e8488a9a29fbd4de75137051))  - (Tonye Jack)

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Update readme ([#333](https://github.com/tj-actions/auto-doc/issues/333)) ([feb3b20](https://github.com/tj-actions/auto-doc/commit/feb3b203847e2893803b123d59981cd93a8abf4a))  - (Tonye Jack)
- Update go version ([8100518](https://github.com/tj-actions/auto-doc/commit/81005188596b54c3d60a927c10273ecedc277f0a))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.3.0 -> v1.3.1
 ([979e1fe](https://github.com/tj-actions/auto-doc/commit/979e1fe7fb197510d7022f546b99d94c00e99cb8))  - (jackton1)

# [1.3.1](https://github.com/tj-actions/auto-doc/compare/v1.3.0...v1.3.1) - (2022-07-28)

## <!-- 30 -->📝 Other

- PR [#328](https://github.com/tj-actions/auto-doc/pull/328): to v1.3.0 ([270305b](https://github.com/tj-actions/auto-doc/commit/270305b00f5d04f12c9eda103d27271dd5d9d2e2))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.15 -> v1.3.0
 ([a296adb](https://github.com/tj-actions/auto-doc/commit/a296adb9ae21dda9a647955a272dce12e5f88355))  - (jackton1)

# [1.3.0](https://github.com/tj-actions/auto-doc/compare/v1.2.15...v1.3.0) - (2022-07-28)

## <!-- 0 -->🚀 Features

- Add support for customizing input and output columns ([1060d69](https://github.com/tj-actions/auto-doc/commit/1060d69274c274dd61ed7b18b8f1f60a1035f2c9))  - (Tonye Jack)

## <!-- 1 -->🐛 Bug Fixes

- Fix lint error
 ([9f44190](https://github.com/tj-actions/auto-doc/commit/9f44190c0539d2addf9f8f10766e2fa05d5ebd53))  - (Tonye Jack)

## <!-- 13 -->📦 Bumps

- Bump gopkg.in/yaml.v3 from 3.0.0 to 3.0.1

Bumps [gopkg.in/yaml.v3](https://github.com/go-yaml/yaml) from 3.0.0 to 3.0.1.
- [Release notes](https://github.com/go-yaml/yaml/releases)
- [Commits](https://github.com/go-yaml/yaml/compare/v3.0.0...v3.0.1)

---
updated-dependencies:
- dependency-name: gopkg.in/yaml.v3
  dependency-type: direct:production
  update-type: version-update:semver-patch
...

Signed-off-by: dependabot[bot] <support@github.com> ([d61c0dd](https://github.com/tj-actions/auto-doc/commit/d61c0dd961e73720e5de23025c69684a41b82453))  - (dependabot[bot])

## <!-- 16 -->➕ Add

- Added debug message
 ([33e422b](https://github.com/tj-actions/auto-doc/commit/33e422b453e213b504c3a4988f550100d8cf1d91))  - (Tonye Jack)
- Added doc string
 ([4e865b6](https://github.com/tj-actions/auto-doc/commit/4e865b677f2e2857f1fd337d0a823abad1297f82))  - (Tonye Jack)
- Added doc string
 ([be21baf](https://github.com/tj-actions/auto-doc/commit/be21bafa13a7067d2e220e42aa8fc5fbc4fff828))  - (Tonye Jack)
- Added doc string
 ([2923c93](https://github.com/tj-actions/auto-doc/commit/2923c93262a6380bc07bdd3fdb4e9f4a738ce50e))  - (Tonye Jack)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([98b23cc](https://github.com/tj-actions/auto-doc/commit/98b23cccdda3704c9b43122af25ef17017f8c83e))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Updated README.
 ([4458a77](https://github.com/tj-actions/auto-doc/commit/4458a77fb9c7a9ae3d9943efe0fa69c8a8e0b956))  - (github-actions[bot])
- Update action.yml ([6ed00fd](https://github.com/tj-actions/auto-doc/commit/6ed00fd9a16cecab1dc9343e27d31b675ef3d4a8))  - (Tonye Jack)
- Update action.yml ([29c1812](https://github.com/tj-actions/auto-doc/commit/29c18126e3e54c68b4a8b7b6d2123eb35a30371e))  - (Tonye Jack)
- Update test
 ([9e5afb6](https://github.com/tj-actions/auto-doc/commit/9e5afb629d8262661c2ce03fed043bfab6728ab5))  - (Tonye Jack)
- Updated README.
 ([b25ca68](https://github.com/tj-actions/auto-doc/commit/b25ca68d76e3cbee1253089f959cbc7c03092183))  - (github-actions[bot])
- Update test
 ([b41f1af](https://github.com/tj-actions/auto-doc/commit/b41f1afd63550ab55c12438221e1a3dbabcf53f2))  - (Tonye Jack)
- Update test
 ([1383e18](https://github.com/tj-actions/auto-doc/commit/1383e18b67763885684c716a24291ee77c6525ac))  - (Tonye Jack)
- Update test
 ([55734ad](https://github.com/tj-actions/auto-doc/commit/55734adda02c3c07a9a49045a1931fed566b46a7))  - (Tonye Jack)
- Update test
 ([d97d4a8](https://github.com/tj-actions/auto-doc/commit/d97d4a8811dc4ec2ee2f3d8809ec2a32deeb043a))  - (Tonye Jack)
- Update test
 ([aa2a38a](https://github.com/tj-actions/auto-doc/commit/aa2a38a61fa390640d213be32d9d096a9f7ab44c))  - (Tonye Jack)
- Update test
 ([8c5d0c2](https://github.com/tj-actions/auto-doc/commit/8c5d0c26d3defdb9909b7349861001c582fe0b84))  - (Tonye Jack)
- Update test
 ([3fb10ff](https://github.com/tj-actions/auto-doc/commit/3fb10ff75a52d8728f3bac9788614d746c909328))  - (Tonye Jack)
- Update test
 ([3c75f4f](https://github.com/tj-actions/auto-doc/commit/3c75f4f8e5a532dc8ed0ee3dfaf865262c5df636))  - (Tonye Jack)
- Update action.yml
 ([5ca2941](https://github.com/tj-actions/auto-doc/commit/5ca294115b1849de085fa8327a0ca98ef135c9a5))  - (Tonye Jack)
- Update action.yml
 ([0442370](https://github.com/tj-actions/auto-doc/commit/0442370ef49db019faadf44f831f2b5de1b1ede8))  - (Tonye Jack)
- Update action.yml
 ([79d3840](https://github.com/tj-actions/auto-doc/commit/79d38407f9bd9fd82c9d38c2ad329e255dac3967))  - (Tonye Jack)
- Update action.yml
 ([590c426](https://github.com/tj-actions/auto-doc/commit/590c426748762e4c0e91b6f57b64b35b8e36797f))  - (Tonye Jack)
- Update tj-actions/github-changelog-generator action to v1.14 ([0bdd8ef](https://github.com/tj-actions/auto-doc/commit/0bdd8efd706699210acd865d842848ba55027488))  - (renovate[bot])
- Update tj-actions/verify-changed-files action to v10 ([f3f255e](https://github.com/tj-actions/auto-doc/commit/f3f255ef22b619a09df1ff5021ee310fb5dd1889))  - (renovate[bot])
- Update module github.com/spf13/cobra to v1.5.0 ([acf3dce](https://github.com/tj-actions/auto-doc/commit/acf3dce55564e7a84f4f074b8e1e16cdeefcc8ad))  - (renovate[bot])

## <!-- 30 -->📝 Other

- PR [#327](https://github.com/tj-actions/auto-doc/pull/327): ([160af83](https://github.com/tj-actions/auto-doc/commit/160af8367b7e8261d90f584f140b9ebfff32eb8b))  - (Tonye Jack)
- Merge 5aa504770e06ca8b01bdec1cb5a0fd1aa8949954 into 6ed00fd9a16cecab1dc9343e27d31b675ef3d4a8
 ([3473987](https://github.com/tj-actions/auto-doc/commit/34739871ef4ae2aaf5e9cdaebeaa449e6a0ec5e1))  - (Tonye Jack)
- PR [#326](https://github.com/tj-actions/auto-doc/pull/326): ([0aaacc5](https://github.com/tj-actions/auto-doc/commit/0aaacc519159cc0c1dd9a332282e78cd9de38f60))  - (Tonye Jack)
- Merge 9f44190c0539d2addf9f8f10766e2fa05d5ebd53 into 0bdd8efd706699210acd865d842848ba55027488
 ([ce3ce7d](https://github.com/tj-actions/auto-doc/commit/ce3ce7d11c2a3ec5210b18e5c2825a65f5b6672f))  - (Tonye Jack)
- PR [#322](https://github.com/tj-actions/auto-doc/pull/322): module github.com/spf13/cobra to v1.5.0 ([ec77ae5](https://github.com/tj-actions/auto-doc/commit/ec77ae5682f6ce7c2bfe6b0431c7fe97f2026f5a))  - (Tonye Jack)
- Merge acf3dce55564e7a84f4f074b8e1e16cdeefcc8ad into d61c0dd961e73720e5de23025c69684a41b82453
 ([2a45bf4](https://github.com/tj-actions/auto-doc/commit/2a45bf4d4e25d78bb56d859bb3a442147b318829))  - (renovate[bot])
- PR [#318](https://github.com/tj-actions/auto-doc/pull/318): to v1.2.15 ([bd89440](https://github.com/tj-actions/auto-doc/commit/bd8944050c0999d55db44ac37afab55bef5e14de))  - (Tonye Jack)

## <!-- 7 -->⚙️ Miscellaneous Tasks

- Update readme ([5aa5047](https://github.com/tj-actions/auto-doc/commit/5aa504770e06ca8b01bdec1cb5a0fd1aa8949954))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.14 -> v1.2.15
 ([0d8d332](https://github.com/tj-actions/auto-doc/commit/0d8d332b5fc36eb7a0e3e711a24ff60c03df18c6))  - (jackton1)

# [1.2.15](https://github.com/tj-actions/auto-doc/compare/v1.2.14...v1.2.15) - (2022-05-23)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([f93e042](https://github.com/tj-actions/auto-doc/commit/f93e0428c6d1d251670055abec8f908e88457161))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Update module gopkg.in/yaml.v3 to v3.0.0
 ([ee2fe08](https://github.com/tj-actions/auto-doc/commit/ee2fe08fd26f22f7903dab472a27b19bd97a2c73))  - (Renovate Bot)
- Update goreleaser/goreleaser-action action to v3
 ([50a8371](https://github.com/tj-actions/auto-doc/commit/50a8371401bd6dc0777163097d1f225b3590a45b))  - (Renovate Bot)
- Update pascalgn/automerge-action action to v0.15.3
 ([4f08cf7](https://github.com/tj-actions/auto-doc/commit/4f08cf77a4cf9ba3dc99d7a27c5e0ccc0788068c))  - (Renovate Bot)
- Update actions/checkout action to v3.0.2
 ([a3b1aa9](https://github.com/tj-actions/auto-doc/commit/a3b1aa978c5b506f7af6bff98a438e8b6a68e2a1))  - (Renovate Bot)
- Update actions/checkout action to v3.0.1
 ([0f469b1](https://github.com/tj-actions/auto-doc/commit/0f469b1a41f31c5eb306508e84672a95d2788cfb))  - (Renovate Bot)
- Update pascalgn/automerge-action action to v0.15.2
 ([549d3b6](https://github.com/tj-actions/auto-doc/commit/549d3b6243a241e9a17150953e970e63bdb4fd08))  - (Renovate Bot)
- Update pascalgn/automerge-action action to v0.15.1
 ([f09e18a](https://github.com/tj-actions/auto-doc/commit/f09e18a42938ecc2014a9ec8367c7296af8bbac5))  - (Renovate Bot)
- Update pascalgn/automerge-action action to v0.14.4
 ([e02e27a](https://github.com/tj-actions/auto-doc/commit/e02e27af2640de672a13d54de015bbdb45b5cf35))  - (Renovate Bot)
- Update peter-evans/create-pull-request action to v4
 ([b2bea3d](https://github.com/tj-actions/auto-doc/commit/b2bea3d56042aed8ace48f96f36cd3c533e9a30c))  - (Renovate Bot)
- Update actions/cache action to v3
 ([89ea3af](https://github.com/tj-actions/auto-doc/commit/89ea3afa7968436f4943a5f0c7ca63720c773754))  - (Renovate Bot)

## <!-- 30 -->📝 Other

- PR [#316](https://github.com/tj-actions/auto-doc/pull/316): ([e27016a](https://github.com/tj-actions/auto-doc/commit/e27016a553ca1fa38c0ec160987194b4685ae06d))  - (Tonye Jack)
- Merge ee2fe08fd26f22f7903dab472a27b19bd97a2c73 into 50a8371401bd6dc0777163097d1f225b3590a45b
 ([e2c82a8](https://github.com/tj-actions/auto-doc/commit/e2c82a8c0774418fb629100ab548a276e92a6898))  - (renovate[bot])
- PR [#307](https://github.com/tj-actions/auto-doc/pull/307): to v1.2.14 ([8911132](https://github.com/tj-actions/auto-doc/commit/89111325daad625868bf99a6c1aaaccc3fc7307f))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.13 -> v1.2.14
 ([8075f01](https://github.com/tj-actions/auto-doc/commit/8075f0178ab6b696d7efc757318b8719ee5b1933))  - (jackton1)

# [1.2.14](https://github.com/tj-actions/auto-doc/compare/v1.2.13...v1.2.14) - (2022-03-19)

## <!-- 13 -->📦 Bumps

- Bump tj-actions/verify-changed-files from 8.8 to 9

Bumps [tj-actions/verify-changed-files](https://github.com/tj-actions/verify-changed-files) from 8.8 to 9.
- [Release notes](https://github.com/tj-actions/verify-changed-files/releases)
- [Changelog](https://github.com/tj-actions/verify-changed-files/blob/main/HISTORY.md)
- [Commits](https://github.com/tj-actions/verify-changed-files/compare/v8.8...v9)

---
updated-dependencies:
- dependency-name: tj-actions/verify-changed-files
  dependency-type: direct:production
  update-type: version-update:semver-major
...

Signed-off-by: dependabot[bot] <support@github.com> ([b2c7d64](https://github.com/tj-actions/auto-doc/commit/b2c7d64632efc5be69546d31ee3e7ce48e4120d3))  - (dependabot[bot])
- Bump github.com/spf13/cobra from 1.3.0 to 1.4.0

Bumps [github.com/spf13/cobra](https://github.com/spf13/cobra) from 1.3.0 to 1.4.0.
- [Release notes](https://github.com/spf13/cobra/releases)
- [Changelog](https://github.com/spf13/cobra/blob/master/CHANGELOG.md)
- [Commits](https://github.com/spf13/cobra/compare/v1.3.0...v1.4.0)

---
updated-dependencies:
- dependency-name: github.com/spf13/cobra
  dependency-type: direct:production
  update-type: version-update:semver-minor
...

Signed-off-by: dependabot[bot] <support@github.com> ([76f7fdc](https://github.com/tj-actions/auto-doc/commit/76f7fdcfae92db0e9d4afd785499fa66581c455a))  - (dependabot[bot])

## <!-- 26 -->🔄 Update

- Updated README.
 ([8a6f7f6](https://github.com/tj-actions/auto-doc/commit/8a6f7f668029e1e00a8bb44135f177d2db98bb41))  - (github-actions[bot])
- Update tj-actions/remark action to v3
 ([6f2718d](https://github.com/tj-actions/auto-doc/commit/6f2718dd0ffeb781af182d3244cedb97a07b6fac))  - (Renovate Bot)
- Update tj-actions/github-changelog-generator action to v1.13
 ([699796f](https://github.com/tj-actions/auto-doc/commit/699796ff0bc9fb564a823c956ea4e0871d819b5f))  - (Renovate Bot)
- Update actions/upload-artifact action to v3
 ([2db2656](https://github.com/tj-actions/auto-doc/commit/2db2656ee3581a3dc3bb244a2254303b54c27a58))  - (Renovate Bot)
- Update actions/checkout action
 ([31aa3f4](https://github.com/tj-actions/auto-doc/commit/31aa3f49490f086469f743c1ca6d43e9d6c8e480))  - (Renovate Bot)
- Update actions/setup-go action to v3
 ([2eb58e5](https://github.com/tj-actions/auto-doc/commit/2eb58e55b17c340e5e1e31ace10d017e88c34872))  - (Renovate Bot)
- Update tj-actions/github-changelog-generator action to v1.12
 ([4b52180](https://github.com/tj-actions/auto-doc/commit/4b5218031855398ebc5863fadc87fad67032688b))  - (Renovate Bot)
- Update tj-actions/sync-release-version action to v11
 ([7475295](https://github.com/tj-actions/auto-doc/commit/747529506ab2e0ce112848a8d8adbb683b100633))  - (Renovate Bot)
- Updated .github/workflows/greetings.yml ([fffbb0f](https://github.com/tj-actions/auto-doc/commit/fffbb0f203817e5d76d943615b244b76598d193f))  - (Tonye Jack)
- Updated .github/workflows/greetings.yml ([15fe1f7](https://github.com/tj-actions/auto-doc/commit/15fe1f7a81f8bca0d8ff46c3c9ea26091a985d27))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#305](https://github.com/tj-actions/auto-doc/pull/305): tj-actions/remark action to v3 ([359cb6b](https://github.com/tj-actions/auto-doc/commit/359cb6b98e7328872d715a6a38a8d5bc09b1f4c8))  - (Tonye Jack)
- Merge 6f2718dd0ffeb781af182d3244cedb97a07b6fac into b2c7d64632efc5be69546d31ee3e7ce48e4120d3
 ([8544d97](https://github.com/tj-actions/auto-doc/commit/8544d9763bffeb1b2366d4cad65aeda520e6b4bb))  - (renovate[bot])
- PR [#294](https://github.com/tj-actions/auto-doc/pull/294): to v1.2.13 ([0496402](https://github.com/tj-actions/auto-doc/commit/0496402e73a1b2ca3fc654b06895b0e4dff19f5e))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.12 -> v1.2.13
 ([d97528d](https://github.com/tj-actions/auto-doc/commit/d97528d36bcfca4860b9f73690a1407b352c18ba))  - (jackton1)

# [1.2.13](https://github.com/tj-actions/auto-doc/compare/v1.2.12...v1.2.13) - (2022-02-05)

## <!-- 1 -->🐛 Bug Fixes

- Rendering pipe separator ([ba65038](https://github.com/tj-actions/auto-doc/commit/ba65038a68f9a0265a9197f00b44ae2f31e06591))  - (Tonye Jack)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([dbac49d](https://github.com/tj-actions/auto-doc/commit/dbac49d4566a944072bfd5cf31fd8198c23fe496))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Updated README.
 ([e6610de](https://github.com/tj-actions/auto-doc/commit/e6610dea2f59033f986073ca3f2988441a4e3a38))  - (github-actions[bot])
- Update root.go ([e3e562f](https://github.com/tj-actions/auto-doc/commit/e3e562fa3a4b16505125d418b21d766af533cfc2))  - (Tonye Jack)
- Updated README.
 ([f5aa843](https://github.com/tj-actions/auto-doc/commit/f5aa843c9f30878c704af2bd48e90e430394e6b9))  - (github-actions[bot])
- Update sync-release-version.yml ([a55e48e](https://github.com/tj-actions/auto-doc/commit/a55e48ee2c1149393b6bd5a3ebc07a81803f02c0))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#293](https://github.com/tj-actions/auto-doc/pull/293): rendering pipe separator ([22fe32f](https://github.com/tj-actions/auto-doc/commit/22fe32f2945946ac5d0aeb20964e03c8c91e2598))  - (Tonye Jack)
- Merge e3e562fa3a4b16505125d418b21d766af533cfc2 into 643614e4b595cd40bc6ebcda07d0ed201e44d663
 ([f77c1f4](https://github.com/tj-actions/auto-doc/commit/f77c1f4778eb0b361e4a26dce1b8fb0bb082f426))  - (Tonye Jack)
- Merge dbac49d4566a944072bfd5cf31fd8198c23fe496 into 643614e4b595cd40bc6ebcda07d0ed201e44d663
 ([720506d](https://github.com/tj-actions/auto-doc/commit/720506de76044955e00b1bff4c7c2ffbb68a1d56))  - (Tonye Jack)
- Merge ba65038a68f9a0265a9197f00b44ae2f31e06591 into 643614e4b595cd40bc6ebcda07d0ed201e44d663
 ([ef6f251](https://github.com/tj-actions/auto-doc/commit/ef6f251a23938a0ef28fc2e371774c2920edc304))  - (Tonye Jack)
- PR [#292](https://github.com/tj-actions/auto-doc/pull/292): to v1.2.12 ([643614e](https://github.com/tj-actions/auto-doc/commit/643614e4b595cd40bc6ebcda07d0ed201e44d663))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.11 -> v1.2.12
 ([914bdd5](https://github.com/tj-actions/auto-doc/commit/914bdd564e5955718d8106af4a856f9d38660925))  - (jackton1)

# [1.2.12](https://github.com/tj-actions/auto-doc/compare/v1.2.11...v1.2.12) - (2022-02-05)

## <!-- 1 -->🐛 Bug Fixes

- Bug with escaping defaults ([d2e3694](https://github.com/tj-actions/auto-doc/commit/d2e3694e967422da7355d69c802a8b02280ba42e))  - (Tonye Jack)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([f64d692](https://github.com/tj-actions/auto-doc/commit/f64d692ff94c86a19b1a3407965fa9ba76389e1c))  - (github-actions[bot])
- Formatted code.
 ([d4fa941](https://github.com/tj-actions/auto-doc/commit/d4fa94184f0ddaaab30d01a67b193b96e6236f64))  - (github-actions[bot])
- Formatted code.
 ([dd71091](https://github.com/tj-actions/auto-doc/commit/dd71091fa441b7655eb5ae9b9e65462b9b012f6a))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Updated README.
 ([dd4777d](https://github.com/tj-actions/auto-doc/commit/dd4777d9e2c7118ea21a94bedb68849d2a4e843b))  - (github-actions[bot])
- Update root.go ([63ca8b7](https://github.com/tj-actions/auto-doc/commit/63ca8b7ac6e6ca2b9043f53cbe29d2ee221fedcd))  - (Tonye Jack)
- Updated README.
 ([d620add](https://github.com/tj-actions/auto-doc/commit/d620add6907c2b251c0fc4a4b4804ce575b96071))  - (github-actions[bot])
- Update root.go ([e3b9e5b](https://github.com/tj-actions/auto-doc/commit/e3b9e5b23fd6b4e5e2fd543b070eabbccb4f830c))  - (Tonye Jack)
- Updated README.
 ([84d2be2](https://github.com/tj-actions/auto-doc/commit/84d2be27d113367f4397209efa9497c2cb2c4555))  - (github-actions[bot])
- Update README.md ([eb898d6](https://github.com/tj-actions/auto-doc/commit/eb898d6580f0848cdff71e3a69c381596d36cb52))  - (Tonye Jack)
- Update action.yml ([0d607ec](https://github.com/tj-actions/auto-doc/commit/0d607ecaf26f680477ff619657aa452961a611c7))  - (Tonye Jack)
- Updated README.
 ([140c313](https://github.com/tj-actions/auto-doc/commit/140c31349fd45aac2b8d0604b28a209707e48c0f))  - (github-actions[bot])
- Update root.go ([b3fbb69](https://github.com/tj-actions/auto-doc/commit/b3fbb69e10f5daa60cd752c53f6df8d05e57f62c))  - (Tonye Jack)
- Updated README.
 ([b6970c3](https://github.com/tj-actions/auto-doc/commit/b6970c3072d51afe304f09cf10c9bdd12b959931))  - (github-actions[bot])
- Update root.go ([adc39fb](https://github.com/tj-actions/auto-doc/commit/adc39fb398ec395fedc41b2006d0e50efd2994c6))  - (Tonye Jack)
- Updated README.
 ([210eeea](https://github.com/tj-actions/auto-doc/commit/210eeeaf76f37d04e73e3d4ebbb6c0b0aedbf46b))  - (github-actions[bot])
- Update root.go ([1dcf652](https://github.com/tj-actions/auto-doc/commit/1dcf652bbc11b24254195879606ed158ad1a04ab))  - (Tonye Jack)
- Updated README.
 ([cdd4f1b](https://github.com/tj-actions/auto-doc/commit/cdd4f1b7e54acad5933d40a2c4a15699094a02d6))  - (github-actions[bot])
- Update root.go ([f4792d2](https://github.com/tj-actions/auto-doc/commit/f4792d20f24a5f4d96bbd291fb8f13718af7f2d0))  - (Tonye Jack)
- Updated README.
 ([5653712](https://github.com/tj-actions/auto-doc/commit/5653712ade52f7d2441b2c0f656dd9d1f8e81fab))  - (github-actions[bot])
- Update root.go ([a1a82a9](https://github.com/tj-actions/auto-doc/commit/a1a82a9a64afceb5620359bf9eaecd5c37fb4fba))  - (Tonye Jack)
- Update root.go ([b608026](https://github.com/tj-actions/auto-doc/commit/b60802648754b842a2e74ace8cf5d61c54937164))  - (Tonye Jack)
- Update root.go ([58da3e8](https://github.com/tj-actions/auto-doc/commit/58da3e8e08b07b217c1cd14429a844dfacbe9388))  - (Tonye Jack)
- Updated README.
 ([7133a5d](https://github.com/tj-actions/auto-doc/commit/7133a5d5f236239dfd9ccefa506adee2522879d4))  - (github-actions[bot])
- Update action.yml ([7de1ac6](https://github.com/tj-actions/auto-doc/commit/7de1ac61c9bbec5d7d05a1e16dd562b37a28960e))  - (Tonye Jack)
- Updated README.
 ([95ae9fb](https://github.com/tj-actions/auto-doc/commit/95ae9fbbe522802653506809df0c594cfcc4f715))  - (github-actions[bot])
- Update root.go ([a436dca](https://github.com/tj-actions/auto-doc/commit/a436dca9632b89570c826c6e351132f5713b4d97))  - (Tonye Jack)
- Update action.yml ([94d4fa3](https://github.com/tj-actions/auto-doc/commit/94d4fa3e4e2340f70faa06900d939782d00f306d))  - (Tonye Jack)
- Update action.yml ([bc54e3c](https://github.com/tj-actions/auto-doc/commit/bc54e3c572d83fa62227afa1b1a4c5f2d5100955))  - (Tonye Jack)
- Update README.md ([37c97c7](https://github.com/tj-actions/auto-doc/commit/37c97c778c88a78dfcccfb982c6b74c15ddfe8f9))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#291](https://github.com/tj-actions/auto-doc/pull/291): README.md ([8d64371](https://github.com/tj-actions/auto-doc/commit/8d6437199ac4a25ddac9fc8342fc09a2745f3825))  - (Tonye Jack)
- Merge 63ca8b7ac6e6ca2b9043f53cbe29d2ee221fedcd into 0d607ecaf26f680477ff619657aa452961a611c7
 ([3b17040](https://github.com/tj-actions/auto-doc/commit/3b170408dc993d699552b25d882d1284c4b425cf))  - (Tonye Jack)
- Merge e3b9e5b23fd6b4e5e2fd543b070eabbccb4f830c into 0d607ecaf26f680477ff619657aa452961a611c7
 ([33f237c](https://github.com/tj-actions/auto-doc/commit/33f237c84d2698764d37851c6b14d68624eef907))  - (Tonye Jack)
- Merge eb898d6580f0848cdff71e3a69c381596d36cb52 into 0d607ecaf26f680477ff619657aa452961a611c7
 ([cfeb9dd](https://github.com/tj-actions/auto-doc/commit/cfeb9dd78ddf843f1b6bc4314bca3560b60352c9))  - (Tonye Jack)
- PR [#290](https://github.com/tj-actions/auto-doc/pull/290): bug escaping defaults ([3707a84](https://github.com/tj-actions/auto-doc/commit/3707a84911a42f54f793020b66813c78b62699dc))  - (Tonye Jack)
- Merge b3fbb69e10f5daa60cd752c53f6df8d05e57f62c into 28b9e67f6995ef85aaf468ff536f46b7e3e3cc6f
 ([971fefc](https://github.com/tj-actions/auto-doc/commit/971fefc63afa2844bd3b2c05e818c3f2da1ef6d3))  - (Tonye Jack)
- Merge adc39fb398ec395fedc41b2006d0e50efd2994c6 into 28b9e67f6995ef85aaf468ff536f46b7e3e3cc6f
 ([1f416bb](https://github.com/tj-actions/auto-doc/commit/1f416bb0dd67be906b912a8c2b81056ea6621ea7))  - (Tonye Jack)
- Merge f64d692ff94c86a19b1a3407965fa9ba76389e1c into 28b9e67f6995ef85aaf468ff536f46b7e3e3cc6f
 ([98ca6f0](https://github.com/tj-actions/auto-doc/commit/98ca6f008fea2fbba8692f847a37fc3daa026a79))  - (Tonye Jack)
- Merge 1dcf652bbc11b24254195879606ed158ad1a04ab into 28b9e67f6995ef85aaf468ff536f46b7e3e3cc6f
 ([070cff5](https://github.com/tj-actions/auto-doc/commit/070cff5a879ef9ef1c579ab1e9b1914fb16643ac))  - (Tonye Jack)
- Merge d4fa94184f0ddaaab30d01a67b193b96e6236f64 into 28b9e67f6995ef85aaf468ff536f46b7e3e3cc6f
 ([53cddc8](https://github.com/tj-actions/auto-doc/commit/53cddc808905bb98a8358ced12652439816895f9))  - (Tonye Jack)
- Merge f4792d20f24a5f4d96bbd291fb8f13718af7f2d0 into 28b9e67f6995ef85aaf468ff536f46b7e3e3cc6f
 ([ba86cea](https://github.com/tj-actions/auto-doc/commit/ba86cea1d0541d26575f4ba6f55195f35d036728))  - (Tonye Jack)
- Merge dd71091fa441b7655eb5ae9b9e65462b9b012f6a into 28b9e67f6995ef85aaf468ff536f46b7e3e3cc6f
 ([576951a](https://github.com/tj-actions/auto-doc/commit/576951a4768e97a31f23cc1324e6423c3863c4af))  - (Tonye Jack)
- Merge a1a82a9a64afceb5620359bf9eaecd5c37fb4fba into 28b9e67f6995ef85aaf468ff536f46b7e3e3cc6f
 ([93ab539](https://github.com/tj-actions/auto-doc/commit/93ab539924b498c8e5b0c0b2b628e3f06829151c))  - (Tonye Jack)
- Merge 7de1ac61c9bbec5d7d05a1e16dd562b37a28960e into 28b9e67f6995ef85aaf468ff536f46b7e3e3cc6f
 ([dbc69ab](https://github.com/tj-actions/auto-doc/commit/dbc69ab1fd085f1a17cafdcedf9f03a08e62bd1d))  - (Tonye Jack)
- Merge a436dca9632b89570c826c6e351132f5713b4d97 into 28b9e67f6995ef85aaf468ff536f46b7e3e3cc6f
 ([7ffa621](https://github.com/tj-actions/auto-doc/commit/7ffa6210c933cfd0b3c9349e9ea188173724b8eb))  - (Tonye Jack)
- PR [#288](https://github.com/tj-actions/auto-doc/pull/288): binary regardless of exit code ([28b9e67](https://github.com/tj-actions/auto-doc/commit/28b9e67f6995ef85aaf468ff536f46b7e3e3cc6f))  - (Tonye Jack)
- Clean binary regardless of exit code

Closes: #281 ([0e0df3d](https://github.com/tj-actions/auto-doc/commit/0e0df3dfc0b961989f527852be461b05a605ddb6))  - (Tonye Jack)
- PR [#287](https://github.com/tj-actions/auto-doc/pull/287): to v1.2.11 ([364a643](https://github.com/tj-actions/auto-doc/commit/364a643ef9e163269390b3f8abe0d1f3f6a0fa88))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.10 -> v1.2.11
 ([9c182dd](https://github.com/tj-actions/auto-doc/commit/9c182dd059b231c022ffb20dd007eac6c9d2671f))  - (jackton1)

# [1.2.11](https://github.com/tj-actions/auto-doc/compare/v1.2.10...v1.2.11) - (2022-02-01)

## <!-- 1 -->🐛 Bug Fixes

- Ensure the latest version is always installed ([4962bb2](https://github.com/tj-actions/auto-doc/commit/4962bb25951a467d6c31d7f2c2327bf7362a3f92))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#286](https://github.com/tj-actions/auto-doc/pull/286): Ensure the latest version is always installed ([b6ab9e2](https://github.com/tj-actions/auto-doc/commit/b6ab9e24ca77073d4eaaed314d2bce423c70f436))  - (Tonye Jack)
- PR [#285](https://github.com/tj-actions/auto-doc/pull/285): to v1.2.10 ([267e3bc](https://github.com/tj-actions/auto-doc/commit/267e3bc19d8ab33adfb6570f2b8b63bb3721b673))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.9 -> v1.2.10
 ([8808f9d](https://github.com/tj-actions/auto-doc/commit/8808f9d244688d677b92352a4cb237d56c01aeca))  - (jackton1)

# [1.2.10](https://github.com/tj-actions/auto-doc/compare/v1.2.9...v1.2.10) - (2022-02-01)

## <!-- 26 -->🔄 Update

- Update README.md ([6b1ba61](https://github.com/tj-actions/auto-doc/commit/6b1ba61bb0b8255cecef3f867ca77108e58234b4))  - (Tonye Jack)
- Updated README.
 ([5a15f45](https://github.com/tj-actions/auto-doc/commit/5a15f458f3d8bd6affeeb7b0a922573eba8e7afd))  - (github-actions[bot])
- Updated action.yml spacing.
 ([02712e9](https://github.com/tj-actions/auto-doc/commit/02712e96ccc5221225f3f6c1797214549b446617))  - (Tonye Jack)
- Updated README.
 ([1928b88](https://github.com/tj-actions/auto-doc/commit/1928b88837461038098daa52d61bcd9bd7be55ec))  - (github-actions[bot])

## <!-- 30 -->📝 Other

- PR [#284](https://github.com/tj-actions/auto-doc/pull/284): output default ([59a5ca4](https://github.com/tj-actions/auto-doc/commit/59a5ca4b5b995fe8e43111275e3e6e2cde776810))  - (Tonye Jack)
- Merge 79861d36238b75653f582444cad7d1511d089544 into 02712e96ccc5221225f3f6c1797214549b446617
 ([767b697](https://github.com/tj-actions/auto-doc/commit/767b697a997423167be1a574acea31ba40aaef74))  - (Tonye Jack)
- Escaped output default ([79861d3](https://github.com/tj-actions/auto-doc/commit/79861d36238b75653f582444cad7d1511d089544))  - (Tonye Jack)
- PR [#283](https://github.com/tj-actions/auto-doc/pull/283): to v1.2.9 ([07c15c3](https://github.com/tj-actions/auto-doc/commit/07c15c3a9a3ebf2e02847d0bd455afa2f8b1979c))  - (Tonye Jack)
- Merge e0d335df9a1946890a5d53063ec9c1a2df442ba2 into 703b6f22978e8b1a43d5158894575bf45160e0c5
 ([6dcc468](https://github.com/tj-actions/auto-doc/commit/6dcc46846fd069a6e6991fd83e1b39dae9853aa2))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.8 -> v1.2.9
 ([e0d335d](https://github.com/tj-actions/auto-doc/commit/e0d335df9a1946890a5d53063ec9c1a2df442ba2))  - (jackton1)

# [1.2.9](https://github.com/tj-actions/auto-doc/compare/v1.2.8...v1.2.9) - (2022-01-05)

## <!-- 26 -->🔄 Update

- Update action.yml ([703b6f2](https://github.com/tj-actions/auto-doc/commit/703b6f22978e8b1a43d5158894575bf45160e0c5))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#282](https://github.com/tj-actions/auto-doc/pull/282): to v1.2.8 ([2b5e785](https://github.com/tj-actions/auto-doc/commit/2b5e78512f25bdf8d438a5958a6aaaa24db204ce))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.7 -> v1.2.8
 ([576bbe7](https://github.com/tj-actions/auto-doc/commit/576bbe7afad9a41321dfe35945bd67effd9b4261))  - (jackton1)

# [1.2.8](https://github.com/tj-actions/auto-doc/compare/v1.2.7...v1.2.8) - (2022-01-05)

## <!-- 26 -->🔄 Update

- Update README.md ([2b6ee8b](https://github.com/tj-actions/auto-doc/commit/2b6ee8baeb32120f5e6eb64dac93d0278ab22a9a))  - (Tonye Jack)
- Update README.md ([079fb17](https://github.com/tj-actions/auto-doc/commit/079fb17fb406e57f78a7485577df55eab14809ee))  - (Tonye Jack)
- Update root.go ([b7ce5f9](https://github.com/tj-actions/auto-doc/commit/b7ce5f9a26f6454b35d1031ebfd3f56567efe3d2))  - (Tonye Jack)
- Update root.go ([fc907c0](https://github.com/tj-actions/auto-doc/commit/fc907c0a038a55e4aa0584563bc6407b7f314784))  - (Tonye Jack)
- Update tj-actions/github-changelog-generator action to v1.11
 ([8eea901](https://github.com/tj-actions/auto-doc/commit/8eea901babfa81a3e86d8fb1966f744101e80c53))  - (Renovate Bot)

## <!-- 30 -->📝 Other

- PR [#278](https://github.com/tj-actions/auto-doc/pull/278): to v1.2.7 ([63d550b](https://github.com/tj-actions/auto-doc/commit/63d550b3abc1c2064c85ab4e0a847e2b4838bf31))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.6 -> v1.2.7
 ([30ea74d](https://github.com/tj-actions/auto-doc/commit/30ea74db1f39e7b387a9c81b139f8524b009799e))  - (jackton1)

# [1.2.7](https://github.com/tj-actions/auto-doc/compare/v1.2.6...v1.2.7) - (2022-01-03)

## <!-- 16 -->➕ Add

- Add support for creating snapshots ([d228eec](https://github.com/tj-actions/auto-doc/commit/d228eec53f04d947796263d1569e12ba8f165977))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Update README.md ([dfc8660](https://github.com/tj-actions/auto-doc/commit/dfc86607485a787c43a67c046210b7089fd6b8c3))  - (Tonye Jack)
- Update test.yml ([f54c623](https://github.com/tj-actions/auto-doc/commit/f54c62313d1d1ff067bf7dbd1308564c86b7f82c))  - (Tonye Jack)
- Update format-tidy.yml ([f986dae](https://github.com/tj-actions/auto-doc/commit/f986daeef7704f2c711f905d4f2b67d936c810e9))  - (Tonye Jack)
- Update lint.yml ([b387d7d](https://github.com/tj-actions/auto-doc/commit/b387d7d07a82d94fe52135e9fe4a28125c3a4874))  - (Tonye Jack)
- Update test.yml ([54f864f](https://github.com/tj-actions/auto-doc/commit/54f864f8212552a425ad6d2a312864aedede9bf2))  - (Tonye Jack)
- Update release.yml ([ae497e0](https://github.com/tj-actions/auto-doc/commit/ae497e039696b31cfd8616bf8e9baf34535783a9))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#277](https://github.com/tj-actions/auto-doc/pull/277): support for creating snapshots ([da44773](https://github.com/tj-actions/auto-doc/commit/da44773120935fcc1d7a53377a5059d8db306d0f))  - (Tonye Jack)
- PR [#276](https://github.com/tj-actions/auto-doc/pull/276): to v1.2.6 ([6b45612](https://github.com/tj-actions/auto-doc/commit/6b456128bc592db8739b482fda01f8104db2d4d8))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.5 -> v1.2.6
 ([79ec8b4](https://github.com/tj-actions/auto-doc/commit/79ec8b437af5ecae545c7f3e84fe20526579afb0))  - (jackton1)

# [1.2.6](https://github.com/tj-actions/auto-doc/compare/v1.2.5...v1.2.6) - (2022-01-03)

## <!-- 17 -->➖ Remove

- Delete update-readme.yml ([200cc08](https://github.com/tj-actions/auto-doc/commit/200cc081245551e0d71692442236a78a252b2366))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Update .goreleaser.yaml ([af3ad50](https://github.com/tj-actions/auto-doc/commit/af3ad503a10d0d18defa68ae9d563078435828d8))  - (Tonye Jack)
- Update test.yml ([5eb565b](https://github.com/tj-actions/auto-doc/commit/5eb565b38470e90b91d129b976fef87d41d1af04))  - (Tonye Jack)
- Updated README.
 ([f495a8a](https://github.com/tj-actions/auto-doc/commit/f495a8a7c646fa980980052da258f67a223a41fd))  - (github-actions[bot])
- Update README.md ([38ff4d4](https://github.com/tj-actions/auto-doc/commit/38ff4d40c5013eaff0d2223df7999eb89eb48297))  - (Tonye Jack)
- Update README.md ([9bf3b83](https://github.com/tj-actions/auto-doc/commit/9bf3b836577315f94efc62cc26173b17b63c43e8))  - (Tonye Jack)
- Updated README.
 ([efedc2e](https://github.com/tj-actions/auto-doc/commit/efedc2e9f948ee3c6917904df4610ff52f9b48d7))  - (github-actions[bot])
- Update test.yml ([f462122](https://github.com/tj-actions/auto-doc/commit/f462122e878ba36a10a7fcc2fcaaf7cd4212796c))  - (Tonye Jack)
- Update README.md ([785f588](https://github.com/tj-actions/auto-doc/commit/785f588eaf967c5264715c5a64751c0aed4fd03c))  - (Tonye Jack)
- Update README.md ([3433be9](https://github.com/tj-actions/auto-doc/commit/3433be9e358b02e201aaddc85e0c8784f1ad2e40))  - (Tonye Jack)
- Update README.md ([84e0fd4](https://github.com/tj-actions/auto-doc/commit/84e0fd47146f781679b88b44b0c8246293390e90))  - (Tonye Jack)
- Update README.md ([90314da](https://github.com/tj-actions/auto-doc/commit/90314da0371b5121b448e1d18d8612a1c03c6d36))  - (Tonye Jack)
- Update tj-actions/github-changelog-generator action to v1.10
 ([4803f52](https://github.com/tj-actions/auto-doc/commit/4803f520a325ec1a4b2c97a0588c7ff5318bf743))  - (Renovate Bot)

## <!-- 30 -->📝 Other

- PR [#274](https://github.com/tj-actions/auto-doc/pull/274): ([fb02554](https://github.com/tj-actions/auto-doc/commit/fb02554d8f78074884aca9b62ffc1da8f6783772))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.4 -> v1.2.5
 ([81714a0](https://github.com/tj-actions/auto-doc/commit/81714a0f8c5a711ab81f4c19fd28ae306ec60750))  - (jackton1)

# [1.2.5](https://github.com/tj-actions/auto-doc/compare/v1.2.4...v1.2.5) - (2021-12-30)

## <!-- 26 -->🔄 Update

- Update README.md ([209d51b](https://github.com/tj-actions/auto-doc/commit/209d51b029d5da9f8c5a514ff2ccb69741481cb7))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#273](https://github.com/tj-actions/auto-doc/pull/273): ([eb28ac3](https://github.com/tj-actions/auto-doc/commit/eb28ac350cef42cafce727e848589463cf1f20aa))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.3 -> v1.2.4
 ([c0dcc89](https://github.com/tj-actions/auto-doc/commit/c0dcc89cbb8ce938488b51e39a46af76a81dc262))  - (jackton1)

# [1.2.4](https://github.com/tj-actions/auto-doc/compare/v1.2.3...v1.2.4) - (2021-12-30)

## <!-- 11 -->⏪ Reverts

- Revert "Remove goreleaser"
 ([8f19459](https://github.com/tj-actions/auto-doc/commit/8f19459c03711d10ce816c026624a193d6eca61d))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Update release.yml ([7883650](https://github.com/tj-actions/auto-doc/commit/7883650045ffc37858a03fc29bd7026597f676f0))  - (Tonye Jack)
- Update README.md ([6b6d184](https://github.com/tj-actions/auto-doc/commit/6b6d184db6e732d97ed0bae44f2965a0afb757d5))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#272](https://github.com/tj-actions/auto-doc/pull/272): ([f1e99ef](https://github.com/tj-actions/auto-doc/commit/f1e99efb1c159744bb223c79b94abedd43cedd96))  - (Tonye Jack)
- PR [#271](https://github.com/tj-actions/auto-doc/pull/271): ([843b022](https://github.com/tj-actions/auto-doc/commit/843b022fcf98fff3e2927e729bea1be00a066175))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.2 -> v1.2.3
 ([a2d2d57](https://github.com/tj-actions/auto-doc/commit/a2d2d574cc7cb9d53ce1f0895b611f0b395b4954))  - (jackton1)

# [1.2.3](https://github.com/tj-actions/auto-doc/compare/v1.2.2...v1.2.3) - (2021-12-30)

## <!-- 26 -->🔄 Update

- Update action.yml ([431bd21](https://github.com/tj-actions/auto-doc/commit/431bd21ac49d59adbcfc4821b7cdf17a47000cb6))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#270](https://github.com/tj-actions/auto-doc/pull/270): ([754f248](https://github.com/tj-actions/auto-doc/commit/754f248688763e62d54ce41ca5de6220325343cb))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.2.1 -> v1.2.2
 ([05b3480](https://github.com/tj-actions/auto-doc/commit/05b3480f29262ebbed401fa24f5118fa26e1dd73))  - (jackton1)

# [1.2.2](https://github.com/tj-actions/auto-doc/compare/v1.2.1...v1.2.2) - (2021-12-30)

## <!-- 1 -->🐛 Bug Fixes

- Fixed test target.
 ([33328fd](https://github.com/tj-actions/auto-doc/commit/33328fdc2fb18ed787999ef9d44ea4c88c1540f5))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Update action.yml ([71152ef](https://github.com/tj-actions/auto-doc/commit/71152ef8f028418602f360b18c3925a1d88ceb07))  - (Tonye Jack)
- Update README.md ([e1ac8ee](https://github.com/tj-actions/auto-doc/commit/e1ac8ee1e1bfc6fd5ac8d8d7668d62dbc1beab09))  - (Tonye Jack)
- Updated test.
 ([9d7ce8e](https://github.com/tj-actions/auto-doc/commit/9d7ce8e2eb59c823c4eb653b9c63e164bf3054b3))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- Merge branch 'main' of github.com:tj-actions/auto-doc
 ([8d06f85](https://github.com/tj-actions/auto-doc/commit/8d06f85d5568976f005d1f7d72a09689ddd35e02))  - (Tonye Jack)

# [1.2.1](https://github.com/tj-actions/auto-doc/compare/v1.2.0...v1.2.1) - (2021-12-28)

## <!-- 26 -->🔄 Update

- Update README.md ([fcd19ef](https://github.com/tj-actions/auto-doc/commit/fcd19efd080d583471c08881ed728e7b10c64915))  - (Tonye Jack)
- Update test.yml ([89d4fba](https://github.com/tj-actions/auto-doc/commit/89d4fbab18f29fcb050007a2769578b04f382f3c))  - (Tonye Jack)
- Update action.yml ([9c35767](https://github.com/tj-actions/auto-doc/commit/9c35767b36ed220dfc0e8ca31ee79abaeaea2a39))  - (Tonye Jack)
- Update action.yml ([da7470f](https://github.com/tj-actions/auto-doc/commit/da7470f71cef1474d9ab33cb8ca394c35b42ea70))  - (Tonye Jack)
- Update README.md ([2fe43c9](https://github.com/tj-actions/auto-doc/commit/2fe43c998ff5d4d625a9ff5d6154bee1cdb16915))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#252](https://github.com/tj-actions/auto-doc/pull/252): ([e1f8b21](https://github.com/tj-actions/auto-doc/commit/e1f8b21c8ede3ada50687c8fb1a4b1a305893995))  - (Tonye Jack)
- PR [#251](https://github.com/tj-actions/auto-doc/pull/251): ([69b8556](https://github.com/tj-actions/auto-doc/commit/69b8556ff0dc8ed92b808ad556132b3da5a8c5ea))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.1.9 -> v1.2.0
 ([d52e007](https://github.com/tj-actions/auto-doc/commit/d52e007ecfa4f0df222c4f501647fbe1553c7780))  - (jackton1)

# [1.2.0](https://github.com/tj-actions/auto-doc/compare/v1.1.9...v1.2.0) - (2021-12-28)

## <!-- 26 -->🔄 Update

- Update action.yml ([9ffeea1](https://github.com/tj-actions/auto-doc/commit/9ffeea16d398783568516c12192657d90757439b))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#250](https://github.com/tj-actions/auto-doc/pull/250): ([7454914](https://github.com/tj-actions/auto-doc/commit/74549144e243d9fd252b0af94b3cfa4aa2552b30))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.1.8 -> v1.1.9
 ([1765ca3](https://github.com/tj-actions/auto-doc/commit/1765ca32437a60275e3dfb78344713f090918534))  - (jackton1)

# [1.1.9](https://github.com/tj-actions/auto-doc/compare/v1.1.8...v1.1.9) - (2021-12-28)

## <!-- 1 -->🐛 Bug Fixes

- Fix ineffassign ([162adab](https://github.com/tj-actions/auto-doc/commit/162adabde19f33790b35bd39eb40520a1a014514))  - (Tonye Jack)
- Fix golint errors ([32c9609](https://github.com/tj-actions/auto-doc/commit/32c960905ce6cdc0c6bf1ceb2ecda747ba09ecef))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Update Makefile ([6789a5a](https://github.com/tj-actions/auto-doc/commit/6789a5ac8ecbe42dd300ab7f689c8e643589b89a))  - (Tonye Jack)
- Update Makefile ([e2bf70a](https://github.com/tj-actions/auto-doc/commit/e2bf70adb4246b0440144d9f0883b2ada8cd76ba))  - (Tonye Jack)
- Update root.go ([0fb9754](https://github.com/tj-actions/auto-doc/commit/0fb97544e10b70c708ee632f86d3d89923d3face))  - (Tonye Jack)
- Update root.go ([1c610fe](https://github.com/tj-actions/auto-doc/commit/1c610fed9b68009dbab9420f127fb5ed02fa7d65))  - (Tonye Jack)
- Update README.md ([12371ba](https://github.com/tj-actions/auto-doc/commit/12371ba5c5ea7df39cdfc4574742a9c8c605db98))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#249](https://github.com/tj-actions/auto-doc/pull/249): ([b74cfca](https://github.com/tj-actions/auto-doc/commit/b74cfca9c587719f9292339f5df801bb7b2ec47c))  - (Tonye Jack)
- PR [#248](https://github.com/tj-actions/auto-doc/pull/248): ([540930c](https://github.com/tj-actions/auto-doc/commit/540930cbab67936037d77cbfbaa5a75ac4e9cbdb))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.1.7 -> v1.1.8
 ([42c27e7](https://github.com/tj-actions/auto-doc/commit/42c27e7ff0705e9b06e2477960b19d7f6772c5fe))  - (jackton1)

# [1.1.8](https://github.com/tj-actions/auto-doc/compare/v1.1.7...v1.1.8) - (2021-12-27)

## <!-- 16 -->➕ Add

- Added support for configuring the max words per line in a column.
 ([9cdb41e](https://github.com/tj-actions/auto-doc/commit/9cdb41e993bc718dadda63e9ef4aa70d38990c33))  - (Tonye Jack)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([aa08c04](https://github.com/tj-actions/auto-doc/commit/aa08c04f4c829bab3797e13899bb75fa1b4c3ae1))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Update README.md ([27726b7](https://github.com/tj-actions/auto-doc/commit/27726b7f0f8eac84a913d10bd35379cfbfecce42))  - (Tonye Jack)
- Updated README.
 ([45826fb](https://github.com/tj-actions/auto-doc/commit/45826fb422b16ec781d8216cc9dc5fad0e881f79))  - (github-actions[bot])
- Update root.go ([7637480](https://github.com/tj-actions/auto-doc/commit/763748038b7b11592f8e4fc132310fb3447f6bf0))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#247](https://github.com/tj-actions/auto-doc/pull/247): support for configuring the max words per line in a column. ([ce2d032](https://github.com/tj-actions/auto-doc/commit/ce2d032da1b4df44e7c0facb8d8701801b55ee26))  - (Tonye Jack)
- Merge 763748038b7b11592f8e4fc132310fb3447f6bf0 into 6624525af4163455d00f0f0a638ff0dad38fffab
 ([02d2751](https://github.com/tj-actions/auto-doc/commit/02d2751d06b57a33893c2ab89f7bfe1dbd917206))  - (Tonye Jack)
- Merge 9cdb41e993bc718dadda63e9ef4aa70d38990c33 into 6624525af4163455d00f0f0a638ff0dad38fffab
 ([ea27b4b](https://github.com/tj-actions/auto-doc/commit/ea27b4b37738569e3612001495b8035343d67525))  - (Tonye Jack)
- PR [#246](https://github.com/tj-actions/auto-doc/pull/246): ([6624525](https://github.com/tj-actions/auto-doc/commit/6624525af4163455d00f0f0a638ff0dad38fffab))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.1.6 -> v1.1.7
 ([4bc6bc5](https://github.com/tj-actions/auto-doc/commit/4bc6bc5d1cf23b0495ef9a94bfae8c0cda4af5fa))  - (jackton1)

# [1.1.7](https://github.com/tj-actions/auto-doc/compare/v1.1.6...v1.1.7) - (2021-12-23)

## <!-- 13 -->📦 Bumps

- Bump github.com/spf13/cobra from 1.2.1 to 1.3.0

Bumps [github.com/spf13/cobra](https://github.com/spf13/cobra) from 1.2.1 to 1.3.0.
- [Release notes](https://github.com/spf13/cobra/releases)
- [Changelog](https://github.com/spf13/cobra/blob/master/CHANGELOG.md)
- [Commits](https://github.com/spf13/cobra/compare/v1.2.1...v1.3.0)

---
updated-dependencies:
- dependency-name: github.com/spf13/cobra
  dependency-type: direct:production
  update-type: version-update:semver-minor
...

Signed-off-by: dependabot[bot] <support@github.com> ([8d207ed](https://github.com/tj-actions/auto-doc/commit/8d207ed82746d6a9915c7c08a29699cfb4ed7fc8))  - (dependabot[bot])
- Bump tj-actions/verify-changed-files from 8.6 to 8.7

Bumps [tj-actions/verify-changed-files](https://github.com/tj-actions/verify-changed-files) from 8.6 to 8.7.
- [Release notes](https://github.com/tj-actions/verify-changed-files/releases)
- [Changelog](https://github.com/tj-actions/verify-changed-files/blob/main/HISTORY.md)
- [Commits](https://github.com/tj-actions/verify-changed-files/compare/v8.6...v8.7)

---
updated-dependencies:
- dependency-name: tj-actions/verify-changed-files
  dependency-type: direct:production
  update-type: version-update:semver-minor
...

Signed-off-by: dependabot[bot] <support@github.com> ([cd607db](https://github.com/tj-actions/auto-doc/commit/cd607db9bf778f3f02c3d4e9c4b0910d35a617dc))  - (dependabot[bot])

## <!-- 16 -->➕ Add

- Added support for wrapping long words ([f0ffecf](https://github.com/tj-actions/auto-doc/commit/f0ffecf308db479af997674914ba3d3801b64792))  - (Tonye Jack)
- Add Codacy badge ([241bcbf](https://github.com/tj-actions/auto-doc/commit/241bcbf18688672b5b18ffdb016879ec47d1a276))  - (The Codacy Badger)

## <!-- 17 -->➖ Remove

- Remove unused file ([0f32710](https://github.com/tj-actions/auto-doc/commit/0f327106d09d00876998308174c1d49ec7aecc64))  - (Tonye Jack)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([94f7884](https://github.com/tj-actions/auto-doc/commit/94f78848609b9e1ce23c933e96f0dd1e86c82e15))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Update README.md ([0f36219](https://github.com/tj-actions/auto-doc/commit/0f362193a2e017109b00ccb0d76fbdb06a65532c))  - (Tonye Jack)
- Update README.md ([e9aac72](https://github.com/tj-actions/auto-doc/commit/e9aac7212c96a1031a8d9473394ccf20357bdcf2))  - (Tonye Jack)
- Update README.md ([551086b](https://github.com/tj-actions/auto-doc/commit/551086bc7ec10ecfcce11b9de2249ecd96140116))  - (Tonye Jack)
- Update README.md ([aaca2d4](https://github.com/tj-actions/auto-doc/commit/aaca2d469dbbdc04be143237a1099c5bb85b05af))  - (Tonye Jack)
- Update update-readme.yml ([1806236](https://github.com/tj-actions/auto-doc/commit/18062366241800d74a2253020621f7eaef2d8907))  - (Tonye Jack)
- Updated the test.
 ([4a1f8b7](https://github.com/tj-actions/auto-doc/commit/4a1f8b7b32ed62cc61ff72db81d5bca8e6bb51bb))  - (Tonye Jack)
- Updated README.
 ([cfdb9e9](https://github.com/tj-actions/auto-doc/commit/cfdb9e923546ddd481af1efc66f8979e4f8f46d0))  - (github-actions[bot])
- Updated to fix formatting
 ([c8d90e0](https://github.com/tj-actions/auto-doc/commit/c8d90e067577d559f602cba0f66fa921c84f27ab))  - (Tonye Jack)
- Updated test.yml
 ([92df1f8](https://github.com/tj-actions/auto-doc/commit/92df1f875b42dcf39774f3020b70ef499eb4989d))  - (Tonye Jack)
- Updated README.
 ([740e75a](https://github.com/tj-actions/auto-doc/commit/740e75aaa37911e342d7486391df997433a93012))  - (github-actions[bot])
- Updated type column for outputs
 ([4797c6f](https://github.com/tj-actions/auto-doc/commit/4797c6f4c026eebeb9a13edb715dfbf342958d09))  - (Tonye Jack)
- Updated README.
 ([0664b6a](https://github.com/tj-actions/auto-doc/commit/0664b6a136748bd81ba20c9979b9596401366085))  - (github-actions[bot])
- Updated formatting of defaults.
 ([1949906](https://github.com/tj-actions/auto-doc/commit/19499064332ca0cb355db17dc6bef19511eeb44b))  - (Tonye Jack)
- Updated README.
 ([c8b3add](https://github.com/tj-actions/auto-doc/commit/c8b3add26166fd33b4b3f65fc3f23aa11e8c30f0))  - (github-actions[bot])
- Updated to include setting the column max width
 ([3d60605](https://github.com/tj-actions/auto-doc/commit/3d606054d5d2d93eb2e0f0c5686a94b3973ac7e6))  - (Tonye Jack)
- Updated test
 ([55bbd3d](https://github.com/tj-actions/auto-doc/commit/55bbd3d09a2e62a70fd99db78cd3f94469002009))  - (Tonye Jack)
- Update root.go ([4b92088](https://github.com/tj-actions/auto-doc/commit/4b92088367e08215913cbacc1ff915671b12fb1a))  - (Tonye Jack)
- Update root.go ([a1c9b31](https://github.com/tj-actions/auto-doc/commit/a1c9b31155e8f8ef858405304ec769f2619c3510))  - (Tonye Jack)
- Update root.go ([ba18621](https://github.com/tj-actions/auto-doc/commit/ba1862127eccd955e36becd731944457124dc628))  - (Tonye Jack)
- Updated README.md
 ([26f31db](https://github.com/tj-actions/auto-doc/commit/26f31db604a4af3bde97432aeb29b74039529e8c))  - (jackton1)
- Update tj-actions/verify-changed-files action to v8.8
 ([187527a](https://github.com/tj-actions/auto-doc/commit/187527a0c69ab9a9e72a0b5b6e7defb7be04aa86))  - (Renovate Bot)
- Updated .github/workflows/auto-merge.yml ([ec3defa](https://github.com/tj-actions/auto-doc/commit/ec3defa5791b3160fd87bfe025c8047f46a783eb))  - (Tonye Jack)
- Updated test ([5fff814](https://github.com/tj-actions/auto-doc/commit/5fff814856a3f7d3113356d2b1bfb7f72a6c56d5))  - (Tonye Jack)
- Update README.md ([08651df](https://github.com/tj-actions/auto-doc/commit/08651dfbff5a177d6358f62eb7ac3bd2d0738a61))  - (Tonye Jack)
- Update README.md ([53ecef6](https://github.com/tj-actions/auto-doc/commit/53ecef65f4ff9fc4081f7e099a1c38327b19ba11))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#244](https://github.com/tj-actions/auto-doc/pull/244): ([7d1932b](https://github.com/tj-actions/auto-doc/commit/7d1932b7c92c65715d3c7a7ee831b55d89be066a))  - (Tonye Jack)
- Merge branch 'chore/wrap-long-words' of github.com:tj-actions/auto-doc into chore/wrap-long-words
 ([d22fbc8](https://github.com/tj-actions/auto-doc/commit/d22fbc8bb5f3ff6c805e6d9de1f8dae592555b52))  - (Tonye Jack)
- Merge c8d90e067577d559f602cba0f66fa921c84f27ab into 792d1aa62e0459cdb999bef9bf1839d96ef423ed
 ([a570770](https://github.com/tj-actions/auto-doc/commit/a5707704e86d2974384aeb3a830223e005f2841c))  - (Tonye Jack)
- PR [#242](https://github.com/tj-actions/auto-doc/pull/242): ([792d1aa](https://github.com/tj-actions/auto-doc/commit/792d1aa62e0459cdb999bef9bf1839d96ef423ed))  - (Tonye Jack)
- Merge branch 'chore/wrap-long-words' of github.com:tj-actions/auto-doc into chore/wrap-long-words
 ([3f92876](https://github.com/tj-actions/auto-doc/commit/3f92876c531c8daf4f62adcc9cf62292beba5039))  - (Tonye Jack)
- Merge 05f0f1ff57789b8b78eb45aac4dc01244513b5c1 into 8d207ed82746d6a9915c7c08a29699cfb4ed7fc8
 ([fe1042c](https://github.com/tj-actions/auto-doc/commit/fe1042cb98a13bffde2681177664917acac693ab))  - (Tonye Jack)
- Merge branch 'chore/wrap-long-words' of github.com:tj-actions/auto-doc into chore/wrap-long-words
 ([05f0f1f](https://github.com/tj-actions/auto-doc/commit/05f0f1ff57789b8b78eb45aac4dc01244513b5c1))  - (Tonye Jack)
- Merge 5a8606101ce15ed47f895192bbe970c8dab7dca8 into 8d207ed82746d6a9915c7c08a29699cfb4ed7fc8
 ([ba3d75f](https://github.com/tj-actions/auto-doc/commit/ba3d75fd8e927a83a48a8ed89202a720b44fc539))  - (Tonye Jack)
- Merge branch 'chore/wrap-long-words' of github.com:tj-actions/auto-doc into chore/wrap-long-words
 ([5a86061](https://github.com/tj-actions/auto-doc/commit/5a8606101ce15ed47f895192bbe970c8dab7dca8))  - (Tonye Jack)
- Merge 3d606054d5d2d93eb2e0f0c5686a94b3973ac7e6 into 8d207ed82746d6a9915c7c08a29699cfb4ed7fc8
 ([0329609](https://github.com/tj-actions/auto-doc/commit/0329609a4cdc129e976b7aeb3ef5ae5b2bca73b6))  - (Tonye Jack)
- Merge f0ffecf308db479af997674914ba3d3801b64792 into 8d207ed82746d6a9915c7c08a29699cfb4ed7fc8
 ([c304b97](https://github.com/tj-actions/auto-doc/commit/c304b97654a8e3572daeb4ef03a5fee5d11460a6))  - (Tonye Jack)
- PR [#238](https://github.com/tj-actions/auto-doc/pull/238): bug copying new table data ([3d943dd](https://github.com/tj-actions/auto-doc/commit/3d943ddbffe05662485f73d4f7c91d44e580144d))  - (Tonye Jack)
- PR [#237](https://github.com/tj-actions/auto-doc/pull/237): unused file ([96af244](https://github.com/tj-actions/auto-doc/commit/96af244cc6b9df54ad87259fbdcb308b7011c4c6))  - (Tonye Jack)
- PR [#235](https://github.com/tj-actions/auto-doc/pull/235): README.md ([d843239](https://github.com/tj-actions/auto-doc/commit/d8432390ef3cc12c5f7456fe5050c9eb7967c11b))  - (Tonye Jack)
- PR [#234](https://github.com/tj-actions/auto-doc/pull/234): a Codacy badge to README.md ([0384b1e](https://github.com/tj-actions/auto-doc/commit/0384b1e40c65b3b80fb9d11362941424486a8a58))  - (Tonye Jack)
- PR [#232](https://github.com/tj-actions/auto-doc/pull/232): test ([77ec7f7](https://github.com/tj-actions/auto-doc/commit/77ec7f728508fd4b668308224ef65c9d939e9ac4))  - (Tonye Jack)
- PR [#230](https://github.com/tj-actions/auto-doc/pull/230): to v1.1.6 ([3631653](https://github.com/tj-actions/auto-doc/commit/3631653fab45b9017506e6fc78e8212fb8633d24))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.1.5 -> v1.1.6
 ([0cc8e5b](https://github.com/tj-actions/auto-doc/commit/0cc8e5bfee04573d70efb5f60b92d14181f658b0))  - (jackton1)

# [1.1.6](https://github.com/tj-actions/auto-doc/compare/v1.1.5...v1.1.6) - (2021-11-07)

## <!-- 30 -->📝 Other

- PR [#229](https://github.com/tj-actions/auto-doc/pull/229): unused code ([5634737](https://github.com/tj-actions/auto-doc/commit/563473756de17c6a9206699053ca7443dc725b69))  - (Tonye Jack)

# [1.1.5](https://github.com/tj-actions/auto-doc/compare/v1.1.4...v1.1.5) - (2021-11-07)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([e8cba34](https://github.com/tj-actions/auto-doc/commit/e8cba34d3373784567b6940a9bd06544a8a6976b))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Update root.go ([98a8f8e](https://github.com/tj-actions/auto-doc/commit/98a8f8e8f21494aeb8bf1154281ae68eeabd144d))  - (Tonye Jack)
- Update README.md ([ba63469](https://github.com/tj-actions/auto-doc/commit/ba634696af0535331723d6569ec9e041a8f78eba))  - (Tonye Jack)
- Updated README.md
 ([60a78b5](https://github.com/tj-actions/auto-doc/commit/60a78b5285be4bca357fd9d124ba83c3a051f68f))  - (jackton1)
- Update README.md ([b48ced7](https://github.com/tj-actions/auto-doc/commit/b48ced76d44e3287308f2dcc89a7604b65c70fb9))  - (Tonye Jack)
- Update test.yml ([2033609](https://github.com/tj-actions/auto-doc/commit/203360924d84a7e90aae0d2a949d4b0eee99d3e4))  - (Tonye Jack)
- Updated README.
 ([84164fc](https://github.com/tj-actions/auto-doc/commit/84164fc93bbabeb12656dcf224ef1f3a1212e09a))  - (github-actions[bot])
- Update action.yml ([598deb5](https://github.com/tj-actions/auto-doc/commit/598deb5e24b7653965abcfefcdfa77d21b5e5821))  - (Tonye Jack)
- Updated README.
 ([10eee3f](https://github.com/tj-actions/auto-doc/commit/10eee3fa243e69a91bbdb1d0d8ae470839bd7a59))  - (github-actions[bot])
- Update README.md ([1df882b](https://github.com/tj-actions/auto-doc/commit/1df882b06538b21b776af885c6e0378a4f2b1dd9))  - (Tonye Jack)
- Updated README.
 ([d357de9](https://github.com/tj-actions/auto-doc/commit/d357de9d147b24fdd9b4bdc5f85bc251415e743d))  - (github-actions[bot])
- Updated README.
 ([4bb81e5](https://github.com/tj-actions/auto-doc/commit/4bb81e5f077167bd8f26cd789656566b500a80d2))  - (github-actions[bot])
- Update test.yml ([b8b37d5](https://github.com/tj-actions/auto-doc/commit/b8b37d5741fe3618428264b8aa21b61ca1282337))  - (Tonye Jack)
- Update test.yml ([2be762e](https://github.com/tj-actions/auto-doc/commit/2be762e36f1e394477138cdca24934d8b5c9ef41))  - (Tonye Jack)
- Updated output columns ([6d010e0](https://github.com/tj-actions/auto-doc/commit/6d010e0e26289bedda96ebbc6803d37eed841a18))  - (Tonye Jack)
- Updated README.md
 ([15a48da](https://github.com/tj-actions/auto-doc/commit/15a48dafa8c126e25b7dc5fe88d83a55ba98b43d))  - (jackton1)
- Update root.go ([a8ad424](https://github.com/tj-actions/auto-doc/commit/a8ad424f55f2444612a6bb237c1b5425177aff11))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#228](https://github.com/tj-actions/auto-doc/pull/228): to 1.1.5 ([3460517](https://github.com/tj-actions/auto-doc/commit/346051744f0192fda5a145a5b3320cf78680a298))  - (Tonye Jack)
- PR [#226](https://github.com/tj-actions/auto-doc/pull/226): README.md ([5259f85](https://github.com/tj-actions/auto-doc/commit/5259f851777600874b3e06f38388c45f2633116b))  - (Tonye Jack)
- PR [#224](https://github.com/tj-actions/auto-doc/pull/224): output columns ([8d91719](https://github.com/tj-actions/auto-doc/commit/8d91719b7c3a5d9bc0b23db6816462b9bfb33100))  - (Tonye Jack)
- Merge 598deb5e24b7653965abcfefcdfa77d21b5e5821 into 2c9ba9b2fee2a3ace825b90d47d3c5a99bbd0f50
 ([dbe9572](https://github.com/tj-actions/auto-doc/commit/dbe95728d5cfda3e956e5404d6478ad19675f9c2))  - (Tonye Jack)
- Merge 1df882b06538b21b776af885c6e0378a4f2b1dd9 into 2c9ba9b2fee2a3ace825b90d47d3c5a99bbd0f50
 ([82921d2](https://github.com/tj-actions/auto-doc/commit/82921d2e3ec53cf4602806a39acbe40fb5328a35))  - (Tonye Jack)
- Merge 4bb81e5f077167bd8f26cd789656566b500a80d2 into 2c9ba9b2fee2a3ace825b90d47d3c5a99bbd0f50
 ([fa74e17](https://github.com/tj-actions/auto-doc/commit/fa74e17c3fb77e323ff3ea56f918220d674da43a))  - (Tonye Jack)
- Merge b8b37d5741fe3618428264b8aa21b61ca1282337 into 2c9ba9b2fee2a3ace825b90d47d3c5a99bbd0f50
 ([4f43171](https://github.com/tj-actions/auto-doc/commit/4f43171cf1395bb9711ce25150be066cfad7a596))  - (Tonye Jack)
- PR [#221](https://github.com/tj-actions/auto-doc/pull/221): README.md ([2c9ba9b](https://github.com/tj-actions/auto-doc/commit/2c9ba9b2fee2a3ace825b90d47d3c5a99bbd0f50))  - (Tonye Jack)
- PR [#222](https://github.com/tj-actions/auto-doc/pull/222): to v1.1.4 ([57ce307](https://github.com/tj-actions/auto-doc/commit/57ce307aa5cd0e5730ec166255b8fede4c3cfdcc))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.1.4 -> 1.1.5
 ([e87e5f2](https://github.com/tj-actions/auto-doc/commit/e87e5f2a3637b39c228f819b5a637b8ba8a23ed2))  - (jackton1)
- Upgraded from v1.1.3 -> v1.1.4
 ([a362779](https://github.com/tj-actions/auto-doc/commit/a362779dfe2c8216f9adf6530ca328eb41d1e8e4))  - (jackton1)

# [1.1.4](https://github.com/tj-actions/auto-doc/compare/v1.1.3...v1.1.4) - (2021-11-06)

## <!-- 16 -->➕ Add

- Added .gitattributes
 ([8608d93](https://github.com/tj-actions/auto-doc/commit/8608d93206d3abc976571c1813af53b242f255f0))  - (Tonye Jack)

## <!-- 17 -->➖ Remove

- Removed unused code.
 ([6b997b9](https://github.com/tj-actions/auto-doc/commit/6b997b959d07a385e105936e790fe07b8fab0091))  - (Jack tonye)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([6c0c1f5](https://github.com/tj-actions/auto-doc/commit/6c0c1f53e1f8727daba52aa7acb7b7f56d88d16e))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Update README.md ([855a448](https://github.com/tj-actions/auto-doc/commit/855a448ab30435facd0450e93ef9518d06bd5f1e))  - (Tonye Jack)
- Update README.md ([7b8d9fc](https://github.com/tj-actions/auto-doc/commit/7b8d9fcaf3e4d422c6920b1f93e7bb36dc679953))  - (Tonye Jack)
- Updated README
 ([39334d6](https://github.com/tj-actions/auto-doc/commit/39334d672101b847006a8ec4d746e54a4fcf62a2))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#219](https://github.com/tj-actions/auto-doc/pull/219): to v1.1.3 ([943f029](https://github.com/tj-actions/auto-doc/commit/943f029befd6322369235a78fe3f604fd2252031))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.1.2 -> v1.1.3
 ([b75868f](https://github.com/tj-actions/auto-doc/commit/b75868f4364e524736e08af44f2340e0d1ce48b8))  - (jackton1)

# [1.1.3](https://github.com/tj-actions/auto-doc/compare/v1.1.2...v1.1.3) - (2021-11-06)

## <!-- 1 -->🐛 Bug Fixes

- Fixed spacing bug.
 ([961b2be](https://github.com/tj-actions/auto-doc/commit/961b2befd34c3d8f50acad2e427b0765d67e9f11))  - (Tonye Jack)
- Fixed spacing bug.
 ([268a1f6](https://github.com/tj-actions/auto-doc/commit/268a1f6a839e6d70dcf97f8faf673d9f419ea701))  - (Tonye Jack)
- Fixed formatting.
 ([0ab1b58](https://github.com/tj-actions/auto-doc/commit/0ab1b58e763954e9e5ba71c30150c9a7cc490466))  - (Tonye Jack)

# [1.1.2](https://github.com/tj-actions/auto-doc/compare/v1.1.1...v1.1.2) - (2021-11-06)

## <!-- 1 -->🐛 Bug Fixes

- Fixed spacing
 ([9561a1b](https://github.com/tj-actions/auto-doc/commit/9561a1b13c25e23a4dcda8350e8446d0148d1acf))  - (Tonye Jack)
- Fixed bug with README.md
 ([b37e297](https://github.com/tj-actions/auto-doc/commit/b37e29727299a169b6838580a2455998113b0ee0))  - (Tonye Jack)
- Fixed new line bug
 ([64c773c](https://github.com/tj-actions/auto-doc/commit/64c773cc14772e89408097e58718553f98701616))  - (Tonye Jack)
- Fixed github action.
 ([459bc8e](https://github.com/tj-actions/auto-doc/commit/459bc8e5ca88bb4939c416d2108dafc411410a11))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Updated README.
 ([60d777e](https://github.com/tj-actions/auto-doc/commit/60d777ee0d5821bea32ce878e2085271cb7efde3))  - (github-actions[bot])
- Updated README.
 ([cd65d2d](https://github.com/tj-actions/auto-doc/commit/cd65d2d927679e2b8ffac6ed1150bcd27ed857d5))  - (github-actions[bot])
- Updated README.
 ([9a1cd02](https://github.com/tj-actions/auto-doc/commit/9a1cd020e25746c3bad4dc313cfbfacc2462b504))  - (github-actions[bot])
- Updated github action
 ([1099529](https://github.com/tj-actions/auto-doc/commit/109952999b2bda06c94c59a50158974092181862))  - (Tonye Jack)
- Updated README.
 ([27350d5](https://github.com/tj-actions/auto-doc/commit/27350d56df6aedc9e1d5c099edd828cef6a5b72f))  - (github-actions[bot])

## <!-- 30 -->📝 Other

- Merge b37e29727299a169b6838580a2455998113b0ee0 into 27350d56df6aedc9e1d5c099edd828cef6a5b72f
 ([5c16f72](https://github.com/tj-actions/auto-doc/commit/5c16f7229f7bc2f6af451113d30343541a228eac))  - (Tonye Jack)
- Merge 9a1cd020e25746c3bad4dc313cfbfacc2462b504 into 27350d56df6aedc9e1d5c099edd828cef6a5b72f
 ([117e289](https://github.com/tj-actions/auto-doc/commit/117e28914215afb3ea17e2c2df2c27ec07810959))  - (Tonye Jack)
- Merge 109952999b2bda06c94c59a50158974092181862 into 27350d56df6aedc9e1d5c099edd828cef6a5b72f
 ([135a182](https://github.com/tj-actions/auto-doc/commit/135a182f9e5f217b13bc67f42cc421f3ca241c97))  - (Tonye Jack)
- PR [#214](https://github.com/tj-actions/auto-doc/pull/214): to v1.1.1 ([a78915d](https://github.com/tj-actions/auto-doc/commit/a78915dfc5bfc11ae6c49ca1d7ba3ccba8186a8f))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.1.0 -> v1.1.1
 ([6308e87](https://github.com/tj-actions/auto-doc/commit/6308e870e88bf0ecbe07eb31ab14e1e462b1d2e6))  - (jackton1)

# [1.1.1](https://github.com/tj-actions/auto-doc/compare/v1.1.0...v1.1.1) - (2021-11-06)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([6dac490](https://github.com/tj-actions/auto-doc/commit/6dac490aa2471b1b29d1143c4e58b54166905c43))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Updated version update action.
 ([e809080](https://github.com/tj-actions/auto-doc/commit/e80908044196e97fbe19dce0c0ecdd9022bd079d))  - (Tonye Jack)
- Updated platform support.
 ([60fd320](https://github.com/tj-actions/auto-doc/commit/60fd3208f95ff79a85ebac45783f561eb3353dae))  - (Tonye Jack)
- Updated platform support.
 ([7db37bb](https://github.com/tj-actions/auto-doc/commit/7db37bb96c54ed80ff23d0b52c1c9f2aff3b2dc6))  - (Tonye Jack)
- Updated README.
 ([7e41b95](https://github.com/tj-actions/auto-doc/commit/7e41b95a579982a145b9fa8704847dd545a2b247))  - (github-actions[bot])
- Updated README.
 ([1bc1235](https://github.com/tj-actions/auto-doc/commit/1bc1235cbefffe70f7d79b53118584d4f9c81180))  - (github-actions[bot])
- Update action.yml ([f95dd9b](https://github.com/tj-actions/auto-doc/commit/f95dd9b1e0432ac18eb5530de2610f91bbc6cf8e))  - (Tonye Jack)
- Updated README.
 ([6180977](https://github.com/tj-actions/auto-doc/commit/61809778dd9359ceea05cb243a2f6c4d51d234a2))  - (github-actions[bot])

## <!-- 30 -->📝 Other

- PR [#213](https://github.com/tj-actions/auto-doc/pull/213): to v1.1.0 ([cf2e14b](https://github.com/tj-actions/auto-doc/commit/cf2e14b9e02c02c15eee13fd1c35befe592a2bd3))  - (Tonye Jack)
- Merge 236b4103a011e543a4d0fe29afe3ea368a4537ae into f95dd9b1e0432ac18eb5530de2610f91bbc6cf8e
 ([6fb9ee0](https://github.com/tj-actions/auto-doc/commit/6fb9ee00c7e27d854cdf643b7101990fd10c1ade))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.0.1 -> v1.1.0
 ([236b410](https://github.com/tj-actions/auto-doc/commit/236b4103a011e543a4d0fe29afe3ea368a4537ae))  - (jackton1)

# [1.1.0](https://github.com/tj-actions/auto-doc/compare/v1.0.1...v1.1.0) - (2021-11-06)

## <!-- 10 -->⬇️ Downgrades

- Downgraded github action.
 ([483ef0f](https://github.com/tj-actions/auto-doc/commit/483ef0f64a9ccd68598a417dbd2377c5bdc0ab65))  - (Tonye Jack)

## <!-- 16 -->➕ Add

- Added support for windows.
 ([520e1b3](https://github.com/tj-actions/auto-doc/commit/520e1b391d5fa72f7f29cec215370cfc14e404b3))  - (Tonye Jack)

## <!-- 17 -->➖ Remove

- Removed duplicate code.
 ([a09a892](https://github.com/tj-actions/auto-doc/commit/a09a8928369565781eadc502cc78f583582964ed))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Updated README.
 ([efa45e8](https://github.com/tj-actions/auto-doc/commit/efa45e8cb4c6a4e52c8537d7a832417e29d8eaa7))  - (github-actions[bot])
- Updated README.
 ([1f12c24](https://github.com/tj-actions/auto-doc/commit/1f12c24c113394c71b97ef9b9c65b421a4b3abea))  - (github-actions[bot])
- Updated README.
 ([e8174ce](https://github.com/tj-actions/auto-doc/commit/e8174ce2dbb67e7a2e480a76c599c78177300bd3))  - (github-actions[bot])
- Updated README.
 ([411fe9c](https://github.com/tj-actions/auto-doc/commit/411fe9ca66e1085456b9ad1c658f09a67e386993))  - (github-actions[bot])
- Updated README.
 ([5a87e9e](https://github.com/tj-actions/auto-doc/commit/5a87e9ebff07faf14ceee366bedb242aa34340b6))  - (github-actions[bot])
- Updated README.
 ([01be86d](https://github.com/tj-actions/auto-doc/commit/01be86d4755d2e2d5aea89a107815975bc169641))  - (github-actions[bot])
- Updated README.
 ([47e177f](https://github.com/tj-actions/auto-doc/commit/47e177fa5f3fe75c82a7090e495a1b553fd53843))  - (github-actions[bot])
- Updated README.
 ([10feae2](https://github.com/tj-actions/auto-doc/commit/10feae290a1375a8d400f2951f7b5c644bc333e9))  - (github-actions[bot])
- Updated README.
 ([501d95b](https://github.com/tj-actions/auto-doc/commit/501d95b03571b8e2bb058ce2b6521b90fddc070c))  - (github-actions[bot])
- Updated README.
 ([1a2d443](https://github.com/tj-actions/auto-doc/commit/1a2d4432dc803190d57582c44fc8154c0f28f879))  - (github-actions[bot])
- Updated README.
 ([9a1e118](https://github.com/tj-actions/auto-doc/commit/9a1e118f2f7f3bff5f7b3ab676949d5e5ce50b66))  - (github-actions[bot])
- Updated README.
 ([b9f8d92](https://github.com/tj-actions/auto-doc/commit/b9f8d927b0844a55b8b61d5b01a4e08b8aa02f84))  - (github-actions[bot])
- Updated README.
 ([99fedda](https://github.com/tj-actions/auto-doc/commit/99feddac1f6a8d2f3a8ff0b767d0b291017c7ab0))  - (github-actions[bot])
- Updated README.
 ([76f5e48](https://github.com/tj-actions/auto-doc/commit/76f5e485ce908651592f036012684cc82ace3f62))  - (github-actions[bot])
- Updated README.
 ([f936209](https://github.com/tj-actions/auto-doc/commit/f936209fdd0cea2a7c849ce35c99bd85f6c7225d))  - (github-actions[bot])
- Updated README.
 ([a7791df](https://github.com/tj-actions/auto-doc/commit/a7791dfe6642cd6110f1965082acc27845938568))  - (github-actions[bot])
- Updated README.
 ([24f9706](https://github.com/tj-actions/auto-doc/commit/24f97062d465a5cc2c41f3e8b5e85b13ceaebb2e))  - (github-actions[bot])
- Updated README.
 ([7dad197](https://github.com/tj-actions/auto-doc/commit/7dad1972c22756dfb934add9aef5b9bdb80ad2bb))  - (github-actions[bot])
- Update tj-actions/verify-changed-files action to v8.7
 ([7eea5f5](https://github.com/tj-actions/auto-doc/commit/7eea5f555038893c2024e09e6b64a6f05339b0cf))  - (Renovate Bot)
- Updated README.
 ([d902e99](https://github.com/tj-actions/auto-doc/commit/d902e99bb049d2bb846406f80097e9e21832e116))  - (github-actions[bot])
- Updated README.
 ([9292170](https://github.com/tj-actions/auto-doc/commit/9292170f21cc209ef2225995760e76ac0f63de36))  - (github-actions[bot])
- Updated README.
 ([53c41b8](https://github.com/tj-actions/auto-doc/commit/53c41b8a7a0aee6b1ab0f8a8e55a342a1d4772a8))  - (github-actions[bot])
- Updated test.
 ([8c5c1b1](https://github.com/tj-actions/auto-doc/commit/8c5c1b1c4d93fb07d885786f11b045d926bb14c7))  - (Tonye Jack)
- Updated README.
 ([6052a1a](https://github.com/tj-actions/auto-doc/commit/6052a1a8d3532e04879e2fdca0386c1009bfbbf5))  - (github-actions[bot])
- Updated README.
 ([17ecad8](https://github.com/tj-actions/auto-doc/commit/17ecad8a4e2b90579953966eaba715853bbb57f6))  - (github-actions[bot])
- Updated README.
 ([fd72963](https://github.com/tj-actions/auto-doc/commit/fd72963d133156c992e055fabbc4da9e86915b38))  - (github-actions[bot])
- Updated test.
 ([e559737](https://github.com/tj-actions/auto-doc/commit/e559737f8558bc8b04c5fa0a09213aaef9fb8478))  - (Tonye Jack)
- Updated README.
 ([5f1f725](https://github.com/tj-actions/auto-doc/commit/5f1f725795f8a109de9d7303dee2db7b47c90187))  - (github-actions[bot])
- Updated README.
 ([79fd5cd](https://github.com/tj-actions/auto-doc/commit/79fd5cdfde6452cb56e4f7f547e7ac081e676107))  - (github-actions[bot])
- Updated README.
 ([77788de](https://github.com/tj-actions/auto-doc/commit/77788ded383477dccb0b187251a022f3dd314dc4))  - (github-actions[bot])
- Updated README.
 ([bdbaf25](https://github.com/tj-actions/auto-doc/commit/bdbaf2578a708027b281aa31ca7f464145e0dec8))  - (github-actions[bot])
- Updated action.
 ([5f964cb](https://github.com/tj-actions/auto-doc/commit/5f964cbac316794a15bf38619b5091523d600c83))  - (Tonye Jack)
- Updated README.
 ([bd1da4f](https://github.com/tj-actions/auto-doc/commit/bd1da4f4bab110591cfc4ced5acb89f243459fc5))  - (github-actions[bot])
- Updated action.
 ([658bfb0](https://github.com/tj-actions/auto-doc/commit/658bfb0f8cba93e241d10fc17674f5c017efffab))  - (Tonye Jack)
- Updated action.
 ([d24cac6](https://github.com/tj-actions/auto-doc/commit/d24cac615797e6ddee3acdf9c02c1df509124403))  - (Tonye Jack)
- Updated action.
 ([86bf66c](https://github.com/tj-actions/auto-doc/commit/86bf66c87f52d6842f916f4517110ed9bf608f3f))  - (Tonye Jack)
- Updated action.
 ([1c8e2cd](https://github.com/tj-actions/auto-doc/commit/1c8e2cd8e586d05fac11a6c6fae51616f382b285))  - (Tonye Jack)
- Update README.md ([dc8ab5c](https://github.com/tj-actions/auto-doc/commit/dc8ab5cfc261315bbd4d2a49616a2b3331efd4bc))  - (Tonye Jack)
- Update greetings.yml ([f09a462](https://github.com/tj-actions/auto-doc/commit/f09a462dbb208985ce4eaefeaad824a4dd70e8f2))  - (Tonye Jack)
- Update update-readme.yml ([169632f](https://github.com/tj-actions/auto-doc/commit/169632feea9e6e3de250930097b0f1c8adbc518f))  - (Tonye Jack)
- Update test.yml ([f15232b](https://github.com/tj-actions/auto-doc/commit/f15232b4fd0a59a6d9120274c6eaf2e4c12a55ed))  - (Tonye Jack)
- Update README.md ([3a82cc1](https://github.com/tj-actions/auto-doc/commit/3a82cc1228ef82995a136dfbeec274fbfff3c8b9))  - (Tonye Jack)
- Update test.yml ([e5fdb7c](https://github.com/tj-actions/auto-doc/commit/e5fdb7cd54f9eccf2e55ddba89e8a519674aeae6))  - (Tonye Jack)
- Update test.yml ([0056dc2](https://github.com/tj-actions/auto-doc/commit/0056dc2eca9e519b9c8721472c3dfc80fcc71bdd))  - (Tonye Jack)
- Update test.yml ([590d169](https://github.com/tj-actions/auto-doc/commit/590d16937252b023a2f86b82f29c6980c7d77d12))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#88](https://github.com/tj-actions/auto-doc/pull/88): to v1.0.1 ([5b2e403](https://github.com/tj-actions/auto-doc/commit/5b2e4034ca6ca08175387debcf411a4fcfb2d87d))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.1 -> v1.0.1
 ([7a8a107](https://github.com/tj-actions/auto-doc/commit/7a8a107bd4ddba71c4e805842c62661e5a5fc79a))  - (jackton1)

# [1.0.1](https://github.com/tj-actions/auto-doc/compare/v1.1...v1.0.1) - (2021-11-05)

## <!-- 26 -->🔄 Update

- Updated test.
 ([67391e6](https://github.com/tj-actions/auto-doc/commit/67391e6d12e8ee4761b2aeb35197f17899689464))  - (Tonye Jack)
- Updated test.
 ([981abdd](https://github.com/tj-actions/auto-doc/commit/981abddd4450ec6f00fcc1edb1f60be5a9e1ca17))  - (Tonye Jack)
- Update README.md ([cd31f30](https://github.com/tj-actions/auto-doc/commit/cd31f307fee43b3f3f426577d0edeec491fcc0a3))  - (Tonye Jack)
- Update test.yml ([63f3e08](https://github.com/tj-actions/auto-doc/commit/63f3e08a1f4f7b17fa300c953fcd0fd246a26991))  - (Tonye Jack)
- Update README.md ([18d2734](https://github.com/tj-actions/auto-doc/commit/18d27346d8de1580bc98edfebf3a53500a5eb8e2))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#38](https://github.com/tj-actions/auto-doc/pull/38): to v1.1 ([b32ca86](https://github.com/tj-actions/auto-doc/commit/b32ca86dd423f7b1ca2981b22958224b04b0bccb))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1.0.0 -> v1.1
 ([2e90c28](https://github.com/tj-actions/auto-doc/commit/2e90c28f2a8568c7cdbece3d7aa730230daa644b))  - (jackton1)

# [1.1](https://github.com/tj-actions/auto-doc/compare/v1.0.0...v1.1) - (2021-11-05)

## <!-- 1 -->🐛 Bug Fixes

- Fixed typo.
 ([bd3c46c](https://github.com/tj-actions/auto-doc/commit/bd3c46c13a8a1cddfcc98b4bd133c499fcf11133))  - (Tonye Jack)

## <!-- 26 -->🔄 Update

- Update action.yml ([b16467f](https://github.com/tj-actions/auto-doc/commit/b16467fa624dda0c1c5ffd8fdc7012aa2c5c9c70))  - (Tonye Jack)
- Updated test.
 ([95c414c](https://github.com/tj-actions/auto-doc/commit/95c414cad26f4f701346e32d4d8aeb66734ad906))  - (Tonye Jack)
- Updated action.
 ([36f21cd](https://github.com/tj-actions/auto-doc/commit/36f21cdf8437c60ce2219a42d4fe5391b28a5c27))  - (Tonye Jack)
- Update README.md ([c037210](https://github.com/tj-actions/auto-doc/commit/c0372101e39d56fc3b180d3ad1721cec47db8795))  - (Tonye Jack)
- Updated README.md
 ([4c98561](https://github.com/tj-actions/auto-doc/commit/4c9856198be08ed856887fa7e0ac064803c35601))  - (jackton1)
- Update README.md ([8b40470](https://github.com/tj-actions/auto-doc/commit/8b40470b9007556c677ca02226610091da14a40f))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#37](https://github.com/tj-actions/auto-doc/pull/37): action. ([f10f986](https://github.com/tj-actions/auto-doc/commit/f10f986741993f2c0083c6d4a44f7e8e90715887))  - (Tonye Jack)
- PR [#36](https://github.com/tj-actions/auto-doc/pull/36): README.md ([6076112](https://github.com/tj-actions/auto-doc/commit/6076112f944974a870c5edec7cb50f2fc1452f48))  - (Tonye Jack)
- PR [#35](https://github.com/tj-actions/auto-doc/pull/35): to v1.0.0 ([e1806ac](https://github.com/tj-actions/auto-doc/commit/e1806ac69122112a3ea79ec3253038f82bce4939))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1 -> v1.0.0
 ([e15cfef](https://github.com/tj-actions/auto-doc/commit/e15cfef611e8986c5c370b90b0ea5880d2e4ead4))  - (jackton1)

# [1.0.0](https://github.com/tj-actions/auto-doc/tree/v1.0.0) - (2021-11-05)

## <!-- 1 -->🐛 Bug Fixes

- Fixed test.
 ([1d607ea](https://github.com/tj-actions/auto-doc/commit/1d607eafa17b351ee2f2ffda9119654083bef2c5))  - (Tonye Jack)
- Fixed lint error.
 ([0ffbac4](https://github.com/tj-actions/auto-doc/commit/0ffbac45e2ea164e8e388e30f3e6349504bd2299))  - (Tonye Jack)
- Fixed lint error.
 ([7f9423b](https://github.com/tj-actions/auto-doc/commit/7f9423b87cbb9c3a55052eced3f16562ce2032f7))  - (Tonye Jack)
- Fixed line ending.
 ([dbf3bde](https://github.com/tj-actions/auto-doc/commit/dbf3bdef3ad096b224059f393dfc7aecd4748729))  - (Tonye Jack)
- Fixed test.
 ([122d91b](https://github.com/tj-actions/auto-doc/commit/122d91bf39ace56c169f1abce940a3e6a316e51b))  - (Tonye Jack)
- Fixed test.
 ([45e81fa](https://github.com/tj-actions/auto-doc/commit/45e81fa37b2525aec2b6cff4678fa899c9660dd5))  - (Tonye Jack)
- Fixed test.
 ([1f4b96c](https://github.com/tj-actions/auto-doc/commit/1f4b96c600cf0349d06523e909fa0e52dbfe9ff5))  - (Tonye Jack)

## <!-- 16 -->➕ Add

- Create test.yml ([ed4b547](https://github.com/tj-actions/auto-doc/commit/ed4b54728882db24804a6bbdb81a63b2f9c4ce50))  - (Tonye Jack)
- Added go modules
 ([5d1d69d](https://github.com/tj-actions/auto-doc/commit/5d1d69d19df6866a4f6d4d59badfa8140183288e))  - (Tonye Jack)
- Create README.md ([24baad9](https://github.com/tj-actions/auto-doc/commit/24baad9f1a429ca238c7f4e99386727df00089ad))  - (Tonye Jack)
- Added .github/workflows/codacy-analysis.yml ([aa5b41b](https://github.com/tj-actions/auto-doc/commit/aa5b41bfc5f1a3599be4e108ce483a1cd3c98263))  - (Tonye Jack)
- Added .github/workflows/greetings.yml ([22cd046](https://github.com/tj-actions/auto-doc/commit/22cd046c9ad01cbb0ab449f5abd10aa80c13d904))  - (Tonye Jack)
- Added CODE_OF_CONDUCT.md ([690b574](https://github.com/tj-actions/auto-doc/commit/690b574858b3fa4bfec3f5e6c300e05faf753cc2))  - (Tonye Jack)
- Added .github/ISSUE_TEMPLATE/feature_request.yaml ([e0948d0](https://github.com/tj-actions/auto-doc/commit/e0948d07cacef8116c8cfb291c610fa9ab4b10d9))  - (Tonye Jack)
- Added .github/ISSUE_TEMPLATE/bug_report.yaml ([7c10da6](https://github.com/tj-actions/auto-doc/commit/7c10da60a7f14721c594189f8348851fe2ee0c61))  - (Tonye Jack)
- Added .whitesource ([3ea9bca](https://github.com/tj-actions/auto-doc/commit/3ea9bca82c7d0baf00dd29d516f58a2169d39fbc))  - (Tonye Jack)
- Added .github/workflows/update-readme.yml ([a6d73b0](https://github.com/tj-actions/auto-doc/commit/a6d73b058c76eb6060bcfca7514125d3c9bd87f8))  - (Tonye Jack)
- Added .github/workflows/auto-merge.yml ([2accd71](https://github.com/tj-actions/auto-doc/commit/2accd71f04a568ec7a1ff5157ce734681e382593))  - (Tonye Jack)
- Added renovate.json ([a7e2483](https://github.com/tj-actions/auto-doc/commit/a7e248322ce15c5d4caddec4bfa3ebc39f109ae3))  - (Tonye Jack)
- Added .github/ISSUE_TEMPLATE/feature_request.md ([b586992](https://github.com/tj-actions/auto-doc/commit/b586992e2830b985edd5a4f0f5516f2a312148f5))  - (Tonye Jack)
- Added .github/ISSUE_TEMPLATE/bug_report.md ([0858101](https://github.com/tj-actions/auto-doc/commit/085810145ce14d15c16049a6668875ab71583911))  - (Tonye Jack)
- Added .github/workflows/sync-release-version.yml ([8a069f1](https://github.com/tj-actions/auto-doc/commit/8a069f18408fbb5ac2712c87fcc53b4ff5802d8f))  - (Tonye Jack)
- Create generator.py ([a587ed6](https://github.com/tj-actions/auto-doc/commit/a587ed68332bc713a6a3b57c7eebb0b7cbbd0309))  - (Tonye Jack)
- Added .github/auto-approve.yml ([a3af724](https://github.com/tj-actions/auto-doc/commit/a3af724a912fba57dd4af82b4c366bd13320f2f3))  - (Tonye Jack)
- Added .github/FUNDING.yml ([4d324ac](https://github.com/tj-actions/auto-doc/commit/4d324ac1b5f589f36a19aab84a1400c27a2d5804))  - (Tonye Jack)
- Added .github/workflows/auto-approve.yml ([8fbee17](https://github.com/tj-actions/auto-doc/commit/8fbee177e53ccd82f491e583fdab3be79758a5d4))  - (Tonye Jack)

## <!-- 17 -->➖ Remove

- Removed unused code.
 ([ccb1b18](https://github.com/tj-actions/auto-doc/commit/ccb1b18f3bb8aef72c10e28e753eb9f0c0b2112d))  - (Tonye Jack)
- Deleted .github/workflows/codacy-analysis.yml ([b2796b9](https://github.com/tj-actions/auto-doc/commit/b2796b96a96c1d891617b459726646c781fd789d))  - (Tonye Jack)
- Deleted .github/ISSUE_TEMPLATE/feature_request.md ([23536b8](https://github.com/tj-actions/auto-doc/commit/23536b800a6398fc0e03f8ff0a9744e8cf73ac42))  - (Tonye Jack)
- Deleted .github/ISSUE_TEMPLATE/bug_report.md ([7f6e709](https://github.com/tj-actions/auto-doc/commit/7f6e70925d612409e713f9d73e9eaa2375def012))  - (Tonye Jack)
- Deleted .github/auto-approve.yml ([6b22219](https://github.com/tj-actions/auto-doc/commit/6b222191acc37b19d45d1e83bbe7c359943bea42))  - (Tonye Jack)

## <!-- 25 -->🎨 Format

- Formatted code.
 ([f54565b](https://github.com/tj-actions/auto-doc/commit/f54565bdf4fb59fe90c06f8ebc6e31ccc04a1c71))  - (github-actions[bot])
- Formatted code.
 ([3ffa875](https://github.com/tj-actions/auto-doc/commit/3ffa87544bc8776374f6fae67047da5d8028cd85))  - (github-actions[bot])
- Formatted code.
 ([881a08b](https://github.com/tj-actions/auto-doc/commit/881a08bdd920b4d6bf44174d7399ec5d6ca9aa9e))  - (github-actions[bot])

## <!-- 26 -->🔄 Update

- Updated README.md
 ([62c5c0b](https://github.com/tj-actions/auto-doc/commit/62c5c0b997e1f015676bcb848a3444f6fd162c61))  - (Tonye Jack)
- Updated action.
 ([ab9881c](https://github.com/tj-actions/auto-doc/commit/ab9881cbcf7907e72ee174e9291ab61a1e01c9ea))  - (Tonye Jack)
- Update update-readme.yml ([b023913](https://github.com/tj-actions/auto-doc/commit/b023913a398d7086ada0cc2cd8a25d28d29dbb67))  - (Tonye Jack)
- Update test.yml ([e45ad90](https://github.com/tj-actions/auto-doc/commit/e45ad908f8d0727e3ecf165799113d647675715d))  - (Tonye Jack)
- Updated README.md
 ([34857a6](https://github.com/tj-actions/auto-doc/commit/34857a697dd24f205473a83f93f36da43b5e5802))  - (Tonye Jack)
- Updated action.
 ([08c9dc4](https://github.com/tj-actions/auto-doc/commit/08c9dc47c33ff5ca82c8849ac489d703775bc01d))  - (Tonye Jack)
- Update README.md ([d9330a7](https://github.com/tj-actions/auto-doc/commit/d9330a7be21b057147c2e86a91c6f78f3966dcc5))  - (Tonye Jack)
- Updated command
 ([02a025c](https://github.com/tj-actions/auto-doc/commit/02a025c5ab488e625eaed3b166f39f75883cf4ea))  - (Tonye Jack)
- Updated command
 ([8b19848](https://github.com/tj-actions/auto-doc/commit/8b19848c76b5a691183a03b12ada916703919c46))  - (Tonye Jack)
- Update root.go ([d10a482](https://github.com/tj-actions/auto-doc/commit/d10a4820a989d017c0d15f7bf32376f75ef60a81))  - (Tonye Jack)
- Update root.go ([b3d5291](https://github.com/tj-actions/auto-doc/commit/b3d5291aef2455e62d0e425a27e15cbea97271d9))  - (Tonye Jack)
- Update root.go ([ea4c804](https://github.com/tj-actions/auto-doc/commit/ea4c80472168df57288e53c66bedc8669e8a6a6d))  - (Tonye Jack)
- Update root.go ([2a2649f](https://github.com/tj-actions/auto-doc/commit/2a2649fa5f217bb5d79d957938cbf8a8c07662c1))  - (Tonye Jack)
- Update root.go ([e4b1f0b](https://github.com/tj-actions/auto-doc/commit/e4b1f0bd6d876a35f13fcfb2717ccf9e529d29c1))  - (Tonye Jack)
- Updated github action.
 ([4693765](https://github.com/tj-actions/auto-doc/commit/46937652c88e9de0e7b7bf4bbada9e9b54667cf9))  - (Tonye Jack)
- Update README.md ([c13bcfd](https://github.com/tj-actions/auto-doc/commit/c13bcfde6e3c98728728331dd1e7f06629a60884))  - (Tonye Jack)
- Update root.go ([e208d95](https://github.com/tj-actions/auto-doc/commit/e208d952183c161c577bb8d5a5bb464aabb84554))  - (Tonye Jack)
- Update README.md ([99f2bef](https://github.com/tj-actions/auto-doc/commit/99f2befdc0f91e21cc7b3158e85eb8875825b4a1))  - (Tonye Jack)
- Updated formatting
 ([0937f4c](https://github.com/tj-actions/auto-doc/commit/0937f4cc12b4345ab0f93cc80e42aa13e558afa2))  - (Tonye Jack)
- Updated README.md
 ([13b390d](https://github.com/tj-actions/auto-doc/commit/13b390d5103f585eaff2f421950cb9d0226cd6c5))  - (jackton1)
- Update README.md ([c25d7c9](https://github.com/tj-actions/auto-doc/commit/c25d7c9044f4fcf550d420aaa6b20157d6259f47))  - (Tonye Jack)
- Updated README.md
 ([38b157d](https://github.com/tj-actions/auto-doc/commit/38b157d12995afbe1f90ad43b373da1ef7fe7718))  - (jackton1)
- Updated action.
 ([682eb3e](https://github.com/tj-actions/auto-doc/commit/682eb3ec37a977bd0366d2e7355c986d155a2afc))  - (Tonye Jack)
- Update actions/checkout action to v2.4.0
 ([39006f1](https://github.com/tj-actions/auto-doc/commit/39006f1802595327359deb95cb9f43a46c4ba1f8))  - (Renovate Bot)
- Updated action.
 ([bf46941](https://github.com/tj-actions/auto-doc/commit/bf46941bcd8ae848f3d3f6aed5c8a84a2e465af6))  - (Tonye Jack)
- Updated action.
 ([3992307](https://github.com/tj-actions/auto-doc/commit/39923077e4ba2cd102c5887fead7fa0caeed4d6d))  - (Tonye Jack)
- Updated action.
 ([147e877](https://github.com/tj-actions/auto-doc/commit/147e877166d6b590c1a063c0158f64bb247c4f1e))  - (Tonye Jack)
- Updated action.
 ([c4e8aa6](https://github.com/tj-actions/auto-doc/commit/c4e8aa69470b939f6868a8776f179779971fda3b))  - (Tonye Jack)
- Updated action.
 ([0b42f1e](https://github.com/tj-actions/auto-doc/commit/0b42f1eec989e97f090818078f0b4c64561397e5))  - (Tonye Jack)
- Updated action.
 ([75770fc](https://github.com/tj-actions/auto-doc/commit/75770fc0ecc901e3da01fbb21741b4e908b65b53))  - (Tonye Jack)
- Updated action.
 ([ff8056a](https://github.com/tj-actions/auto-doc/commit/ff8056a9db772c6c294439bbb1af3b44a7cdcaf2))  - (Tonye Jack)
- Updated action.
 ([3a2099e](https://github.com/tj-actions/auto-doc/commit/3a2099eb754dffaa046a75d3d76da72466eb2d62))  - (Tonye Jack)
- Updated action.
 ([72e147a](https://github.com/tj-actions/auto-doc/commit/72e147a8211b0594dd4088441d644a4ad07f9860))  - (Tonye Jack)
- Updated action.
 ([131d59f](https://github.com/tj-actions/auto-doc/commit/131d59f091e6a1a46bba915d6f4763b4d6148c52))  - (Tonye Jack)
- Updated action.
 ([8481d30](https://github.com/tj-actions/auto-doc/commit/8481d30d76c5bdbb2009c29f022766a0ba24b860))  - (Tonye Jack)
- Updated action.
 ([d156799](https://github.com/tj-actions/auto-doc/commit/d1567996da9c1adf480af96b52992072ff2226fe))  - (Tonye Jack)
- Updated action.
 ([66dde34](https://github.com/tj-actions/auto-doc/commit/66dde345b0e8f92753be5e07dab320c74b2c42d5))  - (Tonye Jack)
- Updated action.
 ([7f9da53](https://github.com/tj-actions/auto-doc/commit/7f9da53cb15795dc7f7840f3cf482924a8135686))  - (Tonye Jack)
- Updated action.
 ([f796c81](https://github.com/tj-actions/auto-doc/commit/f796c81e413be4442c6cc2d185bc81378ecae0e1))  - (Tonye Jack)
- Updated action.
 ([d3f7b6a](https://github.com/tj-actions/auto-doc/commit/d3f7b6a614c67333f61d45f7985e0a08035a652a))  - (Tonye Jack)
- Updated action.
 ([7ac6a9b](https://github.com/tj-actions/auto-doc/commit/7ac6a9b9ba6eedabbf02799f79e4e17f5a775bee))  - (Tonye Jack)
- Updated action.
 ([02b91ba](https://github.com/tj-actions/auto-doc/commit/02b91baed8e1382316548105e374dcc5a1189c62))  - (Tonye Jack)
- Updated action.
 ([81a8851](https://github.com/tj-actions/auto-doc/commit/81a88519710776a388ca36e57087ffedb3ac9ff5))  - (Tonye Jack)
- Updated action.
 ([4c5fe27](https://github.com/tj-actions/auto-doc/commit/4c5fe275b0a09a6857a2ecb75008b9f9ac122e46))  - (Tonye Jack)
- Updated github action.
 ([f3f2402](https://github.com/tj-actions/auto-doc/commit/f3f24022e3e10af24c3a5b899a4dd45ae3c6d87e))  - (Tonye Jack)
- Update root.go ([f48018c](https://github.com/tj-actions/auto-doc/commit/f48018cc8350888cd8fecfda035dddb759d100ca))  - (Tonye Jack)
- Update root.go ([f634ea3](https://github.com/tj-actions/auto-doc/commit/f634ea39875bc1058dce42ce393094175fd45ffc))  - (Tonye Jack)
- Update root.go ([f29a77c](https://github.com/tj-actions/auto-doc/commit/f29a77cd4297643ab9ff64cf5acc8f12f52bfb67))  - (Tonye Jack)
- Update action.yml ([92df792](https://github.com/tj-actions/auto-doc/commit/92df79272b4f95c297b91cedb21d6c02d97f0b2f))  - (Tonye Jack)
- Update test.yml ([ae8582f](https://github.com/tj-actions/auto-doc/commit/ae8582fb48891abfa9b6f9f6d3f980a481ae09ae))  - (Tonye Jack)
- Update test.yml ([f2fc474](https://github.com/tj-actions/auto-doc/commit/f2fc47424c4a602e09c7b76acf2390799dd6fc19))  - (Tonye Jack)
- Updated the root command.
 ([d3ab089](https://github.com/tj-actions/auto-doc/commit/d3ab089b4942e14ee79724541c8f73299ab3df25))  - (Tonye Jack)
- Updated .github/ISSUE_TEMPLATE/feature_request.yaml ([66fc7fd](https://github.com/tj-actions/auto-doc/commit/66fc7fde1511c5301ae2c156ac586acfd581859b))  - (Tonye Jack)
- Updated .github/ISSUE_TEMPLATE/bug_report.yaml ([fa7082c](https://github.com/tj-actions/auto-doc/commit/fa7082cebffc8dec7eafaf08029953b47d804295))  - (Tonye Jack)
- Updated README.md
 ([04fe705](https://github.com/tj-actions/auto-doc/commit/04fe70561b033279a851b2b5c00016d79e541526))  - (jackton1)
- Updated README.md
 ([b270880](https://github.com/tj-actions/auto-doc/commit/b270880bb604100396d8c6ca6f29c5069c9060a5))  - (Tonye Jack)
- Update tj-actions/sync-release-version action to v9
 ([36ac4cc](https://github.com/tj-actions/auto-doc/commit/36ac4ccf42b3c64d147371ed71f8b33bbe07367e))  - (Renovate Bot)
- Updated renovate.json ([b98e8b6](https://github.com/tj-actions/auto-doc/commit/b98e8b649c657c22b62ed4dd9582fe42121fec18))  - (Tonye Jack)
- Updated renovate.json ([5c70d17](https://github.com/tj-actions/auto-doc/commit/5c70d1753dfd2dae5fd2b455fcfaa7d0e97e097b))  - (Tonye Jack)
- Update actions/checkout action to v2.3.5
 ([78f5a76](https://github.com/tj-actions/auto-doc/commit/78f5a76cd80c257fb6950fb166aae00879079a52))  - (Renovate Bot)
- Update tj-actions/verify-changed-files action to v8
 ([99e175f](https://github.com/tj-actions/auto-doc/commit/99e175f9ba5bdc8add9aac5dbc712275053bc546))  - (Renovate Bot)
- Updated the command.
 ([a114b4a](https://github.com/tj-actions/auto-doc/commit/a114b4a0faae43a9b2bec17bdaadb1d323f979a5))  - (Tonye Jack)
- Update main.go ([2f59ef7](https://github.com/tj-actions/auto-doc/commit/2f59ef77e231fba96eb465c6815b9b2a832a0551))  - (Tonye Jack)
- Update root.go ([680b013](https://github.com/tj-actions/auto-doc/commit/680b01370fe773af80fed457313c0d6e58b32144))  - (Tonye Jack)
- Updated to parse arguments.
 ([8156f11](https://github.com/tj-actions/auto-doc/commit/8156f11e217ddec120c33c2cba926ec9d75f4c0e))  - (Tonye Jack)
- Update entrypoint.sh ([6e09134](https://github.com/tj-actions/auto-doc/commit/6e09134646166e21553b53c8412ce1a6c10411a1))  - (Tonye Jack)
- Update Makefile ([0b67cb5](https://github.com/tj-actions/auto-doc/commit/0b67cb56740758896352c5c5d3230ed9669c64d7))  - (Tonye Jack)
- Update test.yml ([68d8236](https://github.com/tj-actions/auto-doc/commit/68d8236d82501d4e8779d4b4f4167413d80f037e))  - (Tonye Jack)
- Update root.go ([df636ac](https://github.com/tj-actions/auto-doc/commit/df636ac0bc286f51085cb49f82b8a075697b6267))  - (Tonye Jack)
- Updated README.md
 ([e32f743](https://github.com/tj-actions/auto-doc/commit/e32f7431d6baf60349b881a69e2a61f39674f097))  - (jackton1)
- Update README.md ([8db41ca](https://github.com/tj-actions/auto-doc/commit/8db41caeae9458b18a4f7d020dcd2b9aa1925ff3))  - (Tonye Jack)
- Update action.yml ([984cec8](https://github.com/tj-actions/auto-doc/commit/984cec897bb422938a9b4df00ecf605b7a6d08af))  - (Tonye Jack)
- Updated action.
 ([cac64ec](https://github.com/tj-actions/auto-doc/commit/cac64ec8f82e8c063057f85dcbac5cee28438b50))  - (Tonye Jack)
- Updated action.
 ([38d21a6](https://github.com/tj-actions/auto-doc/commit/38d21a6aaa52c932572c7121da752cafbf97b936))  - (Tonye Jack)
- Updated action.
 ([c7d3e13](https://github.com/tj-actions/auto-doc/commit/c7d3e130b8c420ec723ae2ce918568e2b36a3974))  - (Tonye Jack)
- Updated README.md
 ([89dc3e0](https://github.com/tj-actions/auto-doc/commit/89dc3e04e930e574ed6ef1ee4370b74588011fdd))  - (jackton1)
- Update Dockerfile ([c787dfa](https://github.com/tj-actions/auto-doc/commit/c787dfad9f59e971daea15fb9a9a9f68d4ed2d3f))  - (Tonye Jack)
- Updated action.
 ([c62571c](https://github.com/tj-actions/auto-doc/commit/c62571cf1eb55e74f37b59bf2e8d8899907b74d8))  - (Tonye Jack)
- Updated action.
 ([eacb078](https://github.com/tj-actions/auto-doc/commit/eacb07861dcfbd3553999a8826e76e6f105357c7))  - (Tonye Jack)
- Updated action.
 ([39c26a0](https://github.com/tj-actions/auto-doc/commit/39c26a0a6cb02922c597e4a683e243e10f887fb9))  - (Tonye Jack)
- Update golang Docker tag to v1.17
 ([8f41d8f](https://github.com/tj-actions/auto-doc/commit/8f41d8fb8c36ef6f5830f9cc1184915c77b03ead))  - (Renovate Bot)
- Updated action.
 ([9d01cb6](https://github.com/tj-actions/auto-doc/commit/9d01cb63cff27cde5cbc8d8b72b9d505923b5e7c))  - (Tonye Jack)
- Updated action
 ([ef57853](https://github.com/tj-actions/auto-doc/commit/ef57853fc06065e994970c06da4efe8d80692441))  - (Tonye Jack)
- Update README.md ([baaaeb3](https://github.com/tj-actions/auto-doc/commit/baaaeb3953df44d0d8873f1487cb58ebf9eea40d))  - (Tonye Jack)
- Update README.md ([3d8d610](https://github.com/tj-actions/auto-doc/commit/3d8d61058d0604af708869895642da8a25109c52))  - (Tonye Jack)
- Update README.md ([1c54f5f](https://github.com/tj-actions/auto-doc/commit/1c54f5fcc3929ed9e766bb02a95555919c3a54c5))  - (Tonye Jack)
- Update module github.com/spf13/viper to v1.9.0
 ([c83267e](https://github.com/tj-actions/auto-doc/commit/c83267ec7d562024e1d1114e96d4afda0885d6dc))  - (Renovate Bot)
- Update README.md ([4dcf34d](https://github.com/tj-actions/auto-doc/commit/4dcf34dad46398e8f4f4bf3c7e19c80ba0b9587f))  - (Tonye Jack)
- Update README.md ([5a05f7c](https://github.com/tj-actions/auto-doc/commit/5a05f7c60c35afe0a715e707d4fa107edc99c9f6))  - (Tonye Jack)
- Update README.md ([e8425e7](https://github.com/tj-actions/auto-doc/commit/e8425e7bbfeb37704a2ce15b006811de7c96112d))  - (Tonye Jack)
- Update pascalgn/automerge-action action to v0.14.3
 ([937ce3c](https://github.com/tj-actions/auto-doc/commit/937ce3cf4f19cb40a03179a304f01a5b070b34c9))  - (Renovate Bot)
- Update tj-actions/remark action to v1.7
 ([0fa6861](https://github.com/tj-actions/auto-doc/commit/0fa68616da3ef2719711b29d2ddbec72ddf1a4de))  - (Renovate Bot)
- Updated .github/ISSUE_TEMPLATE/bug_report.yaml ([497e9e0](https://github.com/tj-actions/auto-doc/commit/497e9e0745f7bd37bcc835e275f01b5d541f29fc))  - (Tonye Jack)
- Updated .github/workflows/codacy-analysis.yml ([7375ac0](https://github.com/tj-actions/auto-doc/commit/7375ac0f4ed011388004ab061924ba437799cad4))  - (Tonye Jack)
- Updated .github/ISSUE_TEMPLATE/bug_report.yaml ([2f18517](https://github.com/tj-actions/auto-doc/commit/2f18517a1866e664b498a9278f8aee167a612696))  - (Tonye Jack)
- Updated .github/ISSUE_TEMPLATE/bug_report.yaml ([bc01c4d](https://github.com/tj-actions/auto-doc/commit/bc01c4d133c786f74651e6ba2b7829ec41598ed3))  - (Tonye Jack)
- Update tj-actions/verify-changed-files action to v7
 ([1fe0ec4](https://github.com/tj-actions/auto-doc/commit/1fe0ec4f599b7b589b9a96f6547e465310fd81cf))  - (Renovate Bot)
- Updated .github/workflows/auto-merge.yml ([a76a2ed](https://github.com/tj-actions/auto-doc/commit/a76a2edf403c5b8dbf5f014382132689c6afa3bc))  - (Tonye Jack)
- Updated .github/workflows/auto-merge.yml ([d81c9a4](https://github.com/tj-actions/auto-doc/commit/d81c9a430c289653b2174393cb27f3d8242f5943))  - (Tonye Jack)
- Updated .github/workflows/auto-approve.yml ([3b2df2f](https://github.com/tj-actions/auto-doc/commit/3b2df2fd06ef88df906fb8c5cdb3c333462979bd))  - (Tonye Jack)
- Updated .github/workflows/auto-merge.yml ([e6d22a7](https://github.com/tj-actions/auto-doc/commit/e6d22a711494d544f7f3d5ef47dde40523314253))  - (Tonye Jack)
- Update pascalgn/automerge-action action to v0.14.2
 ([d10efff](https://github.com/tj-actions/auto-doc/commit/d10efff76f80990a962e38d01f1f6b6c7d7e6bde))  - (Renovate Bot)
- Updated renovate.json ([959d402](https://github.com/tj-actions/auto-doc/commit/959d4023f52378b31e08545a66a6780d87d7df9d))  - (Tonye Jack)
- Updated .github/workflows/auto-approve.yml ([4c135d8](https://github.com/tj-actions/auto-doc/commit/4c135d8d66b516854a5900a1ceb529f2abb4d3c0))  - (Tonye Jack)
- Updated .github/workflows/auto-merge.yml ([7ab1662](https://github.com/tj-actions/auto-doc/commit/7ab1662799ba5e2575c49d62fccebf67ca039b1e))  - (Tonye Jack)
- Updated .github/workflows/auto-merge.yml ([a5a9e1b](https://github.com/tj-actions/auto-doc/commit/a5a9e1b5957b75aff179f56090b829f9ee550bc3))  - (Tonye Jack)
- Updated .github/workflows/auto-merge.yml ([7166a3f](https://github.com/tj-actions/auto-doc/commit/7166a3fb6fdda69cd0fa3f6e543f20c081885e3f))  - (Tonye Jack)
- Updated .github/workflows/auto-merge.yml ([5d9d210](https://github.com/tj-actions/auto-doc/commit/5d9d210b18bda3db345ec4ffb4d8eb674d22a615))  - (Tonye Jack)
- Update tj-actions/github-changelog-generator action to v1.8
 ([2e1beae](https://github.com/tj-actions/auto-doc/commit/2e1beae1d2d7eba73511b18e0da08be49d66cd41))  - (Renovate Bot)
- Updated .github/workflows/auto-approve.yml ([91b2844](https://github.com/tj-actions/auto-doc/commit/91b28444b5011af9fe7bad2599e6a6764c9569e3))  - (Tonye Jack)
- Updated renovate.json ([f9b31ea](https://github.com/tj-actions/auto-doc/commit/f9b31ea07845334f6c9319ff2c76376659ae7b00))  - (Tonye Jack)
- Update tj-actions/github-changelog-generator action to v1.6
 ([8ecad96](https://github.com/tj-actions/auto-doc/commit/8ecad96124c7c30013c9f17f2ecf6c6f00b86915))  - (Renovate Bot)
- Updated .github/workflows/sync-release-version.yml ([f9bd08e](https://github.com/tj-actions/auto-doc/commit/f9bd08efdd357a24c636814c153c835402d3bd6a))  - (Tonye Jack)
- Update generator.py ([7225399](https://github.com/tj-actions/auto-doc/commit/72253994cc4166e3e9350cc42887b0ecff46dd53))  - (Tonye Jack)
- Updated .github/workflows/auto-approve.yml ([8a28885](https://github.com/tj-actions/auto-doc/commit/8a28885a3a73c450cad00cc3debb5cc21a9ac53f))  - (Tonye Jack)
- Updated .github/FUNDING.yml ([f764402](https://github.com/tj-actions/auto-doc/commit/f764402c2329fad6512b7aa79c18a6613bfc11c6))  - (Tonye Jack)

## <!-- 30 -->📝 Other

- PR [#25](https://github.com/tj-actions/auto-doc/pull/25): to v1 ([885fa34](https://github.com/tj-actions/auto-doc/commit/885fa3496cf97a861b5744b9cda41615a1bce365))  - (Tonye Jack)
- PR [#23](https://github.com/tj-actions/auto-doc/pull/23): README.md ([53dd3e5](https://github.com/tj-actions/auto-doc/commit/53dd3e5b1bb60fbab1d96fad72d9173e34202a89))  - (Tonye Jack)
- PR [#22](https://github.com/tj-actions/auto-doc/pull/22): README.md ([ebe4fa7](https://github.com/tj-actions/auto-doc/commit/ebe4fa72bc5d34f3f33ab17fd847a40fbfb027c0))  - (Tonye Jack)
- PR [#20](https://github.com/tj-actions/auto-doc/pull/20): unused code. ([0eb5e9b](https://github.com/tj-actions/auto-doc/commit/0eb5e9b1fb8ac0794304e9c5215bab9164a44076))  - (Tonye Jack)
- PR [#19](https://github.com/tj-actions/auto-doc/pull/19): README.md ([c78175a](https://github.com/tj-actions/auto-doc/commit/c78175aecbead238ea5e71b438097f64edd86955))  - (Tonye Jack)
- Merge branch 'main' of github.com:tj-actions/auto-doc
 ([aaa55ad](https://github.com/tj-actions/auto-doc/commit/aaa55ad78a6f581367c78d9d6be9b789bd8cb4f9))  - (Tonye Jack)
- PR [#15](https://github.com/tj-actions/auto-doc/pull/15): README.md ([51a30cc](https://github.com/tj-actions/auto-doc/commit/51a30ccde112b07049ddb859d09446dea9459f6c))  - (Tonye Jack)
- PR [#14](https://github.com/tj-actions/auto-doc/pull/14): README.md ([c39e101](https://github.com/tj-actions/auto-doc/commit/c39e101c32af6af2fece26baf9deb341533cec18))  - (Tonye Jack)
- PR [#13](https://github.com/tj-actions/auto-doc/pull/13): golang Docker tag to v1.17 ([afd8885](https://github.com/tj-actions/auto-doc/commit/afd88850edc93f8c8784325e080920225cd86278))  - (Tonye Jack)
- PR [#12](https://github.com/tj-actions/auto-doc/pull/12): module github.com/spf13/viper to v1.9.0 ([6dd7e8f](https://github.com/tj-actions/auto-doc/commit/6dd7e8f6e8b70592a7d401ad96018eeda9ce9513))  - (Tonye Jack)

## <!-- 9 -->⬆️ Upgrades

- Upgraded from v1 -> v1
 ([fbdeac3](https://github.com/tj-actions/auto-doc/commit/fbdeac3339d38ddee349e6d9e3bda320ab646ac6))  - (jackton1)

<!-- generated by git-cliff -->
