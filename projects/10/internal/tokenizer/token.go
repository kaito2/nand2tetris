package tokenizer

import "github.com/kaito2/nand2tetris/internal/types"

type Token struct {
	String string
	Type   types.Token
}

// TODO: intVal の場合にも string になってしまう（どのレイヤーで対応するかは要検討）。
func NewToken(text string) Token {
	tokenType := types.CheckTokenType(text)
	switch tokenType {
	// TODO: check unknown type?
	// NOTE: 文字列に加工が必要なものを分岐で処理
	case types.STRING_CONST:
		return Token{
			String: types.GetString(text),
			Type:   tokenType,
		}
	default:
		return Token{
			String: text,
			Type:   tokenType,
		}
	}
}

func (t Token) TypeString() string {
	return string(t.Type)
}
