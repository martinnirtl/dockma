package emojis

var ErrorEmojis []string = []string{"🥵", "⛑", "🐙", "🐞", "🔥", "💥", "🍎", "🌶", "🚒", "🧨", "⛔️", "🟥"}
var SuccessEmojis []string = []string{"🦖", "🐢", "🌳", "🍏", "🥦", "✅", "🪀"}

func GetErrorEmoji() string {
	return ErrorEmojis[random(0, len(ErrorEmojis))]
}

func GetSuccessEmoji() string {
	return SuccessEmojis[random(0, len(SuccessEmojis))]
}
