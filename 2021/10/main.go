package main

import (
	"fmt"
	"sort"
)

func main() {
	var sum int
	var scores []int
nextLine:
	for _, line := range input {
		var stack []rune
		for _, r := range line {
			switch r {
			case '(', '[', '{', '<':
				stack = append(stack, r)
			case ')', ']', '}', '>':
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				switch r {
				case ')':
					if pop != '(' {
						sum += 3
						continue nextLine
					}
				case ']':
					if pop != '[' {
						sum += 57
						continue nextLine
					}
				case '}':
					if pop != '{' {
						sum += 1197
						continue nextLine
					}
				case '>':
					if pop != '<' {
						sum += 25137
						continue nextLine
					}
				}
			}
		}
		if len(stack) > 0 { // incomplete
			var score int
			for i := len(stack) - 1; i >= 0; i-- {
				r := stack[i]
				score *= 5
				switch r {
				case '(':
					score += 1
				case '[':
					score += 2
				case '{':
					score += 3
				case '<':
					score += 4
				}
			}
			scores = append(scores, score)
		}
	}
	fmt.Println(sum)
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}

var input = []string{
	"(<{([{([[[<{(<(){}>({}()))}[(<[]<>>{<>})[{{}()}<{}[]>]]><[[<{}{}>([]{})]<<<>{}>[<>{}]}]{{([]{})({}())}<",
	"<[<<{<[(([[[[<[][]>][(<>()){<>()}]]][<{<<>{}>(()<>)}><([<>{}]{()()})([()()]([]{}))>]]){{<{[[<><>]({}{})]",
	"<(({(<{<({[<[[[]()][<>[]]]{{{}()}({}[])}>]{[{<<>[]>[()[]]}][[[{}[]](<>{})]{[()()]((){})}]}}(",
	"{{[{[[[{{{{{[([]())(<>())](<[]<>>{[]{}})}([[(){}]<<>()>]<<()()><{}{}>>)}}((<[{{}()}<()>]<({}{",
	"({([({([[({([{[]<>}]<<<>{}>[<>[]]>)})({{{[()[]]<()>}(([]<>)<[]()>>}({[<>{}]}([{}<>]([]())))})]][{<",
	"<(((<<<{([[[<{()<>}{{}()}>[({}())([][])]]{{(()())(<>{})}(<{}><[]{}>)}]<[{([]())<{}>}<[[][]](<>)>]<{",
	"(([{[{([(<{{[<{}><<><>>]<<{}<>>>}(({<>{}}(<>[])){{<>[]}})}>)({(<<([]<>){[][]}>([{}<>]({}()))>)<(([<>{}]<{}[]>",
	"<((<[[<({[<({{<>[]}}{<<>{}><{}[]>}){([(){}])<(()[])<<>[]>>}>(<((<>[])[<>[]]){{<>{}}{{}()}}>)]}[<",
	"<[{<[<(<[[[{[<[][]>][[<>()}({}())]}]<<((<>[]))[((){}){[]()}]><((()<>){{}{}})[(()[])<<><>>]>>]]{<[[<<<",
	"[{(<(({{{(<(({{}{}}))>{<[(<>[])<(){}>]<<()<>>{[][]}>>((([]()){[]})({<>[]}<()[]>))})({<<({}())>[(()())",
	"<{[<<[<[([{{<<(){}><()[]>>([{}<>][{}<>])}}[[[[(){}]({}())][(()[])<[]()>]]]]({([<<>>[<>()]]([[]<>]<{}{}>))((<{",
	"({{{<(([{(([{{()()}{<>{}}}[<<>>[()()]]]))<[<[[{}{}]{[]()}]<([]<>){[]()}>}{{{[]}[<>()]}}]>}]",
	"[(([({[<([{<<({}{})<{}[]>>><<{()()}([])>{{[]<>}[()()]}>}{{<([]()){{}{}}>[{<>()}{{}()}]}[({<>()}{(){}})<[{}(",
	"{([{({[<<<<<[{{}[]}](<{}<>>)>>({(<{}()>[[][]])})>><<[[({<><>}({}[]})]([{{}()}[[]()]][({}{})({}{})])",
	"<{([{[<{<<{<<{<>[]}<()<>>>[<<>()>(()[])]>{<({}())([]())>}}<({{{}{}}([]{})}{{{}<>}({}<>)})[{<{}()>(()[",
	"[<{{(<{<<{<(<{[]<>}{(){}}>{{()<>}[[][]]}){[{<>[]]<{}<>>]<({}{})>}>}{<(<{[]{}}(<>())>)[{{[][]}[()]}<<()><",
	"[[[<[[[[<{(((<[][]>(()<>)))<<<<>{}>{<><>}}{({}())(<>{})}>)[({[<><>][{}()]}[([]{})])({(<><>)[(){}]}(<{",
	"({[<[<<([<{<{(<>{}){[]>}<{{}[]}(<><>)>>}><[(<({}())><<<>()>[()[]]>)([{<>{}}({}{})])]({{{<>{}}{{}<>}}[[<",
	"[(<[{<[[(<[<{[()()][{}[]]}{[(){}][<><>]}>]{({<[][]>[()<>]}[([][])<{}>])}><[(<([][])<<>()>>[[<>()](<><>)]",
	"[({(<{[<{[(<<[()<>]([]{})>[<[][]>({}{})]>(<{[]<>}<<>[]>>((()[]){{}[]}))){[<({}{})<()<>>><<()[]>([])>]<<<",
	"<(<((({{<[{<[<()[]><()<>>]{<<>{}><()[]>]>{<{()[]}{<>()}>({()<>}[<>{}])}}<<[[()()]<[][]>]<<()<>><{}{}>>>",
	"<<{{[{([{(({[<()<>>[{}()]]{<()()><[][]>}}(((()<>]{{}{}})([<>[]]{<>[]})))[<[<{}()>({}())]<(()<>)[[]()]>",
	"<<[[(<[<[[<{(<<>{}>{[]<>})}<<[[]<>](<>())><{{}<>}[()()]>>>((({{}()}<()()>))[[(<>{})<<>{}>]{(()){{}<>}}]",
	"[{<{([<<{{[[[<[]{}>{(){}}][{{}()}{[][]}>][({{}()}[{}{}]){({}<>)}]]{{[({}())([])][<{}><{}()>",
	"[{((({([{((<<[[]()]<{}>><<[][]><[]{}>>>[[[<>()]{[]{}}][{()<>}[(){}]]]))<<{[<[]{}>({}{})]}>(<{[[]",
	"{[[[{[{{((<[(<<>[]><{}()>)[{<>()}[()<>]]]{({{}{}}({}{})){[()()](())>}>)({(<<[]()><[]>>)}{[[({}[]){()<",
	"<<<((<<({({<<[[]][[][])>>[({<><>}[<>()])<[{}{}]([][])>]})([{(<[][]>[{}<>])<[<><>]>}<<([]{})(()<>)>[[{",
	"[({[([{[{<(({([][])<{}[]>}({(){}}[()()]))<([{}[])(<>[]))[({}()){[]<>}]>){<{[()[]]<{}{}>}>{<[()()]<[]",
	"{<<[<<[((([{[{[]()}<<>>][<()[]>]}[[{{}()}{[][]}]<({}{})(<>{})>]]){({[{()<>}<{}<>}]{(()<>)}}<([[]][[](",
	"<<({{[{<<{<[{[()()][[][]]}]>(([{[]}(()())]<{()[]}<<>[]>>)(((()[])[()()])))}{((<{{}()}{{}{}}>{[[]",
	"<({[{[(([{([[<{}<>>]]<[[()<>](<>())]>){[[[(){}]]{<{}<>>{()<>}}]{<[[]{}]<{}{}>>{[()<>]}}>}]{(([{([]{}",
	"[<{{{<<[{(({[({}()){{}[]}](({}{}){[]{}})}(<{{}[]>>[<()<>>{{}<>}]))[[<<[][]>><(()())[<>()]>][({()[]}(()",
	"(<{{{{[({[{{<[{}<>]((){})><[[]{}](<><>)>}(({[][])<{}[]>)[<()()>[<>[]]])}]})]<{{[<{(([]{})<()()>){",
	"(<<<<[{((<{({({}[])[[]()]}[<{}{}>([]<>)])(({[]()})({(){}}(<><>)))}[[[<[]{}><()()>]<{{}()}{{}",
	"(([([[<{{{([<(<><>){(){}}>[[[]{}]]]({[[]{}]<()()>}<<<><>>>))}}}]<[[<({[{<><>}{<>()}]{([]<>){()<>}}",
	"(<(<{<(<(<(({{<>}<[]()>}([[]{}])){<(<>())[{}()]>([()[]](()))})[[[{()()}({}{})]<<{}>([]{})>]([<()[]><{}<>>",
	"[(<[[{{((<[<<<[]>[{}{}]}<[[]<>]<[]{}>>>{{[[]()][()<>]}}]>{({([{}()]<()()>)[<()[]><(){}>]}[",
	"{[<{<<<<<[{(<{[][]}({}())>(((){}){()()}))}]>>{(<({<(<>{}){[]{}}>{((){})<<>()>}}<<[<><>][<>()]>[<[][]><<><>>",
	"<{[({<[{{{<{<(()[])><{[]()}<(){}>>}{[([][])[()[]]]}>}({{({{}<>}<[]()>)}}<([{{}<>}]<({}<>)<()<>>>)<<([](",
	"<[<{(<({[<[({<()<>>{(){}}}<{{}[]}(()<>)>)]{<(<<><>>(<>))<{{}<>}({}<>)>>{{<()[]>{{}{}}}{<[]<>>}}}>]}{<[",
	"{{[<({{{{[{({[{}[]]{{}[]}})(({[][]}<<>()>)<[<>{}]{[]<>}>)}{{<[<>{}][()()]>([(){}](<>{}))}}]",
	"([{<[(((({[({<()<>><<>{}>})]{<[([])[<>()]]{(<>)[[]{}]}>[<(<>)({}[])>[({}){<><>}]]}}{[[<(<>)(()[])>",
	"{(<{<{<({[({{[{}[]]}})[({(<><>){[][]}})<([()[]](<>{}))<[{}()]((){})>]]]{{{<(()<>){()<>}>}[(({}())[",
	"<(<[<{{{[[<{({()()}([]{})){[<>]<(){}>}}>[[[{[]<>}{<>[]}][(<>{}){<><>}]](({(){}}<[]<>>))]][<({{{}(",
	"<{((([<{<[[(({()<>}(()<>)))[<{[][]}[<><>]>]]]><<(({([][])[{}{}]}[<[]{}>[<><>]])([{(){}}[()()]][[[]{}]({}[])])",
	"{{<[[{<<[<<{{<<>{}>([]())}[[<>[]](()[])]}[<({}{})[<><>]>[[{}]{[]}>]>>[<([([])(()[])])><{(<(){}>{()[]})",
	"[<[[<<(<{<({{<<><>>([]<>)}}{<[{}{}]{[][]}>[{[]{}}<<><>>]})({{<[]<>>(()[])}(({}[])[()<>])})",
	"(([<{[{[[<{(({<>[]}{()[]}))[{[()[]]{<><>}}(({}{})[{}[]])]}[{(<{}<>>(<><>))<(()<>)([][])>}{<{[]()}<(){}>",
	"(<<{{<{(<{<<[{{}{}}{{}()>]{{<>[]}<<>[]>}>>}>)<[({[{([]<>)(())}[({}{}){(){}}]]{[[(){}]]<([])([][",
	"({{(<[<[<<<{[{<>[]}]<{{}<>}>}>{<<[[]()]((){})><<[]()>[(){}]>>{{([]()){[]()}}(<()()>)}}>({[([<><>][<>()]){(()(",
	"<[{(<(<{([[[<([]{})<<><>>>[[()]([]<>)]]]{[<[<>{}]([]())}([<>()][{}<>])]<<([]{}){{}<>}><{{}()}{<>()}>>}])<",
	"(({{({[[[<({[({})<<>[]>]{<{}()>{()[]}}}){[[(<>{})([]{})]{[{}()](()[])}]([<()()>]{((){})<{}()>})}><[[",
	"[[[(({<{[([(({()<>}<[][]>))][[(([]())[[][]]){<[]>}]([[[][]]<<>[]>]([{}]{()()}))])]{{[({[<>[]][[]]}{(<",
	"([([{<[(({[{{<(){}>([]<>)}(<[]<>><(){}>)}]({([{}{}][()<>])}{(([]<>){[]<>}){[<>[]]<{}<>>}})}))<((<",
	"{<(<[{<<<(({<{()<>}<[]<>>>}{<[<>{}>{{}<>}>({(){}}{()})}){[{(())(()())}({()[]}(()<>))]<[{(){}}({}<>)]<[<>()]>",
	"{{[<[{([<(<<<<[]{}>[[][]]>{[()[]]<<>[]>}>{(({}())([][]))({[]()}<()<>>)}>[(({[]<>}(()[]))[<{}<>>{()[]}])",
	"<{<{{{{<[{{(<{{}{}}{()()}>[([]{}){[]()}])<{[[]<>]<<>()>}[<{}[]>{[]()}]>}<((<()()>[()[]]))>}][{[{<<{}",
	"({(([({[[(<<([<>()]{{}<>})(<()()>[[]()])>(<[()[]]{[]()}>{[()()]{()()}})>){[{[[()<>](()<>>]}([[",
	"({{<{[[{{<{<{([]){[]<>}}{{[]{}}<[]<>>}>}(<{{{}[]}<{}>>>)>}(([[[([]())(<>)]{[[]{}][[]]}]<<({}()",
	"[{([({([<<<<{{()[]}<{}{}>}[[[]()]{<>{}}]>[([<>{}]<[][]>){{[]<>}({}{})}]>{[[<[][]>([]())]{[<><>](<>[",
	"([({{<<({(([[<<>()><{}<>>]<(()[])({}())>]){((<[]()>[(){}])(<<><>>{{}[]})){[[[]()][[]<>]]<<[]()>[{",
	"[{(<<({[<[<<{[{}<>}[{}()]}{[[]<>]{()()}}>(<[()[]][()()]>)>(<{[()<>](()<>)}<<<>>({}{})>>([{",
	"(<[[{(({[{{({[()<>]}(({}[])[<><>]))([[[]<>](<>[])]{<(){}>[()<>]})}(((<<><>>({}()))({<>}((){}))",
	"(<{{<<<[({({[{{}<>}[<><>]][(()<>)<<>()>]}>})(<(<((()[])({}))<<<>>([][])>>)(({<()()>[()<>]})(({{}{}}({}()))(",
	"({<{{[<<{([((((){})<[]()>)(<<>[]>({}())))<{({}[])([])}{<{}[]>{()()}}>]<<{[()[]]{[]<>}}>>]}{<{([<[]<>><(){}>",
	"<[<[<<({{<[[{{<><>}<{}<>>}<[[][]>({}<>)>]]>(<<{(<>{}){()[]}}[<()<>>[[]]]>>[{<<<><>>((){})>[<()()>[{}<>]]}<([",
	"{<<(<[([<[[{<{[]<>}{{}[]}>{[<>()]{{}{}}}}<{<()<>>[()[]]}{(<>{})({}())}>]]><[([<<()<>>{()[]}>(<<>>",
	"{<({<[{{({(((({}[])(()<>)){(()[])})<[<(){}>(()<>)][<{}<>><[][]>]>){<{([]()][()<>]}<(<>())<()",
	"{{(([<[({<[(((<>()){()[]}){(<>{}){[]<>}}]<{{{}()}<{}{}>}[[{}()]{{}<>}]>][<{[[][]]{[]<>}}<<(){}>",
	"{[{(<{{([[{<{{<>}[<>{}]}{<[][]>{[]()}}>(((()[])({}()))<<()()>{[][]}>)}[<(((){}){{}[]})[({}[])]",
	"{[<([[({{({{[<{}[]>{[]()}])}<[[[[]{}]][<()[]>({}[])]]<<[[]<>]<<>{}>><{<>{}}(<>{})>>>)}}){{<([",
	"[([<(([[[{{<{<[]{}>{[][]}}[<[]()>(()[])]>{[{<>{}><()[]>][[[]]]}}<[[((){})[[]<>]][(()[])([][])]]>}(<{[[[]()",
	"[([{[({[{{{{<{()[]}>{<<><>>}}[[(<><>)([]())][([][]){<>[]}]]}}}]})]{{{[<{{{<{{}()}(<>{})>}{<({",
	"[<[{<[[[[<<<[<()[]>[()()]]<[{}()](<>)>><([<>{}]({}[]))>><<[<{}{}>[<>()]]((<>{}){[]()})>[([",
	"[<{{[{([[<<[<(<>{})({}<>)>]{<[[]{}]>{(()<>)[<>[]]}}>{[((()[])(()[])){({}{}}[[]{}]}]{([{}[]]",
	"{<<(<{<<([(<{<{}>[()()]}{{{}{}}[{}()]}>[{{[]()}{{}<>}>{{()()}([]{})}]){(<[()<>]<[]{}>>[{(){}}([]<>)])}]",
	"[<[<([<[{{{(<({}<>)[{}()]>[[<>()]])(({()[]}){({}){[][]}})}}}(<([<{{}[]}[{}{}]>{(()())}])[((({}[])[",
	"{(<(([(([[{<{<()[]>(()[])}(<<><>>(<><>))><[[[]{}][()()]]{([]{})({})}>}]({<[{()()}(<>()))<[[]](()())>>}<({{{}{",
	"{(({{<{([<{<<<(){}><{}()>>{[<><>]<{}()>}><[<{}{}>]<{<>[]}<(){}>>>}{{({[]<>}{()[]})}}><{(<{{}()}{{}()}>({(){",
	"{<<{(([<(({{{((){})<{}()>}}<(<()<>>([]<>)){{<><>}{()()}}>}<{{(())[<>[]]}}<<<(){}}<[][]>>[<{}()>[<><>]]>>))>",
	"<[{{{<{{[[[[({()<>}([]()))<(())([]())>]{{{<>()}{(){}}}{{{}[]}}}]][({{<<>()>}[([][])<()>)}((([",
	"[(({[<<[[({({{(){}}<[]()>})})][<[[([{}{}][[]()])]([(<>{})<{}()>][{{}{}}([]())])][[{[()()]<<>{}>}<<<>",
	"([(<<[(<[<<{[((){})(<>[]}]({<>}[<>[]])}(({{}[]}((){})))>{({{[]<>}[[]<>]}<{{}}<{}>>){(<[]()",
	"[[{{<[<{<{{{{([])[[]()]}((<>)(()[]))}(<<{}{}>([]{})>)}[{(<()<>>([]{}))<(<>())[()()]>}{<{[]()}[",
	"<<(<<[(<(([{({()<>}({}<>))[{()[]}[[][]]]}]((<{<>{}}[<>[]]><(()<>){()<>}>)[(<<>>{()<>})<([]()}[[]()]",
	"[<<<{[<<<[(<<{{}()>{[][]}>[[(){}](<>[])]>{[[[]<>]]<({}){()()}>})][{({<{}[]>}){[{[]()}[{}[]]]{[[][",
	"[{([<<[{((<<((<>[])<()<>>)>>)(<<(<()()>)(({})<<>>)>>))}]>(<{{{[((<()()>[{}{}])[<[][]>([]()}])[<[{}[]]{{}{}}>",
	"[{{{[(<([([{{{<>{}}}{[{}[]]<[]()>}}{{[(){}][[]{}]}((()<>)(()()))}]<<{{()<>}<()()>}>>){[[{(<>",
	"<<[[(<([({([{<{}[]>(<><>)}])}{(({[{}]({}{})}[[<>{}]])[{[()[]]{{}()]}{[{}[]][[]<>]}])<<(<[]()><(",
	"(({[<[[[[{<{<<()()>[<>[]]>{(()<>){()}}}{<(()())([]<>)>{[()<>]<<>{}>}}>}](({([[{}<>]<(){}>])<([[",
	"(<<[{[[<{<[(({<><>}[<><>])({(){}}(<>[])))<[([]<>){{}{}}]>>>}[<{({[<>()](()())}{([]())[<>[]]})[[(<>[])",
	"<[<(({([(((<[<{}[]>{<>{}}]{{<><>}<[]<>>}><([[]<>]{[]{}})[<{}>{()}]>)){[[{([]<>)}{<{}()>{<>{}}}]{[{[]<",
	"{{((<<{((([(<[()[]]([][])>)[{{(){}}[{}()]}<([]{}){[]{}}>]]{[[({}())[()]]][{[[]{}]<[][]>}{<{}[]>[<",
	"([(([{{[({<[[<(){}>[[]<>]]]<[{()<>}<(){}>]>>[{((<>[]))}]}(<[[{()()}{{}<>}][(()[]>{{}{}}]][<({}[])>[",
}