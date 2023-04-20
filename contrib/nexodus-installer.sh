#!/bin/bash

# Provide a "yes" response to the warning prompt as a command line arg.
YES=""

OS="$(uname -s)"
info_message() {
    if [ -z "${1}" ]; then
        echo "info_message() requires a message"
        exit 1
    fi
    echo -e "[\033[0;34m ACTION \033[0m] ${1}"
}

pass_message() {
    if [ -z "${1}" ]; then
        echo "pass_message() requires a message"
        exit 1
    fi
    echo -e "[\033[0;32m PASSED \033[0m] ${1}"
}

error_message() {
    if [ -z "${1}" ]; then
        echo "error_message() requires a message"
        exit 1
    fi
    if [ -n "$1" ]; then
        echo -e "[\033[0;31m FAILED \033[0m] ] ${1}"
    fi
}

function setup() {
    if [ $OS != "Darwin" ] && [ $OS != "Linux" ]; then
        error_message "Currently supports Linux and Darwin."
        exit 1
    fi

    echo "WARNING -- This script will install software on your system. It will:"
    if [ $OS == "Darwin" ]; then
        cat <<EOF
  * Install Homebrew (brew) if not already installed.
  * Install wireguard-tools with brew.
  * Download the Nexodus agent and install it to /usr/local/bin/nexd
EOF
    elif [ $OS == "Linux" ]; then
        cat <<EOF
  * Use sudo to uninstall wireguard-tools using the system's package manager.
  * Use sudo to download the Nexodus agent and install it to /usr/local/sbin/nexd
EOF
    else
        echo "Please add warning message text for $OS"
        exit 1
    fi
    if [ "${YES}" != "y" ]; then
        echo -n "Continue? (y/n): "
        read ANSWER
        if ! [[ "${ANSWER}" =~ ^(y|Y)$ ]]; then
            echo "Aborting based on response."
            exit 1
        fi
    fi

    if [ $OS == "Darwin" ]; then
        if ! [ -x "$(command -v brew)" ]; then
            info_message "Brew is not installed. Installing Brew..."
            /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
            if [ "$?" == "0" ]; then 
                pass_message "Brew is installed successfully."
            else 
                error_message "Brew installation failed."
                exit 1
            fi
        fi
        if ! [ -x "$(command -v wg)" ]; then
            info_message "Wireguard is not installed. Installing WireGuard..."
            HOMEBREW_NO_AUTO_UPDATE=1 HOMEBREW_NO_INSTALL_CLEANUP=1 HOMEBREW_NO_ENV_HINTS=1 brew install wireguard-tools --quiet
            if [ "$?" == "0" ]; then 
                pass_message "WireGuard is installed successfully."
                wg --version
            else 
                error_message "WireGuard installation failed."
                exit 1
            fi
        fi
        info_message "Installing Nexodus..."
        arch=amd64
        if [[ "$(uname -a)" = *ARM64* ]]; then
          arch="arm64"
        fi
        sudo curl -fSL https://nexodus-io.s3.amazonaws.com/darwin-${arch}/nexd --output /usr/local/bin/nexd
        sudo chmod +x /usr/local/bin/nexd
        pass_message "Nexodus is installed successfully."
    fi

    if [ $OS == "Linux" ]; then
        . /etc/os-release
        linuxDistro="${NAME}"
        
        if ! [ -x "$(command -v wg)" ]; then
            info_message "Wireguard is not installed. Installing WireGuard..."
            if [ "$linuxDistro" == "Ubuntu" ]; then
                sudo DEBIAN_FRONTEND=noninteractive apt-get update -y
                sudo DEBIAN_FRONTEND=noninteractive apt-get -qq --no-install-recommends install wireguard wireguard-tools -y
            elif [ "$linuxDistro" == "CentOS Stream" ] || [ "$linuxDistro" == "Fedora Linux" ]; then
                sudo dnf -q install wireguard-tools -y
            else
                error_message "This script only support installing wireguard-tools on Ubuntu, Fedora and Centos Stream. Please install wireguard-tools and try again."
                exit 1
            fi

            if [ "$?" == "0" ]; then 
                pass_message "WireGuard is installed successfully."
                wg --version
            else 
                error_message "WireGuard installation failed."
                exit 1
            fi
        fi

        info_message "Installing Nexodus..."
        sudo curl -fsSL https://nexodus-io.s3.amazonaws.com/linux-amd64/nexd --output /usr/local/sbin/nexd
        sudo chmod +x /usr/local/sbin/nexd
        pass_message "Nexodus is installed successfully."

    fi
}

function cleanup() {
    if [ $OS == "Darwin" ]; then
        if [ -x "$(command -v wg)" ]; then
            info_message "Uninstalling WireGuard."
            HOMEBREW_NO_AUTO_UPDATE=1 HOMEBREW_NO_INSTALL_CLEANUP=1 HOMEBREW_NO_ENV_HINTS=1 brew remove wireguard-tools --quiet
            if [ "$?" == "0" ]; then 
                pass_message "WireGuard is uninstalled successfully."
            else 
                error_message "WireGuard uninstallation failed."
            fi
            sudo rm -f /usr/local/etc/wireguard/wg0-latest-rev.conf
            sudo rm -f /usr/local/etc/wireguard/wg0.conf
            sudo ifconfig wg0 down
        fi
            
        if [ -x "$(command -v brew)" ]; then
            info_message "Not uninstalling the Brew. If you would like to uninstall the brew please fire following commands"
            info_message 'echo | /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/uninstall.sh)"'
        fi

        sudo rm -f /usr/local/bin/nexd
        pass_message "Nexodus is uninstalled successfully."
    
    elif [ $OS == "Linux" ]; then
        . /etc/os-release
        linuxDistro="${NAME}"
        if [ -x "$(command -v wg)" ]; then
            info_message "Uninstalling WireGuard."
            if [ "$linuxDistro" == "Ubuntu" ]; then
                sudo DEBIAN_FRONTEND=noninteractive apt-get -qq purge wireguard wireguard-tools -y
            elif [ "$linuxDistro" == "CentOS Stream" ] || [ "$linuxDistro" == "Fedora" ]; then
                sudo dnf -q remove wireguard-tools -y
            else
                error_message "Currently only support Ubuntu, Fedora and Centos Stream."
            fi
            if [ "$?" == "0" ]; then 
                pass_message "WireGuard is uninstalled successfully."
            else 
                error_message "WireGuard uninstallation failed."
            fi
            sudo rm -f /etc/wireguard/wg0-latest-rev.conf
            sudo rm -f /etc/wireguard/wg0.conf
            sudo ip link del wg0
        fi
        sudo rm -f /usr/local/sbin/nexd
        pass_message "Nexodus is uninstalled successfully."
    fi
}

function help() {
    printf "\n"
    printf "Usage: %s [-iuh]\n" "$0"
    printf "\t-i Install Nexodus and all required dependencies.\n"
    printf "\t-y Provide \"yes\" response to install warning prompt in advance.\n"
    printf "\t-u Uninstall Nexodus and it's dependencies. \n"
    printf "\t-v Verbose output. \n"
    printf "\t-h help\n"
    exit 1
}

OP=""
V=""
while getopts "iuyvh" opt; do
    case $opt in
        i ) OP="setup";;
        u ) OP="cleanup";;
        y ) YES="y";;
        v ) V="y";;
        h ) help
        exit 0;;
        *) help
        exit 1;;
    esac
done
if [ $# -eq 0 ]; then 
    help;exit 0 
fi
if [ -n "${V}" ]; then
  set -x
fi
if [ "$OP" == "setup" ]; then
    setup
elif [ "$OP" == "cleanup" ]; then
    cleanup
else
    help
    exit 1
fi