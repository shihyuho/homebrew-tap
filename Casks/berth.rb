cask "berth" do
  version "1.1.0"
  sha256 "276e4fb455af440daf7b9cb7e3872dbef8c2565a0768a1b9bbb138f8b5600e98"

  url "https://github.com/shihyuho/berth/releases/download/#{version}/Berth-arm64.zip"
  name "Berth"
  desc "Keep your macOS Dock on the display you choose"
  homepage "https://github.com/shihyuho/berth"

  depends_on macos: :ventura
  depends_on arch: :arm64

  app "Berth.app"

  caveats do
    unsigned_accessibility
  end
end
