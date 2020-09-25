package emojis

var errorEmojis []string = []string{"⛑", "🐙", "🐞", "🔥", "💥", "🍎", "🌶", "🚒", "🧨", "⛔️"}
var successEmojis []string = []string{"🦖", "🐢", "🌳", "🍏", "🥦", "✅", "🪀"}

// GetErrorEmoji returns negative emoji
func GetErrorEmoji() string {
	return errorEmojis[random(0, len(errorEmojis))]
}

// GetSuccessEmoji returns positive emoji
func GetSuccessEmoji() string {
	return successEmojis[random(0, len(successEmojis))]
}
