# Tree
The assignment is described at https://www.coursera.org/learn/golang-webservices-1/programming/kJ610/utilita-tree.
My examples on 01.11.2021:
### go run main.go . -f
```
gorkostya@DESKTOP-V4AA9I4:/mnt/d/golangWork/tree/hw1_tree$ go run main.go . -f
├─── .git
│       ├─── COMMIT_EDITMSG (31b)
│       ├─── HEAD (21b)
│       ├─── ORIG_HEAD (41b)
│       ├─── branches
│       ├─── config (273b)
│       ├─── description (73b)
│       ├─── hooks
│       │       ├─── applypatch-msg.sample (478b)
│       │       ├─── commit-msg.sample (896b)
│       │       ├─── fsmonitor-watchman.sample (3079b)
│       │       ├─── post-update.sample (189b)
│       │       ├─── pre-applypatch.sample (424b)
│       │       ├─── pre-commit.sample (1638b)
│       │       ├─── pre-merge-commit.sample (416b)
│       │       ├─── pre-push.sample (1348b)
│       │       ├─── pre-rebase.sample (4898b)
│       │       ├─── pre-receive.sample (544b)
│       │       ├─── prepare-commit-msg.sample (1492b)
│       │       └─── update.sample (3610b)
│       ├─── index (2550b)
│       ├─── info
│       │       └─── exclude (240b)
│       ├─── logs
│       │       ├─── HEAD (1219b)
│       │       └─── refs
│       │               ├─── heads
│       │               │       └─── main (869b)
│       │               └─── remotes
│       │                       └─── origin
│       │                               └─── main (616b)
│       ├─── objects
│       │       ├─── 06
│       │       │       └─── fe2f6dfe4ad7f9789f2878a7d205a08596abd4 (182b)
│       │       ├─── 07
│       │       │       └─── c702eba1feaab09da31ca89f53bae5a86c9a72 (87b)
│       │       ├─── 0d
│       │       │       └─── 94cdd5608f8e34a587e6e4d02a680a98640c89 (54b)
│       │       ├─── 14
│       │       │       └─── e512ae6479a2d4dd6aad64aac7951e849e24ec (68192b)
│       │       ├─── 19
│       │       │       └─── f6bfdb70494a163f0f1a63d2c712bd11d8b311 (163b)
│       │       ├─── 20
│       │       │       └─── f9cc96cecf23f4115c2cc6e4af2ab0e153b3e9 (562b)
│       │       ├─── 37
│       │       │       └─── 4db5f4b91b3506cf7cd6775b19dbf3e59d7633 (44b)
│       │       ├─── 38
│       │       │       └─── ca078ed2d3e8d239e07ee69e343035582ddd3a (2120b)
│       │       ├─── 39
│       │       │       └─── d42da89a7ecd175d94459f975fd5faeaa6745b (22b)
│       │       ├─── 3d
│       │       │       └─── 1c2eb7a4770a1833b1eeaf24c6f5ce6ca2b114 (87b)
│       │       ├─── 48
│       │       │       └─── 7f6baf9ea4338ef65055ed11acdd93cc2f17a5 (175b)
│       │       ├─── 57
│       │       │       └─── 966093020a09b1e97a208ffa60adf73c18bcbf (86b)
│       │       ├─── 5e
│       │       │       └─── 38e6599b84da8c1b555cec1a2bdf757c940eaf (55b)
│       │       ├─── 78
│       │       │       └─── bbe91873ba06fa6b3547b7566c182ea1485f74 (243b)
│       │       ├─── 7d
│       │       │       └─── e1a5770d6f78958ce590a2b34261eca33b018f (1231b)
│       │       ├─── 9a
│       │       │       └─── 1910cbfb9e237d0889ac0ed19f6edaefccd856 (1139b)
│       │       ├─── 9e
│       │       │       └─── 0931e6dccccb4566b6837d16c39e3d063ef775 (37b)
│       │       ├─── a9
│       │       │       └─── 5181f9ecc040ea845c7278bbd1a4322863e6df (243b)
│       │       ├─── af
│       │       │       └─── 6e59c52add68ebedff65233755262988ef3a0b (53b)
│       │       ├─── be
│       │       │       ├─── 76ea7d2a0d1cc8c1b8e2a96a403bdd5dfebada (180b)
│       │       │       └─── f9f3af00f03148aab6219d886e32ae0a1fb86d (121b)
│       │       ├─── c8
│       │       │       └─── aa3c66190cd3d16289a2c53e405476b7c65a7d (149b)
│       │       ├─── ce
│       │       │       └─── 5621a062f79ce5378e12bc79421238d911a5ed (52b)
│       │       ├─── d5
│       │       │       └─── 914101d777fb02de0a41c0e345035388f11c7b (26b)
│       │       ├─── da
│       │       │       └─── 80889387a1b1e42e394f774748781bbeab1bc7 (55b)
│       │       ├─── e6
│       │       │       └─── 9de29bb2d1d6434b8b29ae775ad8c2e48c5391 (15b)
│       │       ├─── e7
│       │       │       └─── fd5face9a1aefda26995546cd484a74be1cf15 (61b)
│       │       ├─── f0
│       │       │       └─── 471b247bfa23e70c1beffa367bdfd11f3b75e3 (243b)
│       │       ├─── f3
│       │       │       ├─── 8fa366c1aa01d34490ca02465664cfb44bea98 (137b)
│       │       │       └─── 9a1f3e1cb0874a94417bd3556a62bbbc2615b1 (35b)
│       │       ├─── f9
│       │       │       └─── f4de17b7313dca0021e4f8c9abef7ee84bdfa8 (1153b)
│       │       ├─── info
│       │       └─── pack
│       └─── refs
│               ├─── heads
│               │       └─── main (41b)
│               ├─── remotes
│               │       └─── origin
│               │               └─── main (41b)
│               └─── tags
├─── README.md (7b)
├─── dockerfile (75b)
├─── go.mod (21b)
├─── hw1.md (4621b)
├─── main.go (2921b)
├─── main_test.go (1866b)
└─── testdata
        ├─── project
        │       ├─── file.txt (19b)
        │       └─── gopher.png (70372b)
        ├─── static
        │       ├─── a_lorem
        │       │       ├─── dolor.txt (empty)
        │       │       ├─── gopher.png (70372b)
        │       │       └─── ipsum
        │       │               └─── gopher.png (70372b)
        │       ├─── css
        │       │       └─── body.css (28b)
        │       ├─── empty.txt (empty)
        │       ├─── html
        │       │       └─── index.html (57b)
        │       ├─── js
        │       │       └─── site.js (10b)
        │       └─── z_lorem
        │               ├─── dolor.txt (empty)
        │               ├─── gopher.png (70372b)
        │               └─── ipsum
        │                       └─── gopher.png (70372b)
        ├─── zline
        │       ├─── empty.txt (empty)
        │       └─── lorem
        │               ├─── dolor.txt (empty)
        │               ├─── gopher.png (70372b)
        │               └─── ipsum
        │                       └─── gopher.png (70372b)
        └─── zzfile.txt (empty)
```
### go run main.go .
```
├─── .git
│       ├─── branches
│       ├─── hooks
│       ├─── info
│       ├─── logs
│       │       └─── refs
│       │               ├─── heads
│       │               └─── remotes
│       │                       └─── origin
│       ├─── objects
│       │       ├─── 06
│       │       ├─── 07
│       │       ├─── 0d
│       │       ├─── 14
│       │       ├─── 19
│       │       ├─── 20
│       │       ├─── 37
│       │       ├─── 38
│       │       ├─── 39
│       │       ├─── 3d
│       │       ├─── 48
│       │       ├─── 57
│       │       ├─── 5e
│       │       ├─── 78
│       │       ├─── 7d
│       │       ├─── 9a
│       │       ├─── 9e
│       │       ├─── a9
│       │       ├─── af
│       │       ├─── be
│       │       ├─── c8
│       │       ├─── ce
│       │       ├─── d5
│       │       ├─── da
│       │       ├─── e6
│       │       ├─── e7
│       │       ├─── f0
│       │       ├─── f3
│       │       ├─── f9
│       │       ├─── info
│       │       └─── pack
│       └─── refs
│               ├─── heads
│               ├─── remotes
│               │       └─── origin
│               └─── tags
└─── testdata
        ├─── project
        ├─── static
        │       ├─── a_lorem
        │       │       └─── ipsum
        │       ├─── css
        │       ├─── html
        │       ├─── js
        │       └─── z_lorem
        │               └─── ipsum
        └─── zline
                └─── lorem
                        └─── ipsum
```
