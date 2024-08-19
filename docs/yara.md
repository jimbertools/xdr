# Installing Yara on Windows

To install Yara on Windows, you'll need to build the latest release from VirusTotal. Follow the steps below to ensure a successful installation.

## Prerequisites

1. Download the latest `yara-master-` ZIP folder and the `Source code` ZIP folder from the [VirusTotal Yara releases](https://github.com/VirusTotal/yara/releases) page.  
   **Important:** Only download from the releases tab, not from the latest commit.

## Installation Steps

1. **Extract and Set Up Yara:**
   - Extract the `yara-master-` ZIP folder to `C:/yara`.
   - Add `C:/yara` to your system environment variable `Path`.

2. **Prepare the Source Code:**
   - Extract the `Source code` ZIP folder to `C:/yara-source`.

3. **Install MSYS2:**
   - Download and install [MSYS2](https://www.msys2.org/).
   - **Important:** Close the `MSYS2.exe` window that opens after installation.

4. **Set Up MSYS2:**
   - Navigate to your MSYS2 installation directory and open the `MinGW.exe` terminal.
   - Run the following command in the terminal to install necessary packages:

     ```bash
     pacman -S mingw-w64-x86_64-toolchain mingw-w64-x86_64-gcc mingw-w64-x86_64-make mingw-w64-x86_64-pkg-config base-devel openssl-devel autoconf-wrapper automake libtool git
     ```

5. **Prepare for Compilation:**
   - In the `MinGW.exe` terminal, create a new folder called `Documents` under the root directory and navigate into it:

     ```bash
     mkdir Documents && cd Documents
     ```

   - Move the extracted Yara source code folder to `Documents`:

     ```bash
     mv C:/yara-source .
     ```

   - Change to the Yara source directory:

     ```bash
     cd yara-source
     ```

6. **Compile Yara:**
   - Run the `bootstrap` script:

     ```bash
     ./bootstrap.sh
     ```

   - Configure the build:

     ```bash
     ./configure
     ```

   - Compile the Yara binaries:

     ```bash
     make
     ```

   - **Note:** You may see errors during the make process, but as long as it completes, you can proceed.

7. **Install Yara:**
   - **Disable your antivirus and Windows Security** temporarily to avoid Yara being flagged as a virus.
   - Install Yara by running:

     ```bash
     make install
     ```

   - Add the MSYS2 `MinGW` bin folder (e.g., `C:\msys64\mingw64\bin`) to your system environment variable `Path`.

8. **Final Steps:**
   - Re-enable your antivirus.
   - Restart your computer to ensure all changes are applied.
   - Test the installation by building your Go project.

## Example Test Program

1. Initialize a `go.mod` file and install the Yara dependency:

   ```bash
   go mod init example.com/myyaraapp
   go get github.com/hillu/go-yara/v4
   ```

2. Use the following Go program to test Yara:

   ```go
   package main

   import (
       "fmt"
       "github.com/hillu/go-yara/v4"
   )

   func main() {
       c, err := yara.NewCompiler()
       if c == nil || err != nil {
           fmt.Println("Error creating compiler:", err)
           return
       }
       rule := `rule test {
           meta: 
               author = "Aviad Levy"
           strings:
               $str = "abc"
           condition:
               $str
       }`
       if err = c.AddString(rule, ""); err != nil {
           fmt.Println("Error adding YARA rule:", err)
           return
       }
       r, err := c.GetRules()
       if err != nil {
           fmt.Println("Error getting YARA rule:", err)
           return
       }
       var m yara.MatchRules
       err = r.ScanMem([]byte(" abc "), 0, 0, &m)
       if err != nil {
           fmt.Println("Error matching YARA rule:", err)
           return
       }
       fmt.Printf("Matches: %+v", m)
   }
   ```

3. **Compile and run the program with**:

   ```bash
   go build -ldflags "-extldflags=-static" -tags yara_static main.go
   ```

## Troubleshooting

- **Linker or GCC Errors:**  
  If you encounter errors related to undefined keywords from the Yara library, your Go environment might be misconfigured. Reinstall Go using the [Go installer](https://go.dev/dl/) for Windows. Make sure to build with the recommended build command for the first time.

- **General Issues:**  
  If other issues arise, ensure no conflicting entries exist in your system environment variables for Go, MinGW, or Yara. Uninstall any conflicting versions, then restart the installation process from the beginning.
