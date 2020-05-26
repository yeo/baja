package node

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"

	"github.com/yeo/baja"
)

type CreateCommand struct{}

func (cmd *CreateCommand) ArgDesc() string {
	return "directory title"
}

func (cmd *CreateCommand) Help() string {
	return "Create a new post, directory is post type. title should wrap in quote"
}

func (cmd *CreateCommand) Run(site *baja.Site, args []string) int {
	fmt.Println(args)
	return Create(site, args[0], args[1])
}

func Create(site *baja.Site, dir, title string) int {
	re := regexp.MustCompile(`[^a-zA-Z]+`)

	slug := strings.Replace(title, " ", "-", -1)
	slug = strings.ToLower(slug)
	slug = re.ReplaceAllString(slug, "-")

	file, err := os.Create("content/" + dir + "/" + slug + ".md")
	if err != nil {
		color.Red("Cannot create file in %s. Check directory permission. Err: err", dir, err)
		return 1
	}

	defer file.Close()

	content := `+++
date = "%s"
title = "%s"
draft = true

tags = []
+++`
	fmt.Fprintf(file, fmt.Sprintf(content, time.Now().Format(time.RFC3339), title))

	return 0
}
