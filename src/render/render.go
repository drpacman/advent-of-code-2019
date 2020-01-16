package render

import (
	"strings"
)

func GenerateBoardHtml(board [][]int64, chars map[int64]string) string {
	return generateBoard(board, chars, "<div><pre>", "</pre></div>")

}

func GenerateBoardText(board [][]int64, chars map[int64]string) string {
	return generateBoard(board, chars, "", "\n")
}

func generateBoard(board [][]int64, chars map[int64]string, rowStart, rowEnd string) string {
	var output strings.Builder
	for y := 0; y < len(board); y++ {
		if rowStart != "" {
			output.WriteString(rowStart)
		}
		for x := 0; x < len(board[y]); x++ {
			if chars != nil {
				entry := chars[board[y][x]]
				if entry == "" {
					entry = " "
				}
				output.WriteString(entry)
			} else {
				output.WriteRune(rune(board[y][x]))
			}

		}
		if rowEnd != "" {
			output.WriteString(rowEnd)
		}
	}
	return output.String()
}
