package tests

import (
	"testing"

	"github.com/WindomZ/quizzee"
	"github.com/WindomZ/testify/assert"
)

type QuestionTest struct {
	Text string
	Keys []string
}

var results = []QuestionTest{
	{
		Text: "5. 下列哪种动物不用冬眠？",
		Keys: []string{"动物", "冬眠"},
	},
	{
		Text: "下列名人中最年轻的是？",
		Keys: []string{"名人", "最年轻"},
	},
	{
		Text: "食用适量的海苔可以预防甲状腺肿大，是因为海苔中的哪种元素？",
		Keys: []string{"海苔", "元素"},
	},
	{
		Text: "1.座头鲸的“座头”来源于一座岛屿、一种药物、还是一种乐器？",
		Keys: []string{"座头"},
	},
	{
		Text: "下列哪一项不属于声音的特性？",
		Keys: []string{"声音", "特性"},
	},
	{
		Text: "下列哪一部不是李安导演的作品？",
		Keys: []string{"李安", "作品"},
	},
	{
		Text: "度量衡是我国的计量单位，其中的衡指哪个方面的标准？",
		Keys: []string{"衡指", "标准"},
	},
}

func TestQuestion_Parse(t *testing.T) {
	for _, r := range results {
		q := quizzee.NewQuestion(r.Text)
		assert.NoError(t, q.Parse())
		//t.Logf("%#v", q)
		assert.Equal(t, r.Keys, q.Keys)
	}
}
