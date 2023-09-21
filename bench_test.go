package main

import (
	"strings"
	"testing"
)

const (
	alphabet = "Подстрока может встречаться несколько раз, например, так: подстрока, подстрока, подстрока, подстрока, подстрока, подстрока, подстрока, подстрока, подстрока, подстрока, подстрока, подстрока"
	randStr  = "pEgaRDZUpxUjYQaHJdWdmgPGGUSQDtQEURVWeWDBPkAZotGudKKQYnRvlVzcsmGswtSnPLOTtCrdVxNrFamJTOMIIPvSHExMFtxEVfwpZwBiUyNTpWBfvuAaGmgULhfwCgszyNdCDHeDYORwhNUGMYhcYBZJRvlxhAFJoeOXGMpHKtoupLtYTQJmpkqApfDRkKFaTaBMKEXGMNoSZpFQxHFhJpprclNmjAILmKXpUkXnfsxsTLzuLoUGHZaxZatcapwmJyYtKsaLenPUMVikTUfYLwrHIyEzSyYkyvXuDfSVOEYjvwupdqbHBOhfLUwgMHHzfaDbuqAMDzKGNcDglrBlnHTkiwNITjxnBkGqLmRuXCawCdcamrLgnKdpUrFADRmJovUldAKManhcZfwmyjKUnezOdOLwWAUbQjsBcnMydXWPMIFSWmGQbNkDrEMMdtwsbfLBJIHFsdDxXtxtlNUHRKzafYVdPpmFczllWwkaGJbxsWBTDhpOliTrINSxxTYznfiVgBssAlSKMWksTZMEhjCcpXtyImVRjrNnGPvwVuKKWZYPSknGCrJooQXdIDeoalYnThwIBrxfMjiisGeLBSggzHWHFHEUarmfmaaxwyKhpNkTnoPRHkgBdNjEKgAQQsEbiiHIsPGytVSsGZlPvAfctMFiZHDnvDcZaKpLmvyIHFFYuelecbaTbbqAewEdtcZJHxlQEpjRIcRDwLQZTCqdIAdHSCzcpiTQUERzbiPbyVgVrgoCqlzvkTGxoHudMjttOqbnFmggpHaKyiBRJKKWCQhQnrLIoxvpWGWmFcVvKRgmVNjtJJOelhmrDGFwVZYEUthorlgxXZgBnQoxLhuwqEXKZKDVzfNndIDJcSiSLpkFvpbTUHSJEVkgrsVwYBRsHcMfpnwrYOZyLKNAbLWNvfFLvuINlgmyvDzBZozduQzyEHYZbhCzRQidcaXnUecGGaNXcExiwTgCkofQUlySyUwAGCztDiD"
	article  = "Go реализует семантику «копирования при присваивании», то есть присваивание приводит к созданию копии значения исходной переменной и размещения этой копии в другой переменной, после чего значения переменных являются различными и при изменении одного из них другое не меняется. Однако это верно только для встроенных скалярных типов, структур и массивов с заданной длиной (то есть для типов, значения которых размещаются в стеке). Массивы с неопределённой длиной и отображения размещаются в куче, переменные этих типов фактически содержат ссылки на объекты, при их присваивании копируется только ссылка, но не сам объект. Иногда это может привести к неожиданным эффектам. "
	req      = "AABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACD"
	chinese  = "女孩子们和雪姑娘手挽手跑去，可是一到树林里，一看见浆果，就把什么都忘得一干二净，大家你往东，我往西，只顾采浆果和在树林里“啊呜！啊呜！”地相互召唤。女孩子们采到不少浆果，可是在树林里把雪姑娘给丢了	雪姑娘叫唤着女友们——没有人答应。可怜的雪姑娘哭开了，她寻找回去的路，可怎么也找不到路了。她爬到一棵树上，高声喊着“啊呜！啊呜！”一只熊走过来，把干树枝踩得劈啪响，把灌木丛压得直往下弯。熊说：“美丽的姑娘，什么事儿？什么事儿？”“啊呜！啊呜！我是雪姑娘，我是用春雪滚成的，春天的太阳给我涂上了胭脂。我的女友们求老公公和老婆婆放我出来，他们同意了；女友们把我带到树林里来，可是她们丢下我不管了！”“下来吧！”熊说。“我送你回家去！”“熊呀，我可不干，”雪姑娘回答道。“我不跟你去，我怕你——你会把我吃掉的！”"
)

func BenchmarkReqIndex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Index(req, "CDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDA")
	}
}

func BenchmarkReqStringsIndex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strings.Index(req, "CDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDAABCDBCDAABCAACDA")
	}
}

func BenchmarkRandIndex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Index(randStr, "FtxEVfwpZwBiUyNTpWBfvuAaGmgULhfwCgszyNdCDHeDYORwhNUGMYhcYBZJRvlxhAFJoeOXGMpHKtoupLtYTQJmpkqApfDRkKFaTaBMKEXGMNoSZpFQxHFhJppr")
	}
}

func BenchmarkRandStringsIndex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strings.Index(randStr, "FtxEVfwpZwBiUyNTpWBfvuAaGmgULhfwCgszyNdCDHeDYORwhNUGMYhcYBZJRvlxhAFJoeOXGMpHKtoupLtYTQJmpkqApfDRkKFaTaBMKEXGMNoSZpFQxHFhJppr")
	}
}
