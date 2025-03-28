package utils

func ObterEmojiCategoria(categoria string) string {
	switch categoria {
	case "Estudo":
		return "📚"
	case "Trabalho":
		return "💼"
	case "Exercício":
		return "🏋️"
	case "Meta Pessoal":
		return "🎯"
	case "Casa":
		return "🏠"
	case "Compras":
		return "🛒"
	case "Cozinha":
		return "🍽️"
	case "Música":
		return "🎵"
	case "Entretenimento":
		return "🎮"
	default:
		return "📝"
	}
}
func StatusEmoji(estado bool) string {
	if estado {
		return "✅ Concluído"
	}
	return "❌ Pendente"
}
