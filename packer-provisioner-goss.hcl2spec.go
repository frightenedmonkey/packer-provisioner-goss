// Code generated by "mapstructure-to-hcl2 -type GossConfig"; DO NOT EDIT.

package main

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatGossConfig is an auto-generated flat version of GossConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatGossConfig struct {
	Version       *string           `cty:"version" hcl:"version"`
	Arch          *string           `cty:"arch" hcl:"arch"`
	URL           *string           `cty:"url" hcl:"url"`
	DownloadPath  *string           `mapstructure:"download_path" cty:"download_path" hcl:"download_path"`
	Username      *string           `cty:"username" hcl:"username"`
	Password      *string           `cty:"password" hcl:"password"`
	SkipInstall   *bool             `mapstructure:"skip_install" cty:"skip_install" hcl:"skip_install"`
	Inspect       *bool             `cty:"inspect" hcl:"inspect"`
	TargetOs      *string           `mapstructure:"target_os" cty:"target_os" hcl:"target_os"`
	Tests         []string          `cty:"tests" hcl:"tests"`
	RetryTimeout  *string           `mapstructure:"retry_timeout" cty:"retry_timeout" hcl:"retry_timeout"`
	Sleep         *string           `mapstructure:"sleep" cty:"sleep" hcl:"sleep"`
	UseSudo       *bool             `mapstructure:"use_sudo" cty:"use_sudo" hcl:"use_sudo"`
	SkipSSLChk    *bool             `mapstructure:"skip_ssl" cty:"skip_ssl" hcl:"skip_ssl"`
	GossFile      *string           `mapstructure:"goss_file" cty:"goss_file" hcl:"goss_file"`
	VarsFile      *string           `mapstructure:"vars_file" cty:"vars_file" hcl:"vars_file"`
	VarsInline    map[string]string `mapstructure:"vars_inline" cty:"vars_inline" hcl:"vars_inline"`
	VarsEnv       map[string]string `mapstructure:"vars_env" cty:"vars_env" hcl:"vars_env"`
	RemoteFolder  *string           `mapstructure:"remote_folder" cty:"remote_folder" hcl:"remote_folder"`
	RemotePath    *string           `mapstructure:"remote_path" cty:"remote_path" hcl:"remote_path"`
	Format        *string           `mapstructure:"format" cty:"format" hcl:"format"`
	FormatOptions *string           `mapstructure:"format_options" cty:"format_options" hcl:"format_options"`
}

// FlatMapstructure returns a new FlatGossConfig.
// FlatGossConfig is an auto-generated flat version of GossConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*GossConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatGossConfig)
}

// HCL2Spec returns the hcl spec of a GossConfig.
// This spec is used by HCL to read the fields of GossConfig.
// The decoded values from this spec will then be applied to a FlatGossConfig.
func (*FlatGossConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"version":        &hcldec.AttrSpec{Name: "version", Type: cty.String, Required: false},
		"arch":           &hcldec.AttrSpec{Name: "arch", Type: cty.String, Required: false},
		"url":            &hcldec.AttrSpec{Name: "url", Type: cty.String, Required: false},
		"download_path":  &hcldec.AttrSpec{Name: "download_path", Type: cty.String, Required: false},
		"username":       &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"password":       &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"skip_install":   &hcldec.AttrSpec{Name: "skip_install", Type: cty.Bool, Required: false},
		"inspect":        &hcldec.AttrSpec{Name: "inspect", Type: cty.Bool, Required: false},
		"target_os":      &hcldec.AttrSpec{Name: "target_os", Type: cty.String, Required: false},
		"tests":          &hcldec.AttrSpec{Name: "tests", Type: cty.List(cty.String), Required: false},
		"retry_timeout":  &hcldec.AttrSpec{Name: "retry_timeout", Type: cty.String, Required: false},
		"sleep":          &hcldec.AttrSpec{Name: "sleep", Type: cty.String, Required: false},
		"use_sudo":       &hcldec.AttrSpec{Name: "use_sudo", Type: cty.Bool, Required: false},
		"skip_ssl":       &hcldec.AttrSpec{Name: "skip_ssl", Type: cty.Bool, Required: false},
		"goss_file":      &hcldec.AttrSpec{Name: "goss_file", Type: cty.String, Required: false},
		"vars_file":      &hcldec.AttrSpec{Name: "vars_file", Type: cty.String, Required: false},
		"vars_inline":    &hcldec.AttrSpec{Name: "vars_inline", Type: cty.Map(cty.String), Required: false},
		"vars_env":       &hcldec.AttrSpec{Name: "vars_env", Type: cty.Map(cty.String), Required: false},
		"remote_folder":  &hcldec.AttrSpec{Name: "remote_folder", Type: cty.String, Required: false},
		"remote_path":    &hcldec.AttrSpec{Name: "remote_path", Type: cty.String, Required: false},
		"format":         &hcldec.AttrSpec{Name: "format", Type: cty.String, Required: false},
		"format_options": &hcldec.AttrSpec{Name: "format_options", Type: cty.String, Required: false},
	}
	return s
}
