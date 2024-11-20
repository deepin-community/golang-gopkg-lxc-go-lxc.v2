// Copyright © 2013, 2014, The Go-LXC Authors. All rights reserved.
// Use of this source code is governed by a LGPLv2.1
// license that can be found in the LICENSE file.

//go:build linux && cgo
// +build linux,cgo

package lxc

import (
	"os"
)

// AttachOptions type is used for defining various attach options.
type AttachOptions struct {

	// Specify the namespaces to attach to, as OR'ed list of clone flags (syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS ...).
	Namespaces int

	// Specify the architecture which the kernel should appear to be running as to the command executed.
	Arch Personality

	// Cwd specifies the working directory of the command.
	Cwd string

	// UID specifies the user id to run as.
	UID int

	// GID specifies the group id to run as.
	GID int

	// Groups specifies the list of additional group ids to run with.
	Groups []int

	// If ClearEnv is true the environment is cleared before running the command.
	ClearEnv bool

	// Env specifies the environment of the process.
	Env []string

	// EnvToKeep specifies the environment of the process when ClearEnv is true.
	EnvToKeep []string

	// StdinFd specifies the fd to read input from.
	StdinFd uintptr

	// StdoutFd specifies the fd to write output to.
	StdoutFd uintptr

	// StderrFd specifies the fd to write error output to.
	StderrFd uintptr

	// RemountSysProc remounts /sys and /proc for the executed command.
	// This is required to reflect the container (PID) namespace context
	// if the command does not attach to the container's mount namespace.
	RemountSysProc bool

	// ElevatedPrivileges runs the command with elevated privileges.
	// The capabilities, cgroup and security module restrictions of the container are not applied.
	// WARNING: This may leak privileges into the container.
	ElevatedPrivileges bool
}

// DefaultAttachOptions is a convenient set of options to be used.
var DefaultAttachOptions = AttachOptions{
	Namespaces:         -1,
	Arch:               -1,
	Cwd:                "/",
	UID:                -1,
	GID:                -1,
	Groups:             nil,
	ClearEnv:           false,
	Env:                nil,
	EnvToKeep:          nil,
	StdinFd:            os.Stdin.Fd(),
	StdoutFd:           os.Stdout.Fd(),
	StderrFd:           os.Stderr.Fd(),
	RemountSysProc:     false,
	ElevatedPrivileges: false,
}

// TemplateOptions type is used for defining various template options.
type TemplateOptions struct {

	// Template specifies the name of the template.
	Template string

	// Backend specifies the type of the backend.
	Backend BackendStore

	BackendSpecs *BackendStoreSpecs

	// Distro specifies the name of the distribution.
	Distro string

	// Release specifies the name/version of the distribution.
	Release string

	// Arch specified the architecture of the container.
	Arch string

	// Variant specifies the variant of the image (default: "default").
	Variant string

	// Image server (default: "images.linuxcontainers.org").
	Server string

	// GPG keyid (default: 0x...).
	KeyID string

	// GPG keyserver to use.
	KeyServer string

	// Disable GPG validation (not recommended).
	DisableGPGValidation bool

	// Flush the local copy (if present).
	FlushCache bool

	// Force the use of the local copy even if expired.
	ForceCache bool

	// ExtraArgs provides a way to specify template specific args.
	ExtraArgs []string
}

// BackendStoreSpecs represents a LXC storage backend.
type BackendStoreSpecs struct {
	FSType string
	FSSize uint64
	Dir    *string
	ZFS    struct {
		Root string
	}
	LVM struct {
		VG, LV, Thinpool string
	}
	RBD struct {
		Name, Pool string
	}
}

// DownloadTemplateOptions is a convenient set of options for "download" template.
var DownloadTemplateOptions = TemplateOptions{
	Template: "download",
	Distro:   "ubuntu",
	Release:  "trusty",
	Arch:     "amd64",
}

// BusyboxTemplateOptions is a convenient set of options for "busybox" template.
var BusyboxTemplateOptions = TemplateOptions{
	Template: "busybox",
}

// UbuntuTemplateOptions is a convenient set of options for "ubuntu" template.
var UbuntuTemplateOptions = TemplateOptions{
	Template: "ubuntu",
}

// ConsoleOptions type is used for defining various console options.
type ConsoleOptions struct {

	// Tty number to attempt to allocate, -1 to allocate the first available tty, or 0 to allocate the console.
	Tty int

	// StdinFd specifies the fd to read input from.
	StdinFd uintptr

	// StdoutFd specifies the fd to write output to.
	StdoutFd uintptr

	// StderrFd specifies the fd to write error output to.
	StderrFd uintptr

	// EscapeCharacter (a means <Ctrl a>, b maens <Ctrl b>).
	EscapeCharacter rune
}

// DefaultConsoleOptions is a convenient set of options to be used.
var DefaultConsoleOptions = ConsoleOptions{
	Tty:             -1,
	StdinFd:         os.Stdin.Fd(),
	StdoutFd:        os.Stdout.Fd(),
	StderrFd:        os.Stderr.Fd(),
	EscapeCharacter: 'a',
}

// CloneOptions type is used for defining various clone options.
type CloneOptions struct {

	// Backend specifies the type of the backend.
	Backend BackendStore

	// lxcpath in which to create the new container. If not set the original container's lxcpath will be used.
	ConfigPath string

	// Do not change the hostname of the container (in the root filesystem).
	KeepName bool

	// Use the same MAC address as the original container, rather than generating a new random one.
	KeepMAC bool

	// Create a snapshot rather than copy.
	Snapshot bool
}

// DefaultCloneOptions is a convenient set of options to be used.
var DefaultCloneOptions = CloneOptions{
	Backend: Directory,
}

// CheckpointOptions type is used for defining checkpoint options for CRIU.
type CheckpointOptions struct {
	Directory string
	Stop      bool
	Verbose   bool
}

// RestoreOptions type is used for defining restore options for CRIU.
type RestoreOptions struct {
	Directory string
	Verbose   bool
}

// MigrateOptions type is used for defining migrate options.
type MigrateOptions struct {
	Directory       string
	PredumpDir      string
	ActionScript    string
	Verbose         bool
	Stop            bool
	PreservesInodes bool
	GhostLimit      uint64
	FeaturesToCheck CriuFeatures
}

// ConsoleLogOptions type is used for defining console log options.
type ConsoleLogOptions struct {
	ClearLog       bool
	ReadLog        bool
	ReadMax        uint64
	WriteToLogFile bool
}
