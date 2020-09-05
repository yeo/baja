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
	if len(args) < 2 {
		color.Red("Usage: baja create node-type file-name")
		return 1
	}

	return Create(site, args[0], args[1])
}

func Create(site *baja.Site, dir, title string) int {
	re := regexp.MustCompile(`[^a-zA-Z]+`)

	slug := strings.Replace(title, " ", "-", -1)

	current_time := time.Now()

	slug = strings.ToLower(slug)
	slug = re.ReplaceAllString(slug, "-")

	// filename with date in it to help sorting
	slug = current_time.Format("2006-01-02") + "-" + slug

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
	color.Green("Create file %s", slug)

	return 0
}
