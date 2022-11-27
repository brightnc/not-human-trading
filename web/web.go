package web

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed all:dist
var embebDirDist embed.FS

// //go:embed dist/index.html
// var index embed.FS

func Web(f *fiber.App) {
	f.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embebDirDist),
		PathPrefix: "dist",
		Index:      "index.html",
		// Browse:     true,
	}))
	// f.Use("/dist/assets", filesystem.New(filesystem.Config{
	// 	Root:       http.FS(embebDirDist),
	// 	PathPrefix: "dist/assets",
	// 	Browse:     true,
	// }))
}
