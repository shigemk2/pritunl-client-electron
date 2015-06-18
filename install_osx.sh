# Pritunl
sudo cp -pR build/osx/Applications/Pritunl.app /Applications
sudo chown -R root:wheel /Applications/Pritunl.app

# Service
sudo mkdir -p /usr/local/bin
sudo cp build/osx/usr/local/bin/pritunl-service /usr/local/bin

# Tuntap
sudo mkdir -p /Library/Extensions
sudo cp -pR build/osx/Library/Extensions/tap.kext /Library/Extensions/
sudo chown -R root:wheel /Library/Extensions/tap.kext
sudo cp -pR build/osx/Library/Extensions/tun.kext /Library/Extensions/
sudo chown -R root:wheel /Library/Extensions/tun.kext
sudo mkdir -p /Library/LaunchDaemons
sudo cp build/osx/Library/LaunchDaemons/net.sf.tuntaposx.tap.plist /Library/LaunchDaemons
sudo cp build/osx/Library/LaunchDaemons/net.sf.tuntaposx.tun.plist /Library/LaunchDaemons

# Openvpn
sudo mkdir -p /usr/local/bin
sudo cp build/osx/usr/local/bin/pritunl-openvpn /usr/local/bin
