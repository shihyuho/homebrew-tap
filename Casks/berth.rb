cask "berth" do
  version "1.2.1"
  sha256 "2add4dc78eb322de20c56c3bd8c794225fab80bf51edac14c0bf0cbd6789fa84"

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
