cask "berth" do
  version "1.2.0"
  sha256 "306e4cff9a23af1ef81c7bb604a998f69bac806b9393f50afccbe8ce7776303c"

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
