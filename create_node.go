package baja

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
)

func CreateNode(dir, title string) error {
	re := regexp.MustCompile(`[^a-zA-Z]+`)

	slug := strings.Replace(title, " ", "-", -1)
	slug = strings.ToLower(slug)
	slug = re.ReplaceAllString(slug, "-")

	file, err := os.Create("content/" + dir + "/" + slug + ".md")
	if err != nil {
		color.Red("Cannot create file in %s. Check directory permission. Err: err", dir, err)
		return err
	}

	defer file.Close()

	content := `+++
date = "%s"
title = "%s"
draft = true

tags = []
+++`
	fmt.Fprintf(file, fmt.Sprintf(content, time.Now().Format(time.RFC3339), title))

	return nil
}