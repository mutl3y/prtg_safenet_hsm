# PRTG_dns

Custom sensor for PRTG to allow you to check a safenet HSM device

Tested with PRTG Version 19.3.51.2722

Place binary in PRTG folder on monitored host (this may be different on your install)
- Windows: C:\Program Files (x86)\PRTG Network Monitor\Custom Sensors\EXEXML
- Linux: /var/prtg/scriptsxml

This can also be compiled onto any Golang supported platform

Linux and windows versions will be found in the release pages of Github
## Note on Linux Support:


## To compile this yourself you need to...
-    install Golang
-    download or clone repo
-    run `go get` to download required packages
-    run `go build`
-    move the binary to the correct place
    
There are likely to be other small steps here as things may vary on your systems, If you need a OS binary and 
not in a rush drop me a request    

Add this to PRTG as an advanced custom exe / ssh script

```
prtg_safenet_hsm-windows-amd64.exe
HSM checks for use in PRTG

Usage:
  prtg_safenet_hsm [command]

Available Commands:
  help        Help about any command
  vtl         vtl wrapper for PRTG

Flags:
  -h, --help   help for prtg_safenet_hsm

Use "prtg_safenet_hsm [command] --help" for more information about a command.

```
```

prtg_safenet_hsm-windows-amd64.exe vtl
vtl wrapper for PRTG

Usage:
  prtg_safenet_hsm vtl [command]

Available Commands:
  verify      verify slot is available

Flags:
  -h, --help   help for vtl

Use "prtg_safenet_hsm vtl [command] --help" for more information about a command.
```
```
prtg_safenet_hsm-windows-amd64.exe vtl verify -h
verify slot is available and return slot count
also offers serial number checking

Usage:
  prtg_safenet_hsm vtl verify [flags]

Flags:
  -d, --dir string      path to vtl app (default "/usr/safenet/lunaclient/bin")
  -e, --exe string      filename of vtl app (default "vtl")
  -h, --help            help for verify
  -S, --serial string   serial number to check



```


If you feel like saying thanks    
        XMR: 49QA139gTEVMDV9LrTbx3qGKKEoYJucCtT4t5oUHHWfPBQbKc4MdktXfKSeT1ggoYVQhVsZcPAMphRS8vu8oxTf769NDTMu
	

With thanks to Jetbrains and their support of the open source community
![ https://www.jetbrains.com/?from=JJ-s-XMR-STAK-HashRate-Monitor-and-Restart-Tool](jetbrains-variant-3.png?v=4&s=200)
 
     

	
