// +build linux

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2020 WireGuard LLC. All Rights Reserved.
 */

package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"

	godnschange "github.com/inverse-inc/go-dnschange"
	"github.com/inverse-inc/packetfence/go/sharedutils"
	"github.com/inverse-inc/wireguard-go/binutils"
	"github.com/inverse-inc/wireguard-go/device"
	"github.com/inverse-inc/wireguard-go/ipc"
	"github.com/inverse-inc/wireguard-go/tun"
	"github.com/inverse-inc/wireguard-go/util"
	"github.com/joho/godotenv"

	_ "net/http/pprof"
)

const (
	ExitSetupSuccess = 0
	ExitSetupFailed  = 1
)

const (
	ENV_WG_TUN_FD             = "WG_TUN_FD"
	ENV_WG_UAPI_FD            = "WG_UAPI_FD"
	ENV_WG_PROCESS_FOREGROUND = "WG_PROCESS_FOREGROUND"
)

var ENV_ID = sharedutils.EnvOrDefault("ID", "")

var logger *device.Logger
var DNSChange *godnschange.DNSStruct

func printUsage() {
	fmt.Printf("usage:\n")
	fmt.Printf("%s [-f/--foreground] INTERFACE-NAME\n", os.Args[0])
}

func warning() {
	if runtime.GOOS != "linux" || os.Getenv(ENV_WG_PROCESS_FOREGROUND) == "1" {
		return
	}

	fmt.Fprintln(os.Stderr, "┌───────────────────────────────────────────────────┐")
	fmt.Fprintln(os.Stderr, "│                                                   │")
	fmt.Fprintln(os.Stderr, "│   Running this software on Linux is unnecessary,  │")
	fmt.Fprintln(os.Stderr, "│   because the Linux kernel has built-in first     │")
	fmt.Fprintln(os.Stderr, "│   class support for WireGuard, which will be      │")
	fmt.Fprintln(os.Stderr, "│   faster, slicker, and better integrated. For     │")
	fmt.Fprintln(os.Stderr, "│   information on installing the kernel module,    │")
	fmt.Fprintln(os.Stderr, "│   please visit: <https://wireguard.com/install>.  │")
	fmt.Fprintln(os.Stderr, "│                                                   │")
	fmt.Fprintln(os.Stderr, "└───────────────────────────────────────────────────┘")
}

func main() {
	defer binutils.CapturePanic()

	godotenv.Load(os.Args[1])

	// get log level (default: info)

	logLevel := func() int {
		switch os.Getenv("LOG_LEVEL") {
		case "debug":
			return device.LogLevelDebug
		case "info":
			return device.LogLevelInfo
		case "error":
			return device.LogLevelError
		case "silent":
			return device.LogLevelSilent
		}
		return device.LogLevelInfo
	}()

	if len(os.Args) > 2 && os.Args[2] == "--master" {
		logger = device.NewLogger(
			logLevel,
			fmt.Sprintf("(%s) ", "Master"),
		)
		setMasterProcess()
		DNSChange = StartDNS()
		go checkParentIsAlive()

		for {
			binutils.RunTunnelFG(os.Args[1])
		}
	} else {

		var foreground = true
		var interfaceName = "wg0"

		if !foreground {
			foreground = os.Getenv(ENV_WG_PROCESS_FOREGROUND) == "1"
		}

		// open TUN device (or use supplied fd)

		tun, err := func() (tun.Device, error) {
			tunFdStr := os.Getenv(ENV_WG_TUN_FD)
			if tunFdStr == "" {
				return tun.CreateTUN(interfaceName, device.DefaultMTU)
			}

			// construct tun device from supplied fd

			fd, err := strconv.ParseUint(tunFdStr, 10, 32)
			if err != nil {
				return nil, err
			}

			err = syscall.SetNonblock(int(fd), true)
			if err != nil {
				return nil, err
			}

			file := os.NewFile(uintptr(fd), "")
			return tun.CreateTUNFromFile(file, device.DefaultMTU)
		}()

		if err == nil {
			realInterfaceName, err2 := tun.Name()
			if err2 == nil {
				interfaceName = realInterfaceName
			}
		}

		logger = device.NewLogger(
			logLevel,
			fmt.Sprintf("(%s) ", interfaceName),
		)

		logger.Info.Println("Starting wireguard-go version", device.WireGuardGoVersion)

		logger.Debug.Println("Debug log enabled")

		if err != nil {
			logger.Error.Println("Failed to create TUN device:", err)
			os.Exit(ExitSetupFailed)
		}

		// open UAPI file (or use supplied fd)

		fileUAPI, err := func() (*os.File, error) {
			uapiFdStr := os.Getenv(ENV_WG_UAPI_FD)
			if uapiFdStr == "" {
				return ipc.UAPIOpen(interfaceName)
			}

			// use supplied fd

			fd, err := strconv.ParseUint(uapiFdStr, 10, 32)
			if err != nil {
				return nil, err
			}

			return os.NewFile(uintptr(fd), ""), nil
		}()

		if err != nil {
			logger.Error.Println("UAPI listen error:", err)
			os.Exit(ExitSetupFailed)
			return
		}
		// daemonize the process

		if !foreground {
			env := os.Environ()
			env = append(env, fmt.Sprintf("%s=3", ENV_WG_TUN_FD))
			env = append(env, fmt.Sprintf("%s=4", ENV_WG_UAPI_FD))
			env = append(env, fmt.Sprintf("%s=1", ENV_WG_PROCESS_FOREGROUND))
			files := [3]*os.File{}
			if os.Getenv("LOG_LEVEL") != "" && logLevel != device.LogLevelSilent {
				files[0], _ = os.Open(os.DevNull)
				files[1] = os.Stdout
				files[2] = os.Stderr
			} else {
				files[0], _ = os.Open(os.DevNull)
				files[1], _ = os.Open(os.DevNull)
				files[2], _ = os.Open(os.DevNull)
			}
			attr := &os.ProcAttr{
				Files: []*os.File{
					files[0], // stdin
					files[1], // stdout
					files[2], // stderr
					tun.File(),
					fileUAPI,
				},
				Dir: ".",
				Env: env,
			}

			path, err := os.Executable()
			if err != nil {
				logger.Error.Println("Failed to determine executable:", err)
				os.Exit(ExitSetupFailed)
			}

			process, err := os.StartProcess(
				path,
				os.Args,
				attr,
			)
			if err != nil {
				logger.Error.Println("Failed to daemonize:", err)
				os.Exit(ExitSetupFailed)
			}
			process.Release()
			return
		}

		device := device.NewDevice(tun, logger)

		logger.Info.Println("Device started")

		errs := make(chan error)
		term := make(chan os.Signal, 1)

		uapi, err := ipc.UAPIListen(interfaceName, fileUAPI)
		if err != nil {
			logger.Error.Println("Failed to listen on uapi socket:", err)
			os.Exit(ExitSetupFailed)
		}

		go func() {
			for {
				conn, err := uapi.Accept()
				if err != nil {
					errs <- err
					return
				}
				go device.IpcHandle(conn)
			}
		}()

		logger.Info.Println("UAPI listener started")

		startInverse(interfaceName, device)

		// wait for program to terminate
		signal.Notify(term, syscall.SIGTERM)
		signal.Notify(term, os.Interrupt)
		signal.Notify(term, syscall.SIGPIPE)

		select {
		case <-term:
		case <-errs:
		case <-device.Wait():
		}

		// clean up

		uapi.Close()
		device.Close()

		quit()
	}
}

func checkParentIsAlive() {
	util.CheckGUIIsAliveUNIX(quit)
}
