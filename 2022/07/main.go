package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	root := &Dir{name: "/", files: make(map[string]int), dirs: make(map[string]*Dir)}
	path := []*Dir{root}
	d := root
	for _, act := range input {
		if act.cmd == "ls" {
			for _, out := range act.output {
				if out[0] == "dir" {
					if _, ok := d.dirs[out[1]]; !ok {
						d.dirs[out[1]] = &Dir{name: out[1], dirs: make(map[string]*Dir), files: make(map[string]int)}
					}
				} else {
					d.files[out[1]], _ = strconv.Atoi(out[0])
				}
			}
		} else {
			cmd := act.cmd[3:]
			switch cmd {
			case "/":
				path = path[:1]
				d = path[0]
			case "..":
				path = path[:len(path)-1]
				d = path[len(path)-1]
			default:
				d = d.dirs[cmd]
				path = append(path, d)
			}
		}
	}

	calcSize(root)

	fmt.Println(findSize(root))

	free := 70000000 - root.size
	needed := 30000000 - free

	sizes := allSizes(root)
	sort.Ints(sizes)
	for _, s := range sizes {
		if s < needed {
			continue
		}
		fmt.Println(s)
		break
	}
}

func allSizes(d *Dir) []int {
	var sizes []int
	for _, sd := range d.dirs {
		sizes = append(sizes, allSizes(sd)...)
	}
	sizes = append(sizes, d.size)
	return sizes
}

func findSize(d *Dir) int {
	var sum int
	for _, sd := range d.dirs {
		sum += findSize(sd)
	}
	if d.size <= 100000 {
		sum += d.size
	}
	return sum
}

func calcSize(d *Dir) int {
	var sum int
	for _, s := range d.files {
		sum += s
	}
	for _, sd := range d.dirs {
		sum += calcSize(sd)
	}
	d.size = sum
	return sum
}

type Dir struct {
	name  string
	files map[string]int
	dirs  map[string]*Dir
	size  int
}

type File struct {
	name string
	size int
}

type Command struct {
	cmd    string
	output [][2]string
}

var input = []Command{
	{
		cmd: "cd /",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "drblq"},
			{"133789", "fjf"},
			{"dir", "jpfrhmw"},
			{"dir", "jqfwd"},
			{"dir", "ncgffsr"},
			{"12962", "ntnr.lrq"},
			{"dir", "qnbq"},
			{"dir", "rqdngnrq"},
			{"dir", "shcvnqq"},
			{"dir", "vsd"},
			{"dir", "vtzvf"},
		},
	},
	{
		cmd: "cd drblq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"133843", "bglzqdd"},
			{"dir", "brfnfhj"},
			{"268201", "fbqjmp.jzv"},
			{"80676", "shcvnqq"},
		},
	},
	{
		cmd: "cd brfnfhj",
	},
	{
		cmd: "ls", output: [][2]string{
			{"150447", "jlcg.dsg"},
			{"dir", "nhvgrzs"},
		},
	},
	{
		cmd: "cd nhvgrzs",
	},
	{
		cmd: "ls", output: [][2]string{
			{"282889", "jlcg.dsg"},
			{"19004", "ncgffsr.gwr"},
			{"dir", "vbzr"},
			{"6338", "vpsgdph.gbh"},
			{"dir", "wdcn"},
		},
	},
	{
		cmd: "cd vbzr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"225101", "fbqjmp"},
			{"243277", "vbzr"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd wdcn",
	},
	{
		cmd: "ls", output: [][2]string{
			{"154089", "dlmpbbf.psv"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd jpfrhmw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"87622", "cffdsj.jzf"},
			{"26165", "qnbq.sbm"},
			{"dir", "vbzr"},
		},
	},
	{
		cmd: "cd vbzr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "blhstw"},
			{"16919", "nttftcts"},
			{"dir", "rgdp"},
			{"116477", "shcvnqq"},
			{"242592", "tmjrnqbz.chq"},
			{"dir", "vbzr"},
			{"dir", "wmct"},
		},
	},
	{
		cmd: "cd blhstw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"98023", "jwdv.qct"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd rgdp",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "gcb"},
			{"141507", "shcvnqq"},
			{"dir", "ssvzm"},
		},
	},
	{
		cmd: "cd gcb",
	},
	{
		cmd: "ls", output: [][2]string{
			{"189016", "ncgffsr.rbq"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ssvzm",
	},
	{
		cmd: "ls", output: [][2]string{
			{"82667", "shcvnqq.zjq"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vbzr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"120202", "jlcg.dsg"},
			{"86205", "vbzr.jtr"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd wmct",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "fbsfcgph"},
			{"155709", "hpsftv"},
			{"13636", "lztgs"},
			{"273353", "ncgffsr.jsg"},
			{"dir", "pvwhpfp"},
		},
	},
	{
		cmd: "cd fbsfcgph",
	},
	{
		cmd: "ls", output: [][2]string{
			{"139944", "ncgffsr.gpf"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd pvwhpfp",
	},
	{
		cmd: "ls", output: [][2]string{
			{"111230", "bscrjpzh.glp"},
			{"dir", "dgjsddgq"},
			{"37234", "lwd"},
			{"107139", "lztgs"},
			{"258111", "mgtwwvwz"},
			{"117638", "qpdvnfb.gnf"},
			{"dir", "szrplcdw"},
			{"dir", "vzsl"},
			{"dir", "wsmf"},
		},
	},
	{
		cmd: "cd dgjsddgq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "qnbq"},
		},
	},
	{
		cmd: "cd qnbq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"199119", "jlcg.dsg"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd szrplcdw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"122236", "qclr.cpf"},
			{"269638", "qnbq"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vzsl",
	},
	{
		cmd: "ls", output: [][2]string{
			{"233006", "twpz.tdm"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd wsmf",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "wcnptvtz"},
		},
	},
	{
		cmd: "cd wcnptvtz",
	},
	{
		cmd: "ls", output: [][2]string{
			{"183952", "shcvnqq.lwt"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd jqfwd",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "hqb"},
			{"285121", "jqffsjbs.jrm"},
			{"dir", "nhpqpdn"},
			{"dir", "qnbq"},
			{"dir", "qtrv"},
			{"dir", "wspztvjr"},
		},
	},
	{
		cmd: "cd hqb",
	},
	{
		cmd: "ls", output: [][2]string{
			{"253786", "jwdv.qct"},
			{"dir", "vbzr"},
		},
	},
	{
		cmd: "cd vbzr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"153", "gbh"},
			{"dir", "gqpqqrgl"},
			{"dir", "jzncgd"},
			{"36914", "nvdnsnls.mpd"},
		},
	},
	{
		cmd: "cd gqpqqrgl",
	},
	{
		cmd: "ls", output: [][2]string{
			{"206691", "dmdgcwm.bgh"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd jzncgd",
	},
	{
		cmd: "ls", output: [][2]string{
			{"122640", "vrgmf.tnp"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd nhpqpdn",
	},
	{
		cmd: "ls", output: [][2]string{
			{"86329", "ntnr.lrq"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd qnbq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"76269", "fbqjmp.lbd"},
			{"118968", "fbqjmp.msg"},
			{"190416", "gfwhsb.dpc"},
			{"dir", "lhgjrmj"},
			{"dir", "pbv"},
			{"173541", "pfl"},
			{"141842", "srrmt.ssj"},
		},
	},
	{
		cmd: "cd lhgjrmj",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "ghccnw"},
			{"180420", "ldzcj.rwz"},
			{"149356", "lztgs"},
			{"61792", "ncgffsr"},
			{"dir", "spmbcjhc"},
		},
	},
	{
		cmd: "cd ghccnw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"253233", "lztgs"},
			{"56439", "ntnr.lrq"},
			{"19225", "ntrmjf.gdb"},
			{"31628", "pdhhzjhm.lbd"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd spmbcjhc",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "shcvnqq"},
		},
	},
	{
		cmd: "cd shcvnqq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"122334", "drjbh"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd pbv",
	},
	{
		cmd: "ls", output: [][2]string{
			{"69436", "cctsjqh.wqr"},
			{"285573", "ljtqddz"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd qtrv",
	},
	{
		cmd: "ls", output: [][2]string{
			{"234568", "dmwqfbwd"},
			{"dir", "pwwsrjc"},
			{"245046", "qmcr"},
			{"159151", "qtvdjncm.rdb"},
			{"dir", "swhzds"},
			{"178915", "vbzr.vgn"},
			{"dir", "vcgv"},
		},
	},
	{
		cmd: "cd pwwsrjc",
	},
	{
		cmd: "ls", output: [][2]string{
			{"173975", "bgdj.jnw"},
			{"202714", "jwdv.qct"},
			{"270702", "wggrgcvw.rtp"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd swhzds",
	},
	{
		cmd: "ls", output: [][2]string{
			{"114686", "jwdv.qct"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vcgv",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "fbqjmp"},
			{"dir", "qlsgtfhf"},
			{"dir", "vbzr"},
		},
	},
	{
		cmd: "cd fbqjmp",
	},
	{
		cmd: "ls", output: [][2]string{
			{"73065", "fbqjmp.jfb"},
			{"dir", "shcvnqq"},
		},
	},
	{
		cmd: "cd shcvnqq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"231428", "shcvnqq"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd qlsgtfhf",
	},
	{
		cmd: "ls", output: [][2]string{
			{"75227", "ntnr.lrq"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vbzr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"128050", "ncgffsr.gsj"},
			{"187649", "vbzr"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd wspztvjr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "pntrhtwh"},
			{"dir", "qnbq"},
			{"dir", "zfdzvv"},
		},
	},
	{
		cmd: "cd pntrhtwh",
	},
	{
		cmd: "ls", output: [][2]string{
			{"237258", "cffhtr"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd qnbq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "qnbq"},
		},
	},
	{
		cmd: "cd qnbq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "ccwmftsj"},
		},
	},
	{
		cmd: "cd ccwmftsj",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "mfc"},
			{"dir", "shcvnqq"},
			{"12262", "smpjmn"},
		},
	},
	{
		cmd: "cd mfc",
	},
	{
		cmd: "ls", output: [][2]string{
			{"198047", "fbqjmp.cgh"},
			{"dir", "gghsht"},
			{"205411", "wlclz"},
		},
	},
	{
		cmd: "cd gghsht",
	},
	{
		cmd: "ls", output: [][2]string{
			{"31767", "vbzr.lmb"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd shcvnqq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "lgrghwf"},
		},
	},
	{
		cmd: "cd lgrghwf",
	},
	{
		cmd: "ls", output: [][2]string{
			{"114786", "shcvnqq.vrz"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd zfdzvv",
	},
	{
		cmd: "ls", output: [][2]string{
			{"54298", "sjp"},
			{"60303", "tcmhrll.htm"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ncgffsr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "fqqsqmpr"},
			{"dir", "gfznw"},
			{"dir", "ncdft"},
			{"dir", "pwmppt"},
			{"dir", "shcvnqq"},
			{"196969", "vbzr"},
			{"214841", "vzgvr"},
		},
	},
	{
		cmd: "cd fqqsqmpr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "mcdjcntr"},
		},
	},
	{
		cmd: "cd mcdjcntr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"281856", "ncgffsr.lbm"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd gfznw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"255657", "fzrctbsj.lgf"},
			{"dir", "ltfsndpd"},
			{"175434", "qnbq"},
			{"31794", "qnbq.zhd"},
			{"13366", "shcvnqq.wld"},
			{"dir", "vcspqgn"},
			{"235199", "wmnjjd.bnh"},
			{"dir", "wqpnp"},
		},
	},
	{
		cmd: "cd ltfsndpd",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "ncgffsr"},
			{"dir", "zpzvdhb"},
		},
	},
	{
		cmd: "cd ncgffsr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"9898", "jjbsnj.gcg"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd zpzvdhb",
	},
	{
		cmd: "ls", output: [][2]string{
			{"106139", "lnp"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vcspqgn",
	},
	{
		cmd: "ls", output: [][2]string{
			{"25386", "dgsmmqj"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd wqpnp",
	},
	{
		cmd: "ls", output: [][2]string{
			{"65905", "wjtbfvjp.fmd"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ncdft",
	},
	{
		cmd: "ls", output: [][2]string{
			{"34616", "bzlpmsqc"},
			{"59863", "jlcg.dsg"},
			{"64629", "zpzjcl.fmp"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd pwmppt",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "dwnqgrzm"},
			{"80901", "vbzr.vsg"},
			{"89557", "vbzr.zlz"},
		},
	},
	{
		cmd: "cd dwnqgrzm",
	},
	{
		cmd: "ls", output: [][2]string{
			{"184770", "jwdv.qct"},
			{"dir", "vbzr"},
		},
	},
	{
		cmd: "cd vbzr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"210329", "jlcg.dsg"},
			{"62272", "jwdv.qct"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd shcvnqq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"128433", "gbh"},
			{"30208", "hjbw"},
			{"200071", "jlcg.dsg"},
			{"dir", "sgcz"},
			{"25045", "tbhlwfqg.hts"},
		},
	},
	{
		cmd: "cd sgcz",
	},
	{
		cmd: "ls", output: [][2]string{
			{"193481", "gbh"},
			{"96461", "jwdv.qct"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd qnbq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"236171", "shcvnqq"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd rqdngnrq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "cprnb"},
			{"280135", "hshsfqwm"},
			{"dir", "hwhm"},
			{"245626", "qnbq"},
			{"145502", "qspgdz"},
			{"114231", "rctg.tgt"},
			{"dir", "zgn"},
		},
	},
	{
		cmd: "cd cprnb",
	},
	{
		cmd: "ls", output: [][2]string{
			{"115025", "twwgmmp.wbb"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd hwhm",
	},
	{
		cmd: "ls", output: [][2]string{
			{"229849", "cvm"},
			{"190622", "jwdv.qct"},
			{"dir", "mscztz"},
			{"dir", "ncgffsr"},
		},
	},
	{
		cmd: "cd mscztz",
	},
	{
		cmd: "ls", output: [][2]string{
			{"59743", "bzgpzn.bds"},
			{"75184", "pbdgv"},
			{"181089", "shcvnqq.dhq"},
			{"dir", "zqgtr"},
		},
	},
	{
		cmd: "cd zqgtr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"189142", "ffnznfs.nct"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ncgffsr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "dphrnjl"},
			{"dir", "zzfztql"},
		},
	},
	{
		cmd: "cd dphrnjl",
	},
	{
		cmd: "ls", output: [][2]string{
			{"117317", "vbzr"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd zzfztql",
	},
	{
		cmd: "ls", output: [][2]string{
			{"51096", "lztgs"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd zgn",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "bpbzwgz"},
			{"dir", "gqnw"},
			{"75631", "ljptj"},
			{"283351", "ljzhsw.rbs"},
			{"131158", "lztgs"},
			{"dir", "ncgffsr"},
			{"3136", "nnpl.swf"},
			{"dir", "shcvnqq"},
			{"dir", "vbzr"},
		},
	},
	{
		cmd: "cd bpbzwgz",
	},
	{
		cmd: "ls", output: [][2]string{
			{"29659", "jlcg.dsg"},
			{"15547", "shcvnqq"},
			{"117389", "zprhsdfv"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd gqnw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"117091", "brqwhst.jgb"},
			{"88406", "nzjmbrrm.hmh"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ncgffsr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"195821", "gbh"},
			{"dir", "lbzgc"},
			{"226692", "llqqr.spq"},
			{"247989", "lztgs"},
			{"231909", "vnctc"},
			{"157973", "wqnggh"},
		},
	},
	{
		cmd: "cd lbzgc",
	},
	{
		cmd: "ls", output: [][2]string{
			{"251414", "ffmsbscc.dqg"},
			{"46840", "lztgs"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd shcvnqq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "dvvmhzcq"},
			{"dir", "ncgffsr"},
			{"dir", "sqzzllv"},
		},
	},
	{
		cmd: "cd dvvmhzcq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "qnbq"},
			{"70226", "qvvm.rpp"},
			{"dir", "shcvnqq"},
		},
	},
	{
		cmd: "cd qnbq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"103994", "bfcjrmvr.ltq"},
			{"dir", "fbqjmp"},
			{"dir", "fcs"},
			{"177152", "gjghvvw.bzg"},
			{"dir", "lbfjqh"},
			{"78412", "ntnr.lrq"},
			{"dir", "sgjtm"},
			{"286995", "shcvnqq"},
			{"51750", "wmq.vjj"},
		},
	},
	{
		cmd: "cd fbqjmp",
	},
	{
		cmd: "ls", output: [][2]string{
			{"267212", "qhhb.zvg"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd fcs",
	},
	{
		cmd: "ls", output: [][2]string{
			{"272051", "znhsswwh.mjj"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd lbfjqh",
	},
	{
		cmd: "ls", output: [][2]string{
			{"261487", "jlcg.dsg"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd sgjtm",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "dnznpj"},
			{"dir", "jzsntnbs"},
			{"dir", "nqgcbd"},
			{"dir", "vdg"},
		},
	},
	{
		cmd: "cd dnznpj",
	},
	{
		cmd: "ls", output: [][2]string{
			{"173938", "hrp.cjq"},
			{"180485", "qnbq.thj"},
			{"215400", "ztvt.wnt"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd jzsntnbs",
	},
	{
		cmd: "ls", output: [][2]string{
			{"67448", "gpvgh.psg"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd nqgcbd",
	},
	{
		cmd: "ls", output: [][2]string{
			{"196250", "fbqjmp.qcv"},
			{"198482", "jlcg.dsg"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vdg",
	},
	{
		cmd: "ls", output: [][2]string{
			{"257343", "jwdv.qct"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd shcvnqq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"156769", "fbqjmp.hdb"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ncgffsr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"205473", "fbqjmp"},
			{"113067", "gsvznzz.qtv"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd sqzzllv",
	},
	{
		cmd: "ls", output: [][2]string{
			{"146018", "ddvjgswr.gsq"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vbzr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "vbzr"},
		},
	},
	{
		cmd: "cd vbzr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"266721", "mhlfqpbs.pwr"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd shcvnqq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "slnvdd"},
		},
	},
	{
		cmd: "cd slnvdd",
	},
	{
		cmd: "ls", output: [][2]string{
			{"90875", "pzqv.gnv"},
			{"207484", "rbrj.vcr"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vsd",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "dfb"},
			{"dir", "fqqnsph"},
			{"dir", "gbwdhjr"},
			{"18837", "jwdv.qct"},
			{"dir", "ncgffsr"},
			{"dir", "qnbq"},
			{"dir", "rjzjrbvs"},
		},
	},
	{
		cmd: "cd dfb",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "bpst"},
			{"66174", "jwdv.qct"},
			{"dir", "lcwhfzjw"},
		},
	},
	{
		cmd: "cd bpst",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "nqftnn"},
			{"dir", "pcvgnvnp"},
		},
	},
	{
		cmd: "cd nqftnn",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "bbrsg"},
			{"dir", "gjfc"},
			{"dir", "hfql"},
			{"dir", "shcvnqq"},
			{"139226", "shcvnqq.sbd"},
			{"dir", "ssnjqbg"},
		},
	},
	{
		cmd: "cd bbrsg",
	},
	{
		cmd: "ls", output: [][2]string{
			{"73382", "vjcf"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd gjfc",
	},
	{
		cmd: "ls", output: [][2]string{
			{"164310", "gbh"},
			{"126316", "mmqnrc"},
			{"133899", "ntnr.lrq"},
			{"102615", "rgfhrt"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd hfql",
	},
	{
		cmd: "ls", output: [][2]string{
			{"14685", "jwdv.qct"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd shcvnqq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"119597", "lztgs"},
			{"34165", "shcvnqq.zcg"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ssnjqbg",
	},
	{
		cmd: "ls", output: [][2]string{
			{"77678", "gqdfbqj.tmj"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd pcvgnvnp",
	},
	{
		cmd: "ls", output: [][2]string{
			{"21250", "lhq"},
			{"266619", "qps.crp"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd lcwhfzjw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "bhdnnbvm"},
			{"dir", "fdnsvfh"},
			{"12002", "jlcg.dsg"},
			{"dir", "lfdbzfl"},
			{"46488", "ncgffsr"},
			{"233704", "nthcv.pnc"},
			{"204660", "ntnr.lrq"},
			{"172482", "shcvnqq"},
			{"dir", "tlw"},
		},
	},
	{
		cmd: "cd bhdnnbvm",
	},
	{
		cmd: "ls", output: [][2]string{
			{"37204", "fwrdjw.zvv"},
			{"3248", "ntnr.lrq"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd fdnsvfh",
	},
	{
		cmd: "ls", output: [][2]string{
			{"20765", "jlfgnwb.szl"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd lfdbzfl",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "fspntmld"},
			{"183925", "jlcg.dsg"},
		},
	},
	{
		cmd: "cd fspntmld",
	},
	{
		cmd: "ls", output: [][2]string{
			{"251568", "lztgs"},
			{"146785", "ncgffsr.mmj"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd tlw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "qqn"},
		},
	},
	{
		cmd: "cd qqn",
	},
	{
		cmd: "ls", output: [][2]string{
			{"39232", "lprqfwf"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd fqqnsph",
	},
	{
		cmd: "ls", output: [][2]string{
			{"132318", "lztgs"},
			{"103863", "ntnr.lrq"},
			{"18793", "tngbs"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd gbwdhjr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"253798", "jwdv.qct"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ncgffsr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"110767", "blctz.tqz"},
			{"dir", "csfssn"},
			{"dir", "dbbfz"},
			{"dir", "hjgm"},
			{"dir", "hwd"},
			{"249139", "rgcz.gnz"},
			{"dir", "wgw"},
		},
	},
	{
		cmd: "cd csfssn",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "dlcw"},
			{"dir", "jqspd"},
			{"119066", "mlwlc.mql"},
			{"dir", "ncgffsr"},
			{"203475", "nwnbsc"},
			{"143071", "qnbq"},
			{"116623", "qvw.gjz"},
			{"83637", "whm.cdg"},
		},
	},
	{
		cmd: "cd dlcw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"232066", "gqllsd.qpl"},
			{"1046", "mfsh"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd jqspd",
	},
	{
		cmd: "ls", output: [][2]string{
			{"251070", "mthmm.bmh"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ncgffsr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"83639", "ntnr.lrq"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd dbbfz",
	},
	{
		cmd: "ls", output: [][2]string{
			{"112576", "jgqf.qmj"},
			{"148549", "jlcg.dsg"},
			{"144811", "jwdv.qct"},
			{"23726", "ntnr.lrq"},
			{"123802", "pgdjchrf.vnm"},
			{"dir", "vzfbzbcp"},
		},
	},
	{
		cmd: "cd vzfbzbcp",
	},
	{
		cmd: "ls", output: [][2]string{
			{"39375", "fbqq"},
			{"31914", "jwdv.qct"},
			{"165999", "lztgs"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd hjgm",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "ljqjtdmf"},
			{"100534", "mdw"},
			{"219057", "qnbq"},
			{"97164", "rzjwmvdw.vlv"},
			{"dir", "shcvnqq"},
			{"83034", "vbzr"},
		},
	},
	{
		cmd: "cd ljqjtdmf",
	},
	{
		cmd: "ls", output: [][2]string{
			{"23716", "dmslzv.qns"},
			{"159519", "gbh"},
			{"dir", "hlvbmpg"},
			{"dir", "nlqqshp"},
			{"247315", "vqt"},
			{"dir", "wlsjnthg"},
		},
	},
	{
		cmd: "cd hlvbmpg",
	},
	{
		cmd: "ls", output: [][2]string{
			{"54421", "jlcg.dsg"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd nlqqshp",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "rvzprwhp"},
		},
	},
	{
		cmd: "cd rvzprwhp",
	},
	{
		cmd: "ls", output: [][2]string{
			{"35024", "lztgs"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd wlsjnthg",
	},
	{
		cmd: "ls", output: [][2]string{
			{"29178", "gnrlgb.bgh"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd shcvnqq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"150311", "nvrd"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd hwd",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "jzqtmm"},
		},
	},
	{
		cmd: "cd jzqtmm",
	},
	{
		cmd: "ls", output: [][2]string{
			{"103547", "jtvdt.jtn"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd wgw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "mmhlt"},
		},
	},
	{
		cmd: "cd mmhlt",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "cmwjh"},
		},
	},
	{
		cmd: "cd cmwjh",
	},
	{
		cmd: "ls", output: [][2]string{
			{"243844", "qnbq.shn"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd qnbq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "dhfng"},
			{"dir", "fbqjmp"},
			{"16855", "rgrszmrh.lbl"},
			{"dir", "rqjs"},
			{"dir", "shcvnqq"},
			{"38322", "vhvrmq"},
		},
	},
	{
		cmd: "cd dhfng",
	},
	{
		cmd: "ls", output: [][2]string{
			{"132537", "gwngz.hpt"},
			{"dir", "lbccc"},
			{"182221", "ntnr.lrq"},
		},
	},
	{
		cmd: "cd lbccc",
	},
	{
		cmd: "ls", output: [][2]string{
			{"282448", "fbqjmp.njj"},
			{"267049", "gbh"},
			{"dir", "jtj"},
			{"dir", "ntnn"},
			{"dir", "vbfgmmvw"},
			{"128500", "vbzr"},
		},
	},
	{
		cmd: "cd jtj",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "hvmlh"},
		},
	},
	{
		cmd: "cd hvmlh",
	},
	{
		cmd: "ls", output: [][2]string{
			{"131886", "dmww.sqc"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ntnn",
	},
	{
		cmd: "ls", output: [][2]string{
			{"109064", "lgh.bbf"},
			{"dir", "wfgdd"},
			{"53862", "wflv.ngc"},
		},
	},
	{
		cmd: "cd wfgdd",
	},
	{
		cmd: "ls", output: [][2]string{
			{"58756", "gbh"},
			{"dir", "lgzlndn"},
			{"dir", "qnbq"},
		},
	},
	{
		cmd: "cd lgzlndn",
	},
	{
		cmd: "ls", output: [][2]string{
			{"190415", "dwsqvczd"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd qnbq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"240922", "znjhmhp.ngt"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vbfgmmvw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"271827", "vbzr.dfl"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd fbqjmp",
	},
	{
		cmd: "ls", output: [][2]string{
			{"144993", "gvpnf"},
			{"150786", "jwdv.qct"},
			{"49025", "pdcwwtt.grs"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd rqjs",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "bwnzs"},
			{"119390", "jlcg.dsg"},
			{"172042", "vjzg"},
		},
	},
	{
		cmd: "cd bwnzs",
	},
	{
		cmd: "ls", output: [][2]string{
			{"108537", "hzzgm.zrn"},
			{"38699", "qgfqbfr"},
			{"dir", "vhvcfhvr"},
		},
	},
	{
		cmd: "cd vhvcfhvr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"2783", "jwdv.qct"},
			{"209933", "mgj.nvj"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd shcvnqq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"257312", "fbqjmp"},
			{"193792", "msdqtrpn.grn"},
			{"98165", "rgm"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd rjzjrbvs",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "ftrlfg"},
			{"dir", "mtrnl"},
			{"dir", "rdpbbd"},
			{"dir", "shcvnqq"},
			{"dir", "vztnr"},
		},
	},
	{
		cmd: "cd ftrlfg",
	},
	{
		cmd: "ls", output: [][2]string{
			{"196590", "cjjvwjb"},
			{"dir", "ffsvh"},
			{"70123", "ldnbc"},
			{"dir", "lwnfc"},
			{"106499", "lztgs"},
			{"dir", "ncgffsr"},
			{"dir", "tfdctq"},
			{"dir", "vgthdbf"},
			{"80852", "zndjt.wtl"},
		},
	},
	{
		cmd: "cd ffsvh",
	},
	{
		cmd: "ls", output: [][2]string{
			{"20370", "dvdftpvb.qcj"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd lwnfc",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "fgmd"},
			{"dir", "gmdjt"},
			{"274331", "hmgjmq.vbz"},
			{"9726", "qjfdqbf.dfj"},
			{"dir", "ssnncn"},
		},
	},
	{
		cmd: "cd fgmd",
	},
	{
		cmd: "ls", output: [][2]string{
			{"280608", "jwdv.qct"},
			{"201912", "rqtbw.shd"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd gmdjt",
	},
	{
		cmd: "ls", output: [][2]string{
			{"202107", "jwdv.qct"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ssnncn",
	},
	{
		cmd: "ls", output: [][2]string{
			{"140697", "jwdv.qct"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ncgffsr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"227389", "fpdfqp.fzl"},
			{"164141", "hzhrrvpm.hlf"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd tfdctq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "cttmzlw"},
			{"dir", "ntvtm"},
			{"257094", "qnbq.zjm"},
			{"284928", "shcvnqq"},
		},
	},
	{
		cmd: "cd cttmzlw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"142651", "rptschdv.mgv"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ntvtm",
	},
	{
		cmd: "ls", output: [][2]string{
			{"176269", "dhpj"},
			{"88278", "gbh"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vgthdbf",
	},
	{
		cmd: "ls", output: [][2]string{
			{"130998", "ncgffsr.mnf"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd mtrnl",
	},
	{
		cmd: "ls", output: [][2]string{
			{"86144", "djwnvdj"},
			{"122600", "gsdpwh.cmb"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd rdpbbd",
	},
	{
		cmd: "ls", output: [][2]string{
			{"177384", "gbh"},
			{"dir", "gstfdm"},
			{"dir", "qnbq"},
			{"dir", "qtj"},
			{"260302", "vbzr.dhq"},
		},
	},
	{
		cmd: "cd gstfdm",
	},
	{
		cmd: "ls", output: [][2]string{
			{"23734", "mnwzrm.hzr"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd qnbq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"51705", "gmt"},
			{"205537", "ntnr.lrq"},
			{"94469", "vbzr.bvj"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd qtj",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "tls"},
			{"dir", "zvpcfhg"},
		},
	},
	{
		cmd: "cd tls",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "chvgwnt"},
			{"dir", "jvgnmfjw"},
		},
	},
	{
		cmd: "cd chvgwnt",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "rbw"},
			{"dir", "srhj"},
		},
	},
	{
		cmd: "cd rbw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"174372", "btjd.bvv"},
			{"272995", "cnqqh.dfc"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd srhj",
	},
	{
		cmd: "ls", output: [][2]string{
			{"134054", "qwzpr"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd jvgnmfjw",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "hdcwbwgm"},
			{"236775", "sdc"},
		},
	},
	{
		cmd: "cd hdcwbwgm",
	},
	{
		cmd: "ls", output: [][2]string{
			{"113707", "ntnr.lrq"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd zvpcfhg",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "lsq"},
		},
	},
	{
		cmd: "cd lsq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"220331", "jlcg.dsg"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd shcvnqq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "cmwrqgfq"},
			{"258731", "fbqjmp.fvn"},
			{"277895", "gbh"},
			{"64973", "jlcg.dsg"},
			{"77978", "jwdv.qct"},
			{"dir", "lttjrdn"},
			{"dir", "sqgnhc"},
		},
	},
	{
		cmd: "cd cmwrqgfq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"81199", "gbh"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd lttjrdn",
	},
	{
		cmd: "ls", output: [][2]string{
			{"23355", "gbh"},
			{"148263", "hcgfqdw"},
			{"57338", "hjwr"},
			{"166510", "jbvnmcj"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd sqgnhc",
	},
	{
		cmd: "ls", output: [][2]string{
			{"dir", "glswqrdp"},
			{"dir", "qnbq"},
		},
	},
	{
		cmd: "cd glswqrdp",
	},
	{
		cmd: "ls", output: [][2]string{
			{"225761", "ncgffsr.vct"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd qnbq",
	},
	{
		cmd: "ls", output: [][2]string{
			{"62861", "pdqz.wzs"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vztnr",
	},
	{
		cmd: "ls", output: [][2]string{
			{"189943", "wvtlfsp"},
		},
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd ..",
	},
	{
		cmd: "cd vtzvf",
	},
	{
		cmd: "ls", output: [][2]string{
			{"43248", "jwdv.qct"},
		},
	},
}
