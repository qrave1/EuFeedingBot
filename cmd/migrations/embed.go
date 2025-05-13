package migrations

import "embed"

//go:embed migrate/*
var Embed embed.FS
