// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "660217392cbd84eaa5c92d6371da04a9"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"28b3af7629381a338d41c4d605c1659b": "1f8b08000000000000ff5256d44fcaccd32fcee0e24acf2c5148492d4e2eca4c4a05040000ffff601451af17000000",
		"331fe9d3ed3849a14ce85fecfd59508e": "1f8b08000000000000ff6c534d6fe33614bceb57ccd281373e48b45d2058a830b03db5a7b628722b8a054d51146b8a8fe0875223c97f2f2431ae17e8c5d2d3cc1b0ee73d6f900613612204a218bd5588791cc94192eb8d466facaa36884a6148c9c796736dd290cf8da4917762321d4d460e7c6ddbac6d398864c855930ad1906b71a8843522aad8e2956993ea82b016ec6c1c2f651307f65e6d4039f99c3a135a309a5408a6eb94fbd699c02af58f926d059c451ce62730286ba95eeaf5bd894355019acec6a10ed9ad3c2fe425b4b8f3afe99cfb5e58e20bc6a7e3fa725cf84bff77fcf16a929cf8027c9df6cdbef95298839097efa881b457de2bbe625fa743b36f8eb32d7f4d03b91f50cb3bfb2dde960230a3a79010afb17cf0c1b8f418afb12929ed3e983d9e43566d293f983d5b14d1071acb619f20829eb39f5544d0d33bdb5525c39b8f0de228ac459f9d9ca78744882a650f81a46657225cf1f36fbffff4fc0b248d5e2473b60a9a3a92882a4c2ac0b8683a858ee445850fd9ff5f1e4d5638cd3571136356911f9fbe1c1b3ccf0b29aca597b8881aa767b93c2a97d6ad5a557b0a10d08491ba6c558472e26c55071fe86f25535326d391fc2fdcdbddbe2dc0e30eafb7f4fec427d43dd8c3a37fe9765c533352c7f017b65b283910980a81e6055a10384ae829bb8ee1ededa682e2e7f4f0285e2ef8fceb1fa7d3e175990c1e8eef9f4bf70edbed5dcf9a174276a8eb30a256606bd2279e46cf3531d41e4ffba77d3bffa09eb0fa6c0bce6390fc613d1b6bb4eda1391c9b27f0f91f56460db65c651e47cbb92529ec40312da2dc5ff44d62bb2d83ad67ee6921b03bc7d94595e6bcd6247f2cd07b79ae5fabeadf000000ffffc740440b61040000",
		"a24a41abb2f5d6e7d1ad3c1f697a1e63": "1f8b08000000000000ff24ca410ac2301085e17d4ef1c8013c805b11ba11179e20a54333904e42f26a859abb4bebf6ff3f466dd086491bab8e2bc39804940fdda68cd8775c6ed92846f4eedcab549d41594a0a549b0fe0a3a4943dbe584b917aba2950aec7b4bc9d61c8cb3f88bde187e7e3eed1fb2f0000ffffb84780aa7f000000",
		"e75c7238ac5f4068c84d9a6037c023c6": "1f8b08000000000000ff648fd14ac3401045dff3159745418bc9561095059ffa0fbea7c92459d866642754649c7f97342b14fa78ef9c3970556bdc0d89dbe5c067840f34073e536e47c22f2efdeb0b6ab34a2b0070d24d746a3f294be4d9053c3f6d7d6a8f945c80ebcabb2b871389ac31c0a95ed9cdee0ba2ea771022605a962f09decb1429f5d244f61838a3e3c459b0f3669bf4526cca1a71c0c348572bdef7cdfe1166c71cc7691933d1ac0a4a42b7ec5b617f2825fefec7cc32f56b98fb75bcabacfa0b0000ffff2c4a278229010000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("Summon Box", "../assets")
		b.SetResolver("bin/version.sh", packr.Pointer{ForwardBox: gk, ForwardPath: "28b3af7629381a338d41c4d605c1659b"})
		b.SetResolver("shields.io/summon.json", packr.Pointer{ForwardBox: gk, ForwardPath: "e75c7238ac5f4068c84d9a6037c023c6"})
		b.SetResolver("summon.config.yaml", packr.Pointer{ForwardBox: gk, ForwardPath: "331fe9d3ed3849a14ce85fecfd59508e"})
		b.SetResolver("text.txt", packr.Pointer{ForwardBox: gk, ForwardPath: "a24a41abb2f5d6e7d1ad3c1f697a1e63"})
	}()

	return nil
}()
